package httpserver

import (
	"log"
	"net"
	"net/http"

	"time"

	"github.com/ta01rus/Skill30_8/internal/storage"
)

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

func New(host, port string) *HttpServer {
	connWs := net.JoinHostPort(host, port)
	db, err := storage.NewEnv()
	if err != nil {
		log.Fatal(err)
	}

	return &HttpServer{
		Server: http.Server{
			Addr:           connWs,
			WriteTimeout:   wrTimeOut,
			ReadTimeout:    rdTimeOut,
			IdleTimeout:    idTimeOut,
			MaxHeaderBytes: maxHeaderBytes,
		},
		host: host,
		port: port,
		db:   db,
	}
}

func (hs *HttpServer) Serve() {

}
