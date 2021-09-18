// Gitr is for Github API calls
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// can access token
	_token := os.Getenv("GITHUB_TOKEN")

	log.Println("token exists:", 0 <= len(_token))

	contents := GetContents()

	for idx, ci := range contents {
		log.Println(idx, ci.Path, ci.GitUrl)

		if ci.Path == "README.md" {
			content := GetContent(ci.GitUrl)
			log.Println(idx, "GetContent")
			log.Println(content.Content)

			encoded := content.Content
			decoded, err := base64.StdEncoding.DecodeString(encoded)
			if err != nil {
				fmt.Println("decode error:", err)
				return
			}
			fmt.Println(string(decoded))

		}
	}

	log.Println("done")
}

//!-

func GetContent(content_url string) ContentSingle {
	url := content_url
	resp, err := http.Get(url)
	check(err)

	body, err := io.ReadAll(resp.Body)
	check(err)

	content := ContentSingle{}
	json.Unmarshal([]byte(body), &content)

	return content
}

func GetContents() []Content {
	url := "https://api.github.com/repos/aaronlelevier/d2a-config/contents/"
	resp, err := http.Get(url)
	check(err)

	body, err := io.ReadAll(resp.Body)
	check(err)

	contents := []Content{}
	json.Unmarshal([]byte(body), &contents)

	return contents
}

func GetRepo() Repo {
	url := "https://api.github.com/repos/aaronlelevier/d2a-config"
	resp, err := http.Get(url)
	check(err)

	body, err := io.ReadAll(resp.Body)
	check(err)

	repo := Repo{}
	json.Unmarshal([]byte(body), &repo)

	return repo
}

// check

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// structs

// ContentSingle .
type ContentSingle struct {
	Content string `json:"content"`
}

// Content .
type Content struct {
	Path   string `json:"path"`
	GitUrl string `json:"git_url"`
}

type Repos struct {
	Repos []Repo
}

type Repo struct {
	Id int `json:"id"`
}
