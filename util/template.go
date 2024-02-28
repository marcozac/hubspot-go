package util

import (
	"bytes"
	"go/format"
	"os"
	"text/template"
)

// WriteTemplateToFile writes the template with the given data to the filename.
func WriteTemplateToFile(tmpl *template.Template, filename string, data any) error {
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		return err
	}
	bb, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(bb); err != nil {
		return err
	}
	return nil
}
