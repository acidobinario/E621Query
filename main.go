package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	tags := flag.String("tags", "", "Tags for query: tag+tag")
	pages := flag.Int("pages", 1, "Number of pages to query")
	limit := flag.Int("limit", 1, "Number of pics to display per page")
	flag.Parse()
	for pagenum := 0; pagenum <= *pages; pagenum++ {
		url := "https://e621.net/post/index.json?tags=" + *tags + "&pages=" + strconv.Itoa(pagenum) + "&limit=" + strconv.Itoa(*limit)
		e621CLient := http.Client{
			Timeout: time.Second * 2,
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Println("Request error: ", err)
		}

		req.Header.Set("User-Agent", "learning-go-uwu")

		res, getErr := e621CLient.Do(req)
		if getErr != nil {
			log.Println("Error making the request: ", getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		type Response struct {
			ID      int    `json:"id"`
			Author  string `json:"author"`
			FileUrl string `json:"file_url"`
		}

		response := []Response{}
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Println("Json Unmarshal error: ", jsonErr)
		}
		for _, author := range response {
			download(author.FileUrl)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func listFiles() []string {
	var files []string
	var output []string

	root := "./E621/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for i, file := range files {
		if i != 0 { //remove first line(root element)
			//fmt.Println("file ->", file, "i", i)
			output = append(output, file)
		}
	}
	return output
}

func download(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	exists := false
	for _, str := range listFiles() {
		if str == fileName {
			exists = true
			break
		}
	}
	if !exists {
		fmt.Println("Downloading", url, "to", fileName)

		// TODO: check file existence first with io.IsExist
		output, err := os.Create("E621/" + fileName)
		if err != nil {
			fmt.Println("Error while creating", fileName, "-", err)
			return
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}
		defer response.Body.Close()

		n, err := io.Copy(output, response.Body)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}

		fmt.Println(n, "bytes downloaded.")
	}
}
