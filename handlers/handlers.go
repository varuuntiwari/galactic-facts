package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/varuuntiwari/galactic-facts/data"
)

type Facts struct {
	Data []string `json:"data"`
	Count int `json:"count"`
	Status uint `json:"status"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("API to retrieve facts about space")
	return
}

func GetFacts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	rand.Seed(time.Now().UnixNano())
	if !q.Has("count") {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		n := rand.Intn(len(data.Facts) - 1)

		json.NewEncoder(w).
		Encode(
			Facts{
				Data: []string{ data.Facts[n] },
				Count: 1,
				Status: http.StatusOK,
			})
		return
	} else if q.Get("count") != "" {
		var res []string
		n, err := strconv.ParseInt(q.Get("count"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).
			Encode(
			Facts{
				Data: nil,
				Count: 0,
				Status: http.StatusUnprocessableEntity,
			})
			return
		}
		randNums := rand.Perm(len(data.Facts))
		var i int64
		for i = 0; i < n; i++ {
			res = append(res, data.Facts[randNums[i]])
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).
		Encode(
			Facts{
				Data: res,
				Count: len(res),
				Status: http.StatusOK,
			})
		return
	}
}