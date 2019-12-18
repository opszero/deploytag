package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type AWS struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string

	CloudAwsSecretId string
	AppAwsSecretIds  []string
}

func (c *AWS) SetCloudAwsSecrets() {
	if c.CloudAwsSecretId == "" {
		return
	}

	log.Println("Loading Cloud Secrets")

	svc := secretsmanager.New(session.New())
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(c.CloudAwsSecretId),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if result.SecretString == nil {
		return
	}

	SetEnvCloudSecrets(*result.SecretString)
}

func (c *AWS) AppAwsSecrets() (fileContent string) {
	for _, secretId := range c.AppAwsSecretIds {
		svc := secretsmanager.New(session.New())
		input := &secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secretId),
		}

		result, err := svc.GetSecretValue(input)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		// Decrypts secret using the associated KMS CMK.
		// Depending on whether the secret is a string or binary, one of these fields will be populated.
		if result.SecretString == nil {
			continue
		}

		fileContent += fmt.Sprintf("# %s", secretId)
		fileContent += "\n"
		fileContent += *result.SecretString
		fileContent += "\n\n"
	}

	return
}

func (c *AWS) Init() {
	if c.AccessKeyID == "" || c.SecretAccessKey == "" || c.Region == "" {
		log.Println("Ensure that AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, and AWS_REGION are set")
	}
}
