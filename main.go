package main

import (
	"os"

	eurekaservices "dc_assignment.com/m/v2/eureka_services"
	"dc_assignment.com/m/v2/models"
)

func main() {
	hostName := os.Getenv("POD_NAME")
	app := os.Getenv("APP_NAME")
	ipAddress := os.Getenv("POD_IP")
	status := "UP"
	port := os.Getenv("POD_PORT")
	enabledPort := "true"
	healthCheckUrl := "http://" + ipAddress + ":" + port + "/healthcheck"
	statusCheckUrl := "http://" + ipAddress + ":" + port + "/status"
	homePageUrl := "http://" + ipAddress + ":" + port
	class := "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo"
	name := "MyOwn"
	ins := &models.InstanceModel{
		HostName:  &hostName,
		App:       &app,
		IpAddress: &ipAddress,
		Status:    &status,
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
}
