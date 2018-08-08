package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/bftx-api", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("query", `{tx: queryTransaction(id:"BFTX471d0d4b0941002b2603ce27acfa3109df03e427d8ecf53c07d202cd04bd0b56"){Id Properties{DescOfGoods}}}`)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	bftxResult, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bftxResult)
}
