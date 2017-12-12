package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		location := ""

		if strings.HasPrefix(url, "http://") {
			location = url
		} else {
			location = "http://" + url
		}

		resp, err := http.Get(location)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("\n\n\nStatus: %s", resp.Status)
		fmt.Printf("\n\n\n%d bytes copied\n\n\n", b)
	}
}
