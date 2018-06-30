package model

type ICodeConfiguration struct {
	RepositoryPath string
	SshPath        string
	ICodeSavePath  string
}

func NewIcodeConfiguration() ICodeConfiguration {
	return ICodeConfiguration{
		RepositoryPath: "empty",
		SshPath:        "default", // set ssh path or default. default mean HomeDir/.ssh/id_rsa
		ICodeSavePath:  "./.tmp",
	}
}
