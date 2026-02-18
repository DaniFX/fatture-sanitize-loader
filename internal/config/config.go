// internal/config/config.go
type DBConfig struct {
	DSN string `yaml:"dsn"`
}

type APIDestConfig struct {
	BaseURL string `yaml:"base_url"`
}

type BatchConfig struct {
	Size int `yaml:"size"`
}

type SanitizationConfig struct {
	AttiveRuleset  string `yaml:"attive_ruleset"`
	PassiveRuleset string `yaml:"passive_ruleset"`
}

type Config struct {
	SourceDB     DBConfig           `yaml:"source_db"`
	APIDest      APIDestConfig      `yaml:"api_dest"`
	Batch        BatchConfig        `yaml:"batch"`
	Sanitization SanitizationConfig `yaml:"sanitization"`
}
