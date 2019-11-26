package main

import "errors"

type GCPKmsSecret struct {
	GCPEncryptedSecretsFile string
	GCPPlainTextSecretsFile string
	GCPKMSKey               string
	GCPKeyRingLocation      string
	GCPKeyRingName			string

}

func (c *Config) GCPKmsSecret() error {
	args := []string{
		"gcloud",
		"kms",
		"decrypt",
		"--ciphertext-file",
		c.GCPKms.GCPEncryptedSecretsFile,
		"--plaintext-file",
		c.GCPKms.GCPPlainTextSecretsFile,
	}
	if c.GCPKms.GCPEncryptedSecretsFile == "" || c.GCPKms.GCPPlainTextSecretsFile == "" {
		return errors.New("the encrypted secrets file and the plain text file to write out with are required")
	}

	if c.GCPKms.GCPKMSKey != "" {
		args = append(args, []string{
			"--key",
			c.GCPKms.GCPKMSKey}...
		)
	}

	if c.GCPKms.GCPKeyRingName != "" {
		args = append(args, []string{
			"--keyring",
			c.GCPKms.GCPKeyRingName}...
		)
	}

	if c.GCPKms.GCPKeyRingLocation != "" {
		args = append(args, []string{
			"--location",
			c.GCPKms.GCPKeyRingLocation}...
		)
	}

	return runCmd(args...)
}
