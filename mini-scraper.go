package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

func getMetaValue(property string, body string) string {
	sel, err := css.Compile(fmt.Sprint("meta[property=\"", property, "\"]"))
	if err != nil {
		panic(err)
	}
	node, err := html.Parse(strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	for _, ele := range sel.Select(node) {
		return ele.Attr[1].Val
	}

	return ""
}

func getTitle(body string) string {
	sel, err := css.Compile(fmt.Sprint("title"))
	if err != nil {
		panic(err)
	}
	node, err := html.Parse(strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	for _, ele := range sel.Select(node) {
		return ele.FirstChild.Data
	}

	return ""
}

func getImage(body string) string {
	sel, err := css.Compile("img")
	if err != nil {
		panic(err)
	}
	node, err := html.Parse(strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	for _, ele := range sel.Select(node) {

		for _, element := range ele.Attr {
			if element.Key == "src" {
				return element.Val
			}
		}
	}

	return ""
}

func main() {
	url := "https://borisaeric.io"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	title := getMetaValue("og:title", string(body))
	fmt.Println(title)

	if title == "" {
		title = getTitle(string(body))
	}

	image := getMetaValue("og:image", string(body))

	if image == "" {
		image = url + "/" + getImage(string(body))
	}

	fmt.Println(image)
}
