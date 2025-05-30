package docker

import (
	"fmt"
	"reflect"
	"strings"
)

type MachineOptions struct {
	name             string // use get/set so that this field is not written to the config file
	Enable           bool   `mapstructure:"enable"`
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	User             string `mapstructure:"user"`
	Password         string `mapstructure:"password"`
	RemotePublickey  string `json:"remote_public_key"`
	UsePublicKeyAuth bool   `json:"use_public_key_auth"`
}

func (opts *MachineOptions) Log() string {
	var result strings.Builder
	v := reflect.ValueOf(*opts)
	typeOfS := v.Type()

	result.WriteString("options:\n")
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := typeOfS.Field(i).Name
		tag := typeOfS.Field(i).Tag

		// Skip field
		if fieldName == "RemotePublickey" || fieldName == "Password" {
			continue
		}

		// Check for mapstructure or json tags
		mapstructureTag := tag.Get("mapstructure")
		jsonTag := tag.Get("json")

		var outputName string
		if mapstructureTag != "" {
			outputName = mapstructureTag
		} else if jsonTag != "" {
			outputName = jsonTag
		} else {
			outputName = fieldName
		}

		// Format the output based on the field type
		switch field.Kind() {
		case reflect.String:
			result.WriteString(fmt.Sprintf("  %s: %q\n", outputName, field.String()))
		case reflect.Bool:
			result.WriteString(fmt.Sprintf("  %s: %t\n", outputName, field.Bool()))
		case reflect.Int:
			result.WriteString(fmt.Sprintf("  %s: %d\n", outputName, field.Int()))
		default:
			result.WriteString(fmt.Sprintf("  %s: %v\n", outputName, field.Interface())) // Handle other types if needed
		}
	}
	return strings.TrimSpace(result.String())
}

func (opts *MachineOptions) Name() string {
	return opts.name
}

func (opts *MachineOptions) SetName(name string) {
	opts.name = name
}

func DefaultMachine() *MachineOptions {
	return &MachineOptions{
		Enable:           false,
		Host:             "192.168.1.69",
		Port:             22,
		User:             "test",
		Password:         "",
		RemotePublickey:  "",
		UsePublicKeyAuth: false,
	}
}
