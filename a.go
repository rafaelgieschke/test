package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func process[x interface{ process() (y, error) }, y any](w http.ResponseWriter, r *http.Request) {
	var x1 x
	if err := json.NewDecoder(r.Body).Decode(&x1); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	y1, err := x1.process()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Printf("%#v\n", y1)
	if err := json.NewEncoder(w).Encode(y1); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

type version struct {
	Component string
}

type versionInfo struct {
	Version string
}

func (v version) process() (versionInfo, error) {
	return versionInfo{
		Version: "1234" + v.Component + "-",
	}, nil
}

func main() {
	http.HandleFunc("/version", process[version])
	println("aa")
	http.ListenAndServe(":8080", nil)
}
