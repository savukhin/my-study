package main

import (
	"math/rand"
	"time"

	"coursework/shared/models"
)

func generateRequests(minCount, maxCount, minTimeMs, maxTimeMs, minPriority, maxPriority int) []models.GeneratorRequest {
	countOfRequests := rand.Intn(maxCount-minCount) + maxCount

	requests := make([]models.GeneratorRequest, countOfRequests)

	for i := 0; i < countOfRequests; i++ {
		requests[i] = models.GeneratorRequest{
			ProcessTime:  time.Duration(rand.Intn(maxCount-minCount) + maxCount),
			NeedResponse: rand.Intn(1) == 1,
			Priority:     rand.Intn(maxPriority-minPriority) + maxPriority,
		}
	}

	return requests
}

func main() {
	rand.Seed(time.Now().UnixNano())
	requests := generateRequests(5, 15,
		500, 1200,
		1, 5)
}
