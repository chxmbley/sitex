package sitex

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// GetSiteText attempts to parse all human-readable text from a webpage.
// "Invisible" text such as HTML tags, JavaScript, and CSS are ignored.
func GetSiteText(url, sep string) (text string, err error) {
	// Make the web request
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// We need to extract all text visible from the page, which is typically
	// stored in text between opening and closing tags.
	tokenizer := html.NewTokenizer(resp.Body)

	var parentTag string

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err = tokenizer.Err()

			// End of document, break the loop and process what we have
			if err == io.EOF {
				err = nil
				break
			}

			// Error was not EOF and likely malformed HTML
			return "", err
		}

		token := tokenizer.Token()

		// Parse iframes recursively as their content is seen as part of the parent site
		if token.Data == "iframe" {
			for _, a := range token.Attr {
				if a.Key == "src" {
					var t string

					src := a.Val

					// Iframes with a src beginning with "//" instead of "http://" or "https://"
					// should be requested using the same protocol as the host page
					if strings.HasPrefix(src, "//") {
						src = strings.Split(url, "//")[0] + src
					}

					// Get the site text from the iframed page
					t, err = GetSiteText(src, sep)
					if err != nil {
						return
					}

					// Append the iframe's extracted text to the parent page's text and move on
					text += t + sep

					continue
				}
			}
		}

		// We only care about textual content that appears to users
		if tokenType != html.TextToken {
			// Keep track of the parent tag so we know whether the inner text is meaningful
			parentTag = token.Data
			continue
		}

		// The tags below contain text that is invisible, such as JavaScript, CSS, etc.
		// Text in these tags does not impact the textual content of the page
		if parentTag == "script" || parentTag == "style" || parentTag == "noscript" {
			continue
		}

		// If the text is nothing but whitespace, then it does not contain meaningful info
		line := strings.TrimSpace(token.Data)
		if line == "" {
			continue
		}

		// Append the line to the return text and add a line separator
		text += line + sep
	}

	// Remove any excess whitespace and return the text
	text = strings.TrimSpace(text)

	return
}
