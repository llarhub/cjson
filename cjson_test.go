//go:build llgo

package cjson

import (
	"testing"
	"unsafe"

	"github.com/goplus/lib/c"
	"github.com/llarhub/cjson"
)

func TestParseAndPrint(t *testing.T) {
	const input = `{"name":"llgo","count":42,"items":[1,2]}`
	root := cjson.Parse(c.AllocaCStr(input))
	if root == nil {
		t.Fatal("Parse returned nil")
	}
	defer root.Delete()
	if root.IsObject() == 0 || root.Type != cjson.Object {
		t.Fatal("expected a JSON object")
	}
	name := root.ObjectItemCaseSensitive(c.AllocaCStr("name"))
	if name == nil || name.IsString() == 0 || c.GoString(name.StringValue()) != "llgo" {
		t.Fatal("string value mismatch")
	}
	if c.GoString(name.Valuestring) != "llgo" || c.GoString(name.String) != "name" {
		t.Fatal("JSON struct fields mismatch")
	}
	count := root.ObjectItemCaseSensitive(c.AllocaCStr("count"))
	if count == nil || count.IsNumber() == 0 || count.NumberValue() != 42 || count.Valueint != 42 {
		t.Fatal("number value mismatch")
	}
	items := root.ObjectItemCaseSensitive(c.AllocaCStr("items"))
	if items == nil || items.IsArray() == 0 || items.ArraySize() != 2 || items.ArrayItem(1).NumberValue() != 2 {
		t.Fatal("array value mismatch")
	}
	printed := root.PrintUnformatted()
	if printed == nil {
		t.Fatal("PrintUnformatted returned nil")
	}
	defer cjson.Free(unsafe.Pointer(printed))
	if got := c.GoString(printed); got != input {
		t.Fatalf("PrintUnformatted = %q, want %q", got, input)
	}
}

func TestCreateObject(t *testing.T) {
	root := cjson.CreateObject()
	if root == nil {
		t.Fatal("CreateObject returned nil")
	}
	defer root.Delete()
	if root.AddStringToObject(c.AllocaCStr("name"), c.AllocaCStr("llgo")) == nil ||
		root.AddNumberToObject(c.AllocaCStr("count"), 42) == nil ||
		root.AddBoolToObject(c.AllocaCStr("ready"), 1) == nil {
		t.Fatal("could not add fields")
	}
	printed := root.PrintUnformatted()
	if printed == nil {
		t.Fatal("PrintUnformatted returned nil")
	}
	defer cjson.Free(unsafe.Pointer(printed))
	if got := c.GoString(printed); got != `{"name":"llgo","count":42,"ready":true}` {
		t.Fatalf("unexpected created object: %q", got)
	}
}

func TestParseWithLengthAndErrors(t *testing.T) {
	const input = `{"ok":true}`
	root := cjson.ParseWithLength(c.AllocaCStr(input), c.SizeT(len(input)+1))
	if root == nil {
		t.Fatal("ParseWithLength returned nil")
	}
	defer root.Delete()
	if field := root.ObjectItemCaseSensitive(c.AllocaCStr("ok")); field == nil || field.IsTrue() == 0 {
		t.Fatal("bool value mismatch")
	}
	if bad := cjson.Parse(c.AllocaCStr("{")); bad != nil {
		bad.Delete()
		t.Fatal("invalid JSON parsed successfully")
	}
	if cjson.GetErrorPtr() == nil {
		t.Fatal("missing parse error position")
	}
	if version := c.GoString(cjson.Version()); version != "1.7.19" {
		t.Fatalf("linked cJSON version = %q, want 1.7.19", version)
	}
}
