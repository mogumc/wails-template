package service

import (
	"context"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return greet(name)
}

func (a *App) Flashtime() {
	flashtime(a)
}

func (a *App) Gettestjson() string {
	return getJSONString()
}
