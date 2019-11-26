package main

import "testing"

//this test assumes that there exists a service account valid at tmp/service-account.josn
//and that the Keyrings and keys exist
//and that there is asecrets file at /tmp/test/secrets.txt

func TestLoadGCPSecrets(t *testing.T){
	c := Config{
		GCPKeyRingLocation: "us-west1",
		GCPKMSKey: "deploytag-dev",
		GCPKeyRingName: "deploytag-dev-ring",
		GCPEncryptedSecretsFile: "/tmp/test/secrets.txt",
		GCPPlainTextSecretsFile: "/tmp/test/decrypt.txt",
		Cloud: GcpCloud,
	}
	c.Init()
	if 	err := c.loadGCPSecrets(); err != nil {
		t.Error(err.Error())
	}
}
