package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var backendURL string

func init() {
	// Set the backend URL from an environment variable
	backendURL = os.Getenv("BACKEND_URL")
	if backendURL == "" {
		log.Fatal("BACKEND_URL environment variable not set")
	}
}

func main() {
	http.HandleFunc("/proxy/", handleProxy)
	port := "8080"
	log.Printf("Proxy service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	// Construct the backend URL with the requested path
	url := backendURL + r.URL.Path[len("/proxy/"):]

	// Get an identity token for the backend
	token, err := getIdentityToken()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get identity token", http.StatusInternalServerError)
		return
	}

	// Forward the request to the backend
	proxyReq, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Copy original headers and add authorization
	copyHeaders(proxyReq.Header, r.Header)
	proxyReq.Header.Set("X-Serverless-Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy response back to the client
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
