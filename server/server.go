package server

import (
	"context"
	golog "log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/henglory/Demo_Golang_v0.0.1/config"
	"github.com/henglory/Demo_Golang_v0.0.1/service"
)

type errorResponse struct {
	StatusCode int64  `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
	Request    string `json:"request"`
}

type Server struct {
	srv *http.Server
	s   service.Service
}

func NewServer(s service.Service) *Server {
	server := &Server{
		s: s,
	}
	return server
}

func (server *Server) Start() {
	go server.ginStart()
}

func (server *Server) Close() {
	server.srv.Close()
}

type readiness struct {
	Success bool `json:"success"`
}

func (server Server) ginStart() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/readiness", func(c *gin.Context) {
		c.JSON(200, readiness{Success: true})
	})

	r.POST("/demo1", func(c *gin.Context) {
		demo1(server.s, c)
	})

	r.POST("/demo2", func(c *gin.Context) {
		demo2(server.s, c)
	})

	server.srv = &http.Server{
		Addr:    config.ServicePort,
		Handler: r,
	}

	if err := server.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		golog.Fatalf("listen: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := server.srv.Shutdown(ctx); err != nil {
		golog.Fatal("Server Shutdown:", err)
	}
	defer cancel()
}
