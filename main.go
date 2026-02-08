package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", Url+Path, nil)
	if err != nil {
		panic(err)
	}

	SetDefaultHeaders(req) // ONLY basic headers that i have with Burp, no cookie ect... (like if it was the first request on a website)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	reader, err := gzip.NewReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
