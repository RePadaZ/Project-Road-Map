package main

import (
	"UnitConverter/model"
	"UnitConverter/pkg"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		data, _ := io.ReadAll(request.Body)
		var post model.Request
		err := json.Unmarshal(data, &post)

		if err != nil {
			fmt.Print(err.Error())
		}

		switch post.Flag {
		case "length":
			result := pkg.ConvertLength(post)
			writer.WriteHeader(http.StatusAccepted)
			_ = json.NewEncoder(writer).Encode(result)
		case "weight":
			result := pkg.ConvertWeight(post)
			writer.WriteHeader(http.StatusAccepted)
			_ = json.NewEncoder(writer).Encode(result)
		case "temperature":
			result := pkg.ConvertTemp(post)
			writer.WriteHeader(http.StatusAccepted)
			_ = json.NewEncoder(writer).Encode(result)
		}

	}))
	if err != nil {
		fmt.Print(err.Error())
	}

}
