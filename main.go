    package main

    import(
        "net/http"
        "log"
        "fmt"
    )
    func main()  {
        http.HandleFunc("/", Inicio)
		log.Println("Servidor running...")
		http.ListenAndServe(":8080", nil)
    }
    func Inicio (w http.ResponseWriter,r *http.Request ) {
        fmt.Fprintf(w, "Hola, como estas?")
    }