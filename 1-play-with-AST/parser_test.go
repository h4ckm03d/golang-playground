package main

import (
	"sort"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	models, err := ParseModel("./service.go")
	assert.Nil(t, err)
	sort.Sort(models)
	modelMap := map[string]*Model{}
	for _, m := range models {
		modelMap[m.Name] = m
	}

	for _, model := range models {
		// Check association, stdout "model.Fields[0].Association.Type"
		resolveAssociate(model, modelMap, make(map[string]bool))
	}
	spew.Dump(models)
	assert.Equal(t, 2, len(models))
}
