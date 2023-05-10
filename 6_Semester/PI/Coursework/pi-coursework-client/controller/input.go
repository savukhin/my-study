package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	API_ADDRESS      = GetEnvDefault("API_ADDRESS", "http://localhost:8080")
	QUERY_ADDRESS, _ = url.JoinPath(API_ADDRESS, "api", "v1", "execute-query")
	HELP             = `
CREATE 
	`
)

const (
	RESPONSE_VALUE = "response_value"
)

type Response struct {
	Message string
	Status  http.ConnState
}

type HttpResponse struct {
	Message string `json:"message"`
}

func NewResponse(msg string, stat http.ConnState) Response {
	return Response{
		Message: msg,
		Status:  stat,
	}
}

func processQuery(query string, responseChan chan Response) {
	// time.Sleep(1 * time.Second)

	command := strings.TrimSpace(query)
	command = strings.ToLower(command)
	if command == "help" {
		responseChan <- NewResponse(HELP, 200)
		return
	}

	data, _ := json.Marshal(map[string]string{
		"query": query,
	})
	r := bytes.NewReader(data)

	resp, err := http.Post(QUERY_ADDRESS, "application/json", r)
	if err != nil {
		responseChan <- NewResponse("No server error", -1)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		responseChan <- NewResponse("Client error", -2)
		return
	}

	httpResp := &HttpResponse{}
	// fmt.Println("")
	json.Unmarshal(bodyBytes, httpResp)

	responseChan <- NewResponse(httpResp.Message, http.ConnState(resp.StatusCode))
}

func showLoading(frontText string, responseChan chan Response) {
	states := []string{"/", "-", "\\", "|", "/", "-", "\\", "|"}
	state_ind := 0
	frontText = strings.TrimSpace(frontText)

	for {
		select {
		case response := <-responseChan:
			fmt.Println("^\rProcessed    ")
			fmt.Println("Result:")
			fmt.Println("-------------------------------------------------")
			fmt.Println(response.Message)
			fmt.Println("-------------------------------------------------")
			return
		default:
			fmt.Printf("\r%s %s", frontText, states[state_ind])
			time.Sleep(100 * time.Millisecond)
			state_ind = (state_ind + 1) % len(states)
		}
	}
}

func suggestNewQuery() {
	fmt.Println("  |  Input here several strings then press enter two times")
	fmt.Println("  Y")
}

func EndlessInput() {
	query := ""
	suggestNewQuery()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("| ")
		text, _ := reader.ReadString('\n')

		// fmt.Println([]byte(text))

		if strings.TrimSpace(text) != "" {
			query += text
			continue
		}

		responseChan := make(chan Response)

		go processQuery(query, responseChan)

		showLoading("Processing", responseChan)
		// reader.Reset()

		suggestNewQuery()

		query = ""
	}
}
