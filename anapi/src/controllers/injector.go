package controllers

import "github.com/deltegui/phoenix"

func Register(app phoenix.App) {
	app.MapRoot(NotFoundError)
}
