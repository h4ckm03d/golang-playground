package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/golang/protobuf/proto"
)

type Book struct {
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Pages    int      `json:"num_pages"`
	Chapters []string `json:"chapters"`
}

/*
syntax = "proto2";
package main;

message BookProto {
  required string title = 1;
  required string author = 2;
  optional int64 pages = 3;
  repeated string chapters = 4;
}
*/
// protoc --go_out=. book.proto

//go:generate msgp -tests=false
type BookDef struct {
	Title    string   `msg:"title"`
	Author   string   `msg:"author"`
	Pages    int      `msg:"num_pages"`
	Chapters []string `msg:"chapters"`
}

// go generate

func generateObject() *Book {
	return &Book{
		Title:    "The Art of Computer Programming, Vol. 2",
		Author:   "Donald E. Knuth",
		Pages:    784,
		Chapters: []string{"Random numbers", "Arithmetic"},
	}
}

func generateMessagePackObject() *BookDef {
	obj := generateObject()
	return &BookDef{
		Title:    obj.Title,
		Author:   obj.Author,
		Pages:    obj.Pages,
		Chapters: obj.Chapters,
	}
}

func generateProtoBufObject() *BookProto {
	obj := generateObject()
	return &BookProto{
		Title:    proto.String(obj.Title),
		Author:   proto.String(obj.Author),
		Pages:    proto.Int64(int64(obj.Pages)),
		Chapters: obj.Chapters,
	}
}

func BenchmarkSerializationJSONMarshal(b *testing.B) {
	obj := generateObject()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSerializationJSONUnmarshal(b *testing.B) {
	out, err := json.Marshal(generateObject())
	if err != nil {
		panic(err)
	}

	obj := &Book{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = json.Unmarshal(out, obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSerializationProtoBufMarshal(b *testing.B) {
	obj := generateProtoBufObject()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := proto.Marshal(obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSerializationProtoBufUnmarshal(b *testing.B) {
	out, err := proto.Marshal(generateProtoBufObject())
	if err != nil {
		panic(err)
	}

	obj := &BookProto{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = proto.Unmarshal(out, obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSerializationMessagePackMarshal(b *testing.B) {
	obj := generateMessagePackObject()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := obj.MarshalMsg(nil)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSerializationMessagePackUnmarshal(b *testing.B) {
	obj := generateMessagePackObject()
	msg, err := obj.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	obj = &BookDef{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err = obj.UnmarshalMsg(msg)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSerializationGobMarshal(b *testing.B) {
	obj := generateObject()

	enc := gob.NewEncoder(ioutil.Discard)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := enc.Encode(obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSerializationGobUnmarshal(b *testing.B) {
	obj := generateObject()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(obj)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		err = enc.Encode(obj)
		if err != nil {
			panic(err)
		}
	}

	dec := gob.NewDecoder(&buf)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = dec.Decode(&Book{})
		if err != nil {
			panic(err)
		}
	}
}
