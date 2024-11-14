package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

var backendURL string
var allowedCors string

func init() {
	// Set the backend URL from an environment variable
	backendURL = os.Getenv("GCP_PROXY_BACKEND_URL")
	if backendURL == "" {
		log.Fatal("GCP_PROXY_BACKEND_URL environment variable not set")
	}

	allowedCors = os.Getenv("GCP_PROXY_ALLOW_ORIGINS")
	if allowedCors == "" {
		log.Fatal("GCP_PROXY_ALLOW_ORIGINS environment variable not set")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/proxy/", handleProxy)

	port := "8080"
	log.Printf("Proxy service running on port %s", port)

	allowedOrigins := strings.Split(os.Getenv("GCP_PROXY_ALLOW_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	handler := corsMiddleware.Handler(mux)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Construct the backend URL with the requested path
	url := backendURL + "/" + r.URL.Path[len("/proxy/"):] + "?" + r.URL.RawQuery

	// Forward the request to the backend
	proxyReq, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		log.Println("Failed to create request", err)
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	copyHeaders(proxyReq.Header, r.Header)

	if os.Getenv("IS_HTTPS") == "true" {
		// Get an identity token for the backend
		token, err := getIdentityToken()
		if err != nil {
			log.Println("Failed to get identity token", err)
			http.Error(w, "Failed to get identity token", http.StatusInternalServerError)
			return
		}

		// Copy original headers and add authorization
		proxyReq.Header.Set("X-Serverless-Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		log.Println("Failed to forward request", err)
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy response back to the client
	responseHeaders := w.Header()
	copyHeaders(responseHeaders, resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// getIdentityToken fetches an identity token from the metadata server
func getIdentityToken() (string, error) {
	// Request the token from the metadata server
	identityURL := "http://metadata/computeMetadata/v1/instance/service-accounts/default/identity?audience=" + backendURL
	req, err := http.NewRequest("GET", identityURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Metadata-Flavor", "Google")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	token, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

// copyHeaders copies headers from the original request to the proxy request
func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
