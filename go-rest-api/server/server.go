package server

import (
	"fmt"
	"net/http"
	"wallet-service/handler"
	"wallet-service/repository"
	"wallet-service/service"
)

type HttpServer struct {
	port uint
}

func NewServer(port uint) *HttpServer {
	return &HttpServer{
		port: port,
	}
}

func (s *HttpServer) Start() error {
	walletRepository := repository.NewWalletRepository()
	walletService := service.NewWalletService(walletRepository)
	walletHandler := handler.NewWalletHandler(walletService)
	http.HandleFunc("/", walletHandler.HandleRequest)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}
