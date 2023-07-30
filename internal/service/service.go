package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/ta01rus/Skill30_8/pkg/storage"
	"github.com/ta01rus/Skill30_8/pkg/storage/postgres"
	"github.com/ta01rus/Skill30_8/pkg/tmphtml"
)

var wait time.Duration

type HttpServer struct {
	// моршруты
	Routes *gin.Engine

	Temps map[string]*tmphtml.FileHtmlTemplate

	Db storage.DB

	host string
	port string
}

func NewEnv() *HttpServer {
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")
	return New(host, port)
}

func New(host, port string) *HttpServer {
	db, err := postgres.NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	return &HttpServer{
		Routes: gin.Default(),
		Db:     db,
		host:   host,
		port:   port,
		Temps:  make(map[string]*tmphtml.FileHtmlTemplate),
	}
}

// добавление маршрутов
func (hs *HttpServer) InitRoutes() {
	//load html file

	hs.Routes.LoadHTMLFiles("templates/app.html")

	hs.Routes.GET("/", hs.HomeEndPoint)
	hs.Routes.Match([]string{http.MethodGet, http.MethodPost}, "/task/add", hs.AddTaskEndPoint)

	hs.Routes.POST("/task/:id", hs.TaskEndPoint)
	hs.Routes.POST("/task/:id/edit", hs.EditTaskEndPoint)
	hs.Routes.DELETE("/task/:id/del", hs.DelTaskEndPoint)

	hs.Routes.StaticFS("/static", http.Dir("./web"))
	hs.Routes.StaticFile("/favicon.ico", "./web/favicon.svg")
}

func (hs *HttpServer) Serve() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hs.InitRoutes()

	fmt.Printf(`Server start %s: %s`, hs.host, hs.port)

	srv := &http.Server{
		Addr:           net.JoinHostPort(hs.host, hs.port),
		WriteTimeout:   time.Second * 15,
		ReadTimeout:    time.Second * 15,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
		Handler:        hs.Routes,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	defer srv.Shutdown(ctx)

	c := make(chan os.Signal, 1)

	// https://ru.wikipedia.org/wiki/%D0%A1%D0%B8%D0%B3%D0%BD%D0%B0%D0%BB_(Unix)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)

	<-c

	log.Println("shutting down")
	os.Exit(0)

	return nil

}
