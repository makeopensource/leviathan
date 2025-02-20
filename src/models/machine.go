package models

type MachineOptions struct {
	Name           string `mapstructure:"name"`
	Host           string `mapstructure:"host"`
	Port           int64  `mapstructure:"port"`
	User           string `mapstructure:"user"`
	PrivateKeyFile string `mapstructure:"pvt_key_file"`
}
