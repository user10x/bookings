package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
)

func main() {
	api_key := "api-key"
	endpoint := "https://api.openai.com/v1/engines/davinci-codex/completions"

	data := `{
        "prompt": "Once upon a time",
        "max_tokens": 5,
        "temperature": 0.5
    }`

	/*
		//	endpoint := "https://api.openai.com/v1/edits"
			data := `{
				"model": "text-davinci-edit-001",
					"input": "What day of the wek is it?",
					"instruction": "Fix the spelling mistakes"
			}`
	*/
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read the response body and convert to json

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	var output map[string]interface{}

	err = json.Unmarshal([]byte(body), &output)
	if err != nil {
		fmt.Println("JSON parse error: ", err)
		return
	}

	for key, value := range output {
		color.New(color.FgGreen).Printf("%s: ", key)
		fmt.Printf("%v\n", value)
	}

}
