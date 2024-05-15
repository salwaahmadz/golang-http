package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"

	"go.mod/app"
)

func main() {

	// JSON GET RESPONSE
	http.HandleFunc("/api/v1/prodcuts.json", func(res http.ResponseWriter, req *http.Request) {
		products := []app.Product{}

		product1 := app.Product{
			Price: 500000000,
		}
		product1.SetIDAndName(1, "Lakungan")
		product1.Category = app.Category{}
		product1.Category.SetIDAndName(2, "Ecommerce")

		product2 := app.Product{
			Price: 600000000,
		}
		product2.SetIDAndName(2, "My Score Credit")
		product2.Category = app.Category{}
		product2.Category.SetIDAndName(1, "Project Management")

		products = append(products, product1, product2)
		output, _ := json.MarshalIndent(products, "", " ")

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(output)
	})

	// XML GET RESPONSE
	http.HandleFunc("/api/v1/prodcuts.xml", func(res http.ResponseWriter, req *http.Request) {
		products := []app.Product{}

		product1 := app.Product{
			Price: 500000000,
		}
		product1.SetIDAndName(1, "Lakungan")
		product1.Category = app.Category{}
		product1.Category.SetIDAndName(2, "Ecommerce")

		product2 := app.Product{
			Price: 600000000,
		}
		product2.SetIDAndName(2, "My Score Credit")
		product2.Category = app.Category{}
		product2.Category.SetIDAndName(1, "Project Management")

		products = append(products, product1, product2)
		output, _ := xml.MarshalIndent(products, "", " ")

		res.Header().Set("Content-Type", "text/xml")
		res.WriteHeader(200)
		res.Write(output)
	})

	// JSON POST RESPONSE
	http.HandleFunc("/api/v1/add-prodcuts.json", func(res http.ResponseWriter, req *http.Request) {
		products := app.Product{}
		payload, _ := io.ReadAll(req.Body)
		errUnmarshal := json.Unmarshal(payload, &products)
		if nil != errUnmarshal {
			panic(errUnmarshal)
		}

		output, _ := json.MarshalIndent(products, "", " ")

		log.Println(string(payload))

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(output)
	})

	// XML POST RESPONSE
	http.HandleFunc("/api/v1/add-prodcuts.xml", func(res http.ResponseWriter, req *http.Request) {
		products := app.Product{}
		payload, _ := io.ReadAll(req.Body)
		errUnmarshal := xml.Unmarshal(payload, &products)
		if nil != errUnmarshal {
			panic(errUnmarshal)
		}

		output, _ := xml.MarshalIndent(products, "", " ")

		log.Println(string(payload))

		res.Header().Set("Content-Type", "text/xml")
		res.WriteHeader(200)
		res.Write(output)
	})

	log.Println("Go server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
