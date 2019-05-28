package productlistrenderer

import (
	"fmt"
	"html/template"
	"io"
	"path"
)

func New(funcs template.FuncMap, dirname string, filenames ...string) (*renderer, error) {
	r := &renderer{}

	if funcs == nil {
		funcs = template.FuncMap{}
	}

	// Warning: erase func keys
	for k, v := range r.localFuncs() {
		funcs[k] = v
	}

	fullFilenames := make([]string, 0)
	for _, filename := range filenames {
		fullFilenames = append(fullFilenames, path.Join(dirname, filename))
	}

	tmpl, err := template.New("").Funcs(funcs).ParseFiles(fullFilenames...)
	if err != nil {
		return nil, err
	}

	r.tmpl = tmpl

	return r, nil
}

func (r *renderer) Render(w io.Writer, data interface{}) error {
	return r.tmpl.ExecuteTemplate(w, "layout", data)
}

func (r *renderer) localFuncs() template.FuncMap {
	return template.FuncMap{
		"formatPrice": func(priceCents uint) string {
			v := float64(priceCents) * 0.01
			return fmt.Sprintf("%.2f euros", v)
		},
	}
}

type renderer struct {
	tmpl *template.Template
}
