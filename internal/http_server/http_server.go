package http_server

import (
	"L0/internal/storage"
	"fmt"
	"github.com/go-pg/pg/v10"
	"io"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	http.ServeFile(w, r, "templates/index.html")
}

func getJson(strg *storage.Storage, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uid := string(body)
	var data string
	if data = strg.StorageCache.Storage[uid]; data != "" {
		fmt.Fprintf(w, data)
	} else {
		_, err = strg.StorageDB.QueryOne(pg.Scan(&data), "SELECT data_string FROM json_data WHERE id = ?", uid)
		if err != nil {
			http.Error(w, "Invalid UID", http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, data)
			strg.PutInCache(uid, data)
		}
	}
}

func Start(address string, timeout time.Duration, idleTimeout time.Duration, strg *storage.Storage) *http.Server {
	server := &http.Server{
		Addr:         address,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  idleTimeout,
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)
	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) { getJson(strg, w, r) })

	go func() {
		log.Println("Starting HTTP server on :8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()
	return server
}
