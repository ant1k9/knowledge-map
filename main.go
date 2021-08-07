package main

import (
	_ "embed"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type (
	File struct {
		Name        string
		Link        string
		Description string
		Docs        string
		Labels      []string
		Examples    string

		Path string
	}

	Server struct {
		files []File
	}
)

const (
	ResponseLimit = 20
)

var (
	libPattern         = regexp.MustCompile(`\[([^]]+)]\(([^)]+)\)`)
	descriptionPattern = regexp.MustCompile(`\*Description\*: (.*)`)
	docsPattern        = regexp.MustCompile(`\*Docs\*: (.*)`)
	labelsPattern      = regexp.MustCompile(`\*Labels\*: (.*)`)
	examplesPattern    = regexp.MustCompile(
		`(?m)\*Examples\*:[\s\n]+` + "```" + `.*[\s\n]+` + "([^`]+" + `[\s\n]+)+`,
	)
)

var (
	//go:embed templates/lib.html.tpl
	libTpl string

	//go:embed templates/search.html.tpl
	searchTpl string

	//go:embed templates/tree.html.tpl
	treeTpl string
)

func main() {
	files, err := collect()
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
			files: files,
		},
		ReadTimeout: 5 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func collect() ([]File, error) {
	files := make([]File, 0, 1024)
	return files, filepath.Walk("lib", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		raw, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(raw)

		f := File{Path: "/" + path}
		if m := libPattern.FindStringSubmatch(content); len(m) != 0 {
			f.Name = m[1]
			f.Link = m[2]
		}
		if m := descriptionPattern.FindStringSubmatch(content); len(m) != 0 {
			f.Description = m[1]
		}
		if m := docsPattern.FindStringSubmatch(content); len(m) != 0 {
			f.Docs = m[1]
		}
		if m := labelsPattern.FindStringSubmatch(content); len(m) != 0 {
			f.Labels = strings.Split(m[1], " ")
			for idx := range f.Labels {
				f.Labels[idx] = f.Labels[idx][1:]
			}
		}
		if m := examplesPattern.FindStringSubmatch(content); len(m) != 0 {
			f.Examples = m[1]
		}

		files = append(files, f)
		return nil
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		http.Redirect(w, r, "/search/", http.StatusFound)
		return
	}

	switch {
	case strings.HasPrefix(path, "/lib"):
		s.handleLib(w, r, path)
	case strings.HasPrefix(path, "/search"):
		s.handleSearch(w, r, path)
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

		t, err := template.New("lib").Parse(libTpl)
		if err != nil {
			log.Printf("error parsing lib template: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t.Execute(w, f)
		return
	}

	http.NotFound(w, r)
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request, path string) {
	var label string
	if strings.Trim(path, "/") != "search" {
		label = strings.ToLower(strings.Replace(path, "/search/", "", 1))
	}

	q := strings.ToLower(r.URL.Query().Get("q"))
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

	t, err := template.New("serach").Parse(searchTpl)
	if err != nil {
		log.Printf("error parsing search template: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sort.Slice(responseFiles, func(i, j int) bool {
		return strings.ToLower(responseFiles[i].Name) < strings.ToLower(responseFiles[j].Name)
	})

	t.Execute(w, responseFiles)
}

func (s *Server) handleTree(w http.ResponseWriter, r *http.Request, path string) {
	path = strings.TrimRight(strings.Replace(path, "/tree", "", 1), "/") + "/"
	if path == "/" {
		path = "/lib/"
	}

	responsePaths := make(map[string]File)
	for _, f := range s.files {
		if !strings.Contains(f.Path, path) {
			continue
		}

		nodePath := strings.Replace(f.Path, path, "", 1)
		if idx := strings.Index(nodePath, "/"); idx > -1 {
			nodePath = nodePath[:idx]
		}
		filePath := path + nodePath

		if !strings.HasSuffix(filePath, ".md") {
			responsePaths["/tree"+filePath] = File{
				Name: nodePath,
				Link: "/tree" + filePath,
			}
			continue
		}

		f := f
		responsePaths[filePath] = f
	}

	t, err := template.New("tree").Parse(treeTpl)
	if err != nil {
		log.Printf("error parsing tree template: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	t.Execute(w, responsePaths)
}

func (f *File) hasLabel(label string) bool {
	for _, l := range f.Labels {
		if strings.ToLower(l) == label {
			return true
		}
	}
	return false
}

func (f *File) matchDescription(q string) bool {
	for _, part := range strings.Split(q, " ") {
		if !strings.Contains(strings.ToLower(f.Description), part) &&
			!strings.Contains(strings.ToLower(f.Name), part) {
			return false
		}
	}
	return true
}
