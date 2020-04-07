package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const filePath = "../../extensions.go"
const sourceURL = "http://data.iana.org/TLD/tlds-alpha-by-domain.txt"
const tmplExtentionsMap = `
package finder

// NOTE: THIS FILE WAS CREATED BY
// EXTENTIONS GENERATION TOOL. (github.com/zajann/finder/cmd/extensionGenerator)
// DO NOT EDIT. IF EDIT THIS FILE THE PROGRAM
// MIGHT NOT WORK PROPERLY.

var AllowExtension = map[string]struct{}{
	{{- range .List}}
		"{{.}}": struct{}{},
	{{- end}}
}
`

type TemplateData struct {
	List []string
}

func main() {
	list, err := getExtensionList()
	if err != nil {
		log.Fatal(err)
	}
	dataSource, err := createDataSource(list)
	if err != nil {
		log.Fatal(err)
	}
	os.Remove(filePath)
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.Write(dataSource); err != nil {
		log.Fatal(err)
	}
}

func getExtensionList() ([]string, error) {
	res, err := http.Get(sourceURL)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("Invalid Response: %d", res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	lists := strings.Split(string(data), "\n")
	return lists[1 : len(lists)-1], nil
}

func createDataSource(list []string) ([]byte, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("template").Parse(tmplExtentionsMap))
	if err := t.Execute(&buf, TemplateData{List: list}); err != nil {
		return nil, err
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}
	return src, nil
}
