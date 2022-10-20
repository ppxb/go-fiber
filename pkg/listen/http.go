package listen

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

func Http(options ...func(*HttpOptions)) error {
	ops := getHttpOptions(nil)
	for _, f := range options {
		f(ops)
	}

	host := ops.host
	port := ops.port
	ctx := ops.ctx
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: ops.handler,
	}

	defer func() {
		if err := recover(); err != nil {
			log.WithContext(ctx).WithError(errors.Errorf("%v", err)).Error("server run failed, stack: %s", string(debug.Stack()))
		}
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithContext(ctx).WithError(err).Error("server listen failed")
		}
	}()

	log.WithContext(ctx).Info("running at %s:%d", host, port)

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if ops.exit != nil {
		ops.exit()
	}

	log.WithContext(ctx).Info("shutting down")

	_, cancel := context.WithTimeout(ops.ctx, 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ops.ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("forced to shutdown failed")
	}
	log.WithContext(ctx).Info("exiting")
	return err
}
