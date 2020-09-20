package io

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Key  string `yaml:"key"`
	Text string `yaml:"text"`
}

func TestContainer(t *testing.T) {
	path := "test/test.json"
	pathDir := "test/errors"
	pathWrite := "test/test_write.json"

	files := testReadDir(pathDir, t)
	for i := 0; i < len(files); i++ {
		result := &[]TestStruct{}
		testReadFile(files[i], result, t)
	}

	json := testReadAll(path, t)
	testWriteFile(pathWrite, json, t)

}
func testReadFile(path string, result *[]TestStruct, t *testing.T) error {
	err := ReadYamlFile(path, result)
	if err != nil {
		t.Fatal("read file error")
	}
	return nil
}
func testReadAll(path string, t *testing.T) string {
	result, err := ReadAll(path)
	if err != nil || result == "" {
		t.Fatal("read file error")
	}
	return result
}
func testWriteFile(path string, json string, t *testing.T) {
	err := WriteFile(path, json)
	if err != nil {
		t.Log(err)
		t.Fatal("write file error")
	}
}
func testReadDir(path string, t *testing.T) []string {
	result, err := ReadDir(path)
	fmt.Printf("%+v\n", result)
	if err != nil {
		t.Fatal("read dirctory error")
	}
	return result
}
