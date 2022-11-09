package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	redis "github.com/yunus-mohammed/golang-project/redisclient"
	util "github.com/yunus-mohammed/golang-project/utiliity"
)

type Details struct {
	Name      string `json:"name,omitempty"`
	Job       string `json:"job,omitempty"`
	Id        string `json:"id,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

func GetData(req Details, c *gin.Context) (Details, error) {

	var res Details

	str, redisErr := redis.RedisClientGetVal(string(req.Name))

	if redisErr == nil {
		err1 := json.Unmarshal([]byte(str), &res)
		if err1 != nil {
			log.Fatalf("An Error Occured %v", err1)
		}

		if str != "" || len(str) != 0 {

			return res, nil
		}
	}

	detailJSON, _ := json.Marshal(req)

	reqBody := bytes.NewReader(detailJSON)

	bodyStr := util.ExternalPostAPI("https://reqres.in/api/users", reqBody)

	if len(bodyStr) == 0 {
		return res, fmt.Errorf("No bidy found")
	}

	redis.RedisClientSetVal(req.Name, bodyStr)

	err := json.Unmarshal([]byte(bodyStr), &res)

	if err != nil {
		return res, fmt.Errorf("Unable to parse body")
	}

	return res, nil

}
