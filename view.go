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

// Package view is a manager of the HTML view template engine.
//
// You can register more than one template engine, and get it
// by the extension of the template filename to render HTML.
package view

import (
	"errors"
	"fmt"
	"path/filepath"
)

// ErrNoViewEngine represents that it cannot find the view engine by the extension.
var ErrNoViewEngine = errors.New("not found view engine")

// Engine is the interface which all view engines should be implemented.
type Engine interface {
	// Load (re)loads all the templates.
	Load() error

	// Execute executes and renders a template by its filename.
	Execute(data interface{}, filename string, filenames ...string) ([]byte, error)

	// Ext should return the final file extension which this view engine is responsible to render.
	Ext() string
}

// View is a view engine manager.
type View map[string]Engine

// NewView returns a new view engine manager.
func NewView() View {
	return make(View)
}

// Register registers a view engine.
func (v View) Register(e Engine) View {
	if _, ok := v[e.Ext()]; ok {
		panic(fmt.Errorf("The extension '%s' has been registered", e.Ext()))
	}
	v[e.Ext()] = e
	return v
}

// Find returns the view engine by the a filename extension.
func (v View) Find(filename string) Engine {
	return v[filepath.Ext(filename)]
}

// Execute renders the view by the html template and returns the result.
func (v View) Execute(data interface{}, filename string, filenames ...string) ([]byte, error) {
	if engine := v.Find(filename); engine != nil {
		return engine.Execute(data, filename, filenames...)
	}
	return nil, ErrNoViewEngine
}

// Load (re)loads all the templates.
func (v View) Load() error {
	for ext, engine := range v {
		if err := engine.Load(); err != nil {
			return fmt.Errorf("the %s template engine failed to load: %s", ext, err)
		}
	}
	return nil
}
