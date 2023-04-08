package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/supachai-sukd/sandboxtesting/config"
	"github.com/supachai-sukd/sandboxtesting/router"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func main() {

	cfg := config.New().All()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		logger.Fatal("unable to configure database", zap.Error(err))
	}
	config.InitDB(sql)

	routersInit := router.RegRoute(sql)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: routersInit,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("unexpected shutdown the server", zap.Error(err))
		}
		logger.Info("gracefully shutdown the server")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gCtx := context.Background()
	ctx, cancel := context.WithTimeout(gCtx, 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("unexpected shutdown the server", zap.Error(err))
	}
}
