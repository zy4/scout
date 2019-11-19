// Package scout implement a middleware for report panic to sentry.io.
package scout

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/middleware"
)

// SentryRecovery recover from panic, report the error back to sentry, log it,
// and return an internal server error back to the user.
func SentryRecovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rval := recover(); rval != nil {
				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rval, debug.Stack())
				} else {
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rval)
					debug.PrintStack()
				}
				rvalStr := fmt.Sprint(rval)
				var event = sentry.NewEvent()
				event.Message = rvalStr
				sentry.CaptureEvent(event)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
