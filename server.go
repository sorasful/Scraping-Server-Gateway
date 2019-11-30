package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// queue to store urls to scrape
var queue = list.New()

func main() {
	const PORT int = 8080

	log.Print(fmt.Sprintf("Starting server on port : %d ", PORT))

	// more specific to more global
	http.HandleFunc("/add", addUrlToScrape)
	http.HandleFunc("/to-scrape", toScrape)
	http.HandleFunc("/treat", treat)
	http.HandleFunc("/", index)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", PORT), nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	var text string = `	Available routes : 
		- /add?url=https://google.fr (GET)
		- /to-scrape (GET)
		- /treat (POST) 
	`
	fmt.Fprintf(w, text)

}

func addUrlToScrape(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed, only GET with query param 'url' is allowed", http.StatusMethodNotAllowed)
		return
	}

	keys, ok := r.URL.Query()["url"]

	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Url Param 'url' is missing", http.StatusBadRequest)
		return
	}
	url := keys[0]

	log.Printf("URL : %s", url)

	queue.PushBack(url)
}

func toScrape(w http.ResponseWriter, r *http.Request) {
	/*
		Route to get the next url that needs to be scraped.
	*/

	var remote_ip string = strings.Split(r.RemoteAddr, ":")[0]
	log.Print(fmt.Sprintf("Request received from %s", remote_ip))

	if r.Method == "GET" {

		if queue.Len() > 0 {
			e := queue.Front() // First element
			fmt.Fprintf(w, "%s", e.Value)

			queue.Remove(e) // Dequeue
		} else {
			fmt.Fprint(w, "")
		}

	} else {
		http.Error(w, "Only GET is handled for this route.", http.StatusMethodNotAllowed)
	}

}

func treat(w http.ResponseWriter, r *http.Request) {
	/*
		Route to receive the html and url of a page and then do what must be done.
	*/

	if r.Method == "POST" {
		var page_url string = r.Header.Get("scraped_page")
		if page_url == "" {
			http.Error(w, "Missing scraped_page header in the header", http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		// PERFORM STUFF HERE :
		log.Println("%s", body)

	} else {
		http.Error(w, "Only POST is handled for this route.", http.StatusMethodNotAllowed)
	}

}
