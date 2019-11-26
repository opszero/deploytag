package main

import "errors"

type GCPKmsSecret struct {
	GCPEncryptedSecretsFile string
	GCPPlainTextSecretsFile string
	GCPKMSKey               string
	GCPKeyRingLocation      string
	GCPKeyRingName			string

}

func (c *Config) loadGCPSecrets() error {
	args := []string{
		"gcloud",
		"kms",
		"decrypt",
		"--ciphertext-file",
		c.GCPKmsSecret.GCPEncryptedSecretsFile,
		"--plaintext-file",
		c.GCPKmsSecret.GCPPlainTextSecretsFile,
	}

	if c.GCPKmsSecret.GCPEncryptedSecretsFile == "" || c.GCPKmsSecret.GCPPlainTextSecretsFile == "" {
		return errors.New("the encrypted secrets file and the plain text file to write out with are required")
	}

	if c.GCPKmsSecret.GCPKMSKey != "" {
		args = append(args, []string{
			"--key",
			c.GCPKmsSecret.GCPKMSKey}...
		)
	}

	if c.GCPKmsSecret.GCPKeyRingName != "" {
		args = append(args, []string{
			"--keyring",
			c.GCPKmsSecret.GCPKeyRingName}...
		)
	}

	if c.GCPKmsSecret.GCPKeyRingLocation != "" {
		args = append(args, []string{
			"--location",
			c.GCPKmsSecret.GCPKeyRingLocation}...
		)
	}

	return runCmd(args...)
}
