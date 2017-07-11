# Usage

extract links from html string

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/biolee/linkExtractor"
)

func main() {
	u := "https://github.com/biolee"
	resp, err := http.Get(u)
	h, err := ioutil.ReadAll(resp.Body)
	urls, err := linkExtractor.Extract(string(h), u)
	fmt.Println(urls, err)
}
```