package configs

type DNSProvider interface {
	CreateRecords() error
}

type CIProvider interface {
	BuildDockerImage(cmd RunCommand) ([]string, error)
	UploadDockerImage(cmd RunCommand) error
	//RunTests takes in a list of commands to run
	//it returns out an array of strings which are the artifact paths
	RunTests([]RunCommand) ([]string, error)
}

type RunCommand interface {
	Run() (interface{}, error)
}
