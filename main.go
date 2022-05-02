package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	eurekaservices "dc_assignment.com/m/v2/eureka_services"
	"dc_assignment.com/m/v2/models"
	"dc_assignment.com/m/v2/routes"
)

func main() {
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	randomNumber := strconv.FormatInt(rand.Int63(), 10)
	id := currentTime + randomNumber
	hostName := os.Getenv("POD_NAME")
	app := os.Getenv("APP_NAME")
	ipAddress := os.Getenv("POD_IP")
	port := os.Getenv("POD_PORT")
	// hostName := "PRIMENUMBER"
	// app := "PRIMENUMBERAPP"
	// ipAddress := "localhost"
	// port := "8080"
	status := "UP"
	enabledPort := "true"
	healthCheckUrl := "http://" + ipAddress + ":" + port + "/healthcheck"
	statusCheckUrl := "http://" + ipAddress + ":" + port + "/status"
	homePageUrl := "http://" + ipAddress + ":" + port
	class := "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo"
	name := "MyOwn"
	ins := &models.InstanceModel{
		InstanceId: &id,
		HostName:   &hostName,
		App:        &app,
		IpAddress:  &ipAddress,
		Status:     &status,
		Port: &models.PortModel{
			PortNumber: &port,
			Enabled:    &enabledPort,
		},
		HealthCheckUrl: &healthCheckUrl,
		StatusPageUrl:  &statusCheckUrl,
		HomePageUrl:    &homePageUrl,
		DataCenterInfo: &models.DataCenterInfoModel{
			Class: &class,
			Name:  &name,
		},
	}
	eurekaservices.RegisterInstance(app, ins)
	r := routes.SetupRouter()
	r.Run()

}
