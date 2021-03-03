package config

type Config struct {
	PluginName string `json:"pluginName,omitempty" yaml:"pluginName,omitempty"`
	DeploymentName string `json:"deploymentName,omitempty" yaml:"deploymentName,omitempty"`
	RunTime string `json:"runTime,omitempty" yaml:"runTime,omitempty"`
	Handler string `json:"handler,omitempty" yaml:"handler,omitempty"`
	FromFile string `json:"fromFile,omitempty" yaml:"fromFile,omitempty"`
	Dependencies string `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`

    OutputOption OutputOption `json:"outputOption,omitempty" yaml:"outputOption,omitempty"`
    PayLoad string `json:"payLoad,omitempty" yaml:"payLoad,omitempty"`
	PluginParams map[string]string `json:"pluginParams,omitempty" yaml:"pluginParams,omitempty"`
}

type OutputOption struct {
	KafkaConfig  KafkaConfig  `json:"kafkaConfig,omitempty" yaml:"kafkaConfig,omitempty"`
	MinioConfig  MinioConfig  `json:"minioConfig,omitempty" yaml:"minioConfig,omitempty"`
}

type KafkaConfig struct {
	Topic   string     `json:"topic,omitempty" yaml:"topic,omitempty"`
	ServiceEndpoint   string     `json:"serviceEndpoint,omitempty" yaml:"serviceEndpoint,omitempty"`
}

type MinioConfig struct {
	Bucket string    `json:"bucket,omitempty" yaml:"bucket,omitempty"`
	Path   string    `json:"path,omitempty" yaml:"path,omitempty"`
	Region   string  `json:"region,omitempty" yaml:"region,omitempty"`
	AccessKey   string  `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`
	AccessKeySecret   string  `json:"accessKeySecret,omitempty" yaml:"accessKeySecret,omitempty"`
}
