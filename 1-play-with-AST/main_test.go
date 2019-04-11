package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRename(t *testing.T) {
	src := `package main

type Foo struct{}
`

	result := `package main

type Bar struct{}
`

	res, err := Rename("foo.go", "Bar", src)
	assert.Nil(t, err)
	assert.Equal(t, result, string(res))
}

func TestParser(t *testing.T) {
	models, err := ParseModel("./service.go")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(models))
}
