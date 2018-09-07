package router

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/user/api/server"
	"github.com/user/api/stats"
)

func newreq(t *testing.T, method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return req
}

func TestEndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests in short mode.")
	}

	sStats := stats.New()
	serv := server.New("8080")
	routes := CreateRouter(sStats, serv)
	serv.RegisterRoutes(routes)
	testServer := httptest.NewUnstartedServer(routes)
	testServer.Config = serv.Server
	testServer.Start()
	defer testServer.Close()

	tests := []struct {
		name           string
		request        *http.Request
		expectedStatus int
		expectedBody   string
	}{
		{name: "GET /stats", request: newreq(t, "GET", testServer.URL+"/stats", nil), expectedStatus: 200, expectedBody: `{"total":0,"average":0}`},
		{name: "POST /hash", request: newreq(t, "POST", testServer.URL+"/hash", strings.NewReader("password=angryMonkey")), expectedStatus: 200, expectedBody: `ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==`},
		{name: "POST /hash missingPassword", request: newreq(t, "POST", testServer.URL+"/hash", strings.NewReader("notPassword=angryMonkey")), expectedStatus: 400, expectedBody: "password was not found in form data\n"},
		{name: "POST /hash emptyPassword", request: newreq(t, "POST", testServer.URL+"/hash", strings.NewReader("password=")), expectedStatus: 400, expectedBody: "password was not found in form data\n"},
		{name: "GET /shutdown", request: newreq(t, "GET", testServer.URL+"/shutdown", nil), expectedStatus: 200, expectedBody: "Shutting Down"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(test.request)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			// Check the status code
			if resp.StatusCode != test.expectedStatus {
				t.Errorf("route returned wrong status code: got %v want %v",
					resp.Status, test.expectedStatus)
			}
			// check for expected response here.
			if string(body) != test.expectedBody {
				t.Errorf("route returned body code: got %v want %v",
					string(body), test.expectedBody)
			}
		})
	}

	// Is the server really shutdown?
	finalRequest := newreq(t, "GET", testServer.URL+"/shutdown", nil)
	_, err := http.DefaultClient.Do(finalRequest)
	if err == nil || !strings.Contains(err.Error(), "connect: connection refused") {
		t.Errorf("server didn't refuse connection after second request to shutdown")
	}
}
