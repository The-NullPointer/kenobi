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
		HttpPort = 8000
	}

	AllowedOrigins, ok := conf["AllowedOrigins"]
	if !ok {
		AllowedOrigins = []string{"*"}
	}

	AllowedMethods, ok := conf["AllowedMethods"]
	if !ok {
		AllowedMethods = []string{"GET", "HEAD", "POST", "OPTIONS", "PUT", "PATCH"}
	}

	AllowedHeaders, ok := conf["AllowedHeaders"]
	if !ok {
		AllowedHeaders = []string{"accept", "accept-encoding", "authorization", "content-type", "dnt", "origin", "user-agent", "x-csrftoken", "x-requested-with"}
	}

	cors := handlers.CORS(
		handlers.AllowedOrigins(AllowedOrigins.([]string)),
		handlers.AllowedMethods(AllowedMethods.([]string)),
		handlers.AllowedHeaders(AllowedHeaders.([]string)),
	)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", HttpPort),
		Handler:     cors(r.Router),
		ReadTimeout: 2 * time.Minute,
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
