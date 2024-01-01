package http

import (
	"io"
	"net/http"
	"net/url"

	kong "github.com/gabriel-valin/kongjson"
)

type ForwardClient struct {
	Client *http.Client
	Config *kong.Config
}

func (c *ForwardClient) ForwardRequest(urlBase string, w http.ResponseWriter, r *http.Request) error {
	forwardURL := urlBase + r.URL.Path

	var err error
	r.URL, err = url.Parse(forwardURL)
	if err != nil {
		return err
	}

	r.RequestURI = ""
	resp, err := c.Client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for k := range resp.Header {
		w.Header().Set(k, resp.Header.Get(k))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)

	return nil
}
