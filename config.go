package main

type config struct {
	serverAddress string // format will be 0.0.0.0:8080 with ip & port
}

func newConfig() config {
	return config{
		serverAddress: "0.0.0.0:8080",
	}
}
