package app

import (
	"context"
	"github.com/Slintox/chat_server/pkg/common/closer"
)

type App struct {
	configPath string

	serviceProvider ServiceProvider
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	app := &App{
		configPath: configPath,
	}

	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.configPath)
	return nil
}
