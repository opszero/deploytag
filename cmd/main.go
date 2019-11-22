package cmd

import (
	"fmt"
	"github.com/opszero/deploytag/configs"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	config = &configs.Config{}
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "deploytag",
		Short: "CI /CD Helper for Kubernetes and Serverless Apps",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
	}

	rootCmd.PersistentFlags().StringVar(&config.Cloud, "cloud", "", "aws, gcp, or azure")
	rootCmd.PersistentFlags().StringVar(&config.AWSAccessKeyID, "aws-access-key-id", os.Getenv("AWS_ACCESS_KEY_ID"), "AWS Access Key")
	rootCmd.PersistentFlags().StringVar(&config.AWSSecretAccessKey, "aws-secret-access-key", os.Getenv("AWS_SECRET_ACCESS_KEY"), "AWS Secret Access Key")
	rootCmd.PersistentFlags().StringVar(&config.AWSDefaultRegion, "aws-default-region", os.Getenv("AWS_DEFAULT_REGION"), "AWS Secret Access Key")
	rootCmd.PersistentFlags().StringVar(&config.GCPServiceKeyFile, "gcp-service-key-file", "", "GCP Auth File. ~/gcp.json")
	rootCmd.PersistentFlags().StringVar(&config.GCPServiceKeyBase64, "gcp-service-key-base64", "", "Base64 encoded version of gcp-service-key-base64")

	rootCmd.PersistentFlags().StringVar(&config.CloudAwsSecretId, "cloud-aws-secret-id", "", "Use AWS Secrets Manager for Config. If set it pull the environment variables from aws secrets manager.")
	rootCmd.PersistentFlags().StringArrayVar(&config.AppAwsSecretIds, "app-aws-secret-ids", []string{}, "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com")

	rootCmd.Flags().StringVar(&config.CloudFlareKey, "cloudflare-key", os.Getenv(configs.CloudFlareAPIKey), "api key for cloudflare")
	rootCmd.Flags().StringVar(&config.CloudFlareEmail, "cloudflare-email", os.Getenv(configs.CloudFlareEmail), "email for cloudflare")
	rootCmd.Flags().StringVar(&config.CloudFlareZoneName, "cloudflare-domain", "", "email for cloudflare")

	var runScriptCmd = &cobra.Command{
		Use:   "run-script",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			config.Init()
			config.HelmRunScript()
		},
	}

	runScriptCmd.Flags().StringVar(&config.RunScript.PodAppLabel, "pod-app-label", "", "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com ")
	runScriptCmd.Flags().StringVar(&config.RunScript.Container, "container", "", "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com ")
	runScriptCmd.Flags().StringArrayVar(&config.RunScript.Cmds, "cmds", []string{}, "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com ")

	rootCmd.AddCommand(runScriptCmd)

	// deployCmd represents the deploy command
	var deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			config.Init()
			config.HelmDeploy()
		},
	}

	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVar(&config.Deploy.Env, "env", "", "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com ")
	deployCmd.Flags().StringVar(&config.Deploy.ChartName, "chart-name", "", "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com ")
	deployCmd.Flags().StringVar(&config.Deploy.HelmConfig, "helm-config", "", "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com ")
	deployCmd.Flags().StringArrayVar(&config.Deploy.HelmSet, "helm-set", []string{}, "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com ")

	deployCmd.Flags().StringVar(&config.Build.ContainerRegistry, "container-registry", "", "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com ")
	deployCmd.Flags().StringVar(&config.Build.ProjectId, "project-id", "", "Ex. opszero")
	deployCmd.Flags().StringVar(&config.Build.Image, "image", "", "Ex. deploytag")
	deployCmd.Flags().StringArrayVar(&config.ExternalHostNames, "set-dns", []string{}, "list of external hostnames to resolve against the cluster's load balancer")

	var buildCmd = &cobra.Command{
		Use:   "build",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			config.Init()
			config.DockerBuild()
		},
	}

	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVar(&config.Build.DotEnvFile, "dotenv-file", "", "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com")
	buildCmd.Flags().StringVar(&config.Build.ContainerRegistry, "container-registry", "", "Ex. 1234.dkr.ecr.us-west-2.amazonaws.com ")
	buildCmd.Flags().StringVar(&config.Build.ProjectId, "project-id", "", "Ex. opszero")
	buildCmd.Flags().StringVar(&config.Build.Image, "image", "", "Ex. deploytag")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
