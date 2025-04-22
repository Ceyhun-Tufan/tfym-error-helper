package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var endpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key="

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type RequestBody struct {
	Contents []Content `json:"contents"`
}

type Candidate struct {
	Content struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"content"`
}

type Response struct {
	Candidates []Candidate `json:"candidates"`
}

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Can't find home directory:", err)
		return
	}

	defaultEnvPath := homeDir + "/.config/tfym-cli/.env"

	err = godotenv.Load(defaultEnvPath)
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			fmt.Println("Couldn't load .env file from ~/.config/tfym-cli/ or current directory.")
			return
		}
	}

	//  YOU CAN CHANGE THE PROMPT HERE

	prompt := "You are an error handling robot that talks like a homie. Be short and quick. Explain me what the fuck does this even mean : \n"
	conf := flag.String("prompt", prompt, "Prompt to use Gemini API")
	flag.Parse()
	prompt = *conf

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("You need to give a prompt to say 'The f*ck you mean?'.")
		return
	}

	for _, arg := range args {
		prompt += arg + " "
	}

	prompt = prompt[:len(prompt)-1]

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API Key not found in environment variables")
		return
	}
	endpoint += apiKey

	reqBody := RequestBody{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: prompt},
				},
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Probably an API key problem.")
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		text := response.Candidates[0].Content.Parts[0].Text
		fmt.Println(text)
	} else {
		fmt.Println("Bro couldn't find an answer.")
	}
}
