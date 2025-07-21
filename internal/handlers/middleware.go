package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
)

func (h *Handler) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) RecoverAllPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			logger := h.logger
			if err := recover(); err != nil {
				pc, file, line, _ := runtime.Caller(4)
				funcName := runtime.FuncForPC(pc).Name()
				logger.WithFields(logrus.Fields{
					"panic": err,
					"file":  file,
					"line":  line,
					"func":  funcName,
				}).Error("Паника была обработана")
				http.Error(w, "Серверная ошибка", 500)
			}
			return
		}()
		next.ServeHTTP(w, r)
	})
}
