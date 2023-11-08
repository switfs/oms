package config

type ClientConfig struct {
	Client Client `json:"client"`
}

type Client struct {
	Hosts string `json:"hosts"`
	Ports string `json:"ports"`
}
