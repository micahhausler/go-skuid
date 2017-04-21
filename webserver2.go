package main

import "log"
import "net/http"
import "encoding/json"

// START OMIT
type Page struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func main() {
	http.HandleFunc("/page", hello)
	log.Print("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request: %v", r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	p := Page{Id: 1, Content: `<skuidpage unsavedchangeswarning="yes" personalizationmode="client" hideeditbutton="true" theme="Skuid Public">`}
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, "Error encoding Page", http.StatusInternalServerError)
		return
	}
}

// END OMIT
