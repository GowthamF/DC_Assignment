package eurekaservices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dc_assignment.com/m/v2/models"
)

func RegisterInstance(appName string, instance *models.InstanceModel) {
	instanceJson := map[string]*models.InstanceModel{"instance": instance}

	fmt.Println(instanceJson)

	json_data, err := json.Marshal(instanceJson)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8761/eureka/apps/"+appName, bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 204 {
		log.Println("Successfully Registered")

	} else {
		log.Fatalln("Error during registering")
	}
}

func getInstances() {

}

func updateHeartBeat() {

}
