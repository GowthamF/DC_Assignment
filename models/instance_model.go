package models

type InstanceModel struct {
	InstanceId     *string              `json:"instanceId"`
	HostName       *string              `json:"hostName"`
	App            *string              `json:"app"`
	IpAddress      *string              `json:"ipAddr"`
	Status         *string              `json:"status"`
	HealthCheckUrl *string              `json:"healthCheckUrl"`
	StatusPageUrl  *string              `json:"statusPageUrl"`
	HomePageUrl    *string              `json:"homePageUrl"`
	Port           *PortModel           `json:"port"`
	DataCenterInfo *DataCenterInfoModel `json:"dataCenterInfo"`
}

type PortModel struct {
	PortNumber *string `json:"$"`
	Enabled    *string `json:"@enabled"`
}

type DataCenterInfoModel struct {
	Class *string `json:"@class"`
	Name  *string `json:"name"`
}
