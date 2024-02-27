package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	port          string
	defaultSearch string
	aliases       map[string]string
}

func main() {
	log.Fatal(run())
}

func run() error {
	conf := loadConfig()

	http.HandleFunc("/", handler(conf))

	return http.ListenAndServe(fmt.Sprintf(":%s", conf.port), nil)
}

func loadConfig() Config {
	viper.SetDefault("port", "9090")
	viper.SetDefault("defaultSearch", "https://www.google.com/search?q=%s")
	viper.SetDefault("aliases", map[string]string{})

	return Config{
		port:          "9090",
		defaultSearch: "https://www.google.com/search?q=%s",
		aliases:       map[string]string{},
	}
}

func handler(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alias := strings.Replace(r.URL.Path, "/", "", 1)

		dest := c.aliases[alias]
		if dest == "" {
			http.Redirect(w, r, strings.Replace(c.defaultSearch, "%s", alias, 1), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, dest, http.StatusTemporaryRedirect)
	}
}
