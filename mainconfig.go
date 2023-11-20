package main

import (
	"fmt"
	"net/http"
	"startProject/Camera"
	"startProject/utils"
)

func main() {
	errorConnection := utils.RedisConnection()
	if errorConnection != nil {
		fmt.Println(errorConnection.Error())
	}
	err1 := utils.AddToRedis()
	if err1 != nil {
		return
	}
	mux := http.NewServeMux()
	fmt.Println("start server")
	Camera.CameraRouters(mux)
	err := http.ListenAndServe("localhost:9090", mux)
	if err != nil {
		return
	}

}
