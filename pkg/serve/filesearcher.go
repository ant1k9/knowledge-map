package serve

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	libPattern         = regexp.MustCompile(`\[([^]]+)]\(([^)]+)\)`)
	descriptionPattern = regexp.MustCompile(`\*Description\*: (.*)`)
	docsPattern        = regexp.MustCompile(`\*Docs\*: (.*)`)
	labelsPattern      = regexp.MustCompile(`\*Labels\*: (.*)`)
	examplesPattern    = regexp.MustCompile(
		`(?m)\*Examples\*:[\s\n]+` + "```" + `(.*)[\s\n]+` + "([^`]*" + `[\s\n]+)+`,
	)
)

type (
	FileSearcher interface {
		LibPattern() *regexp.Regexp
		DocsPattern() *regexp.Regexp
		LabelsPattern() *regexp.Regexp
		ExamplesPattern() *regexp.Regexp
		DescriptionPattern() *regexp.Regexp
		Collect() ([]File, error)
	}

	fileSearcher struct {
		libPattern         *regexp.Regexp
		docsPattern        *regexp.Regexp
		labelsPattern      *regexp.Regexp
		examplesPattern    *regexp.Regexp
		descriptionPattern *regexp.Regexp
	}
)

func NewFileSearcher() *fileSearcher {
	return &fileSearcher{
		descriptionPattern: descriptionPattern,
		docsPattern:        docsPattern,
		examplesPattern:    examplesPattern,
		labelsPattern:      labelsPattern,
		libPattern:         libPattern,
	}
}

func (s *fileSearcher) LibPattern() *regexp.Regexp         { return s.libPattern }
func (s *fileSearcher) DocsPattern() *regexp.Regexp        { return s.docsPattern }
func (s *fileSearcher) LabelsPattern() *regexp.Regexp      { return s.labelsPattern }
func (s *fileSearcher) ExamplesPattern() *regexp.Regexp    { return s.examplesPattern }
func (s *fileSearcher) DescriptionPattern() *regexp.Regexp { return s.descriptionPattern }

func (s *fileSearcher) Collect() ([]File, error) {
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

		files = append(files, f)
		return nil
	})
}
