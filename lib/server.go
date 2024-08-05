package lib

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Server struct {
	Address string
	Client *http.Client
}

func NewServer(address string) *Server {
	return &Server{Address: address, Client: http.DefaultClient}
}

func (s *Server) StartServer() error {
	s.bindRoutes()

	err := http.ListenAndServe(s.Address, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) bindRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./site")))

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		s.getProjects(w, r)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
}

func (s *Server) getProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := 
	`{
		user(login: "cameronMcConnell") {
			pinnedItems(first: 6, types: REPOSITORY) {
				nodes {
					... on RepositoryInfo {
						name
						description
						url
					}
				}
			}
		}
	}`

	requestBody := map[string]string{
		"query": query,
	}

	jsonBytes, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, "Error parsing json", http.StatusInternalServerError)
		return
	}

	body := bytes.NewReader(jsonBytes)

	forwardReq, err := http.NewRequest(http.MethodPost, "https://api.github.com/graphql", body)
	if err != nil {
		http.Error(w, "Failed to create github request", http.StatusInternalServerError)
		return
	}

	forwardReq.Header.Set("Authorization", "Bearer " + ReadEnv("GITHUB_API_ACCESS_KEY"))
	forwardReq.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(forwardReq)
	if err != nil {
		http.Error(w, "Failed to request data from github.com", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseBody)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}