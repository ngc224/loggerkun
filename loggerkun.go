package loggerkun

import (
	"io"
	"os"
	"text/template"
)

type Loggerkun struct {
	OUT io.Writer
	tpl *template.Template
}

func New(format string) (*Loggerkun, error) {
	tpl, err := template.New("loggerkun").Parse(format + "\n")

	if err != nil {
		return nil, err
	}

	return &Loggerkun{
		OUT: os.Stdout,
		tpl: tpl,
	}, nil
}

func (l *Loggerkun) PUT(data interface{}) error {
	return l.tpl.Execute(l.OUT, data)
}
