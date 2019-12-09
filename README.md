# jdchum/sitex

Package `jdchum/sitex` reads the text content from a website regardless of styling, behavior, or structure. This package can be used to search site text for key words and phrases as well as monitoring text for changes.

## Install

```sh
go get -u github.com/jdchum/sitex
```

## Example

```go
package main

import (
    "io/ioutil"

    "github.com/jdchum/sitex"
)

const url = "https://en.wikipedia.org/wiki/Go_(programming_language)"

func main() {
    // Get the site's text
    text, err := sitex.GetSiteText(url, " ")
    if err != nil {
        panic(err)
    }

    // Output the text to disk
    err = ioutil.WriteFile("out.txt", []byte(text), 0644)
    if err != nil {
        panic(err)
    }
}

```

## API

### `sitex.GetSiteText(url, sep string) (text string, err error)`

> Attempts to parse all human-readable text from a webpage. "Invisible" text such as HTML tags, JavaScript, and CSS are ignored.

* `url` - URL of the webpage to fetch and parse
* `sep` - Separator to place between chunks of parsed text

Returns the text parsed from the webpage or an error if one occured.

## Limitations

Text is parsed as-is from the initial content returned by the server. This means that content requiring additional network requests or user interactions is not available to the parser.

## Roadmap

* [ ] Unicode support
* [ ] Parse visible text from attributes
* [ ] Follow server redirects
* [x] Parse embedded iframes
* [ ] Parse embedded PDF text

## License

MIT licensed. Copyright (c) 2019-2020 Joshua Chumbley. See the LICENSE file for details.
