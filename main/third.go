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

	"github.com/yunus-mohammed/golang-project/redisclient"
)

type details struct {
	Name      string `json:"name,omitempty"`
	Job       string `json:"job,omitempty"`
	Id        string `json:"id,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

func respHandler(c *gin.Context) {

	var req details
	var res details
	var strDetails details

	err := c.ShouldBind(&req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	str := RedisClientGetVal(string(req.Name))

	err1 := json.Unmarshal([]byte(str), &strDetails)
	if err1 != nil {
		log.Fatalf("An Error Occured %v", err1)
	}

	if str != "" || len(str) != 0 {

		c.JSON(http.StatusOK, strDetails)
		return
	}

	res = externalApi(req)
	c.JSON(http.StatusOK, res)
	return
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

	RedisClientSetVal(string(result.Name), string(body))

	return result

}

/*
func redisClientSet(result details) {

	var ctx = context.Background()

	client := getRedisClient()

	Resultjson, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(ctx, result.Name, Resultjson, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

}
*/
func RedisClientSetVal(key string, str string) {

	var strDetails details

	var ctx = context.Background()

	client := redisclient.GetRedisClient()

	err := json.Unmarshal([]byte(str), &strDetails)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(ctx, key, strDetails, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

}

/*
func redisClientGet(result details) string {

	var ctx = context.Background()

	client := getRedisClient()

	val, err := client.Get(ctx, result.Name).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val

}*/

func RedisClientGetVal(key string) string {

	var ctx = context.Background()

	client := redisclient.GetRedisClient()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val

}

func main() {
	router := gin.Default()
	redisclient.ClientInit()
	router.POST("/test", respHandler)
	router.Run("localhost:8087")
}
