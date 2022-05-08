package services_test

import (
	"sync"
	"testing"

	"github.com/MarceloMPJ/gorreios-etiqueta/app"
	"github.com/MarceloMPJ/gorreios-etiqueta/lib"
	"github.com/MarceloMPJ/gorreios-etiqueta/services"
	"github.com/fiorix/wsdl2go/soap"
)

func TestTrackingCodeGenerator_RequestNewTrackingCode(t *testing.T) {
	app.Init("../config.json")

	correios_api := lib.NewCorreiosSoapService(&soap.Client{
		URL:                    lib.BaseURL,
		Namespace:              lib.Namespace,
		ExcludeActionNamespace: true,
	})

	generator := services.NewTrackingCodeGenerator(correios_api)
	generator.Start()

	var wg sync.WaitGroup

	for i := 0; i < 20000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			generator.Next()
		}()
	}

	wg.Wait()
}
