package main

import (
	"fmt"
	"log"
	"os"
)

const (
	GCPDecryptionKeyVariable = "GCP_KMS_KEY"
)

type GCP struct {
	ServiceKeyFile   string
	ServiceKeyBase64 string

	CloudKmsSecretFile string
	AppKmsSecretFiles  []string

	// KMS Secrets
	KMSKey          string
	KeyRingLocation string
	KeyRingName     string
}

func (c *GCP) SetCloudKmsSecrets() {
	SetEnvCloudSecrets(c.runKmsDecryptCmd(c.CloudKmsSecretFile))
}

func (c *GCP) AppKmsSecrets() (fileContent string) {
	for _, secretFile := range c.AppKmsSecretFiles {
		fileContent += fmt.Sprintf("# %s", secretFile)
		fileContent += "\n"
		fileContent += c.runKmsDecryptCmd(secretFile)
		fileContent += "\n\n"
	}

	return
}

func (c *GCP) runKmsDecryptCmd(fileName string) (output string) {
	return runCmdOutput("gcloud", "kms", "decrypt", "--ciphertext-file", fileName, "--plaintext-file", "-", "--key", c.KMSKey, "--keyring", c.KeyRingName, "--location", c.KeyRingLocation)
}

func (c *GCP) Init() {
	if os.Getenv("GCLOUD_SERVICE_KEY_BASE64") != "" {
		if err := runCmd("bash", "-c", "echo $GCLOUD_SERVICE_KEY_BASE64 | base64 -d > /tmp/gcloud-service-key.json"); err != nil {
			log.Fatalln("the service key given was not base64 encoded")
			return
		}
	} else {
		log.Println("No Google Service Account Key given")
	}
	if err := runCmd("gcloud", "auth", "activate-service-account", "--key-file=/tmp/gcloud-service-key.json"); err != nil {
		log.Fatalln("failed to authenticate gcp from service account")
		return
	}
}
