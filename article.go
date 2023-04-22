package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type article struct {
	out        io.WriteCloser
	paragraphs []string
	pageHeader string
	headline   string
}

func (a *article) fillDefaults() {
	a.pageHeader = "h2"
	a.headline = "h3"
}

func (a *article) handle(tag string, value string) {
	if tag == "h1" {
		a.pageHeader = "h1"
		a.headline = "h2"
	}
	value = strings.TrimSpace(value)
	switch tag {
	case a.pageHeader:
		a.close()
		fmt.Printf("  %s: %s\n", tag, value)
		a.renew(value)
	case a.headline:
		a.flush()
		a.accumulate(value)
	default: // <p>
		a.accumulate(value)
	}
}

func (a *article) renew(value string) {
	fileName := decideFileName(value)
	fmt.Printf("    -> %s\n", fileName)
	a.out = createWriter(fileName)
	fmt.Fprintf(a.out, "%s\n\n", value)
}

var regexDateTitle = regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)

func decideFileName(value string) string {
	tokens := regexDateTitle.FindStringSubmatch(value)
	if len(tokens) != 4 {
		panic(fmt.Sprintf("Illegal date format in page header: %v", value))
	}
	fileName := fmt.Sprintf("%04s%02s%02s.txt", tokens[3], tokens[2], tokens[1])
	return fileName
}

func createWriter(fileName string) io.WriteCloser {
	out, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return out
}

func (a *article) accumulate(value string) {
	if a.out != nil && len(value) > 0 {
		a.paragraphs = append(a.paragraphs, value)
	}
}

var regexExtraSpace = regexp.MustCompile(`  +`)
var regexMissingSpace = regexp.MustCompile(`(\.)([A-ZÄÖÅ])`)
var regexFullStop = regexp.MustCompile(`(\.) ([A-ZÄÖÅ])`)

func (a *article) flush() {
	if a.out == nil || len(a.paragraphs) == 0 {
		return
	}

	fmt.Fprintln(a.out, "----")
	for _, paragraph := range a.paragraphs {
		fmt.Fprintf(a.out, "%s\n\n", paragraph)
	}

	fmt.Fprintln(a.out, "----")
	for _, paragraph := range a.paragraphs {
		paragraph = regexExtraSpace.ReplaceAllString(paragraph, " ")
		paragraph = regexMissingSpace.ReplaceAllString(paragraph, "$1 $2")
		paragraph = regexFullStop.ReplaceAllString(paragraph, "$1\n\n\n$2")
		fmt.Fprintf(a.out, "%s\n\n\n", paragraph)
	}

	a.paragraphs = nil
}

func (a *article) close() {
	a.flush()
	if a.out != nil {
		a.out.Close()
		a.out = nil
	}
}
