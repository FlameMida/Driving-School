package config

type Weather struct {
	AppKey string `json:"app-key" mapstructure:"app-key" yaml:"app-key"`
	Lang   string `json:"lang" mapstructure:"lang" yaml:"lang"`
	Lat    string `json:"lat" mapstructure:"lat" yaml:"lat"`
	Log    string `json:"log" mapstructure:"log" yaml:"log"`
	Units  string `json:"units" mapstructure:"units" yaml:"units"`
}

type Location struct {
	Lat string `json:"lat" mapstructure:"lat" yaml:"lat"`
	Log string `json:"log" mapstructure:"log" yaml:"log"`
}
