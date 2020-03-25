package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/VivaLaPanda/isle-api/api/handlers"
	"github.com/dgraph-io/dgo"
	"github.com/gorilla/mux"
)

type key int

const (
	requestIDKey key = 0
)

var (
	healthy int32
)

// ServeAPI exposes the REST interface for the service.
// Queries will be passed to middleware and then the DAL
func ServeAPI(port int, db *dgo.Dgraph) {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	basePath := "/api"

	// amw, err := NewAuthMiddleware(authCfgFilename, basePath)
	// if err != nil {
	// 	logger.Fatalf("Couldn't find/parse provided auth config file. Err: %v\n", err)
	// }

	// Router setup
	baseRouter := mux.NewRouter()
	router := baseRouter.PathPrefix(basePath).Subrouter()
	//router.Use(amw.Middleware)
	router.Use(headerMiddleware)
	router.Handle("/", index()).
		Methods("GET")
	router.Handle("/content", handlers.NewContentNode(db)).
		Methods("POST")
	router.Handle("/content/{uid}", handlers.ExpandContentNode(db)).
		Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(notFound)

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	// Basic server setup
	listenAddr := fmt.Sprintf("localhost:%d", port)
	server := &http.Server{
		Addr:         listenAddr,
		Handler:      tracing(nextRequestID)(logging(logger)(router)),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server is ready to handle requests at", listenAddr, "/api")
	atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}

// notFound is the function in charge of responding to 404s
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "{\"error\":\"Endpoint not found. Doublecheck your query or take a look at the"+
		"docs: https://github.com/VivaLaPanda/isle-api\"}")
}

// index is a utility function to provide guidance if you hit the root
// TODO: eventually this should list all routes
func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "{\"message\":\"This is the UtaStream client API."+
			"Documentation on routes is at https://github.com/VivaLaPanda/isle-api\"}")
	})
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add headers to all responses
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}
