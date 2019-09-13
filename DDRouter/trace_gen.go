package DDRouter

import (
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

// wrapResponseWriter wraps an underlying http.ResponseWriter so that it can
// trace the http response codes. It also checks for various http interfaces
// (Flusher, Pusher, CloseNotifier, Hijacker) and if the underlying
// http.ResponseWriter implements them it generates an unnamed struct with the
// appropriate fields.
//
// This code is generated because we have to account for all the permutations
// of the interfaces.
func wrapResponseWriter(w http.ResponseWriter, span ddtrace.Span) http.ResponseWriter {
	hFlusher, okFlusher := w.(http.Flusher)
	hPusher, okPusher := w.(http.Pusher)
	hCloseNotifier, okCloseNotifier := w.(http.CloseNotifier)
	hHijacker, okHijacker := w.(http.Hijacker)

	w = newResponseWriter(w, span)
	switch {
	case okFlusher && okPusher && okCloseNotifier && okHijacker:
		w = struct {
			http.ResponseWriter
			http.Flusher
			http.Pusher
			http.CloseNotifier
			http.Hijacker
		}{w, hFlusher, hPusher, hCloseNotifier, hHijacker}
	case okFlusher && okPusher && okCloseNotifier:
		w = struct {
			http.ResponseWriter
			http.Flusher
			http.Pusher
			http.CloseNotifier
		}{w, hFlusher, hPusher, hCloseNotifier}
	case okFlusher && okPusher && okHijacker:
		w = struct {
			http.ResponseWriter
			http.Flusher
			http.Pusher
			http.Hijacker
		}{w, hFlusher, hPusher, hHijacker}
	case okFlusher && okCloseNotifier && okHijacker:
		w = struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
		}{w, hFlusher, hCloseNotifier, hHijacker}
	case okPusher && okCloseNotifier && okHijacker:
		w = struct {
			http.ResponseWriter
			http.Pusher
			http.CloseNotifier
			http.Hijacker
		}{w, hPusher, hCloseNotifier, hHijacker}
	case okFlusher && okPusher:
		w = struct {
			http.ResponseWriter
			http.Flusher
			http.Pusher
		}{w, hFlusher, hPusher}
	case okFlusher && okCloseNotifier:
		w = struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{w, hFlusher, hCloseNotifier}
	case okFlusher && okHijacker:
		w = struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
		}{w, hFlusher, hHijacker}
	case okPusher && okCloseNotifier:
		w = struct {
			http.ResponseWriter
			http.Pusher
			http.CloseNotifier
		}{w, hPusher, hCloseNotifier}
	case okPusher && okHijacker:
		w = struct {
			http.ResponseWriter
			http.Pusher
			http.Hijacker
		}{w, hPusher, hHijacker}
	case okCloseNotifier && okHijacker:
		w = struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
		}{w, hCloseNotifier, hHijacker}
	case okFlusher:
		w = struct {
			http.ResponseWriter
			http.Flusher
		}{w, hFlusher}
	case okPusher:
		w = struct {
			http.ResponseWriter
			http.Pusher
		}{w, hPusher}
	case okCloseNotifier:
		w = struct {
			http.ResponseWriter
			http.CloseNotifier
		}{w, hCloseNotifier}
	case okHijacker:
		w = struct {
			http.ResponseWriter
			http.Hijacker
		}{w, hHijacker}
	}

	return w
}
