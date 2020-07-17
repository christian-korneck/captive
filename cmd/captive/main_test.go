package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

//tests with mocked http responses
func TestMock(t *testing.T) {
	t.Run("Test with a GOOD mocked response", func(t *testing.T) {
		c, err := Iscaptive(&ClientMockGood{})
		if c != true || err != nil {
			t.Errorf("captive status offline or errored with: '%s'", err)
		}
	})
	t.Run("Test with a BAD mocked response", func(t *testing.T) {
		c, err := Iscaptive(&ClientMockBad{})
		if c == true || err != nil {
			t.Errorf("captive status offline or errored with: '%s'", err)
		}
	})
}

//int test against external service
func TestInt(t *testing.T) {
	t.Run("int test against external service", func(t *testing.T) {
		c, err := Iscaptive(&http.Client{})
		if c != true || err != nil {
			t.Errorf("captive status offline or errored with: '%s' - is this machine disconnected from the Internet?", err)
		}
	})
}

//ClientMockGood provides a GOOD canned http client
type ClientMockGood struct {
}

//Do Good provides a GOOD canned http response
func (c *ClientMockGood) Do(req *http.Request) (*http.Response, error) {
	r := &http.Response{}

	h := http.Header{}
	h.Add("Accept-Ranges", "bytes")
	h.Add("Age", "128")
	h.Add("Cache-Control", "max-age=300")
	h.Add("Cdnuuid", "90949e61-b56b-4d6b-8af3-830f063397fe-5207351420")
	h.Add("Connection", "keep-alive")
	h.Add("Content-Length", "69")
	h.Add("Content-Type", "test/html")
	h.Add("Date", "Fri, 17 Jul 2020 01:40:53 GMT")
	h.Add("Etag", "\"41ba060eb1c0898e0a4a0cca36a8ca91\"")
	h.Add("Last-Modified", "Fri, 17 Feb 2017 20:36:28 GMT")
	h.Add("Server", "ATS/8.0.8")
	h.Add("Via",
		"http/1.1 defra3-edge-lx-011.ts.apple.com (ApacheTrafficServer/8.0.8), http/1.1 defra3-edge-bx-011.ts.apple.com (ApacheTrafficServer/8.0.8)")
	h.Add("X-Amz-Id-2", "kWlTdJRjDgiQoW1vwjg877haXnHp5W2uvpaJUy1Q3Hv3UrvlkYXsEiwJjuwAk0ecX1cO+VW+2kc=")
	h.Add("X-Amz-Request-Id", "9BA807BDB0D70C55")
	h.Add("X-Cache", "hit-fresh, hit-fresh")

	r.Status = "200 OK"
	r.StatusCode = 200
	r.Proto = "HTTP/1.1"
	r.ProtoMajor = 1
	r.ProtoMinor = 1
	r.Header = h
	r.Body = ioutil.NopCloser(bytes.NewReader([]byte(
		"<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML>")))
	r.ContentLength = 69
	t := make([]string, 0)
	r.TransferEncoding = t
	r.Close = false
	r.Uncompressed = false
	r.Trailer = http.Header{}
	r.TLS = nil
	r.Request = req
	return r, nil
}

//ClientMockBad provides a BAD canned http client
type ClientMockBad struct {
}

//Do provides a BAD canned http response
func (c *ClientMockBad) Do(req *http.Request) (*http.Response, error) {
	r := &http.Response{}

	h := http.Header{}
	h.Add("Accept-Ranges", "bytes")
	h.Add("Age", "128")
	h.Add("Cache-Control", "max-age=300")
	h.Add("Cdnuuid", "90949e61-b56b-4d6b-8af3-830f063397fe-5207351420")
	h.Add("Connection", "keep-alive")
	h.Add("Content-Length", "69")
	h.Add("Content-Type", "test/html")
	h.Add("Date", "Fri, 17 Jul 2020 01:40:53 GMT")
	h.Add("Etag", "\"41ba060eb1c0898e0a4a0cca36a8ca91\"")
	h.Add("Last-Modified", "Fri, 17 Feb 2017 20:36:28 GMT")
	h.Add("Server", "ATS/8.0.8")
	h.Add("Via",
		"http/1.1 defra3-edge-lx-011.ts.apple.com (ApacheTrafficServer/8.0.8), http/1.1 defra3-edge-bx-011.ts.apple.com (ApacheTrafficServer/8.0.8)")
	h.Add("X-Amz-Id-2", "kWlTdJRjDgiQoW1vwjg877haXnHp5W2uvpaJUy1Q3Hv3UrvlkYXsEiwJjuwAk0ecX1cO+VW+2kc=")
	h.Add("X-Amz-Request-Id", "9BA807BDB0D70C55")
	h.Add("X-Cache", "hit-fresh, hit-fresh")

	r.Status = "200 OK"
	r.StatusCode = 200
	r.Proto = "HTTP/1.1"
	r.ProtoMajor = 1
	r.ProtoMinor = 1
	r.Header = h
	r.Body = ioutil.NopCloser(bytes.NewReader([]byte(
		"<HTML><HEAD><TITLE>Los Pollos Hermanos Hotspot</TITLE></HEAD><BODY>Two hours free surfing after login!</BODY></HTML>")))
	r.ContentLength = 69
	t := make([]string, 0)
	r.TransferEncoding = t
	r.Close = false
	r.Uncompressed = false
	r.Trailer = http.Header{}
	r.TLS = nil
	r.Request = req
	return r, nil
}
