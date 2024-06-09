package serve

import (
	_ "embed"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	ResponseLimit = 36
)

var (
	//go:embed templates/lib.html.tpl
	libTpl string

	//go:embed templates/search.html.tpl
	searchTpl string

	//go:embed templates/tree.html.tpl
	treeTpl string

	//go:embed templates/base.html.tpl
	baseTpl string
)

type (
	Server struct {
		files    []File
		metadata map[string]File
	}
)

func Serve(fs FileSearcher) {
	files, metadata, err := fs.Collect()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("collected %d files\n", len(files))

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	server := http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: &Server{
			files:    files,
			metadata: metadata,
		},
		ReadTimeout: 5 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		http.Redirect(w, r, "/search/", http.StatusFound)
		return
	}

	switch {
	case strings.HasPrefix(path, "/assets"):
		http.ServeFile(w, r, "."+path)
	case strings.HasPrefix(path, "/search") || r.URL.Query().Get("q") != "":
		s.handleSearch(w, r, path)
	case strings.HasPrefix(path, "/lib"):
		s.handleLib(w, r, path)
	case strings.HasPrefix(path, "/tree"):
		s.handleTree(w, r, path)
	default:
		http.NotFound(w, r)
	}
}

func (s *Server) handleLib(w http.ResponseWriter, r *http.Request, path string) {
	for _, f := range s.files {
		if f.Path != path {
			continue
		}

		renderTemplate(w, "lib", libTpl, map[string]any{
			"UpLink": filepath.Join("/tree", filepath.Dir(f.Path)),
			"File":   f,
		})
		return
	}

	http.NotFound(w, r)
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request, path string) {
	q := strings.ToLower(r.URL.Query().Get("q"))
	if strings.HasPrefix(path, "/lib") {
		http.Redirect(w, r, "/search/?q="+q, http.StatusFound)
		return
	}

	var label string
	if strings.HasPrefix(path, "/search") && strings.Trim(path, "/") != "search" {
		label = strings.ToLower(strings.Replace(path, "/search/", "", 1))
	}

	responseFiles := make([]File, 0, ResponseLimit)
	for _, f := range s.files {
		if len(responseFiles) == ResponseLimit {
			break
		}

		f := f
		if label != "" && !f.hasLabel(label) {
			continue
		}
		if q != "" && !f.matchDescription(q) {
			continue
		}
		responseFiles = append(responseFiles, f)
	}
	sort.Slice(responseFiles, func(i, j int) bool {
		return strings.ToLower(responseFiles[i].Name) < strings.ToLower(responseFiles[j].Name)
	})

	renderTemplate(w, "search", searchTpl, map[string]any{
		"UpLink":        "/search",
		"ResponseFiles": responseFiles,
	})
}

func (s *Server) handleTree(w http.ResponseWriter, r *http.Request, path string) {
	path = strings.TrimRight(strings.Replace(path, "/tree", "", 1), "/") + "/"
	if path == "/" {
		path = "/lib/"
	}

	responsePaths := make([]File, 0, 16)
	included := make(map[string]struct{})
	for _, f := range s.files {
		if !strings.Contains(f.Path, path) {
			continue
		}

		nodePath := strings.Replace(f.Path, path, "", 1)
		if idx := strings.Index(nodePath, "/"); idx > -1 {
			nodePath = nodePath[:idx]
		}

		filePath := path + nodePath

		if _, ok := included[filePath]; ok {
			continue
		}
		included[filePath] = struct{}{}

		if !strings.HasSuffix(filePath, ".md") {
			responsePaths = append(responsePaths, File{
				Name:        nodePath,
				Description: s.metadata[filePath+DirectoryDescription].Description,
				Path:        "/tree" + filePath,
			})
			continue
		}

		f := f
		responsePaths = append(responsePaths, f)
	}

	renderTemplate(w, "tree", treeTpl, map[string]any{
		"UpLink":        "/tree" + filepath.Dir(strings.TrimSuffix(path, "/")),
		"ResponseFiles": responsePaths,
	})
}
