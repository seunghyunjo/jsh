package main
 
import (
    "net/http"
    "html/template"

)
 

 
func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/htm")
  	mainTemplate:=template.Must(template.ParseFiles("topics.html"))
	mainTemplate.Execute(w, nil)
}


 
func main() {
			
    http.HandleFunc("/", mainPage)
    http.ListenAndServe(":8080", nil)

	
}

