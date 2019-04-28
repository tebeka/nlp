package main

import (
	"encoding/json"
	"os"
	// _ "expvar"  // import just for side effects (get the /debug/vars http handler)
	"expvar"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"

	"github.com/tebeka/nlp"
)

var (
	tokCounter = expvar.NewInt("num_tokenize")
	// Version is executable version
	Version = "0.2.0"

	options struct {
		Port int
	}
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version and exit")
	flag.Parse()

	if showVersion {
		fmt.Printf("nlpd version %s\n", Version)
		os.Exit(0)
	}

	if err := parseOptions(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/healthz", okHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)
	// routing
	// /x/y - will match only /x/y
	// /x/y/ - will match /x/y/, /x/y/z ....

	addr := fmt.Sprintf(":%d", options.Port)
	log.Printf("nlpd ready on %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func parseOptions() error {
	// Defaults
	options.Port = 8080

	// Get environment
	return envconfig.Process("NLPD", &options)
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	tokCounter.Add(1)
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("can't read request - %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokens := nlp.Tokenize(string(data))
	out, err := json.Marshal(tokens)
	if err != nil {
		log.Printf("can't marshal output - %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Sanity checks
	fmt.Fprintf(w, "OK\n")
}
