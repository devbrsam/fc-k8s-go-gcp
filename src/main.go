package main
import (
  "html/template"
  "net/http"
  "log"
  "fmt"
)

type Content struct {
  Text string
}

func main() {
  log.Print("starting server...")
	http.HandleFunc("/", renderTemplate)
  log.Printf("listening on port %s", "8000")
  http.ListenAndServe(":8000", nil)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8000"), nil))
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template.html")
  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
	
  if err := tmpl.Execute(w, template.HTML(greeting("Code.education Rocks!"))); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}