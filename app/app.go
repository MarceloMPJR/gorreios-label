package app

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	IdServico     int64  `json:"id_servico,omitempty"`
	Identificador string `json:"identificador,omitempty"`
	Usuario       string `json:"usuario,omitempty"`
	Senha         string `json:"senha,omitempty"`
}

var Config Configuration

func Init(config_path string) error {
	setDefaultConfig()

	if config_path == "" {
		return nil
	}

	file, err := os.Open(config_path)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}

func setDefaultConfig() {
	Config.IdServico = 162000
	Config.Identificador = "34028316000103"
	Config.Usuario = "sigep"
	Config.Senha = "n5f9t8"
}
