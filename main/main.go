package main
/*
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

type details struct {
	Name      string `json:"name,omitempty"`
	Job       string `json:"job,omitempty"`
	Id        string `json:"id,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

var redisClient *redis.Client

func detailsPage(c *gin.Context) {

	var detail1 details
	var detail2 details
	var detailResp details

	err := c.ShouldBind(&detail1)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}






	val := redisClientGet(detail2)

	if val != "" || len(val) == 0 {
		val2 := redisClientGet(detail2)
	} else {
		detailResp = externalApi(detail1)
	}

	c.JSON(http.StatusOK, detailResp)
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

func redisClientGet(result details) string {

	var ctx = context.Background()

	client := getRedisClient()

	val, err := client.Get(ctx, result.Name).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val

}

func main() {
	router := gin.Default()
	clientInit()
	router.POST("/test", detailsPage)
	router.Run("localhost:8087")
}
*/