package main

// Generator provider
type Generator interface {
	Gen(tpl string, data interface{}) []byte
}
