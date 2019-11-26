package main

import "errors"

func (c *Config) loadGCPSecrets() error {
	args := []string{"gcloud", "kms", "decrypt", "--ciphertext-file", c.GCPEncryptedSecretsFile, "--plaintext-file", c.GCPPlainTextSecretsFile}
	if c.GCPEncryptedSecretsFile == "" || c.GCPPlainTextSecretsFile == "" {
		return errors.New("the encrypted secrets file and the plain text file to write out with are required")
	}
	if c.GCPKMSKey != "" {
		args = append(args, []string{"--key", c.GCPKMSKey}...)
	}
	if c.GCPKeyRingLocation != "" {
		args = append(args, []string{"--keyring", c.GCPKeyRingLocation}...)
	}

	if c.GCPKeyRingLocation != "" {
		args = append(args, []string{"--location", c.GCPKeyRingLocation}...)
	}
	return runCmd(args...)
}
