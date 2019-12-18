package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func SetEnvCloudSecrets(secretsStr string) {
	secrets, err := godotenv.Parse(strings.NewReader(secretsStr))
	if err != nil {
		log.Println(err)
		return
	}

	if secrets == nil {
		for k := range secrets {
			log.Println("Setting up var", k)
			os.Setenv(k, secrets[k])
		}
	}
}

func (c *Config) writeAppSecrets(fileName string) {
	content := fmt.Sprintf("DEPLOYTAG_BRANCH=%s\n\n", c.Git.DockerBranch())

	content += c.AppDotEnv
	log.Println("Writing .env", content)

	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		log.Println(err)
	}
}

func (c *Config) AppendAppSecrets(content string) {
	c.AppDotEnv += "\n\n"
	c.AppDotEnv += content
	c.AppDotEnv += "\n\n"
}

func (c *Config) SecretsInit() {
	c.AWS.SetCloudAwsSecrets()
	c.GCP.SetCloudKmsSecrets()

	c.AppendAppSecrets(c.AWS.AppAwsSecrets())
	c.AppendAppSecrets(c.GCP.AppKmsSecrets())
}
