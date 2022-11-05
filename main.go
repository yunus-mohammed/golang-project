package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type details struct {
	name      string `json:"name,omitempty"`
	job       string `json:"job,omitempty"`
	id        string `json:"id,omitempty"`
	createdAt string `json:"createdAt,omitempty"`
}

func testPage(c *gin.Context) {

	var detail1 details

	err := c.ShouldBind(&detail1)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	fmt.Println(detail1)

	res := externalApi(detail1)

	fmt.Println(res)

}

func externalApi(res details) details {

	detailJSON, err := json.Marshal(res)

	reqBody := bytes.NewReader(detailJSON)

	resp, err := http.Post("https://reqres.in/api/users", "application/json", reqBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	var result details
	err = json.Unmarshal([]byte(body), &result)
	fmt.Println(err)

	return (result)
}

func main() {
	router := gin.Default()
	router.POST("/test", testPage)
	router.Run("localhost:8085")
}
