package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"coursework/shared/models"
	"coursework/shared/utils"

	"github.com/goccy/go-json"
)

func generateRequests(minCount, maxCount, minTimeMs, maxTimeMs, minPriority, maxPriority int) []models.GeneratorRequest {
	countOfRequests := rand.Intn(maxCount-minCount) + maxCount

	requests := make([]models.GeneratorRequest, countOfRequests)

	for i := 0; i < countOfRequests; i++ {
		requests[i] = models.GeneratorRequest{
			ProcessTimeSec: rand.Intn(maxCount-minCount) + maxCount,
			NeedResponse:   rand.Intn(2) == 1,
			Priority:       rand.Intn(maxPriority-minPriority) + maxPriority,
		}
	}

	return requests
}

func main() {
	serverUrl := utils.GetEnv("SERVER_URL", "localhost:3000")
	sleepTime := time.Duration(utils.GetEnvInt("SLEEP_TIME_MS", 500))

	rand.Seed(time.Now().UnixNano())
	requests := generateRequests(5, 15,
		500, 1200,
		1, 5)

	for _, request := range requests {
		data, err := json.Marshal(request)
		if err != nil {
			panic(err)
		}

		http.Post(serverUrl, "application/json", bytes.NewReader(data))

		time.Sleep(sleepTime * time.Millisecond)
	}

	fmt.Println(requests)
}
