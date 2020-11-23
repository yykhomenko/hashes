package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET %s", html.EscapeString(r.URL.Path))
		ifaces, _ := net.Interfaces()
		// handle err
		for _, i := range ifaces {
			addrs, _ := i.Addrs()
			// handle err
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				// process IP address
				fmt.Fprintf(w, "Hello, %q, IP: %s\n", html.EscapeString(r.URL.Path), ip)
			}
		}
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Println("listen http :80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
