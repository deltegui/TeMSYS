package controllers

import (
	"net/http"

	"github.com/deltegui/phoenix"
)

func NotFoundError() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		phoenix.NewJSONPresenter(w).Present(struct{ Code string }{Code: "404"})
	}
}
