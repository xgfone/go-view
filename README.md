# go-view

It is a manager of the HTML view template engine. You can register more than one template engine, and get it by the extension of the template filename to render HTML.

## Install

```bash
$ go get -u github.com/xgfone/go-view
```

## Example

```html
<!-- django_template.html -->
<html><head></head><body>{{ body }}</body></html>
```

```go
package main

import (
    "fmt"

    view "github.com/xgfone/go-view"
    "github.com/xgfone/go-view/django"
)

func main() {
    views := view.NewView()
    views.Register(django.NewEngine("/path/to", ".html"))
    html, _ := views.Execute(map[string]interface{}{"body": "abc"}, "django_template.html")
    fmt.Println(string(html))
    // <html><head></head><body>abc</body></html>
}
```