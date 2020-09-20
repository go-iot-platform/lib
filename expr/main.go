package expr

import (
	"bytes"
	"os"
	"path/filepath"

	"text/template"
)

// ParseText parse template
func ParseText(name string, body string, data interface{}) (string, error) {
	_template := template.New(name)
	_templateParse, err := _template.Parse(body)
	if err != nil {
		return "", err
	}
	var t = template.Must(_templateParse, nil)
	buf := bytes.Buffer{}
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ParseFile parse template
func ParseFile(path string, data interface{}) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fp := filepath.Join(wd, path)

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
