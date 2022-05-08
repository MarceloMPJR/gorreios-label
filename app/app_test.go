package app_test

import (
	"reflect"
	"testing"

	"github.com/MarceloMPJ/gorreios-etiqueta/app"
)

func TestApp_Init(t *testing.T) {
	t.Run("when config_path is empty should set default config", func(t *testing.T) {
		err := app.Init("")

		expectedNoError(t, err)
		setsDefaultConfig(t)
	})

	t.Run("when config_path is valid should set the config with attributes of file", func(t *testing.T) {
		err := app.Init("../fixtures/config.json")

		expected := app.Configuration{
			IdServico:     1,
			Identificador: "identificador",
			Usuario:       "usuario",
			Senha:         "senha",
		}

		expectedNoError(t, err)

		if !reflect.DeepEqual(app.Config, expected) {
			t.Errorf("expected %v, got %v", expected, app.Config)
		}
	})

	t.Run("when config_path is invalid should set default config and return error", func(t *testing.T) {
		err := app.Init("invalid_path")

		expectedError(t, err)
		setsDefaultConfig(t)
	})
}

func setsDefaultConfig(t *testing.T) {
	t.Helper()

	expected := app.Configuration{
		IdServico:     162000,
		Identificador: "34028316000103",
		Usuario:       "sigep",
		Senha:         "n5f9t8",
	}

	if !reflect.DeepEqual(app.Config, expected) {
		t.Errorf("expected %v, got %v", expected, app.Config)
	}
}

func expectedNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func expectedError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected error, got no error")
	}
}
