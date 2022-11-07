package main

/*
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var num int = 0

type details struct {
	Name      string `json:"name,omitempty"`
	Job       string `json:"job,omitempty"`
	Id        string `json:"id,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

func detailsPage(c *gin.Context) {

	var detail1 details
	var detail2 details

	err := c.ShouldBind(&detail1)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	detail2 = externalApi(detail1)

	val := redisClientGet(detail2)

	c.JSON(http.StatusOK, val)

}

func externalApi(detail1 details) details {

	detailJSON, err := json.Marshal(detail1)

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

	redisClientSet(result)

	return result

}

func redisClientSet(result details) {

	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	Resultjson, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(ctx, result.Name, Resultjson, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

}

func redisClientGet(result details) string {

	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	val, err := client.Get(ctx, result.Name).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val

}

func main() {
	router := gin.Default()
	router.POST("/test", detailsPage)
	router.Run("localhost:8087")
}

*/
