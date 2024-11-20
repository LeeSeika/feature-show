package config

import (
	"fmt"
)

type System struct {
	Host string `json:"host" yaml:"host" `
	Port string `json:"port" yaml:"port"`
	Env  string `json:"env" yaml:"env"`
}

func (s *System) GetAddr() string {
	return fmt.Sprintf(s.Host + ":" + s.Port)
}
