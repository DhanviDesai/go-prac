package config

type WalkConfig struct {
	Method string `json:"method"`
}

var DefaultWalkConfig WalkConfig = WalkConfig{
	Method: "GET",
}
