package io

import (
	// "fmt"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

//ReadYamlFile Read file from path marshal to an interface
func ReadYamlFile(path string, result interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err := yaml.Unmarshal(bytes, result); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}
	return nil
}

//ReadAll Read file from path
func ReadAll(path string) (string, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	result := string(byteValue)
	return result, err
}

// WriteFile write file from json data
func WriteFile(path string, json string) error {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(json)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

// ReadDir read all file from directory
func ReadDir(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	result := []string{}
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, errors.New("file is empty")
	}
	for i := 0; i < len(files); i++ {
		f := files[i]
		result = append(result, path+"/"+f.Name())
	}

	return result, nil
}
