package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})
	http.ListenAndServe(":8000", nil)
}

// I would have to create a semaphore kind of thing (buffered channel as seen in the video). This would be blocked until that is done. I think doing it via a work group would be much better.
// What is a work group in golang? How will that help me in doing whatever it is that I want to do?

// The basic flow of listening to something on a socket (as per my understanding): You register yourself to that socket, by telling kernel that you are interested in that socket. You sleep waiting for some data to be available on that socket. Once a data is available (put on the socket pending queue), kernel notifies you, you wake up and do whatever it is you do.
