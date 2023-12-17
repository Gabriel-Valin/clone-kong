package kong

import (
	"errors"
	"regexp"
)

type Service struct {
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Plugins []Plugin `json:"plugins"`
	Routes  []Route  `json:"routes"`
}

type Plugin struct {
	Name  string         `json:"name"`
	Input map[string]any `json:"input,omitempty"`
}

type Route struct {
	Name       string   `json:"name"`
	Paths      []string `json:"paths"`
	PathRegexp *regexp.Regexp
	Methods    []string `json:"methods"`
}

var ErrPluginNotFound = errors.New("plugin not found")
