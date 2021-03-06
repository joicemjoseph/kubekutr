package utils

import (
	"fmt"
	"io"
	"text/template"

	"github.com/knadh/stuffbin"
)

// parse takes in a template path and the variables to be "applied" on it. The rendered template
// is saved to the destination path.
func parse(src string, fs stuffbin.FileSystem) (*template.Template, error) {
	// read template file
	tmpl := template.New(src)
	// load default templates
	c, err := fs.Read("templates/containers.tmpl")
	if err != nil {
		return nil, fmt.Errorf("error parsing default template: %v", err)
	}
	tmpl, err = tmpl.Parse(string(c))
	if err != nil {
		return nil, fmt.Errorf("error parsing default template: %v", err)
	}
	// load main template
	f, err := fs.Read(src)
	if err != nil {
		return nil, fmt.Errorf("error reading template file %s: %v", src, err)
	}
	return tmpl.Parse(string(f))
}

func writeTemplate(tmpl *template.Template, config map[string]interface{}, dest io.Writer) error {
	// apply the variable and save the rendered template to the file.
	err := tmpl.Execute(dest, config)
	if err != nil {
		return err
	}
	return nil
}

func saveResource(template string, name string, dest io.Writer, config map[string]interface{}, fs stuffbin.FileSystem) error {
	// parse template file
	tmpl, err := parse(template, fs)
	if err != nil {
		return err
	}

	err = writeTemplate(tmpl, config, dest)
	if err != nil {
		return err
	}

	return nil
}
