package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MarceloMPJR/gorreios-etiqueta/app"
	"github.com/MarceloMPJR/gorreios-etiqueta/lib"
)

var (
	tipoDestinatario = "C"
	qtdEtiquetas     = 10000
	multipliers      = [8]int{8, 6, 4, 2, 3, 5, 9, 7}
)

const ()

type TrackingCodeGenerator struct {
	trackingCodeRange   [2]int
	trackingCodePrefix  string
	trackingCodeSufix   string
	currentTrackingCode chan string
	correiosApi         lib.CorreiosSigepService
}

func NewTrackingCodeGenerator(correiosApi lib.CorreiosSigepService) *TrackingCodeGenerator {
	return &TrackingCodeGenerator{
		correiosApi:         correiosApi,
		trackingCodeRange:   [2]int{},
		currentTrackingCode: make(chan string),
	}
}

func (t TrackingCodeGenerator) Start() {
	go func() {
		for {
			err := t.renewTrackingCodeRange()

			if err != nil {
				fmt.Println(err)
				continue
			}

			for i := t.trackingCodeRange[0]; i <= t.trackingCodeRange[1]; i++ {
				trackingCode, err := t.buildTrackingCode(i)
				if err != nil {
					fmt.Println(err)
					continue
				}

				t.currentTrackingCode <- trackingCode
			}
		}
	}()
}

func (t TrackingCodeGenerator) Next() (string, error) {
	select {
	case ret := <-t.currentTrackingCode:
		return ret, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timeout")
	}
}

func (t *TrackingCodeGenerator) renewTrackingCodeRange() error {
	resp, err := t.correiosApi.SolicitaEtiquetas(t.buildSolicitaEtiquetas())

	if err != nil {
		return err
	}

	return t.updateTrackingCodeAttributes(*resp.Return)
}

func (t *TrackingCodeGenerator) buildTrackingCode(trackingCodeNumber int) (string, error) {
	digit, err := t.digitValidator(trackingCodeNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%d%d%s", t.trackingCodePrefix, trackingCodeNumber, digit, t.trackingCodeSufix), nil
}

func (t *TrackingCodeGenerator) digitValidator(trackingCodeNumber int) (int, error) {
	trackingCodeNumberStr := fmt.Sprintf("%08d", trackingCodeNumber)

	sum := 0
	for i := 0; i < 8; i++ {
		digit, err := strconv.Atoi(trackingCodeNumberStr[i : i+1])
		if err != nil {
			return 0, err
		}

		sum += digit * multipliers[i]
	}

	rest := sum % 11

	if rest == 0 {
		return 5, nil
	} else if rest == 1 {
		return 0, nil
	} else {
		return 11 - rest, nil
	}
}

func (t *TrackingCodeGenerator) updateTrackingCodeAttributes(trackingCodeRangeStr string) (err error) {
	codes := strings.Split(trackingCodeRangeStr, ",")

	t.trackingCodePrefix = codes[0][:2]
	t.trackingCodeSufix = codes[0][11:]

	t.trackingCodeRange[0], err = t.trackingCodeNumber(codes[0])

	if err != nil {
		return err
	}

	t.trackingCodeRange[1], err = t.trackingCodeNumber(codes[1])

	return
}

func (*TrackingCodeGenerator) trackingCodeNumber(tracking_code string) (int, error) {
	return strconv.Atoi(tracking_code[2:10])
}

func (*TrackingCodeGenerator) buildSolicitaEtiquetas() *lib.SolicitaEtiquetas {
	return &lib.SolicitaEtiquetas{
		TipoDestinatario: &tipoDestinatario,
		Identificador:    &app.Config.Identificador,
		IdServico:        &app.Config.IdServico,
		QtdEtiquetas:     &qtdEtiquetas,
		Usuario:          &app.Config.Usuario,
		Senha:            &app.Config.Senha,
	}
}
