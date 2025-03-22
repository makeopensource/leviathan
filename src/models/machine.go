package models

type MachineOptions struct {
	Enable    bool   `mapstructure:"enable"`
	Name      string `mapstructure:"name"`
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Publickey string `mapstructure:"host_publickey"`
}
