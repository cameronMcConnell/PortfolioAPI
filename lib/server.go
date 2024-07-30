package lib

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	Address string
}

func NewServer(address string) *Server {
	return &Server{Address: address}
}

func (s *Server) StartServer() error {
	bindRoutes()

	err := http.ListenAndServeTLS(s.Address, "https/cert.pem", "https/key.pem", nil)
	if err != nil {
		return err
	}

	return nil
}

func bindRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./Portfolio/site")))

	http.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
		handleGithub(w, r)
	})
}

func handleGithub(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	apiKey := ReadEnv("GITHUB_API_ACCESS_KEY")

	apiKeyJson := GithubAPIKeyJSON{ApiKey: apiKey}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	err := json.NewEncoder(w).Encode(apiKeyJson)
	if err != nil {
		http.Error(w, "Error writing response body", http.StatusInternalServerError)
	}
}