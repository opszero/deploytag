package main

import "os"

type Azure struct{}

func (c *Azure) Init() {
	runCmd("az", "login", "--service-principal", "--tenant", os.Getenv("AZURE_SERVICE_PRINCIPAL_TENANT"), "--username", os.Getenv("AZURE_SERVICE_PRINCIPAL"), "--password", os.Getenv("AZURE_SERVICE_PRINCIPAL_PASSWORD"))
}
