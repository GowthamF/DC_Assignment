package eurekaservices

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"dc_assignment.com/m/v2/models"
	"github.com/carlescere/scheduler"
)

func RegisterInstance(appName string, instance *models.InstanceModel) {
	instanceJson := map[string]*models.InstanceModel{"instance": instance}

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

func UpdateHeartBeat(appName string, instanceId string) {

	job := func() {
		client := &http.Client{}

		req, err := http.NewRequest("PUT", "http://localhost:8761/eureka/apps/"+appName+"/"+instanceId, bytes.NewBuffer([]byte{}))

		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
		resp, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		if resp.StatusCode == 200 {
			log.Println("Heart beat updated")

		} else {
			log.Fatalln("Heart beat failed")
		}
	}

	scheduler.Every(25).Seconds().Run(job)
	runtime.Goexit()
}

func updateRole() {}