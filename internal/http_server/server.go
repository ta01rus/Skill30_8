package httpserver

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/ta01rus/Skill30_8/internal/storage"
)

var wait time.Duration

const (
	wrTimeOut      = time.Second * 15
	rdTimeOut      = time.Second * 15
	idTimeOut      = time.Second * 15
	maxHeaderBytes = 1 << 20
)

type HttpServer struct {
	http.Server
	host string
	port string
	db   storage.DB
}

func NewEnv() *HttpServer {
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")
	return New(host, port)
}

func New(host, port string) *HttpServer {

	connWs := net.JoinHostPort(host, port)
	db, err := storage.NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	return &HttpServer{
		Server: http.Server{
			Addr:           connWs,
			WriteTimeout:   wrTimeOut,
			ReadTimeout:    rdTimeOut,
			IdleTimeout:    idTimeOut,
			MaxHeaderBytes: maxHeaderBytes,
			Handler:        mux,
		},
		host: host,
		port: port,
		db:   db,
	}
}

// добавление маршрутов
func (hs *HttpServer) InitRoutes() {
	mx := hs.Handler.(*http.ServeMux)
	mx.HandleFunc("/", Home)
}

func (hs *HttpServer) Serve() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hs.BaseContext = func(_ net.Listener) context.Context { return ctx }

	log.Printf("http serve %s:%s", hs.host, hs.port)

	hs.InitRoutes()

	go func() {
		if err := hs.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	defer hs.Shutdown(ctx)

	c := make(chan os.Signal, 1)

	// https://ru.wikipedia.org/wiki/%D0%A1%D0%B8%D0%B3%D0%BD%D0%B0%D0%BB_(Unix)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)

	<-c

	log.Println("shutting down")
	os.Exit(0)

	return nil

}
