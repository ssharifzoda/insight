package api

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"net/http"
	"os"
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

func (h *Handler) TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(w, consts.RepeatSignIn, 401, 0)
			return
		}
		userId, _, sessionId, err := utils.ParseToken(authHeader)
		if err != nil {
			utils.ErrorResponse(w, consts.RepeatSignIn, 401, 0)
			return
		}
		userAuth, err := h.service.Authorization.GetAuthParamsByUserId(userId)
		if err != nil {
			utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
			return
		}
		token, err := jwt.ParseWithClaims(authHeader, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
		})
		if err != nil {
			utils.ErrorResponse(w, consts.InvalidToken, 403, 0)
			return
		}
		if sessionId != userAuth.SessionId {
			utils.ErrorResponse(w, consts.RepeatSignIn, 401, 0)
			return
		}
		if _, ok := token.Claims.(*utils.CustomClaims); ok && token.Valid {
			next.ServeHTTP(w, r)
		} else {
			utils.ErrorResponse(w, consts.InvalidToken, 403, 0)
		}
	})
}
