package model

type RemoteAdd struct {
	RemoteIp string `json:"remote_ip"`
	RemotePort string `json:"remote_port"`
	Url string `json:"url"`
}