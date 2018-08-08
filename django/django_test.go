/*
Copyright 2018 xgfone

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package django

import (
	"os"
	"testing"

	view "github.com/xgfone/go-view"
)

func TestEngine(t *testing.T) {
	htmlData := `<html><head></head><body>{{ data }}</body></html>`
	filename := "_test_django_engine_.html"

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		t.Fail()
		return
	}
	file.WriteString(htmlData)
	file.Close()
	defer os.Remove(filename)

	tmpl := view.NewView()
	tmpl.Register(NewEngine("."))
	data, err := tmpl.Execute(map[string]interface{}{"data": "abc"}, filename)
	if err != nil || string(data) != "<html><head></head><body>abc</body></html>" {
		t.Fail()
	}
}
