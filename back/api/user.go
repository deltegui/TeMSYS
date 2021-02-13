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

func CreateUserHandler(createUserCase temsys.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		var userReq temsys.CreateUserRequest
		if err := json.NewDecoder(req.Body).Decode(&userReq); err != nil {
			presenter.PresentError(temsys.MalformedRequestErr)
			return
		}
		createUserCase.Exec(presenter, userReq)
	}
}
