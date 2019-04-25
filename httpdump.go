package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
)

func dump(w http.ResponseWriter, r *http.Request) {
	var jsonData interface{}

	defer r.Body.Close()

	r.Header.Set("Host", r.Host)

	for _, header := range sortHeaderKeys(r.Header) {
		fmt.Printf("%s %s: %s\n", color.YellowString("[H]"), header, r.Header.Get(header))
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error: unable to read request body: %s", err)
		return
	}

	if getContentType(r) == "application/json" {
		if err := json.Unmarshal(body, &jsonData); err != nil {
			fmt.Printf("%s error: request Content-Type is application/json, but unable to process JSON %s\n", color.RedString("[!]"), err)
			fmt.Printf("%s %s\n", color.CyanString("[B]"), body)
		} else {
			jf := prettyjson.NewFormatter()
			j, _ := jf.Format(body)
			fmt.Printf("%s %s\n", color.CyanString("[B]"), j)
		}
	} else {
		fmt.Printf("%s %s\n", color.CyanString("[B]"), body)
	}
}

func getContentType(input interface{}) string {
	var header http.Header

	switch input.(type) {
	case *http.Request:
		header = input.(*http.Request).Header
	case *http.Response:
		header = input.(*http.Response).Header
	case http.ResponseWriter:
		header = input.(http.ResponseWriter).Header()
	default:
		return ""
	}

	contentType := header.Get("Content-Type")

	index := strings.Index(contentType, ";")
	if index != -1 {
		return contentType[:index]
	}

	return contentType
}

func sortHeaderKeys(headers http.Header) []string {
	var keys []string

	for k := range headers {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
