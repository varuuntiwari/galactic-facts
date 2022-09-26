package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/varuuntiwari/galactic-facts/data"
)

type Facts struct {
	Data []string `json:"data"`
}

func GetFacts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	rand.Seed(time.Now().UnixNano())
	if !q.Has("count") {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		n := rand.Intn(len(data.Facts) - 1)

		json.NewEncoder(w).Encode(Facts{ Data: []string{ data.Facts[n] } })
		return
	} else if q.Get("count") != "" {
		var res []string
		n, err := strconv.ParseInt(q.Get("count"), 10, 64)
		if err != nil {
			log.Panic(err)
		}
		if n > int64(len(data.Facts)) {
			n = int64(len(data.Facts))
		}
		var i int64
		for i = 0; i < n; i++ {
			r := rand.Intn(len(data.Facts) - 1)
			res = append(res, data.Facts[r])
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(Facts{ Data: res })
		return
	}
}