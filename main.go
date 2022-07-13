package main


import (
  "net/http";
  "fmt";
  "log";
)

func postHandler(w http.ResponseWriter, r *http.Request) {
  if err:=r.ParseForm(); err != nil {
    fmt.Fprintf(w, "Parsed Form error: %v", err)

    return
  }

  fmt.Fprintf(w, "POST successful")
  first := r.FormValue("first")
  second := r.FormValue("second")
  
  fmt.Fprintf(w, "First Name = %v", first)
  fmt.Fprintf(w, "Second Name = %v", second)
}


func getHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/get" {
    http.Error(w, "404 not Found", http.StatusNotFound)
    return
  }

  if r.Method != "GET" {
    http.Error(w, "This method is not suppoted \nOnly GET method supported", http.StatusBadRequest)
    return
  }

  fmt.Fprintf(w, "GET Successful")
}

func main(){
  fileServer := http.FileServer(http.Dir("./static-files"))
	http.Handle("/", fileServer)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/get", getHandler)

  fmt.Printf("Starting server on port 3000\n")
  if err := http.ListenAndServe(":3000", nil); err != nil {
    log.Fatal(err)
  }
}
