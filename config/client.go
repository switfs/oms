package config

type ClientConfig struct {
	Client Client `json:"client"`
}

type Client struct {
	Enabled bool   `json:"enabled"`
	Hosts   string `json:"hosts"`
	Ports   string `json:"ports"`
}
