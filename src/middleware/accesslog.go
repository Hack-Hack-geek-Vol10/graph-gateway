package middleware

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type accessLog struct {
	RequestClientIP  string `json:"http_request_client_ip"`
	RequestDuration  int64  `json:"http_request_duration"`
	RequestHost      string `json:"http_request_host"`
	RequestMethod    string `json:"http_request_method"`
	RequestPath      string `json:"http_request_path"`
	RequestProtocol  string `json:"http_request_protocol"`
	RequestSize      int64  `json:"http_request_size"`
	RequestTime      string `json:"http_request_time"`
	RequestUserAgent string `json:"http_request_user_agent"`
	ResponseSize     int    `json:"http_response_size"`
	ResponseStatus   int    `json:"http_response_status"`
}

// AccessLog helps log request and response related data.
func AccessLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		wr := newWriter(w, r, t)

		h.ServeHTTP(wr, r)

		al := &accessLog{
			RequestClientIP:  wr.reqClientIP,
			RequestDuration:  time.Since(t).Milliseconds(),
			RequestHost:      wr.reqHost,
			RequestMethod:    wr.reqMethod,
			RequestPath:      wr.reqPath,
			RequestProtocol:  wr.reqProto,
			RequestSize:      wr.reqSize,
			RequestTime:      wr.reqTime,
			RequestUserAgent: wr.reqUserAgent,
			ResponseSize:     wr.resSize,
			ResponseStatus:   wr.resStatus,
		}

		data, err := json.Marshal(al)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(data))
	})
}

// Acts as an adapter for http.ResponseWriter type to store request and
// response data.
type writer struct {
	http.ResponseWriter

	reqClientIP  string
	reqHost      string
	reqMethod    string
	reqPath      string
	reqProto     string
	reqSize      int64 // bytes
	reqTime      string
	reqUserAgent string

	resStatus int
	resSize   int // bytes
}

func newWriter(w http.ResponseWriter, r *http.Request, t time.Time) *writer {
	return &writer{
		ResponseWriter: w,

		reqClientIP:  r.Header.Get("X-Forwarded-For"),
		reqMethod:    r.Method,
		reqHost:      r.Host,
		reqPath:      r.RequestURI,
		reqProto:     r.Proto,
		reqSize:      r.ContentLength,
		reqTime:      t.Format(time.RFC3339),
		reqUserAgent: r.UserAgent(),
	}
}

// Overrides http.ResponseWriter type.
func (w *writer) WriteHeader(status int) {
	if w.resStatus == 0 {
		w.resStatus = status
		w.ResponseWriter.WriteHeader(status)
	}
}

// Overrides http.ResponseWriter type.
func (w *writer) Write(body []byte) (int, error) {
	if w.resStatus == 0 {
		w.WriteHeader(http.StatusOK)
	}

	var err error
	w.resSize, err = w.ResponseWriter.Write(body)

	return w.resSize, err
}

// Overrides http.Flusher type.
func (w *writer) Flush() {
	if fl, ok := w.ResponseWriter.(http.Flusher); ok {
		if w.resStatus == 0 {
			w.WriteHeader(http.StatusOK)
		}

		fl.Flush()
	}
}

// Overrides http.Hijacker type.
func (w *writer) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the hijacker interface is not supported")
	}

	return hj.Hijack()
}
