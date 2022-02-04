package http_server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type test struct {
	handler      http.HandlerFunc
	method       string
	url          string
	bodyRequest  string
	bodyResponse string
	code         int
}

func TestGetHandler(t *testing.T) {
	store = map[string]Lake{
		"id00": {
			Id:   "id00",
			Name: "Malawi",
			Area: 29500,
		},
		"id01": {
			Id:   "id01",
			Name: "Great Bear Lake",
			Area: 31000,
		},
	}
	cases := []test{
		{
			handler:      GetHandler,
			method:       "GET",
			url:          "/get",
			bodyRequest:  `{"type":"get","payload":"id03"}`,
			bodyResponse: "404 Not Found\n",
			code:         404,
		},
		{
			handler:      GetHandler,
			method:       "GET",
			url:          "/get",
			bodyRequest:  `{"type":"get","payload":"id00"}`,
			bodyResponse: "Malawi\n29500",
			code:         200,
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%s/%s", tt.method, tt.url), func(t *testing.T) {
			// create the request
			req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.bodyRequest))
			if err != nil {
				t.Fatal(err)
			}
			// create a recorder to evaluate bodyResponse
			rec := httptest.NewRecorder()
			tt.handler.ServeHTTP(rec, req)

			// check status code
			if rec.Code != tt.code {
				t.Errorf("got code %v, want %v", rec.Code, tt.code)
			}
			// check bodyRequest bodyResponse
			if rec.Body.String() != tt.bodyResponse {
				t.Errorf("got body %v, want %v", rec.Body.String(), tt.bodyResponse)
			}
		})
	}
}

func TestPostHandler(t *testing.T) {
	store = map[string]Lake{
		"id00": {
			Id:   "id00",
			Name: "Malawi",
			Area: 29500,
		},
		"id01": {
			Id:   "id01",
			Name: "Great Bear Lake",
			Area: 31000,
		},
	}
	cases := []test{
		{
			handler:      PostHandler,
			method:       "POST",
			url:          "/post",
			bodyRequest:  `{"type":"post","payload":"{\"id\":\"id09\",\"name\":\"Mississippi\",\"area\":82000}"}`,
			bodyResponse: "",
			code:         200,
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%s/%s", tt.method, tt.url), func(t *testing.T) {
			// create the request
			req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.bodyRequest))
			if err != nil {
				t.Fatal(err)
			}
			// create a recorder to evaluate bodyResponse
			rec := httptest.NewRecorder()
			tt.handler.ServeHTTP(rec, req)

			// check status code
			if rec.Code != tt.code {
				t.Errorf("got code %v, want %v", rec.Code, tt.code)
			}
			// check bodyRequest bodyResponse
			if rec.Body.String() != tt.bodyResponse {
				t.Errorf("got body %v, want %v", rec.Body.String(), tt.bodyResponse)
			}
			// check the lake is successfully added to storage
			if _, ok := store["id09"]; !ok {
				t.Error("lake not found in store")
			}
		})
	}
}

func TestDeleteHandler(t *testing.T) {
	store = map[string]Lake{
		"id00": {
			Id:   "id00",
			Name: "Malawi",
			Area: 29500,
		},
		"id01": {
			Id:   "id01",
			Name: "Great Bear Lake",
			Area: 31000,
		},
	}
	cases := []test{
		{
			handler:      DeleteHandler,
			method:       "DELETE",
			url:          "/delete",
			bodyRequest:  `{"type":"delete","payload":"id00"}`,
			bodyResponse: "",
			code:         200,
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%s/%s", tt.method, tt.url), func(t *testing.T) {
			// create the request
			req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.bodyRequest))
			if err != nil {
				t.Fatal(err)
			}
			// create a recorder to evaluate bodyResponse
			rec := httptest.NewRecorder()
			tt.handler.ServeHTTP(rec, req)

			// check status code
			if rec.Code != tt.code {
				t.Errorf("got code %v, want %v", rec.Code, tt.code)
			}
			// check bodyRequest bodyResponse
			if rec.Body.String() != tt.bodyResponse {
				t.Errorf("got body %v, want %v", rec.Body.String(), tt.bodyResponse)
			}
			// check the lake is successfully added to storage
			if _, ok := store["id00"]; ok {
				t.Error("lake still in store")
			}
		})
	}
}
