package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mychewcents/hotel_data_server/internal/controller"
	"github.com/mychewcents/hotel_data_server/internal/models"
)

func main() {
	http.HandleFunc("/hotels", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		req := &models.GetHotelsRequest{}
		reqAsBytes := make([]byte, 0)
		_, _ = r.Body.Read(reqAsBytes)
		fmt.Printf("%+v", reqAsBytes)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Printf("\n%+v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hotels, err := controller.GetHotels(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		hotelsAsBytes, _ := json.Marshal(hotels)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		_, _ = w.Write(hotelsAsBytes)
	})

	fmt.Println("\nstarting server on 8080...")

	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		panic(err)
	}
}
