package middleware

import (
	"errors"
	"github.com/kkstas/goapi/api"
	"github.com/kkstas/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var usrname string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {

		}

	})
}
