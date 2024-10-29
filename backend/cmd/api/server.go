package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func (app *application) server() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdowmError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit)
		s := <-quit

		app.logger.PrintInfo("shutting down server", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdowmError <- err
		}

		app.logger.PrintInfo("completting background tasks", map[string]string{
			"addr": srv.Addr,
		})

		app.wg.Wait()
		shutdowmError <- err
	}()

	app.logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  app.cfg.env,
	})

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdowmError
	if err != nil {
		return err
	}

	app.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})

	return nil
}
