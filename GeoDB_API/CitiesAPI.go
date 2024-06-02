package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// url := "https://wft-geo-db.p.rapidapi.com/v1/geo/cities"

	// req, _ := http.NewRequest("GET", url, nil)

	// req.Header.Add("X-RapidAPI-Key", "e0e9295368msh0ba370d88e27ed5p13224djsn9d9f621e92e9")
	// req.Header.Add("X-RapidAPI-Host", "wft-geo-db.p.rapidapi.com")

	// res, _ := http.DefaultClient.Do(req)

	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	url := "https://andruxnet-world-cities-v1.p.rapidapi.com/?query=paris&searchby=city"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "e0e9295368msh0ba370d88e27ed5p13224djsn9d9f621e92e9")
	req.Header.Add("X-RapidAPI-Host", "andruxnet-world-cities-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
