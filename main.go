package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		process("https://yle.fi/uutiset/osasto/selkouutiset/")
	} else {
		for _, target := range os.Args[1:] {
			process(target)
		}
	}
}

func process(target string) {
	fmt.Println(target)

	var in io.ReadCloser
	var err error
	if strings.HasPrefix(target, "https://") {
		resp, err := http.Get(target)
		if err != nil {
			panic(err)
		}
		in = resp.Body
	} else {
		in, err = os.Open(target)
		if err != nil {
			panic(err)
		}
	}

	doc, err := goquery.NewDocumentFromReader(in)
	if err != nil {
		panic(err)
	}
	in.Close()

	article := article{}
	doc.Find("article.content").Find("h1, h3, p").Each(func(i int, s *goquery.Selection) {
		article.handle(goquery.NodeName(s), s.Text())
	})
	article.close()
}
