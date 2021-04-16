package config

type Config struct {
    OutputOption OutputOption `json:"outputOption,omitempty" yaml:"outputOption,omitempty"`
	PluginParams map[string]interface{} `json:"pluginRuntimeParams,omitempty" yaml:"pluginRuntimeParams,omitempty"`
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
