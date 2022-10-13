package listen

import (
	"context"
	"fmt"
	"github.com/ppxb/go-fiber/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Http(options ...func(*HttpOptions)) {
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

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithContext(ctx).WithError(err).Error("[HTTP SERVER] listen failed")
		}
	}()

	log.WithContext(ctx).Info("[HTTP SERVER] running at %s:%d/%s", host, port, ops.urlPrefix)

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if ops.exit != nil {
		ops.exit()
	}

	log.WithContext(ctx).Info("[HTTP SERVER] shutting down...")

	_, cancel := context.WithTimeout(ops.ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ops.ctx); err != nil {
		log.WithContext(ctx).WithError(err).Error("[%s][HTTP SERVER] forced to shutdown failed")
	}

	log.WithContext(ctx).Info("[%s][HTTP SERVER] exiting")
}