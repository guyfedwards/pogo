package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
)

func main() {
	log.Fatal(run())
}

var port = flag.String("Port", "9090", "port to run on")

func run() error {
	flag.Parse()

	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("run: %w", err)
	}
	configFile := fmt.Sprintf("%s/pogo/config.yaml", configDir)

	conf := NewConfig(configFile, *port)
	err = conf.Load()

	cancel := conf.Monitor()
	defer func() {
		cancel <- syscall.SIGINT
	}()

	if err != nil {
		return fmt.Errorf("run: %w", err)
	}

	http.HandleFunc("/", handler(conf))

	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)
}

func handler(c *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alias := strings.Replace(r.URL.Path, "/", "", 1)

		dest := c.Aliases[alias]
		if dest == "" {
			http.Redirect(w, r, strings.Replace(c.DefaultSearch, "%s", alias, 1), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, dest, http.StatusTemporaryRedirect)
	}
}
