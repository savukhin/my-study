package controller

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

func processQuery(query string, cancelFunc context.CancelFunc) {
	time.Sleep(1 * time.Second)
	cancelFunc()
	// fmt.Println("\rWorked")
}

func showLoading(frontText string, ctx context.Context) {
	states := []string{"/", "-", "\\", "|", "/", "-", "\\", "|"}
	state_ind := 0
	frontText = strings.TrimSpace(frontText)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("^\rProcessed    ")
			return
		default:
			fmt.Printf("\r%s %s", frontText, states[state_ind])
			time.Sleep(100 * time.Millisecond)
			state_ind = (state_ind + 1) % len(states)
		}
	}
}

func EndlessInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')

		ctx, cancelFunc := context.WithCancel(context.Background())

		go processQuery(text, cancelFunc)

		showLoading("Processing", ctx)

	}
}
