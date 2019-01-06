package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"to-do-list/models"
	u "to-do-list/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

//JWTAuthentication jwt implementation
var JWTAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//List of endpoints that doesn't require auth
		notAuth := []string{"/api/user/new", "/api/user/login"}
		// Current request path
		requestPath := r.URL.Path
		//check if request does not need authentication, serve the request if it doesn't need it
		for _, v := range notAuth {
			if v == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		// Grab the token from the header
		tokenHeader := r.Header.Get("Authorization")
		// ifToken is missing, returns with error code 403 Unauthorized
		if tokenHeader == "" {
			response = u.Message(false, "Missing Auth Token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		/* The token normally comes in format `Bearer {token-body}`,
		we check if the retrieved token matched this requirement.
		*/
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Grab the token part, what we are truly interested in
		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token password")), nil
		})

		//Malformed token, returns with http code 403 as usual
		if err != nil {
			response := u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Token is invalid, maybe not signed on this server
		if !token.Valid {
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		/*
			Everything went well, proceed with the request and
			set the caller to the user retrieved from the parsed token.
			Useful for monitoring
		*/
		fmt.Sprintf("User %s", tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
