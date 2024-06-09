package serve

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	DirectoryDescription = "/dir.md"
)

var (
	libPattern         = regexp.MustCompile(`\[([^]]+)]\(([^)]+)\)`)
	descriptionPattern = regexp.MustCompile(`\*Description\*: (.*)`)
	docsPattern        = regexp.MustCompile(`\*Docs\*: (.*)`)
	labelsPattern      = regexp.MustCompile(`\*Labels\*: (.*)`)
	linksPattern       = regexp.MustCompile(`(?m)\*Links\*: *\n( *-.*\n)+`)
	linkPattern        = regexp.MustCompile(`(?m)\s*-\s*(.*)\s*`)
	examplesPattern    = regexp.MustCompile(
		`(?m)\*Examples\*:[\s\n]+` + "```" + `(.*)[\s\n]+` + "([^`]*" + `[\s\n]+)+`,
	)
)

type (
	FileSearcher interface {
		LibPattern() *regexp.Regexp
		DocsPattern() *regexp.Regexp
		LabelsPattern() *regexp.Regexp
		LinksPattern() *regexp.Regexp
		ExamplesPattern() *regexp.Regexp
		DescriptionPattern() *regexp.Regexp
		Collect() ([]File, map[string]File, error)
	}

	fileSearcher struct {
		libPattern         *regexp.Regexp
		docsPattern        *regexp.Regexp
		labelsPattern      *regexp.Regexp
		examplesPattern    *regexp.Regexp
		descriptionPattern *regexp.Regexp
		linksPattern       *regexp.Regexp
	}
)

func NewFileSearcher() *fileSearcher {
	return &fileSearcher{
		descriptionPattern: descriptionPattern,
		docsPattern:        docsPattern,
		examplesPattern:    examplesPattern,
		labelsPattern:      labelsPattern,
		libPattern:         libPattern,
		linksPattern:       linksPattern,
	}
}

func (s *fileSearcher) LibPattern() *regexp.Regexp         { return s.libPattern }
func (s *fileSearcher) DocsPattern() *regexp.Regexp        { return s.docsPattern }
func (s *fileSearcher) LabelsPattern() *regexp.Regexp      { return s.labelsPattern }
func (s *fileSearcher) ExamplesPattern() *regexp.Regexp    { return s.examplesPattern }
func (s *fileSearcher) LinksPattern() *regexp.Regexp       { return s.linksPattern }
func (s *fileSearcher) DescriptionPattern() *regexp.Regexp { return s.descriptionPattern }

func (s *fileSearcher) Collect() ([]File, map[string]File, error) {
	files := make([]File, 0, 1024)
	metadata := make(map[string]File, 1024)
	return files, metadata, filepath.Walk("lib",
		func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			raw, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			if strings.HasSuffix(path, DirectoryDescription) {
				file := s.parseFileInfo(string(raw), path)
				metadata[file.Path] = s.parseFileInfo(string(raw), path)
				return nil
			}

			files = append(files, s.parseFileInfo(string(raw), path))
			return nil
		},
	)
}

func (s *fileSearcher) parseFileInfo(content, path string) File {
	f := File{Path: "/" + path}
	if m := s.LibPattern().FindStringSubmatch(content); len(m) != 0 {
		f.Name = m[1]
		f.Link = m[2]
	}
	if m := s.DescriptionPattern().FindStringSubmatch(content); len(m) != 0 {
		f.Description = m[1]
	}
	if m := s.DocsPattern().FindStringSubmatch(content); len(m) != 0 {
		f.Docs = m[1]
	}
	if m := s.LabelsPattern().FindStringSubmatch(content); len(m) != 0 {
		f.Labels = strings.Split(m[1], " ")
		for idx := range f.Labels {
			f.Labels[idx] = f.Labels[idx][1:]
		}
	}
	if m := s.ExamplesPattern().FindStringSubmatch(content); len(m) != 0 {
		f.ExamplesLanguage = m[1]
		f.Examples = m[2]
	}
	if m := s.LinksPattern().FindStringSubmatch(content); len(m) != 0 {
		for _, link := range linkPattern.FindAllStringSubmatch(m[0], -1) {
			f.ExtraLinks = append(f.ExtraLinks, link[1])
		}
	}
	return f
}
