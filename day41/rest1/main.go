package main

import (
	"io"
	"github.com/emicklei/go-restful"
	"fmt"
	"time"
	"net/http"
)


func pingTime(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}


func main() {
	webservice := new(restful.WebService)

	webservice.Route(webservice.GET("/ping").To(pingTime))

	restful.Add(webservice)

	http.ListenAndServe(":8000", nil)
}