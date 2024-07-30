package lib

import (
	"testing"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
)

// Mock transport for GitHub API
type MockTransport struct {
	Response *http.Response
}

func (mt *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return mt.Response, nil
}

func TestNewServer(t *testing.T) {
	got := NewServer(":8080")

	if got == nil {
		t.Error("Got nil from struct initialization")
	}
}

// Test getProjects handler
func TestGetProjects(t *testing.T) {
	// Set up a mock response for the GitHub API
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"data": {"user": {"pinnedRepositories": {"nodes": [{"name": "repo1", "url": "http://example.com", "description": "desc"}]}}}}`))),
	}
	mockTransport := &MockTransport{Response: mockResponse}
	client := &http.Client{Transport: mockTransport}

	server := &Server{Client: client}

	req := httptest.NewRequest(http.MethodGet, "/projects", nil)
	rec := httptest.NewRecorder()

	server.getProjects(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	// Verify the response body
	expectedBody := `{"data": {"user": {"pinnedRepositories": {"nodes": [{"name": "repo1", "url": "http://example.com", "description": "desc"}]}}}}`
	body, _ := io.ReadAll(res.Body)
	if string(body) != expectedBody {
		t.Errorf("expected body %v; got %v", expectedBody, string(body))
	}
}