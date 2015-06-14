package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var reqCount uint32

func startServer(addr string, path string) {
	reqCount = 0
	//runtime.GOMAXPROCS(8)
	mux := http.NewServeMux()
	/*mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint32(&reqCount, 1)
		fmt.Fprintf(w, "Hello World!", html.EscapeString(r.URL.Path))
	})*/
	mux.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		atomic.StoreUint32(&reqCount, 0)
		fmt.Fprintf(w, "reset ok\n")
	})
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		atomic.StoreUint32(&reqCount, 0)
		fmt.Fprintf(w, "ua=%s, cmd=%s, cooke=%s", r.UserAgent(), r.FormValue("cmd"), r.Cookies())
		fmt.Fprintf(w, "reset ok\n")
	})
	// Check the static directory
	/*d, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Fatalln("Static dir not exists or not readable", err)
	}

	names, err := d.Readdirnames(-1)
	if err != nil {
		log.Fatalln("Get name in dir failed", err)
	}
	for n := range names {
		p := fmt.Sprintf("/%s/", names[n])
		fmt.Println("Add Handle: ", p)
		mux.Handle(p, http.FileServer(http.Dir(path)))
	}*/
	mux.Handle("/", http.FileServer(http.Dir(path)))
	s := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go showStat()
	s.SetKeepAlivesEnabled(true)
	log.Fatal(s.ListenAndServe())
}

func showStat() {
	for {
		fmt.Printf("\r%d               ", reqCount)
		time.Sleep(time.Second * 1)
	}
}
