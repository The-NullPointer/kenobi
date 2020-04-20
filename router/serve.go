package router

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (r *Router) Serve(ctx context.Context) {

	conf := r.App.Config
	HttpPort, ok := conf["HttpPort"]

	if !ok || HttpPort.(int) == 0 {
		logrus.Error("HttpPort is required")
	}

	AllowedOrigins, ok := conf["AllowedOrigins"]
	if !ok {
		logrus.Error("AllowedOrigins is required")
	}

	AllowedMethods, ok := conf["AllowedMethods"]
	if !ok {
		logrus.Error("AllowedMethods is required")
	}

	AllowedHeaders, ok := conf["AllowedHeaders"]
	if !ok {
		logrus.Error("AllowedHeaders is required")
	}

	ReadTimeout, ok := conf["ReadTimeout"]
	if !ok {
		logrus.Error("ReadTimeout is required")
	}

	cors := handlers.CORS(
		handlers.AllowedOrigins(AllowedOrigins.([]string)),
		handlers.AllowedMethods(AllowedMethods.([]string)),
		handlers.AllowedHeaders(AllowedHeaders.([]string)),
	)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", HttpPort),
		Handler:     cors(r.Router),
		ReadTimeout: ReadTimeout.(time.Duration),
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			logrus.Error(err)
		}
		close(done)
	}()

	logrus.Infof("starting http server http://127.0.0.1:%d", HttpPort)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}
	<-done

}
