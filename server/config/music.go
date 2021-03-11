package config

type Music struct {
	ListId string `json:"list_id" mapstructure:"list-id" yaml:"list-id"`
	Length int    `json:"length" mapstructure:"length" yaml:"length"`
}
