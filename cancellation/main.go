package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	errGotEven = errors.New("ups we got even")
)

type RequestKey string

const (
	CancelKey RequestKey = "cancel"
)

func cancelContext(ctx context.Context, err error) {
	cancel, ok := ctx.Value(CancelKey).(context.CancelCauseFunc)
	if ok {
		cancel(err)
	}
}

func getOdd(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	if now.Minute()%2 == 1 {
		fmt.Fprintf(w, "now %v \n", now)
		return
	}

	cancelContext(r.Context(), errGotEven)
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			ctx := r.Context()

			msg := "OK"
			if err := context.Cause(ctx); err != nil {
				msg = err.Error()
				fmt.Fprintf(w, "err:%s", msg)
			}

			log.Printf("%s - %s %s %s %s %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI(), msg, time.Since(start))
		}()
		ctx := r.Context()
		ctx, cancel := context.WithCancelCause(ctx)
		ctx = context.WithValue(ctx, CancelKey, cancel)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func main() {
	srv := http.Server{
		Addr:         ":8888",
		WriteTimeout: 5 * time.Second,
		Handler:      http.TimeoutHandler(logger(http.HandlerFunc(getOdd)), 3*time.Second, "Timeout!\n"),
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
