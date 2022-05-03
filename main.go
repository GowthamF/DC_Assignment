package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"dc_assignment.com/m/v2/eurekaservices"
	"dc_assignment.com/m/v2/models"
	"dc_assignment.com/m/v2/routes"
)

var (
	portNumber = flag.String("portNumber", "8080", "Port Number to serve")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	currentTime := fmt.Sprint(time.Now().UnixMilli())
	randomNumber := fmt.Sprint(rand.Int63())

	id := currentTime + randomNumber
	// hostName := os.Getenv("POD_NAME")
	// app := os.Getenv("APP_NAME")
	// ipAddress := os.Getenv("POD_IP")
	// port := os.Getenv("POD_PORT")
	hostName := "PRIMENUMBER"
	app := "PRIMENUMBERAPP"
	ipAddress := "localhost"
	port := *portNumber
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
	go eurekaservices.UpdateHeartBeat(app, id)

	r := routes.SetupRouter()
	r.Run(":" + *portNumber)

}
