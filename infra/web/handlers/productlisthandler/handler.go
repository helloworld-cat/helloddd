package productlisthandler

import (
	"github.com/pagedegeek/helloddd/domain/entities/product"
	"io"
	"log"
	"net/http"
)

func New(src service, r renderer) *handler {
	return &handler{src: src, renderer: r}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	products, err := h.src.List()
	if err != nil {
		h.error(w, err)
		return
	}

	// FIXME
	// ... render template with productLsit
	data := map[string]interface{}{
		"Products": products,
	}
	if err := h.renderer.Render(w, data); err != nil {
		h.error(w, err)
		return
	}
}

func (h *handler) error(w http.ResponseWriter, err error) {
	log.Printf("Cannot serve product list handler: %s", err)
	w.WriteHeader(http.StatusInternalServerError)
}

type (
	handler struct {
		src      service
		renderer renderer
	}

	service interface {
		List() ([]*product.Entity, error)
	}

	renderer interface {
		Render(w io.Writer, data interface{}) error
	}
)
