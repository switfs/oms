package config

type ClientConfig struct {
	Client Client `json:"client"`
	Proxy  Proxy  `json:"proxy"`
}

type Client struct {
	Hosts string `json:"hosts"`
	Ports string `json:"ports"`
}
