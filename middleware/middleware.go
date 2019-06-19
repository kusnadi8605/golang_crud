package middleware

import (
	"encoding/json"
	"fmt"
	conf "golang_crud/config"
	dts "golang_crud/datastruct"
	mdl "golang_crud/models"
	"log"
	"net/http"
	"time"
)

//Middleware ..
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

//Logging ..
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			//defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			log.Printf("%d %s", time.Since(start).Nanoseconds()/1e3, r.URL.Path)
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

//UserAgent ..
func ContentType(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Header.Get("Content-Type"))
			// Do middleware things
			if r.Header.Get("Content-Type") != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

func IsvalidToken(conn *conf.Connection) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			channel := r.Header.Get("channel")
			isValidToken, err := mdl.CheckToken(conn, r, channel)
			var MiddlewareResponse dts.MiddlewareResponse

			if isValidToken == false {
				MiddlewareResponse.ResponseCode = "301"
				MiddlewareResponse.ResponseDesc = err.Error()
				json.NewEncoder(w).Encode(MiddlewareResponse)

				//	logger.Logf("check Token Product : %v", MiddlewareResponse)

				return
			}
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
