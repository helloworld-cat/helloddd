package productlistrenderer

import (
	"encoding/json"
	"io"
)

func New() *renderer {
	return &renderer{}
}

func (r *renderer) Render(w io.Writer, data interface{}) error {
	blob, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if _, err := w.Write(blob); err != nil {
		return err
	}

	return nil
}

type (
	renderer struct {
	}
)
