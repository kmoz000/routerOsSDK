---
title: RouterOS SDK (Rest api client)
language_tabs:
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<h1 id="routeros-rest-schema-v7-12beta3-28240783-all-">RouterOS REST Schema (v7.12beta3.28240783-all) v7.12beta3.28240783-all</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Base URLs:

* <a href="https://{host}:{port}/rest">https://{host}:{port}/rest</a>

    * **host** -  Default: host

    * **port** -  Default: port

* <a href="http://{host}:{port}/rest">http://{host}:{port}/rest</a>

    * **host** -  Default: host

    * **port** -  Default: port

# Authentication

- HTTP Authentication, scheme: basic Mikrotik REST API only supports Basic Authentication, secured by HTTPS

<h1 id="routeros-rest-schema-v7-12beta3-28240783-all--default">Default</h1>

## POST_beep

<a id="opIdPOST_beep"></a>

> Code samples

```shell
# You can also use wget
curl -X POST https://{host}:{port}/rest/beep \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST https://{host}:{port}/rest/beep HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```go
// example.go
package main
import (
	"context"
	"fmt"
	"net/http"
)
var (
    username   string =  "admin"
    password   string = ""
    address    string = "http://192.168.1.1/rest"
)
func main(){
	c, err := NewClient(address, WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}))
	if err != nil {
		fmt.Printf("Error Debugging RouterOS device at %s\nError> \t %v", address, err.Error())
		return
	}
	if res.StatusCode() != 200 {
		fmt.Printf("Error Debugging RouterOS device at %s\nStatusCode> \t", address, res.StatusCode())
		return
	}
	fmt.Printf("results: %v", res)
}
```
