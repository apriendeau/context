package context

import (
	"net/http"
	"sync"
	"time"
)

type context struct {
	// custom values go here
	Created time.Time
}

var (
	mutex sync.RWMutex
	c     = make(map[*http.Request]*context)
)

func Get(r *http.Request) *context {
	mutex.Lock()
	ctx, ok := c[r]
	if !ok {
		ctx = &context{
			Created: time.Now(),
		}
		c[r] = ctx
	}
	mutex.Unlock()
	return ctx
}

func Clear(r *http.Request) {
	mutex.Lock()
	clear(r)
	mutex.Unlock()
}

func clear(r *http.Request) {
	delete(c, r)
}

func ClearHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer Clear(r)
		h.ServeHTTP(w, r)
	})
}
