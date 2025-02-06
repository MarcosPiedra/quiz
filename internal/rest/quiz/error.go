package quiz

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error          string `json:"error"`
	HTTPStatusCode int    `json:"status"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	render.SetContentType(render.ContentTypeJSON)
	return nil
}

func Err(err error, code int) render.Renderer {
	return &ErrorResponse{
		Error:          err.Error(),
		HTTPStatusCode: code,
	}
}
