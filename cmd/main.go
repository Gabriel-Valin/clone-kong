package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	kong "github.com/gabriel-valin/kongjson"
	"github.com/gabriel-valin/kongjson/internal"

	internalhttp "github.com/gabriel-valin/kongjson/http"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(l)

	c := &kong.Config{}
	if err := internal.Loader(c, "config.json", 5*time.Second); err != nil {
		panic(err)
	}

	server := internalhttp.NewServer(c)

	fmt.Println("Server listening on port 8080...")

	http.Handle("/", server)
	http.ListenAndServe(":8080", nil)
}
