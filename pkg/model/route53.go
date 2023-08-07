package model

type HostedZone struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

type ListHostedZonesResponse struct {
	HostedZones []HostedZone `json:"hostedzones" yaml:"hostedzones"`
}
