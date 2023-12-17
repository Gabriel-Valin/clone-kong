package internal

import (
	"testing"
	"time"

	kong "github.com/gabriel-valin/kongjson"
)

func TestEmptyJson(t *testing.T) {
	c := kong.Config{}
	json := []byte(`{}`)

	err := c.Refresh(json, time.Now())

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if len(c.Services) != 0 {
		t.Errorf("expected services to be empty got %v", c.Services)
	}
}

func TestSimpleJson(t *testing.T) {
	c := kong.Config{}
	json := []byte(
		`
{
  "services": [
    {
      "name": "payments",
      "url": "http://localhost:3001",
      "plugins": [
        {
          "name": "jwt_auth",
          "input": {
            "secret": "cloudsecret",
            "key_in_header": true,
            "key_in_query": false,
            "key_name": "Authorization"
          }
        },
        {
          "name": "request_size_limiting",
          "input": {
            "allowed_payload_size": 100
          }
        },
        {
          "name": "http_log"
        }
      ],
      "routes": [
        {
          "name": "create-payment",
          "paths": [
            "/payments"
          ],
          "methods": [
            "POST"
          ]
        },
        {
          "name": "get-payment",
          "paths": [
            "/payments/{id}"
          ],
          "methods": [
            "GET"
          ]
        }
      ]
    }
  ]
}
    `)

	err := c.Refresh(json, time.Now())

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if len(c.Services) != 1 {
		t.Errorf("expected services to have 1 element got %v", len(c.Services))
	}

	if c.Services[0].Name != "payments" || c.Services[0].URL != "http://localhost:3001" ||
		len(c.Services[0].Plugins) != 3 || len(c.Services[0].Routes) != 2 {
		t.Errorf("expected service name to be payments got %v", c.Services[0].Name)
	}
}
