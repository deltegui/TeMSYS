package api

import (
	"encoding/json"
	"net/http"
	"temsys"

	"github.com/deltegui/phoenix"
)

func LoginHandler(loginCase temsys.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		var loginReq temsys.LoginRequest
		if err := json.NewDecoder(req.Body).Decode(&loginReq); err != nil {
			presenter.PresentError(temsys.MalformedRequestErr)
			return
		}
		loginCase.Exec(presenter, loginReq)
	}
}
