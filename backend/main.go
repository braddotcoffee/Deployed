package main

import (
	"deployed/datastore"
	"deployed/utils"
	"fmt"
	"net/http"
)

func main() {
	/*
		r := mux.NewRouter()

		r.HandleFunc("/hello-world", helloWorld)

		c := cors.New(cors.Options{
			AllowedOrigins: []string{"https://app.brad.coffee"},
		})
		handler := c.Handler(r)

		srv := &http.Server{
			Handler: handler,
			Addr:    ":" + os.Getenv("PORT"),
		}

		log.Fatal(srv.ListenAndServe())
	*/
	datastore.Connect()
	deployment := datastore.Deployment{
		Name: "Test Deployment 2",
	}
	datastore.AddDeployment(&deployment)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}
