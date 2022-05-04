package main

import (
	"flag"
	"math/rand"
	"strconv"
	"time"

	"dc_assignment.com/m/v2/controllers"
	"dc_assignment.com/m/v2/eurekaservices"
	"dc_assignment.com/m/v2/models"
	"dc_assignment.com/m/v2/queue"
	"dc_assignment.com/m/v2/routes"
)

var (
	portNumber      = flag.Int("portNumber", 8080, "Port Number to serve")
	isStartElection = flag.Bool("election", false, "To start the election")
	insId           = flag.Int("id", 0, "ID")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	// currentTime := fmt.Sprint(time.Now().UnixMilli())
	// randomNumber := fmt.Sprint(rand.Int31n(10000))

	id := strconv.Itoa(*insId)

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
	healthCheckUrl := "http://" + ipAddress + ":" + strconv.Itoa(port) + "/healthcheck"
	statusCheckUrl := "http://" + ipAddress + ":" + strconv.Itoa(port) + "/status"
	homePageUrl := "http://" + ipAddress + ":" + strconv.Itoa(port)
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
	go queue.ReceiveMessage("masterElection")
	if *isStartElection {
		controllers.GetHigherInstanceIds(id, app)
	}

	go eurekaservices.UpdateHeartBeat(app, id)

	r := routes.SetupRouter(id)
	r.Run(":" + strconv.Itoa(port))

}
