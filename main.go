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

type test struct {
	name      string `json:"name,omitempty"`
	job       string `json:"job,omitempty"`
	id        string `json:"id,omitempty"`
	createdAt string `json:"createdAt,omitempty"`
}

func testPage(c *gin.Context) {

	name := c.Param("name")
	job := c.Param("job")

	test1 := test{name: name, job: job}

	testJSON, err := json.Marshal(test1)

	rb := bytes.NewBuffer(testJSON)

	resp, err := http.Post("https://reqres.in/api/users", "application/json", rb)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result test
	err = json.Unmarshal([]byte(body), &result)
	fmt.Println(err)

	sb := string(body)
	c.JSON(http.StatusOK, sb)

}

func main() {
	router := gin.Default()
	router.POST("/test", testPage)
	router.Run("localhost:8085")
}
