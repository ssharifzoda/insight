package server

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	ctx        context.Context
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1mb
		ReadTimeout:    12 * time.Second,
		WriteTimeout:   12 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) FirebaseConn() *messaging.Client {
	//opt := option.WithCredentialsFile(consts.FirebaseKeyFilePath)
	//app, err := firebase.NewApp(s.ctx, nil, opt)
	//if err != nil {
	//	log.Printf("Ошибка инициализации Firebase: %v", err)
	//	return nil
	//}
	//client, err := app.Messaging(s.ctx)
	//if err != nil {
	//	log.Printf("Ошибка создания клиента сообщений: %v", err)
	//	return nil
	//}
	//return client
	return nil
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
