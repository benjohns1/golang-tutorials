package restapi

import (
	json "encoding/json"
	"io"

	"github.com/benjohns1/golang-tutorials/chi-api/core"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Routes builds and returns the HTTP router
func Routes(unitOfWork core.UnitOfWork) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	reqDecoder := RequestDecoder{jsonDecoder}
	resFormatter := NewResponseFormatter(render.JSON)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo", todoRoutes(unitOfWork, chi.NewRouter(), resFormatter, reqDecoder))
	})

	return router
}

func jsonDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}

func errorResponse(err error) map[string]string {
	response := make(map[string]string)
	response["error"] = err.Error()
	return response
}

func successMessage(message string) map[string]string {
	response := make(map[string]string)
	response["message"] = message
	return response
}
