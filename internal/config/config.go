// internal/config/config.go
type APIDestConfig struct {
	BaseURL string `yaml:"base_url"`
}

type Config struct {
	SourceDB DBConfig      `yaml:"source_db"`
	APIDest  APIDestConfig `yaml:"api_dest"`
	Batch    BatchConfig   `yaml:"batch"`
}
