package main

import "fmt"

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
	c.runKmsDecryptCmd(c.CloudKmsSecretFile)
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
