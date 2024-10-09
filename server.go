package main

import "net/http"

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)

    w.Write([]byte("<h1> Hello World!</h1>"))
}