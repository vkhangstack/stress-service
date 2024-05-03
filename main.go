package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/random"
)

func stress(method, url string, payload io.Reader) string {
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("x-request-id", random.New().String(32, random.Alphabetic))
	req.Header.Set("user-agent", "null")
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	body, readErr := io.ReadAll(res.Body)

	if readErr != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("method", method)
	fmt.Println("url", url)
	fmt.Println("status", res.Status)
	return string(body)
}

func main() {
	envFile, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println(err.Error())
	}

	body := []byte(string(envFile["BODY"]))
	total, _ := strconv.Atoi(envFile["TIMES_REQUEST"])
	n := 1
	for n < total+1 {
		data := stress(envFile["METHOD"], envFile["URL"], bytes.NewBuffer(body))
		fmt.Println("times", n)
		fmt.Println("data", data)
		n += 1
	}

}
