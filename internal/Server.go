// Server triggers the server intialization and listening
package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sapiderman/mock-server/internal/config"
	"github.com/sapiderman/mock-server/internal/logger"
	log "github.com/sirupsen/logrus"
)

var (
	srvLog = log.WithField("module", "server")

	// StartUpTime records first ime up
	startUpTime time.Time

	// HTTPServer object
	HTTPServer *http.Server

	// Address of server
	address string
)

// InitializeServer initializes all server configs
func InitializeServer() error {
	logf := srvLog.WithField("fn", "InitializeServer")

	// load system / env configs
	config.LoadConfig()
	// configure logging
	logger.ConfigureLogging()

	startUpTime = time.Now()

	// t := time.Duration(config.GetInt("server.context.timeout"))
	// ctx, cancel := context.WithTimeout(context.Background(), t*time.Second)
	// defer cancel()

	logf.Info("Initialization done")
	return nil
}

// StartServer starts listening at given address:port
func StartServer() {

	var wait time.Duration
	logf := srvLog.WithField("fn", "StartServer")

	go func() {
		if err := HTTPServer.ListenAndServe(); err != nil {
			logf.Error(err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(gracefulStop, os.Interrupt)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	// Block until we receive our signal.
	<-gracefulStop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	HTTPServer.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logf.Info("shutting down........ bye")

	t := time.Now()
	upTime := t.Sub(startUpTime)
	fmt.Println(" ***** server was up for : ", upTime.String(), " *******")
	os.Exit(0)
}
