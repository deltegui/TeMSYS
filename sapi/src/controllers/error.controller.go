package controllers

import (
	"net/http"

	"github.com/deltegui/phoenix"
)

func NotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		phoenix.NewJSONPresenter(w).Present(struct {
			Code string `json:"code"`
		}{Code: "404"})
	}
}
