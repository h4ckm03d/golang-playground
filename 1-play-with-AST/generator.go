package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	AssociationNone      = 0
	AssociationBelongsTo = 1
	AssociationHasMany   = 2
	AssociationHasOne    = 3
)

type Model struct {
	Name      string
	Group     string
	Namespace string
	Fields    []*Field
}

func (m *Model) AllPreloadAssocs() []string {
	result := []string{}

	for _, field := range m.Fields {
		result = append(result, field.PreloadAssocs()...)
	}

	return result
}

type Models []*Model // implements Sort interface

func (m Models) Len() int {
	return len(m)
}

func (m Models) Less(i, j int) bool {
	return m[i].Name < m[j].Name
}

func (m Models) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type Field struct {
	Name        string
	JSONName    string
	Type        string
	Tag         string
	Association *Association
}

func (f *Field) PreloadAssocs() []string {
	if f.Association == nil || f.Association.Type == AssociationNone {
		return []string{}
	}

	result := []string{
		f.Name,
	}

	for _, field := range f.Association.Model.Fields {
		if field.Association == nil || field.Association.Type == AssociationNone {
			continue
		}

		result = append(result, f.Name+"."+field.Name)
	}

	return result
}

func (f *Field) IsAssociation() bool {
	return f.Association != nil && f.Association.Type != AssociationNone
}

func (f *Field) IsBelongsTo() bool {
	return f.Association != nil && f.Association.Type == AssociationBelongsTo
}

type Association struct {
	Type  int
	Model *Model
}

// CollectModels is a function to get all models in outModelDir
// It will return array of model and error
func CollectModels(outModelDir string) (Models, error) {
	files, err := ioutil.ReadDir(outModelDir)
	if err != nil {
		return nil, err
	}

	var models Models
	var wg sync.WaitGroup
	var mu sync.Mutex
	errCh := make(chan error, 1)
	done := make(chan bool, 1)

	for _, file := range files {
		wg.Add(1)
		go func(f os.FileInfo) {
			defer wg.Done()
			if f.IsDir() {
				return
			}

			if !strings.HasSuffix(f.Name(), ".go") {
				return
			}

			modelPath := filepath.Join(outModelDir, f.Name())
			ms, err := ParseModel(modelPath)
			if err != nil {
				errCh <- err
			}

			mu.Lock()

			for _, m := range ms {
				models = append(models, m)
			}

			mu.Unlock()
		}(file)
	}

	wg.Wait()
	close(done)

	select {
	case <-done:
	case err := <-errCh:
		if err != nil {
			return nil, err
		}
	}

	return models, nil
}
