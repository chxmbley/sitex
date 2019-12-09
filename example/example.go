package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jdchum/sitex"
)

var url = flag.String("u", "https://en.wikipedia.org/wiki/Go_(programming_language)", "URL of the webpage to read")
var out = flag.String("o", "./out.txt", "File path to output result")

func main() {
	flag.Parse()

	// Prepend protocol prefix if not supplied
	if !strings.HasPrefix(*url, "http://") && !strings.HasPrefix(*url, "https://") {
		*url = "https://" + *url
	}

	// Get the site's text
	text, err := sitex.GetSiteText(*url, " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output the text to disk
	err = ioutil.WriteFile(*out, []byte(text), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
