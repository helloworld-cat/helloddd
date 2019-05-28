package main

import (
	"log"
	"net/http"

	"github.com/pagedegeek/helloddd/infra/web/handlers/productlisthandler"
	// "github.com/pagedegeek/helloddd/infra/web/renderers/html/productlistrenderer"
	"github.com/pagedegeek/helloddd/infra/web/renderers/json/productlistrenderer"

	"github.com/pagedegeek/helloddd/domain/services/productlistsrc"
)

func main() {
	router := http.NewServeMux()

	// HTML
	// productListRenderer, err := productlistrenderer.New(
	// 	nil,
	// 	"infra/web/templates",
	// 	"layout.tmpl",
	// 	"products.tmpl",
	// )
	// if err != nil {
	// 	log.Fatalf("cannot load products tempalte: %s", err)
	// }

	// JSON
	productListRenderer := productlistrenderer.New()

	productListSrc := productlistsrc.New()
	productListHandler := productlisthandler.New(productListSrc, productListRenderer)

	router.Handle("/", productListHandler)

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal(err)
	}
}
