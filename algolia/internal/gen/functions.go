//+build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"trimPrefix": strings.TrimPrefix,
	"title":      strings.Title,
}

func createTemplate(filename string) *template.Template {
	name := path.Base(filename)
	return template.Must(template.New(name).Funcs(funcMap).ParseFiles(filename))
}

func generateFile(tmpl *template.Template, data interface{}, filepath string) {
	var (
		b       bytes.Buffer
		content []byte
	)

	err := tmpl.Execute(&b, data)
	if err != nil {
		fmt.Printf("cannot execute template %s: %v", filepath, err)
		os.Exit(1)
	}

	content, err = format.Source(b.Bytes())
	if err != nil {
		fmt.Printf("cannot format generated code from template %s: %v", filepath, err)
		os.Exit(1)
	}

	os.Remove(filepath)

	if err = ioutil.WriteFile(filepath, content, 0644); err != nil {
		fmt.Printf("cannot write generated file from template %s: %v", filepath, err)
		os.Exit(1)
	}
}

func shouldBeGenerated(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return true
	}

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("cannot detect if file %s should be generated: cannot open file: %v\n", filepath, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "// Code generated by go generate. DO NOT EDIT.") {
			return true
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("cannot detect if file %s should be generated: error while reading lines: %v\n", filepath, err)
		os.Exit(1)
	}

	return false
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func generateFilename(jsonName string) (filename string) {
	filename = jsonName
	filename = matchFirstCap.ReplaceAllString(filename, "${1}_${2}")
	filename = matchAllCap.ReplaceAllString(filename, "${1}_${2}")
	filename = strings.ToLower(filename) + ".go"
	return
}

func convertInterfaceToString(defaultValue interface{}) string {
	if defaultValue == nil {
		return "nil"
	}

	var s string
	switch v := defaultValue.(type) {
	case bool:
		s = fmt.Sprintf("%t", v)
	case int:
		s = fmt.Sprintf("%d", v)
	case string:
		s = fmt.Sprintf("%q", v)
	case []string, map[string]string, map[string][]string:
		s = fmt.Sprintf("%#v", v)
	default:
		fmt.Printf("cannot convert interface to string: unhandled type %#v\n", defaultValue)
		os.Exit(1)
	}
	return s
}
