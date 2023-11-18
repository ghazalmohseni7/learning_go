package main

import (
	"fmt"
	"net/http"
	camera "startProject/Camera"
)

func main() {
	mux := http.NewServeMux()
	fmt.Println("start server")
	camera.CameraRouters(mux)
	err := http.ListenAndServe("localhost:9090", mux)
	if err != nil {
		return
	}
}
