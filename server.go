package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type profile struct {
	Ip       string `json:"ipaddress"`
	Language string `json:"language"`
	Software string `json:"software"`
}

func (p profile) writeJson(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getSoftware(r *http.Request) (u string) {
	u = r.UserAgent()
	u = strings.Split(u, "(")[1]
	u = strings.Split(u, ")")[0]
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := &profile{
		Ip:       getIPAdress(r),
		Language: r.Header.Get("accept-language"),
		Software: getSoftware(r),
	}
	p.writeJson(w, r)
}

func main() {
	http.HandleFunc("/api/whoami", handler)
	p := os.Getenv("PORT")
	if p == "" {
		p = "3750"
	}
	http.ListenAndServe(":"+p, nil)
}
