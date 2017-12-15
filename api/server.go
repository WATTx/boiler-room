package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

type Server struct {
	port           string
	livelinkClient *LivelinkClient
}

func NewServer(port string) *Server {
	livelinkClient := NewLivelinkClient()
	return &Server{
		port:           port,
		livelinkClient: livelinkClient,
	}
}

func (s *Server) ListenAndServe() error {
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})
	r.Use(c.Handler)

	r.Patch("/{livelinkHost}/lights/{lightID}", s.patchLight)

	return http.ListenAndServe(":"+s.port, r)
}

func (s *Server) patchLight(w http.ResponseWriter, r *http.Request) {
	livelinkHost := chi.URLParam(r, "livelinkHost")
	lightID := chi.URLParam(r, "lightID")
	msg, _ := ioutil.ReadAll(r.Body)

	err := s.livelinkClient.SetLevel(livelinkHost, lightID, msg)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type LivelinkClient struct {
	httpc http.Client
}

func NewLivelinkClient() *LivelinkClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := http.Client{
		Transport: tr,
		Timeout:   40 * time.Second,
	}

	return &LivelinkClient{
		httpc: httpClient,
	}
}

type levelMessage struct {
	Level int `json:"level"`
}

func (l *LivelinkClient) SetLevel(livelinkHost, lightID string, msg []byte) error {
	url := fmt.Sprintf("https://%s.local:8443/rest/devices/lights/%s/levelcontrol", livelinkHost, lightID)

	log.Println(url)
	body := ioutil.NopCloser(bytes.NewBuffer(msg))
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return err
	}

	_, err = l.httpc.Do(request)
	return err
}
