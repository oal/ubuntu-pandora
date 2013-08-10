package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		artistAlbum := fmt.Sprintf("by %v (on %v)", req.PostForm["artist"][0], req.PostForm["album"][0])

		cmd := exec.Command("notify-send", "-i", "gnome-audio", req.PostForm["song"][0], artistAlbum)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	})

	http.ListenAndServe("0.0.0.0:10002", nil)
}
