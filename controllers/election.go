package controllers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"dc_assignment.com/m/v2/eurekaservices"
	"github.com/gin-gonic/gin"
)

func StartElection(c *gin.Context) {
	instanceId := c.GetString("instanceId")
	GetHigherInstanceIds(instanceId, "PRIMENUMBERAPP")
}

func StopElection(c *gin.Context) {}

func RequestElection(c *gin.Context) {
	requestInstanceId, _ := strconv.ParseInt(c.Param("requestInstanceId"), 0, 64)
	instanceId, _ := strconv.ParseInt(c.GetString("instanceId"), 0, 64)
	// fmt.Println(instanceId)
	// fmt.Println(requestInstanceId)
	if instanceId < requestInstanceId {
		fmt.Println(requestInstanceId)
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusNotAcceptable)
		GetHigherInstanceIds(c.GetString("instanceId"), "PRIMENUMBERAPP")
	}

}

func GetHigherInstanceIds(myId string, appName string) {
	instances := eurekaservices.GetInstances(appName)
	for _, instance := range *instances.Application.Instance {
		if instance != nil {
			insId, _ := strconv.ParseInt(*instance.InstanceId, 0, 64)
			myAppId, _ := strconv.ParseInt(myId, 0, 64)

			if myAppId < insId {
				SendElectionRequest(instance.HomePageUrl, &myId)
			} else {
				log.Println(myAppId, " is the leader")
			}
		}
	}
}

func SendElectionRequest(url *string, myAppId *string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", *url+"/requestElection/"+*myAppId, bytes.NewBuffer([]byte{}))

	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 200 {
		log.Println("Continue the election")
		GetHigherInstanceIds(*myAppId, "PRIMENUMBERAPP")

	} else if resp.StatusCode == 406 {
		log.Println("Do not continue")
	}
}
