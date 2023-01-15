package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Start program")
	workDone := make(chan bool)
	url := "https://dummyjson.com/products/1"
	go checkChannel(workDone)
	time.Sleep(1 * time.Second)
	go makeNetworkCall(url, workDone)
	status := <-workDone
	if status {
		fmt.Println("Successfully did things")
	} else {
		fmt.Println("There was some network error, but gracefully exited")
	}
}

func checkChannel(workDone <-chan bool) {
	for {
		select {
		default:
			showProgress()
		case <-workDone:
			fmt.Println("Done so moving on")
			fmt.Println("Finishing my work")
			return
		}
	}
}

func showProgress() {
	fmt.Print("\033[?25l")
	chars := [4]byte{'|', '/', '-', '\\'}
	for i := 0; i < 4; i++ {
		fmt.Print("\b" + string(chars[i]))
		time.Sleep(100 * time.Millisecond)
	}

}

func makeNetworkCall(url string, workDone chan bool) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in making network call")
		workDone <- false
		return
	}
	fmt.Println(resp.StatusCode)
	readBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(readBody))
	workDone <- true
}
