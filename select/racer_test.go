package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0)
		defer fastServer.Close()

		var (
			slowURL = slowServer.URL
			fastURL = fastServer.URL

			want     = fastURL
			got, err = Racer(slowURL, fastURL)
		)

		if want != got {
			t.Errorf("wanted: %q, got: %q", want, got)
		}
		if err != nil {
			t.Errorf("didn't want an error, but got one: %v", err)
		}
	})

	t.Run("return an error if none of the servers respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error, but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
	}))
}
