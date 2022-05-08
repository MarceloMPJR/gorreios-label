package lib

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://cliente.bean.master.sigep.bsb.correios.com.br/"
var BaseURL = "https://apphom.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente?wsdl"

// NewAtendeCliente creates an initializes a AtendeCliente.
func NewCorreiosSoapService(cli *soap.Client) CorreiosSigepService {
	return &atendeCliente{cli}
}

// AtendeCliente was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type CorreiosSigepService interface {
	// VerificaSeTodosObjetosCancelados was auto-generated from WSDL.
	VerificaSeTodosObjetosCancelados(VerificaSeTodosObjetosCancelados *VerificaSeTodosObjetosCancelados) (*VerificaSeTodosObjetosCanceladosResponse, error)

	// AtualizaPagamentoNaEntrega was auto-generated from WSDL.
	AtualizaPagamentoNaEntrega(AtualizaPagamentoNaEntrega *AtualizaPagamentoNaEntrega) (*AtualizaPagamentoNaEntregaResponse, error)

	// AtualizaRemessaAgrupada was auto-generated from WSDL.
	AtualizaRemessaAgrupada(AtualizaRemessaAgrupada *AtualizaRemessaAgrupada) (*AtualizaRemessaAgrupadaResponse, error)

	// BloquearObjeto was auto-generated from WSDL.
	BloquearObjeto(BloquearObjeto *BloquearObjeto) (*BloquearObjetoResponse, error)

	// BuscaCliente was auto-generated from WSDL.
	BuscaCliente(BuscaCliente *BuscaCliente) (*BuscaClienteResponse, error)

	// BuscaContrato was auto-generated from WSDL.
	BuscaContrato(BuscaContrato *BuscaContrato) (*BuscaContratoResponse, error)

	// BuscaDataAtual was auto-generated from WSDL.
	BuscaDataAtual(BuscaDataAtual *BuscaDataAtual) (*BuscaDataAtualResponse, error)

	// BuscaOpcoes was auto-generated from WSDL.
	BuscaOpcoes(BuscaOpcoes *BuscaOpcoes) (*BuscaOpcoesResponse, error)

	// BuscaPagamentoEntrega was auto-generated from WSDL.
	BuscaPagamentoEntrega(BuscaPagamentoEntrega *BuscaPagamentoEntrega) (*BuscaPagamentoEntregaResponse, error)

	// BuscaServicos was auto-generated from WSDL.
	BuscaServicos(BuscaServicos *BuscaServicos) (*BuscaServicosResponse, error)

	// BuscaServicosAdicionaisAtivos was auto-generated from WSDL.
	BuscaServicosAdicionaisAtivos(BuscaServicosAdicionaisAtivos *BuscaServicosAdicionaisAtivos) (*BuscaServicosAdicionaisAtivosResponse, error)

	// BuscaServicosValorDeclarado was auto-generated from WSDL.
	BuscaServicosValorDeclarado(BuscaServicosValorDeclarado *BuscaServicosValorDeclarado) (*BuscaServicosValorDeclaradoResponse, error)

	// BuscaServicosXServicosAdicionais was auto-generated from WSDL.
	BuscaServicosXServicosAdicionais(BuscaServicosXServicosAdicionais *BuscaServicosXServicosAdicionais) (*BuscaServicosXServicosAdicionaisResponse, error)

	// BuscaTarifaVale was auto-generated from WSDL.
	BuscaTarifaVale(BuscaTarifaVale *BuscaTarifaVale) (*BuscaTarifaValeResponse, error)

	// CalculaTarifaServico was auto-generated from WSDL.
	CalculaTarifaServico(CalculaTarifaServico *CalculaTarifaServico) (*CalculaTarifaServicoResponse, error)

	// CancelarObjeto was auto-generated from WSDL.
	CancelarObjeto(CancelarObjeto *CancelarObjeto) (*CancelarObjetoResponse, error)

	// CancelarPedidoScol was auto-generated from WSDL.
	CancelarPedidoScol(CancelarPedidoScol *CancelarPedidoScol) (*CancelarPedidoScolResponse, error)

	// ConsultaCEP was auto-generated from WSDL.
	ConsultaCEP(ConsultaCEP *ConsultaCEP) (*ConsultaCEPResponse, error)

	// FechaPlp was auto-generated from WSDL.
	FechaPlp(FechaPlp *FechaPlp) (*FechaPlpResponse, error)

	// FechaPlpVariosServicos was auto-generated from WSDL.
	FechaPlpVariosServicos(FechaPlpVariosServicos *FechaPlpVariosServicos) (*FechaPlpVariosServicosResponse, error)

	// GeraDigitoVerificadorEtiquetas was auto-generated from WSDL.
	GeraDigitoVerificadorEtiquetas(GeraDigitoVerificadorEtiquetas *GeraDigitoVerificadorEtiquetas) (*GeraDigitoVerificadorEtiquetasResponse, error)

	// GetStatusCartaoPostagem was auto-generated from WSDL.
	GetStatusCartaoPostagem(GetStatusCartaoPostagem *GetStatusCartaoPostagem) (*GetStatusCartaoPostagemResponse, error)

	// GetStatusPLP was auto-generated from WSDL.
	GetStatusPLP(GetStatusPLP *GetStatusPLP) (*GetStatusPLPResponse, error)

	// IntegrarUsuarioScol was auto-generated from WSDL.
	IntegrarUsuarioScol(IntegrarUsuarioScol *IntegrarUsuarioScol) (*IntegrarUsuarioScolResponse, error)

	// ObterClienteAtualizacao was auto-generated from WSDL.
	ObterClienteAtualizacao(ObterClienteAtualizacao *ObterClienteAtualizacao) (*ObterClienteAtualizacaoResponse, error)

	// ObterEmbalagemLRS was auto-generated from WSDL.
	ObterEmbalagemLRS(ObterEmbalagemLRS *ObterEmbalagemLRS) (*ObterEmbalagemLRSResponse, error)

	// ObterMensagemParametrizada was auto-generated from WSDL.
	ObterMensagemParametrizada(ObterMensagemParametrizada *ObterMensagemParametrizada) (*ObterMensagemParametrizadaResponse, error)

	// PesquisarDimensoesServico was auto-generated from WSDL.
	PesquisarDimensoesServico(PesquisarDimensoesServico *PesquisarDimensoesServico) (*PesquisarDimensoesServicoResponse, error)

	// PesquisarEmbalagensPorServico was auto-generated from WSDL.
	PesquisarEmbalagensPorServico(PesquisarEmbalagensPorServico *PesquisarEmbalagensPorServico) (*PesquisarEmbalagensPorServicoResponse, error)

	// PesquisarParametrosPorDescricao was auto-generated from WSDL.
	PesquisarParametrosPorDescricao(PesquisarParametrosPorDescricao *PesquisarParametrosPorDescricao) (*PesquisarParametrosPorDescricaoResponse, error)

	// PesquisarServicosAdicionais was auto-generated from WSDL.
	PesquisarServicosAdicionais(PesquisarServicosAdicionais *PesquisarServicosAdicionais) (*PesquisarServicosAdicionaisResponse, error)

	// SolicitaEtiquetas was auto-generated from WSDL.
	SolicitaEtiquetas(SolicitaEtiquetas *SolicitaEtiquetas) (*SolicitaEtiquetasResponse, error)

	// SolicitaPLP was auto-generated from WSDL.
	SolicitaPLP(SolicitaPLP *SolicitaPLP) (*SolicitaPLPResponse, error)

	// SolicitaXmlPlp was auto-generated from WSDL.
	SolicitaXmlPlp(SolicitaXmlPlp *SolicitaXmlPlp) (*SolicitaXmlPlpResponse, error)

	// SolicitarPostagemScol was auto-generated from WSDL.
	SolicitarPostagemScol(SolicitarPostagemScol *SolicitarPostagemScol) (*SolicitarPostagemScolResponse, error)

	// ValidaEtiquetaPLP was auto-generated from WSDL.
	ValidaEtiquetaPLP(ValidaEtiquetaPLP *ValidaEtiquetaPLP) (*ValidaEtiquetaPLPResponse, error)

	// ValidaPlp was auto-generated from WSDL.
	ValidaPlp(ValidaPlp *ValidaPlp) (*ValidaPlpResponse, error)

	// ValidarPostagemReversa was auto-generated from WSDL.
	ValidarPostagemReversa(ValidarPostagemReversa *ValidarPostagemReversa) (*ValidarPostagemReversaResponse, error)

	// ValidarPostagemSimultanea was auto-generated from WSDL.
	ValidarPostagemSimultanea(ValidarPostagemSimultanea *ValidarPostagemSimultanea) (*ValidarPostagemSimultaneaResponse, error)

	// VerificaDisponibilidadeServico was auto-generated from WSDL.
	VerificaDisponibilidadeServico(VerificaDisponibilidadeServico *VerificaDisponibilidadeServico) (*VerificaDisponibilidadeServicoResponse, error)

	// VerificaModalTransporte was auto-generated from WSDL.
	VerificaModalTransporte(VerificaModalTransporte *VerificaModalTransporte) (*VerificaModalTransporteResponse, error)
}

// DateTime in WSDL format.
type DateTime string

// Acao was auto-generated from WSDL.
type Acao string

// Validate validates Acao.
func (v Acao) Validate() bool {
	for _, vv := range []string{
		"DEVOLVIDO_AO_REMETENTE",
		"ENCAMINHADO_PARA_REFUGO",
		"REINTEGRADO_E_DEVOLVIDO_AO_REMETENTE",
		"DESBLOQUEADO",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CategoriaServico was auto-generated from WSDL.
type CategoriaServico string

// Validate validates CategoriaServico.
func (v CategoriaServico) Validate() bool {
	for _, vv := range []string{
		"SEM_CATEGORIA",
		"PAC",
		"SEDEX",
		"CARTA",
		"GRANDES_FORMATOS",
		"REVERSO",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// SimNao was auto-generated from WSDL.
type SimNao string

// Validate validates SimNao.
func (v SimNao) Validate() bool {
	for _, vv := range []string{
		"S",
		"N",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatusCartao was auto-generated from WSDL.
type StatusCartao string

// Validate validates StatusCartao.
func (v StatusCartao) Validate() bool {
	for _, vv := range []string{
		"Desconhecido",
		"Normal",
		"Suspenso",
		"Cancelado",
		"Irregular",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatusGerente was auto-generated from WSDL.
type StatusGerente string

// Validate validates StatusGerente.
func (v StatusGerente) Validate() bool {
	for _, vv := range []string{
		"Ativo",
		"Inativo",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatusObjetoPostal was auto-generated from WSDL.
type StatusObjetoPostal string

// Validate validates StatusObjetoPostal.
func (v StatusObjetoPostal) Validate() bool {
	for _, vv := range []string{
		"EmBranco",
		"Postado",
		"Cancelado",
		"Estorno",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatusPlp was auto-generated from WSDL.
type StatusPlp string

// Validate validates StatusPlp.
func (v StatusPlp) Validate() bool {
	for _, vv := range []string{
		"Aberta",
		"Fechada",
		"Postada",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatusUsuario was auto-generated from WSDL.
type StatusUsuario string

// Validate validates StatusUsuario.
func (v StatusUsuario) Validate() bool {
	for _, vv := range []string{
		"Ativo",
		"Inativo",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// TipoBloqueio was auto-generated from WSDL.
type TipoBloqueio string

// Validate validates TipoBloqueio.
func (v TipoBloqueio) Validate() bool {
	for _, vv := range []string{
		"FRAUDE_BLOQUEIO",
		"EXTRAVIO_VAREJO_PRE_INDENIZADO",
		"EXTRAVIO_VAREJO_POS_INDENIZADO",
		"EXTRAVIO_CORPORATIVO",
		"INTERNACIONAL_LDI",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// TipoEmbalagem was auto-generated from WSDL.
type TipoEmbalagem string

// Validate validates TipoEmbalagem.
func (v TipoEmbalagem) Validate() bool {
	for _, vv := range []string{
		"DE",
		"PD",
		"RL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// TipoGerente was auto-generated from WSDL.
type TipoGerente string

// Validate validates TipoGerente.
func (v TipoGerente) Validate() bool {
	for _, vv := range []string{
		"GerenteConta",
		"GerenteContaMaster",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// TipoMensagem was auto-generated from WSDL.
type TipoMensagem string

// Validate validates TipoMensagem.
func (v TipoMensagem) Validate() bool {
	for _, vv := range []string{
		"A",
		"E",
		"S",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ErroMontagemRelatorio was auto-generated from WSDL.
type ErroMontagemRelatorio struct {
	Message *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// Exception was auto-generated from WSDL.
type Exception struct {
	Message *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// SQLException was auto-generated from WSDL.
type SQLException struct {
	ErrorCode *int    `xml:"errorCode,omitempty" json:"errorCode,omitempty" yaml:"errorCode,omitempty"`
	SQLState  *string `xml:"sQLState,omitempty" json:"sQLState,omitempty" yaml:"sQLState,omitempty"`
	Message   *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// VerificaSeTodosObjetosCancelados was auto-generated from WSDL.
type VerificaSeTodosObjetosCancelados struct {
	Arg0 []*ObjetoPostal `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// VerificaSeTodosObjetosCanceladosResponse was auto-generated
// from WSDL.
type VerificaSeTodosObjetosCanceladosResponse struct {
	Return *bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// AtualizaPagamentoNaEntrega was auto-generated from WSDL.
type AtualizaPagamentoNaEntrega struct {
	Usuario *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha   *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// AtualizaPagamentoNaEntregaResponse was auto-generated from WSDL.
type AtualizaPagamentoNaEntregaResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// AtualizaRemessaAgrupada was auto-generated from WSDL.
type AtualizaRemessaAgrupada struct {
	Usuario *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha   *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// AtualizaRemessaAgrupadaResponse was auto-generated from WSDL.
type AtualizaRemessaAgrupadaResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BloquearObjeto was auto-generated from WSDL.
type BloquearObjeto struct {
	NumeroEtiqueta *string       `xml:"numeroEtiqueta,omitempty" json:"numeroEtiqueta,omitempty" yaml:"numeroEtiqueta,omitempty"`
	IdPlp          *int64        `xml:"idPlp,omitempty" json:"idPlp,omitempty" yaml:"idPlp,omitempty"`
	TipoBloqueio   *TipoBloqueio `xml:"tipoBloqueio,omitempty" json:"tipoBloqueio,omitempty" yaml:"tipoBloqueio,omitempty"`
	Acao           *Acao         `xml:"acao,omitempty" json:"acao,omitempty" yaml:"acao,omitempty"`
	Usuario        *string       `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha          *string       `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// BloquearObjetoResponse was auto-generated from WSDL.
type BloquearObjetoResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaCliente was auto-generated from WSDL.
type BuscaCliente struct {
	IdContrato       *string `xml:"idContrato,omitempty" json:"idContrato,omitempty" yaml:"idContrato,omitempty"`
	IdCartaoPostagem *string `xml:"idCartaoPostagem,omitempty" json:"idCartaoPostagem,omitempty" yaml:"idCartaoPostagem,omitempty"`
	Usuario          *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha            *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// BuscaClienteResponse was auto-generated from WSDL.
type BuscaClienteResponse struct {
	Return *ClienteERP `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaContrato was auto-generated from WSDL.
type BuscaContrato struct {
	Numero    *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Diretoria *int64  `xml:"diretoria,omitempty" json:"diretoria,omitempty" yaml:"diretoria,omitempty"`
	Usuario   *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha     *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// BuscaContratoResponse was auto-generated from WSDL.
type BuscaContratoResponse struct {
	Return *ContratoERP `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaDataAtual was auto-generated from WSDL.
type BuscaDataAtual struct {
}

// BuscaDataAtualResponse was auto-generated from WSDL.
type BuscaDataAtualResponse struct {
	Return *DateTime `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaOpcoes was auto-generated from WSDL.
type BuscaOpcoes struct {
	ListaObjetos  []*string `xml:"listaObjetos,omitempty" json:"listaObjetos,omitempty" yaml:"listaObjetos,omitempty"`
	TipoResultado *string   `xml:"tipoResultado,omitempty" json:"tipoResultado,omitempty" yaml:"tipoResultado,omitempty"`
	Usuario       *string   `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha         *string   `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// BuscaOpcoesResponse was auto-generated from WSDL.
type BuscaOpcoesResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaPagamentoEntrega was auto-generated from WSDL.
type BuscaPagamentoEntrega struct {
	Usuario    *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha      *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
	Contrato   *string `xml:"contrato,omitempty" json:"contrato,omitempty" yaml:"contrato,omitempty"`
	DataInicio *string `xml:"dataInicio,omitempty" json:"dataInicio,omitempty" yaml:"dataInicio,omitempty"`
	DataFim    *string `xml:"dataFim,omitempty" json:"dataFim,omitempty" yaml:"dataFim,omitempty"`
	Etiqueta   *string `xml:"etiqueta,omitempty" json:"etiqueta,omitempty" yaml:"etiqueta,omitempty"`
}

// BuscaPagamentoEntregaResponse was auto-generated from WSDL.
type BuscaPagamentoEntregaResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaServicos was auto-generated from WSDL.
type BuscaServicos struct {
	IdContrato       *string `xml:"idContrato,omitempty" json:"idContrato,omitempty" yaml:"idContrato,omitempty"`
	IdCartaoPostagem *string `xml:"idCartaoPostagem,omitempty" json:"idCartaoPostagem,omitempty" yaml:"idCartaoPostagem,omitempty"`
	Usuario          *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha            *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// BuscaServicosAdicionaisAtivos was auto-generated from WSDL.
type BuscaServicosAdicionaisAtivos struct {
	Usuario *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha   *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// BuscaServicosAdicionaisAtivosResponse was auto-generated from
// WSDL.
type BuscaServicosAdicionaisAtivosResponse struct {
	Return []*ServicoAdicionalXML `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaServicosResponse was auto-generated from WSDL.
type BuscaServicosResponse struct {
	Return []*ServicoERP `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaServicosValorDeclarado was auto-generated from WSDL.
type BuscaServicosValorDeclarado struct {
	Usuario *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha   *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// BuscaServicosValorDeclaradoResponse was auto-generated from
// WSDL.
type BuscaServicosValorDeclaradoResponse struct {
	Return []*string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaServicosXServicosAdicionais was auto-generated from WSDL.
type BuscaServicosXServicosAdicionais struct {
	Usuario *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha   *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// BuscaServicosXServicosAdicionaisResponse was auto-generated
// from WSDL.
type BuscaServicosXServicosAdicionaisResponse struct {
	Return []*string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// BuscaTarifaVale was auto-generated from WSDL.
type BuscaTarifaVale struct {
	CodAdministrativo *string  `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	Usuario           *string  `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha             *string  `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
	CodServico        *string  `xml:"codServico,omitempty" json:"codServico,omitempty" yaml:"codServico,omitempty"`
	CepOrigem         *string  `xml:"cepOrigem,omitempty" json:"cepOrigem,omitempty" yaml:"cepOrigem,omitempty"`
	CepDestino        *string  `xml:"cepDestino,omitempty" json:"cepDestino,omitempty" yaml:"cepDestino,omitempty"`
	Peso              *string  `xml:"peso,omitempty" json:"peso,omitempty" yaml:"peso,omitempty"`
	CodFormato        *int     `xml:"codFormato,omitempty" json:"codFormato,omitempty" yaml:"codFormato,omitempty"`
	Comprimento       *float64 `xml:"comprimento,omitempty" json:"comprimento,omitempty" yaml:"comprimento,omitempty"`
	Altura            *float64 `xml:"altura,omitempty" json:"altura,omitempty" yaml:"altura,omitempty"`
	Largura           *float64 `xml:"largura,omitempty" json:"largura,omitempty" yaml:"largura,omitempty"`
	ValorDeclarado    *float64 `xml:"valorDeclarado,omitempty" json:"valorDeclarado,omitempty" yaml:"valorDeclarado,omitempty"`
	ServicoAdicional  *string  `xml:"servicoAdicional,omitempty" json:"servicoAdicional,omitempty" yaml:"servicoAdicional,omitempty"`
}

// BuscaTarifaValeResponse was auto-generated from WSDL.
type BuscaTarifaValeResponse struct {
	Return *ValePostal `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// CalculaTarifaServico was auto-generated from WSDL.
type CalculaTarifaServico struct {
	CodAdministrativo   *string  `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	Usuario             *string  `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha               *string  `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
	CodServico          *string  `xml:"codServico,omitempty" json:"codServico,omitempty" yaml:"codServico,omitempty"`
	CepOrigem           *string  `xml:"cepOrigem,omitempty" json:"cepOrigem,omitempty" yaml:"cepOrigem,omitempty"`
	CepDestino          *string  `xml:"cepDestino,omitempty" json:"cepDestino,omitempty" yaml:"cepDestino,omitempty"`
	Peso                *string  `xml:"peso,omitempty" json:"peso,omitempty" yaml:"peso,omitempty"`
	CodFormato          *int     `xml:"codFormato,omitempty" json:"codFormato,omitempty" yaml:"codFormato,omitempty"`
	Comprimento         *float64 `xml:"comprimento,omitempty" json:"comprimento,omitempty" yaml:"comprimento,omitempty"`
	Altura              *float64 `xml:"altura,omitempty" json:"altura,omitempty" yaml:"altura,omitempty"`
	Largura             *float64 `xml:"largura,omitempty" json:"largura,omitempty" yaml:"largura,omitempty"`
	Diametro            *float64 `xml:"diametro,omitempty" json:"diametro,omitempty" yaml:"diametro,omitempty"`
	CodMaoPropria       *string  `xml:"codMaoPropria,omitempty" json:"codMaoPropria,omitempty" yaml:"codMaoPropria,omitempty"`
	ValorDeclarado      *float64 `xml:"valorDeclarado,omitempty" json:"valorDeclarado,omitempty" yaml:"valorDeclarado,omitempty"`
	CodAvisoRecebimento *string  `xml:"codAvisoRecebimento,omitempty" json:"codAvisoRecebimento,omitempty" yaml:"codAvisoRecebimento,omitempty"`
}

// CalculaTarifaServicoResponse was auto-generated from WSDL.
type CalculaTarifaServicoResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// CancelarObjeto was auto-generated from WSDL.
type CancelarObjeto struct {
	IdPlp          *int64  `xml:"idPlp,omitempty" json:"idPlp,omitempty" yaml:"idPlp,omitempty"`
	NumeroEtiqueta *string `xml:"numeroEtiqueta,omitempty" json:"numeroEtiqueta,omitempty" yaml:"numeroEtiqueta,omitempty"`
	Usuario        *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha          *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// CancelarObjetoResponse was auto-generated from WSDL.
type CancelarObjetoResponse struct {
	Return *bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// CancelarPedidoScol was auto-generated from WSDL.
type CancelarPedidoScol struct {
	CodAdministrativo *string `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	IdPostagem        *string `xml:"idPostagem,omitempty" json:"idPostagem,omitempty" yaml:"idPostagem,omitempty"`
	Tipo              *string `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Usuario           *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha             *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// CancelarPedidoScolResponse was auto-generated from WSDL.
type CancelarPedidoScolResponse struct {
	Return *RetornoCancelamento `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// CartaoPostagemERP was auto-generated from WSDL.
type CartaoPostagemERP struct {
	CodigoAdministrativo             *string               `xml:"codigoAdministrativo,omitempty" json:"codigoAdministrativo,omitempty" yaml:"codigoAdministrativo,omitempty"`
	Contratos                        []*ContratoERP        `xml:"contratos,omitempty" json:"contratos,omitempty" yaml:"contratos,omitempty"`
	DataAtualizacao                  *DateTime             `xml:"dataAtualizacao,omitempty" json:"dataAtualizacao,omitempty" yaml:"dataAtualizacao,omitempty"`
	DataVigenciaFim                  *DateTime             `xml:"dataVigenciaFim,omitempty" json:"dataVigenciaFim,omitempty" yaml:"dataVigenciaFim,omitempty"`
	DataVigenciaInicio               *DateTime             `xml:"dataVigenciaInicio,omitempty" json:"dataVigenciaInicio,omitempty" yaml:"dataVigenciaInicio,omitempty"`
	DatajAtualizacao                 *int                  `xml:"datajAtualizacao,omitempty" json:"datajAtualizacao,omitempty" yaml:"datajAtualizacao,omitempty"`
	DatajVigenciaFim                 *int                  `xml:"datajVigenciaFim,omitempty" json:"datajVigenciaFim,omitempty" yaml:"datajVigenciaFim,omitempty"`
	DatajVigenciaInicio              *int                  `xml:"datajVigenciaInicio,omitempty" json:"datajVigenciaInicio,omitempty" yaml:"datajVigenciaInicio,omitempty"`
	DescricaoStatusCartao            *string               `xml:"descricaoStatusCartao,omitempty" json:"descricaoStatusCartao,omitempty" yaml:"descricaoStatusCartao,omitempty"`
	DescricaoUnidadePostagemGenerica *string               `xml:"descricaoUnidadePostagemGenerica,omitempty" json:"descricaoUnidadePostagemGenerica,omitempty" yaml:"descricaoUnidadePostagemGenerica,omitempty"`
	HorajAtualizacao                 *int                  `xml:"horajAtualizacao,omitempty" json:"horajAtualizacao,omitempty" yaml:"horajAtualizacao,omitempty"`
	Numero                           *string               `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Servicos                         []*ServicoERP         `xml:"servicos,omitempty" json:"servicos,omitempty" yaml:"servicos,omitempty"`
	StatusCartaoPostagem             *string               `xml:"statusCartaoPostagem,omitempty" json:"statusCartaoPostagem,omitempty" yaml:"statusCartaoPostagem,omitempty"`
	StatusCodigo                     *string               `xml:"statusCodigo,omitempty" json:"statusCodigo,omitempty" yaml:"statusCodigo,omitempty"`
	UnidadeGenerica                  *string               `xml:"unidadeGenerica,omitempty" json:"unidadeGenerica,omitempty" yaml:"unidadeGenerica,omitempty"`
	UnidadesPostagem                 []*UnidadePostagemERP `xml:"unidadesPostagem,omitempty" json:"unidadesPostagem,omitempty" yaml:"unidadesPostagem,omitempty"`
}

// ChancelaMaster was auto-generated from WSDL.
type ChancelaMaster struct {
	Ativo           *SimNao         `xml:"ativo,omitempty" json:"ativo,omitempty" yaml:"ativo,omitempty"`
	Chancela        *[]byte         `xml:"chancela,omitempty" json:"chancela,omitempty" yaml:"chancela,omitempty"`
	DataAtualizacao *DateTime       `xml:"dataAtualizacao,omitempty" json:"dataAtualizacao,omitempty" yaml:"dataAtualizacao,omitempty"`
	Descricao       *string         `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	Id              *int64          `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	ServicosSigep   []*ServicoSigep `xml:"servicosSigep,omitempty" json:"servicosSigep,omitempty" yaml:"servicosSigep,omitempty"`
}

// ClienteERP was auto-generated from WSDL.
type ClienteERP struct {
	Cnpj                   *string         `xml:"cnpj,omitempty" json:"cnpj,omitempty" yaml:"cnpj,omitempty"`
	Contratos              []*ContratoERP  `xml:"contratos,omitempty" json:"contratos,omitempty" yaml:"contratos,omitempty"`
	DataAtualizacao        *DateTime       `xml:"dataAtualizacao,omitempty" json:"dataAtualizacao,omitempty" yaml:"dataAtualizacao,omitempty"`
	DatajAtualizacao       *int            `xml:"datajAtualizacao,omitempty" json:"datajAtualizacao,omitempty" yaml:"datajAtualizacao,omitempty"`
	DescricaoStatusCliente *string         `xml:"descricaoStatusCliente,omitempty" json:"descricaoStatusCliente,omitempty" yaml:"descricaoStatusCliente,omitempty"`
	GerenteConta           []*GerenteConta `xml:"gerenteConta,omitempty" json:"gerenteConta,omitempty" yaml:"gerenteConta,omitempty"`
	HorajAtualizacao       *int64          `xml:"horajAtualizacao,omitempty" json:"horajAtualizacao,omitempty" yaml:"horajAtualizacao,omitempty"`
	Id                     *int64          `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	InscricaoEstadual      *string         `xml:"inscricaoEstadual,omitempty" json:"inscricaoEstadual,omitempty" yaml:"inscricaoEstadual,omitempty"`
	Nome                   *string         `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	StatusCodigo           *string         `xml:"statusCodigo,omitempty" json:"statusCodigo,omitempty" yaml:"statusCodigo,omitempty"`
}

// Coleta was auto-generated from WSDL.
type Coleta struct {
	Cklist          *string    `xml:"cklist,omitempty" json:"cklist,omitempty" yaml:"cklist,omitempty"`
	Descricao       *string    `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	Documento       []*string  `xml:"documento,omitempty" json:"documento,omitempty" yaml:"documento,omitempty"`
	Id_cliente      *string    `xml:"id_cliente,omitempty" json:"id_cliente,omitempty" yaml:"id_cliente,omitempty"`
	Produto         []*Produto `xml:"produto,omitempty" json:"produto,omitempty" yaml:"produto,omitempty"`
	Remetente       *Remetente `xml:"remetente,omitempty" json:"remetente,omitempty" yaml:"remetente,omitempty"`
	Tipo            *string    `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Valor_declarado *string    `xml:"valor_declarado,omitempty" json:"valor_declarado,omitempty" yaml:"valor_declarado,omitempty"`
}

// ColetaReversa was auto-generated from WSDL.
type ColetaReversa struct {
	Cklist            *string    `xml:"cklist,omitempty" json:"cklist,omitempty" yaml:"cklist,omitempty"`
	Descricao         *string    `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	Documento         []*string  `xml:"documento,omitempty" json:"documento,omitempty" yaml:"documento,omitempty"`
	Id_cliente        *string    `xml:"id_cliente,omitempty" json:"id_cliente,omitempty" yaml:"id_cliente,omitempty"`
	Produto           []*Produto `xml:"produto,omitempty" json:"produto,omitempty" yaml:"produto,omitempty"`
	Remetente         *Remetente `xml:"remetente,omitempty" json:"remetente,omitempty" yaml:"remetente,omitempty"`
	Tipo              *string    `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Valor_declarado   *string    `xml:"valor_declarado,omitempty" json:"valor_declarado,omitempty" yaml:"valor_declarado,omitempty"`
	Ag                *string    `xml:"ag,omitempty" json:"ag,omitempty" yaml:"ag,omitempty"`
	Ar                *int       `xml:"ar,omitempty" json:"ar,omitempty" yaml:"ar,omitempty"`
	Cartao            *string    `xml:"cartao,omitempty" json:"cartao,omitempty" yaml:"cartao,omitempty"`
	Numero            *int       `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Obj_col           []*Objeto  `xml:"obj_col,omitempty" json:"obj_col,omitempty" yaml:"obj_col,omitempty"`
	Servico_adicional *string    `xml:"servico_adicional,omitempty" json:"servico_adicional,omitempty" yaml:"servico_adicional,omitempty"`
	TypeAttrXSI       string     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ColetaReversa) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:coletaReversa"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://cliente.bean.master.sigep.bsb.correios.com.br/"
	}
}

// ColetaSimultanea was auto-generated from WSDL.
type ColetaSimultanea struct {
	Cklist          *string    `xml:"cklist,omitempty" json:"cklist,omitempty" yaml:"cklist,omitempty"`
	Descricao       *string    `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	Documento       []*string  `xml:"documento,omitempty" json:"documento,omitempty" yaml:"documento,omitempty"`
	Id_cliente      *string    `xml:"id_cliente,omitempty" json:"id_cliente,omitempty" yaml:"id_cliente,omitempty"`
	Produto         []*Produto `xml:"produto,omitempty" json:"produto,omitempty" yaml:"produto,omitempty"`
	Remetente       *Remetente `xml:"remetente,omitempty" json:"remetente,omitempty" yaml:"remetente,omitempty"`
	Tipo            *string    `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Valor_declarado *string    `xml:"valor_declarado,omitempty" json:"valor_declarado,omitempty" yaml:"valor_declarado,omitempty"`
	Obj             *string    `xml:"obj,omitempty" json:"obj,omitempty" yaml:"obj,omitempty"`
	Obs             *string    `xml:"obs,omitempty" json:"obs,omitempty" yaml:"obs,omitempty"`
	TypeAttrXSI     string     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ColetaSimultanea) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:coletaSimultanea"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://cliente.bean.master.sigep.bsb.correios.com.br/"
	}
}

// ConsultaCEP was auto-generated from WSDL.
type ConsultaCEP struct {
	Cep *string `xml:"cep,omitempty" json:"cep,omitempty" yaml:"cep,omitempty"`
}

// ConsultaCEPResponse was auto-generated from WSDL.
type ConsultaCEPResponse struct {
	Return *EnderecoERP `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ContratoERP was auto-generated from WSDL.
type ContratoERP struct {
	CartoesPostagem            []*CartaoPostagemERP `xml:"cartoesPostagem,omitempty" json:"cartoesPostagem,omitempty" yaml:"cartoesPostagem,omitempty"`
	Cliente                    *ClienteERP          `xml:"cliente,omitempty" json:"cliente,omitempty" yaml:"cliente,omitempty"`
	CodigoCliente              *int64               `xml:"codigoCliente,omitempty" json:"codigoCliente,omitempty" yaml:"codigoCliente,omitempty"`
	CodigoDiretoria            *string              `xml:"codigoDiretoria,omitempty" json:"codigoDiretoria,omitempty" yaml:"codigoDiretoria,omitempty"`
	ContratoPK                 *ContratoERPPK       `xml:"contratoPK,omitempty" json:"contratoPK,omitempty" yaml:"contratoPK,omitempty"`
	DataAtualizacao            *DateTime            `xml:"dataAtualizacao,omitempty" json:"dataAtualizacao,omitempty" yaml:"dataAtualizacao,omitempty"`
	DataAtualizacaoDDMMYYYY    *string              `xml:"dataAtualizacaoDDMMYYYY,omitempty" json:"dataAtualizacaoDDMMYYYY,omitempty" yaml:"dataAtualizacaoDDMMYYYY,omitempty"`
	DataVigenciaFim            *DateTime            `xml:"dataVigenciaFim,omitempty" json:"dataVigenciaFim,omitempty" yaml:"dataVigenciaFim,omitempty"`
	DataVigenciaFimDDMMYYYY    *string              `xml:"dataVigenciaFimDDMMYYYY,omitempty" json:"dataVigenciaFimDDMMYYYY,omitempty" yaml:"dataVigenciaFimDDMMYYYY,omitempty"`
	DataVigenciaInicio         *DateTime            `xml:"dataVigenciaInicio,omitempty" json:"dataVigenciaInicio,omitempty" yaml:"dataVigenciaInicio,omitempty"`
	DataVigenciaInicioDDMMYYYY *string              `xml:"dataVigenciaInicioDDMMYYYY,omitempty" json:"dataVigenciaInicioDDMMYYYY,omitempty" yaml:"dataVigenciaInicioDDMMYYYY,omitempty"`
	DatajAtualizacao           *int                 `xml:"datajAtualizacao,omitempty" json:"datajAtualizacao,omitempty" yaml:"datajAtualizacao,omitempty"`
	DatajVigenciaFim           *int                 `xml:"datajVigenciaFim,omitempty" json:"datajVigenciaFim,omitempty" yaml:"datajVigenciaFim,omitempty"`
	DatajVigenciaInicio        *int                 `xml:"datajVigenciaInicio,omitempty" json:"datajVigenciaInicio,omitempty" yaml:"datajVigenciaInicio,omitempty"`
	DescricaoDiretoriaRegional *string              `xml:"descricaoDiretoriaRegional,omitempty" json:"descricaoDiretoriaRegional,omitempty" yaml:"descricaoDiretoriaRegional,omitempty"`
	DescricaoStatus            *string              `xml:"descricaoStatus,omitempty" json:"descricaoStatus,omitempty" yaml:"descricaoStatus,omitempty"`
	DiretoriaRegional          *UnidadePostagemERP  `xml:"diretoriaRegional,omitempty" json:"diretoriaRegional,omitempty" yaml:"diretoriaRegional,omitempty"`
	HorajAtualizacao           *int                 `xml:"horajAtualizacao,omitempty" json:"horajAtualizacao,omitempty" yaml:"horajAtualizacao,omitempty"`
	StatusCodigo               *string              `xml:"statusCodigo,omitempty" json:"statusCodigo,omitempty" yaml:"statusCodigo,omitempty"`
}

// ContratoERPPK was auto-generated from WSDL.
type ContratoERPPK struct {
	Diretoria *int64  `xml:"diretoria,omitempty" json:"diretoria,omitempty" yaml:"diretoria,omitempty"`
	Numero    *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
}

// DimensaoTO was auto-generated from WSDL.
type DimensaoTO struct {
	Altura      *MedidaTO `xml:"altura,omitempty" json:"altura,omitempty" yaml:"altura,omitempty"`
	Comprimento *MedidaTO `xml:"comprimento,omitempty" json:"comprimento,omitempty" yaml:"comprimento,omitempty"`
	Diametro    *MedidaTO `xml:"diametro,omitempty" json:"diametro,omitempty" yaml:"diametro,omitempty"`
	Largura     *MedidaTO `xml:"largura,omitempty" json:"largura,omitempty" yaml:"largura,omitempty"`
	Peso        *MedidaTO `xml:"peso,omitempty" json:"peso,omitempty" yaml:"peso,omitempty"`
	Soma        *MedidaTO `xml:"soma,omitempty" json:"soma,omitempty" yaml:"soma,omitempty"`
}

// EmbalagemLRSMaster was auto-generated from WSDL.
type EmbalagemLRSMaster struct {
	Codigo *string `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
	Nome   *string `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	Tipo   *string `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
}

// EnderecoERP was auto-generated from WSDL.
type EnderecoERP struct {
	Bairro           *string               `xml:"bairro,omitempty" json:"bairro,omitempty" yaml:"bairro,omitempty"`
	Cep              *string               `xml:"cep,omitempty" json:"cep,omitempty" yaml:"cep,omitempty"`
	Cidade           *string               `xml:"cidade,omitempty" json:"cidade,omitempty" yaml:"cidade,omitempty"`
	Complemento2     *string               `xml:"complemento2,omitempty" json:"complemento2,omitempty" yaml:"complemento2,omitempty"`
	End              *string               `xml:"end,omitempty" json:"end,omitempty" yaml:"end,omitempty"`
	Uf               *string               `xml:"uf,omitempty" json:"uf,omitempty" yaml:"uf,omitempty"`
	UnidadesPostagem []*UnidadePostagemERP `xml:"unidadesPostagem,omitempty" json:"unidadesPostagem,omitempty" yaml:"unidadesPostagem,omitempty"`
}

// FechaPlp was auto-generated from WSDL.
type FechaPlp struct {
	Xml            *string `xml:"xml,omitempty" json:"xml,omitempty" yaml:"xml,omitempty"`
	IdPlpCliente   *int64  `xml:"idPlpCliente,omitempty" json:"idPlpCliente,omitempty" yaml:"idPlpCliente,omitempty"`
	CartaoPostagem *string `xml:"cartaoPostagem,omitempty" json:"cartaoPostagem,omitempty" yaml:"cartaoPostagem,omitempty"`
	FaixaEtiquetas *string `xml:"faixaEtiquetas,omitempty" json:"faixaEtiquetas,omitempty" yaml:"faixaEtiquetas,omitempty"`
	Usuario        *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha          *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// FechaPlpResponse was auto-generated from WSDL.
type FechaPlpResponse struct {
	Return *int64 `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// FechaPlpVariosServicos was auto-generated from WSDL.
type FechaPlpVariosServicos struct {
	Xml            *string   `xml:"xml,omitempty" json:"xml,omitempty" yaml:"xml,omitempty"`
	IdPlpCliente   *int64    `xml:"idPlpCliente,omitempty" json:"idPlpCliente,omitempty" yaml:"idPlpCliente,omitempty"`
	CartaoPostagem *string   `xml:"cartaoPostagem,omitempty" json:"cartaoPostagem,omitempty" yaml:"cartaoPostagem,omitempty"`
	ListaEtiquetas []*string `xml:"listaEtiquetas,omitempty" json:"listaEtiquetas,omitempty" yaml:"listaEtiquetas,omitempty"`
	Usuario        *string   `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha          *string   `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// FechaPlpVariosServicosResponse was auto-generated from WSDL.
type FechaPlpVariosServicosResponse struct {
	Return *int64 `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GeraDigitoVerificadorEtiquetas was auto-generated from WSDL.
type GeraDigitoVerificadorEtiquetas struct {
	Etiquetas []*string `xml:"etiquetas,omitempty" json:"etiquetas,omitempty" yaml:"etiquetas,omitempty"`
	Usuario   *string   `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha     *string   `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// GeraDigitoVerificadorEtiquetasResponse was auto-generated from
// WSDL.
type GeraDigitoVerificadorEtiquetasResponse struct {
	Return []*int `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GerenteConta was auto-generated from WSDL.
type GerenteConta struct {
	ClientesVisiveis   []*ClienteERP        `xml:"clientesVisiveis,omitempty" json:"clientesVisiveis,omitempty" yaml:"clientesVisiveis,omitempty"`
	DataAtualizacao    *DateTime            `xml:"dataAtualizacao,omitempty" json:"dataAtualizacao,omitempty" yaml:"dataAtualizacao,omitempty"`
	DataInclusao       *DateTime            `xml:"dataInclusao,omitempty" json:"dataInclusao,omitempty" yaml:"dataInclusao,omitempty"`
	DataSenha          *DateTime            `xml:"dataSenha,omitempty" json:"dataSenha,omitempty" yaml:"dataSenha,omitempty"`
	Login              *string              `xml:"login,omitempty" json:"login,omitempty" yaml:"login,omitempty"`
	Matricula          *string              `xml:"matricula,omitempty" json:"matricula,omitempty" yaml:"matricula,omitempty"`
	Senha              *string              `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
	Status             *StatusGerente       `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	TipoGerente        *TipoGerente         `xml:"tipoGerente,omitempty" json:"tipoGerente,omitempty" yaml:"tipoGerente,omitempty"`
	UsuariosInstalacao []*UsuarioInstalacao `xml:"usuariosInstalacao,omitempty" json:"usuariosInstalacao,omitempty" yaml:"usuariosInstalacao,omitempty"`
	Validade           *string              `xml:"validade,omitempty" json:"validade,omitempty" yaml:"validade,omitempty"`
}

// GetStatusCartaoPostagem was auto-generated from WSDL.
type GetStatusCartaoPostagem struct {
	NumeroCartaoPostagem *string `xml:"numeroCartaoPostagem,omitempty" json:"numeroCartaoPostagem,omitempty" yaml:"numeroCartaoPostagem,omitempty"`
	Usuario              *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha                *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// GetStatusCartaoPostagemResponse was auto-generated from WSDL.
type GetStatusCartaoPostagemResponse struct {
	Return *StatusCartao `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetStatusPLP was auto-generated from WSDL.
type GetStatusPLP struct {
	Arg0 []*ObjetoPostal `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
	Arg1 *string         `xml:"arg1,omitempty" json:"arg1,omitempty" yaml:"arg1,omitempty"`
}

// GetStatusPLPResponse was auto-generated from WSDL.
type GetStatusPLPResponse struct {
	Return *StatusPlp `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// IntegrarUsuarioScol was auto-generated from WSDL.
type IntegrarUsuarioScol struct {
	CodAdministrativo *int    `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	Usuario           *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha             *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// IntegrarUsuarioScolResponse was auto-generated from WSDL.
type IntegrarUsuarioScolResponse struct {
	Return *bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// MedidaTO was auto-generated from WSDL.
type MedidaTO struct {
	Maximo *float64 `xml:"maximo,omitempty" json:"maximo,omitempty" yaml:"maximo,omitempty"`
	Minimo *float64 `xml:"minimo,omitempty" json:"minimo,omitempty" yaml:"minimo,omitempty"`
}

// MensagemParametrizadaTO was auto-generated from WSDL.
type MensagemParametrizadaTO struct {
	Mensagem *string       `xml:"mensagem,omitempty" json:"mensagem,omitempty" yaml:"mensagem,omitempty"`
	Tipo     *TipoMensagem `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Titulo   *string       `xml:"titulo,omitempty" json:"titulo,omitempty" yaml:"titulo,omitempty"`
}

// Objeto was auto-generated from WSDL.
type Objeto struct {
	Desc    *string `xml:"desc,omitempty" json:"desc,omitempty" yaml:"desc,omitempty"`
	Entrega *string `xml:"entrega,omitempty" json:"entrega,omitempty" yaml:"entrega,omitempty"`
	Id      *string `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Item    *string `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
	Num     *string `xml:"num,omitempty" json:"num,omitempty" yaml:"num,omitempty"`
}

// ObjetoPostal was auto-generated from WSDL.
type ObjetoPostal struct {
	CodigoEtiqueta           *string             `xml:"codigoEtiqueta,omitempty" json:"codigoEtiqueta,omitempty" yaml:"codigoEtiqueta,omitempty"`
	DataAtualizacaoCliente   *DateTime           `xml:"dataAtualizacaoCliente,omitempty" json:"dataAtualizacaoCliente,omitempty" yaml:"dataAtualizacaoCliente,omitempty"`
	DataBloqueioObjeto       *DateTime           `xml:"dataBloqueioObjeto,omitempty" json:"dataBloqueioObjeto,omitempty" yaml:"dataBloqueioObjeto,omitempty"`
	DataCancelamentoEtiqueta *DateTime           `xml:"dataCancelamentoEtiqueta,omitempty" json:"dataCancelamentoEtiqueta,omitempty" yaml:"dataCancelamentoEtiqueta,omitempty"`
	DataInclusao             *DateTime           `xml:"dataInclusao,omitempty" json:"dataInclusao,omitempty" yaml:"dataInclusao,omitempty"`
	ObjetoPostalPK           *ObjetoPostalPK     `xml:"objetoPostalPK,omitempty" json:"objetoPostalPK,omitempty" yaml:"objetoPostalPK,omitempty"`
	PlpNu                    *int64              `xml:"plpNu,omitempty" json:"plpNu,omitempty" yaml:"plpNu,omitempty"`
	PreListaPostagem         *PreListaPostagem   `xml:"preListaPostagem,omitempty" json:"preListaPostagem,omitempty" yaml:"preListaPostagem,omitempty"`
	RestricaoAerea           *SimNao             `xml:"restricaoAerea,omitempty" json:"restricaoAerea,omitempty" yaml:"restricaoAerea,omitempty"`
	StatusBloqueio           *string             `xml:"statusBloqueio,omitempty" json:"statusBloqueio,omitempty" yaml:"statusBloqueio,omitempty"`
	StatusEtiqueta           *StatusObjetoPostal `xml:"statusEtiqueta,omitempty" json:"statusEtiqueta,omitempty" yaml:"statusEtiqueta,omitempty"`
}

// ObjetoPostalPK was auto-generated from WSDL.
type ObjetoPostalPK struct {
	CodigoEtiqueta *string `xml:"codigoEtiqueta,omitempty" json:"codigoEtiqueta,omitempty" yaml:"codigoEtiqueta,omitempty"`
	PlpNu          *int64  `xml:"plpNu,omitempty" json:"plpNu,omitempty" yaml:"plpNu,omitempty"`
}

// ObjetoSimplificado was auto-generated from WSDL.
type ObjetoSimplificado struct {
	Datahora_cancelamento *string `xml:"datahora_cancelamento,omitempty" json:"datahora_cancelamento,omitempty" yaml:"datahora_cancelamento,omitempty"`
	Numero_pedido         *int    `xml:"numero_pedido,omitempty" json:"numero_pedido,omitempty" yaml:"numero_pedido,omitempty"`
	Status_pedido         *string `xml:"status_pedido,omitempty" json:"status_pedido,omitempty" yaml:"status_pedido,omitempty"`
}

// ObterClienteAtualizacao was auto-generated from WSDL.
type ObterClienteAtualizacao struct {
	CnpjCliente *string `xml:"cnpjCliente,omitempty" json:"cnpjCliente,omitempty" yaml:"cnpjCliente,omitempty"`
	Usuario     *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha       *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// ObterClienteAtualizacaoResponse was auto-generated from WSDL.
type ObterClienteAtualizacaoResponse struct {
	Return *DateTime `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ObterEmbalagemLRS was auto-generated from WSDL.
type ObterEmbalagemLRS struct {
}

// ObterEmbalagemLRSResponse was auto-generated from WSDL.
type ObterEmbalagemLRSResponse struct {
	Return []*EmbalagemLRSMaster `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ObterMensagemParametrizada was auto-generated from WSDL.
type ObterMensagemParametrizada struct {
	Id *int64 `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
}

// ObterMensagemParametrizadaResponse was auto-generated from WSDL.
type ObterMensagemParametrizadaResponse struct {
	Return *MensagemParametrizadaTO `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ParametroMaster was auto-generated from WSDL.
type ParametroMaster struct {
	PrmCoParametro *int64  `xml:"prmCoParametro,omitempty" json:"prmCoParametro,omitempty" yaml:"prmCoParametro,omitempty"`
	PrmTxParametro *string `xml:"prmTxParametro,omitempty" json:"prmTxParametro,omitempty" yaml:"prmTxParametro,omitempty"`
	PrmTxValor     *string `xml:"prmTxValor,omitempty" json:"prmTxValor,omitempty" yaml:"prmTxValor,omitempty"`
}

// PesquisarDimensoesServico was auto-generated from WSDL.
type PesquisarDimensoesServico struct {
	Codigo    *string `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
	Embalagem *string `xml:"embalagem,omitempty" json:"embalagem,omitempty" yaml:"embalagem,omitempty"`
}

// PesquisarDimensoesServicoResponse was auto-generated from WSDL.
type PesquisarDimensoesServicoResponse struct {
	Return *DimensaoTO `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// PesquisarEmbalagensPorServico was auto-generated from WSDL.
type PesquisarEmbalagensPorServico struct {
	Codigo *string `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
}

// PesquisarEmbalagensPorServicoResponse was auto-generated from
// WSDL.
type PesquisarEmbalagensPorServicoResponse struct {
	Return []*TipoEmbalagem `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// PesquisarParametrosPorDescricao was auto-generated from WSDL.
type PesquisarParametrosPorDescricao struct {
	Prefix *string `xml:"prefix,omitempty" json:"prefix,omitempty" yaml:"prefix,omitempty"`
}

// PesquisarParametrosPorDescricaoResponse was auto-generated from
// WSDL.
type PesquisarParametrosPorDescricaoResponse struct {
	Return *DimensaoTO `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// PesquisarServicosAdicionais was auto-generated from WSDL.
type PesquisarServicosAdicionais struct {
	Codigo *string `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
}

// PesquisarServicosAdicionaisResponse was auto-generated from
// WSDL.
type PesquisarServicosAdicionaisResponse struct {
	Return []*ServicoAdicionalTO `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Pessoa was auto-generated from WSDL.
type Pessoa struct {
	Bairro      *string `xml:"bairro,omitempty" json:"bairro,omitempty" yaml:"bairro,omitempty"`
	Cep         *string `xml:"cep,omitempty" json:"cep,omitempty" yaml:"cep,omitempty"`
	Cidade      *string `xml:"cidade,omitempty" json:"cidade,omitempty" yaml:"cidade,omitempty"`
	Complemento *string `xml:"complemento,omitempty" json:"complemento,omitempty" yaml:"complemento,omitempty"`
	Ddd         *string `xml:"ddd,omitempty" json:"ddd,omitempty" yaml:"ddd,omitempty"`
	Email       *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Logradouro  *string `xml:"logradouro,omitempty" json:"logradouro,omitempty" yaml:"logradouro,omitempty"`
	Nome        *string `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	Numero      *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Referencia  *string `xml:"referencia,omitempty" json:"referencia,omitempty" yaml:"referencia,omitempty"`
	Telefone    *string `xml:"telefone,omitempty" json:"telefone,omitempty" yaml:"telefone,omitempty"`
	Uf          *string `xml:"uf,omitempty" json:"uf,omitempty" yaml:"uf,omitempty"`
}

// PreListaPostagem was auto-generated from WSDL.
type PreListaPostagem struct {
	CartaoPostagem         *CartaoPostagemERP `xml:"cartaoPostagem,omitempty" json:"cartaoPostagem,omitempty" yaml:"cartaoPostagem,omitempty"`
	ConteudoProibido       *SimNao            `xml:"conteudoProibido,omitempty" json:"conteudoProibido,omitempty" yaml:"conteudoProibido,omitempty"`
	DataAtualizacaoCliente *DateTime          `xml:"dataAtualizacaoCliente,omitempty" json:"dataAtualizacaoCliente,omitempty" yaml:"dataAtualizacaoCliente,omitempty"`
	DataAtualizacaoSara    *DateTime          `xml:"dataAtualizacaoSara,omitempty" json:"dataAtualizacaoSara,omitempty" yaml:"dataAtualizacaoSara,omitempty"`
	DataFechamento         *DateTime          `xml:"dataFechamento,omitempty" json:"dataFechamento,omitempty" yaml:"dataFechamento,omitempty"`
	DataPostagem           *DateTime          `xml:"dataPostagem,omitempty" json:"dataPostagem,omitempty" yaml:"dataPostagem,omitempty"`
	DataPostagemSara       *DateTime          `xml:"dataPostagemSara,omitempty" json:"dataPostagemSara,omitempty" yaml:"dataPostagemSara,omitempty"`
	ObjetosPostais         []*ObjetoPostal    `xml:"objetosPostais,omitempty" json:"objetosPostais,omitempty" yaml:"objetosPostais,omitempty"`
	PlpCliente             *int64             `xml:"plpCliente,omitempty" json:"plpCliente,omitempty" yaml:"plpCliente,omitempty"`
	PlpNu                  *int64             `xml:"plpNu,omitempty" json:"plpNu,omitempty" yaml:"plpNu,omitempty"`
	PlpXml                 []*int64           `xml:"plpXml,omitempty" json:"plpXml,omitempty" yaml:"plpXml,omitempty"`
	PlpXmlRetorno          []*int64           `xml:"plpXmlRetorno,omitempty" json:"plpXmlRetorno,omitempty" yaml:"plpXmlRetorno,omitempty"`
	Status                 *StatusPlp         `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
}

// Produto was auto-generated from WSDL.
type Produto struct {
	Codigo *string `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
	Qtd    *string `xml:"qtd,omitempty" json:"qtd,omitempty" yaml:"qtd,omitempty"`
	Tipo   *string `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
}

// Remetente was auto-generated from WSDL.
type Remetente struct {
	Bairro        *string `xml:"bairro,omitempty" json:"bairro,omitempty" yaml:"bairro,omitempty"`
	Cep           *string `xml:"cep,omitempty" json:"cep,omitempty" yaml:"cep,omitempty"`
	Cidade        *string `xml:"cidade,omitempty" json:"cidade,omitempty" yaml:"cidade,omitempty"`
	Complemento   *string `xml:"complemento,omitempty" json:"complemento,omitempty" yaml:"complemento,omitempty"`
	Ddd           *string `xml:"ddd,omitempty" json:"ddd,omitempty" yaml:"ddd,omitempty"`
	Email         *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Logradouro    *string `xml:"logradouro,omitempty" json:"logradouro,omitempty" yaml:"logradouro,omitempty"`
	Nome          *string `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	Numero        *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Referencia    *string `xml:"referencia,omitempty" json:"referencia,omitempty" yaml:"referencia,omitempty"`
	Telefone      *string `xml:"telefone,omitempty" json:"telefone,omitempty" yaml:"telefone,omitempty"`
	Uf            *string `xml:"uf,omitempty" json:"uf,omitempty" yaml:"uf,omitempty"`
	Celular       *string `xml:"celular,omitempty" json:"celular,omitempty" yaml:"celular,omitempty"`
	Ddd_celular   *string `xml:"ddd_celular,omitempty" json:"ddd_celular,omitempty" yaml:"ddd_celular,omitempty"`
	Identificacao *string `xml:"identificacao,omitempty" json:"identificacao,omitempty" yaml:"identificacao,omitempty"`
	Sms           *string `xml:"sms,omitempty" json:"sms,omitempty" yaml:"sms,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Remetente) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:remetente"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://cliente.bean.master.sigep.bsb.correios.com.br/"
	}
}

// RetornoCancelamento was auto-generated from WSDL.
type RetornoCancelamento struct {
	Cod_erro              *string             `xml:"cod_erro,omitempty" json:"cod_erro,omitempty" yaml:"cod_erro,omitempty"`
	Codigo_administrativo *string             `xml:"codigo_administrativo,omitempty" json:"codigo_administrativo,omitempty" yaml:"codigo_administrativo,omitempty"`
	Data                  *string             `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Hora                  *string             `xml:"hora,omitempty" json:"hora,omitempty" yaml:"hora,omitempty"`
	Msg_erro              *string             `xml:"msg_erro,omitempty" json:"msg_erro,omitempty" yaml:"msg_erro,omitempty"`
	Objeto_postal         *ObjetoSimplificado `xml:"objeto_postal,omitempty" json:"objeto_postal,omitempty" yaml:"objeto_postal,omitempty"`
}

// ServicoAdicionalERP was auto-generated from WSDL.
type ServicoAdicionalERP struct {
	Categoria        *string         `xml:"categoria,omitempty" json:"categoria,omitempty" yaml:"categoria,omitempty"`
	Codigo           *string         `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
	DataAtualizacao  *DateTime       `xml:"dataAtualizacao,omitempty" json:"dataAtualizacao,omitempty" yaml:"dataAtualizacao,omitempty"`
	DatajAtualizacao *int            `xml:"datajAtualizacao,omitempty" json:"datajAtualizacao,omitempty" yaml:"datajAtualizacao,omitempty"`
	Descricao        *string         `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	HorajAtualizacao *int            `xml:"horajAtualizacao,omitempty" json:"horajAtualizacao,omitempty" yaml:"horajAtualizacao,omitempty"`
	Id               *int            `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Sigla            *string         `xml:"sigla,omitempty" json:"sigla,omitempty" yaml:"sigla,omitempty"`
	ValorDeclarado   *ValorDeclarado `xml:"valorDeclarado,omitempty" json:"valorDeclarado,omitempty" yaml:"valorDeclarado,omitempty"`
}

// ServicoAdicionalTO was auto-generated from WSDL.
type ServicoAdicionalTO struct {
	Categoria *string `xml:"categoria,omitempty" json:"categoria,omitempty" yaml:"categoria,omitempty"`
	Codigo    *string `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
	Sigla     *string `xml:"sigla,omitempty" json:"sigla,omitempty" yaml:"sigla,omitempty"`
	Tipo      *string `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
}

// ServicoAdicionalXML was auto-generated from WSDL.
type ServicoAdicionalXML struct {
	Categoria            *string  `xml:"categoria,omitempty" json:"categoria,omitempty" yaml:"categoria,omitempty"`
	Codigo               *string  `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
	Descricao            *string  `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	MaximoValorDeclarado *float64 `xml:"maximoValorDeclarado,omitempty" json:"maximoValorDeclarado,omitempty" yaml:"maximoValorDeclarado,omitempty"`
	MinimoValorDeclarado *float64 `xml:"minimoValorDeclarado,omitempty" json:"minimoValorDeclarado,omitempty" yaml:"minimoValorDeclarado,omitempty"`
	Sigla                *string  `xml:"sigla,omitempty" json:"sigla,omitempty" yaml:"sigla,omitempty"`
}

// ServicoERP was auto-generated from WSDL.
type ServicoERP struct {
	Codigo             *string                `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
	DataAtualizacao    *DateTime              `xml:"dataAtualizacao,omitempty" json:"dataAtualizacao,omitempty" yaml:"dataAtualizacao,omitempty"`
	DatajAtualizacao   *int                   `xml:"datajAtualizacao,omitempty" json:"datajAtualizacao,omitempty" yaml:"datajAtualizacao,omitempty"`
	Descricao          *string                `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	HorajAtualizacao   *int                   `xml:"horajAtualizacao,omitempty" json:"horajAtualizacao,omitempty" yaml:"horajAtualizacao,omitempty"`
	Id                 *int64                 `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	ServicoSigep       *ServicoSigep          `xml:"servicoSigep,omitempty" json:"servicoSigep,omitempty" yaml:"servicoSigep,omitempty"`
	ServicosAdicionais []*ServicoAdicionalERP `xml:"servicosAdicionais,omitempty" json:"servicosAdicionais,omitempty" yaml:"servicosAdicionais,omitempty"`
	Tipo1Codigo        *string                `xml:"tipo1Codigo,omitempty" json:"tipo1Codigo,omitempty" yaml:"tipo1Codigo,omitempty"`
	Tipo1Descricao     *string                `xml:"tipo1Descricao,omitempty" json:"tipo1Descricao,omitempty" yaml:"tipo1Descricao,omitempty"`
	Tipo2Codigo        *string                `xml:"tipo2Codigo,omitempty" json:"tipo2Codigo,omitempty" yaml:"tipo2Codigo,omitempty"`
	Tipo2Descricao     *string                `xml:"tipo2Descricao,omitempty" json:"tipo2Descricao,omitempty" yaml:"tipo2Descricao,omitempty"`
	Vigencia           *VigenciaERP           `xml:"vigencia,omitempty" json:"vigencia,omitempty" yaml:"vigencia,omitempty"`
}

// ServicoSigep was auto-generated from WSDL.
type ServicoSigep struct {
	CategoriaServico  *CategoriaServico `xml:"categoriaServico,omitempty" json:"categoriaServico,omitempty" yaml:"categoriaServico,omitempty"`
	Chancela          *ChancelaMaster   `xml:"chancela,omitempty" json:"chancela,omitempty" yaml:"chancela,omitempty"`
	Descricao         *string           `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	ExigeDimensoes    *bool             `xml:"exigeDimensoes,omitempty" json:"exigeDimensoes,omitempty" yaml:"exigeDimensoes,omitempty"`
	ExigeValorCobrar  *bool             `xml:"exigeValorCobrar,omitempty" json:"exigeValorCobrar,omitempty" yaml:"exigeValorCobrar,omitempty"`
	Imitm             *int64            `xml:"imitm,omitempty" json:"imitm,omitempty" yaml:"imitm,omitempty"`
	PagamentoEntrega  *string           `xml:"pagamentoEntrega,omitempty" json:"pagamentoEntrega,omitempty" yaml:"pagamentoEntrega,omitempty"`
	RemessaAgrupada   *string           `xml:"remessaAgrupada,omitempty" json:"remessaAgrupada,omitempty" yaml:"remessaAgrupada,omitempty"`
	Restricao         *SimNao           `xml:"restricao,omitempty" json:"restricao,omitempty" yaml:"restricao,omitempty"`
	Servico           *int64            `xml:"servico,omitempty" json:"servico,omitempty" yaml:"servico,omitempty"`
	ServicoERP        *ServicoERP       `xml:"servicoERP,omitempty" json:"servicoERP,omitempty" yaml:"servicoERP,omitempty"`
	SsiCoCodigoPostal *string           `xml:"ssiCoCodigoPostal,omitempty" json:"ssiCoCodigoPostal,omitempty" yaml:"ssiCoCodigoPostal,omitempty"`
}

// SolicitaEtiquetas was auto-generated from WSDL.
type SolicitaEtiquetas struct {
	TipoDestinatario *string `xml:"tipoDestinatario,omitempty" json:"tipoDestinatario,omitempty" yaml:"tipoDestinatario,omitempty"`
	Identificador    *string `xml:"identificador,omitempty" json:"identificador,omitempty" yaml:"identificador,omitempty"`
	IdServico        *int64  `xml:"idServico,omitempty" json:"idServico,omitempty" yaml:"idServico,omitempty"`
	QtdEtiquetas     *int    `xml:"qtdEtiquetas,omitempty" json:"qtdEtiquetas,omitempty" yaml:"qtdEtiquetas,omitempty"`
	Usuario          *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha            *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// SolicitaEtiquetasResponse was auto-generated from WSDL.
type SolicitaEtiquetasResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// SolicitaPLP was auto-generated from WSDL.
type SolicitaPLP struct {
	IdPlpMaster *int64  `xml:"idPlpMaster,omitempty" json:"idPlpMaster,omitempty" yaml:"idPlpMaster,omitempty"`
	NumEtiqueta *string `xml:"numEtiqueta,omitempty" json:"numEtiqueta,omitempty" yaml:"numEtiqueta,omitempty"`
	Usuario     *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha       *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// SolicitaPLPResponse was auto-generated from WSDL.
type SolicitaPLPResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// SolicitaXmlPlp was auto-generated from WSDL.
type SolicitaXmlPlp struct {
	IdPlpMaster *int64  `xml:"idPlpMaster,omitempty" json:"idPlpMaster,omitempty" yaml:"idPlpMaster,omitempty"`
	Usuario     *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha       *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

type Plp struct {
	Text                *string `xml:",chardata"`
	IDPlp               *string `xml:"id_plp"`
	ValorGlobal         *string `xml:"valor_global"`
	McuUnidadePostagem  *string `xml:"mcu_unidade_postagem"`
	NomeUnidadePostagem *string `xml:"nome_unidade_postagem"`
	CartaoPostagem      *string `xml:"cartao_postagem"`
}

type Correioslog struct {
	Plp *Plp `xml:"plp"`
}

type Return struct {
	Correioslog *Correioslog `xml:"correioslog"`
}

// SolicitaXmlPlpResponse was auto-generated from WSDL.
type SolicitaXmlPlpResponse struct {
	*Return `xml:"return"`
}

func (s *SolicitaXmlPlpResponse) GetShippingPrice() (float64, error) {
	if s == nil || s.Correioslog == nil || s.Correioslog.Plp == nil || s.Correioslog.Plp.ValorGlobal == nil {
		return 0, fmt.Errorf("plp does not have a shipping price")
	}

	return strconv.ParseFloat(*s.Correioslog.Plp.ValorGlobal, 64)
}

// SolicitarPostagemScol was auto-generated from WSDL.
type SolicitarPostagemScol struct {
	CodAdministrativo *int    `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	Xml               *string `xml:"xml,omitempty" json:"xml,omitempty" yaml:"xml,omitempty"`
	Usuario           *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha             *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// SolicitarPostagemScolResponse was auto-generated from WSDL.
type SolicitarPostagemScolResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// UnidadePostagemERP was auto-generated from WSDL.
type UnidadePostagemERP struct {
	DiretoriaRegional *string      `xml:"diretoriaRegional,omitempty" json:"diretoriaRegional,omitempty" yaml:"diretoriaRegional,omitempty"`
	Endereco          *EnderecoERP `xml:"endereco,omitempty" json:"endereco,omitempty" yaml:"endereco,omitempty"`
	Id                *string      `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Nome              *string      `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	Status            *string      `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Tipo              *string      `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
}

// UsuarioInstalacao was auto-generated from WSDL.
type UsuarioInstalacao struct {
	DataAtualizacao *DateTime          `xml:"dataAtualizacao,omitempty" json:"dataAtualizacao,omitempty" yaml:"dataAtualizacao,omitempty"`
	DataInclusao    *DateTime          `xml:"dataInclusao,omitempty" json:"dataInclusao,omitempty" yaml:"dataInclusao,omitempty"`
	DataSenha       *DateTime          `xml:"dataSenha,omitempty" json:"dataSenha,omitempty" yaml:"dataSenha,omitempty"`
	GerenteMaster   *GerenteConta      `xml:"gerenteMaster,omitempty" json:"gerenteMaster,omitempty" yaml:"gerenteMaster,omitempty"`
	HashSenha       *string            `xml:"hashSenha,omitempty" json:"hashSenha,omitempty" yaml:"hashSenha,omitempty"`
	Id              *int64             `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Login           *string            `xml:"login,omitempty" json:"login,omitempty" yaml:"login,omitempty"`
	Nome            *string            `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	Parametros      []*ParametroMaster `xml:"parametros,omitempty" json:"parametros,omitempty" yaml:"parametros,omitempty"`
	Senha           *string            `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
	Status          *StatusUsuario     `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Validade        *string            `xml:"validade,omitempty" json:"validade,omitempty" yaml:"validade,omitempty"`
}

// ValePostal was auto-generated from WSDL.
type ValePostal struct {
	CidNoCidade            *string   `xml:"cidNoCidade,omitempty" json:"cidNoCidade,omitempty" yaml:"cidNoCidade,omitempty"`
	CtcCoAadministrativo   *string   `xml:"ctcCoAadministrativo,omitempty" json:"ctcCoAadministrativo,omitempty" yaml:"ctcCoAadministrativo,omitempty"`
	CtcNuContrato          *int64    `xml:"ctcNuContrato,omitempty" json:"ctcNuContrato,omitempty" yaml:"ctcNuContrato,omitempty"`
	CtcNuContratoEct       *int64    `xml:"ctcNuContratoEct,omitempty" json:"ctcNuContratoEct,omitempty" yaml:"ctcNuContratoEct,omitempty"`
	CvpEdBairro            *string   `xml:"cvpEdBairro,omitempty" json:"cvpEdBairro,omitempty" yaml:"cvpEdBairro,omitempty"`
	CvpEdCliente           *string   `xml:"cvpEdCliente,omitempty" json:"cvpEdCliente,omitempty" yaml:"cvpEdCliente,omitempty"`
	CvpEdComplemento       *string   `xml:"cvpEdComplemento,omitempty" json:"cvpEdComplemento,omitempty" yaml:"cvpEdComplemento,omitempty"`
	CvpEdNumero            *string   `xml:"cvpEdNumero,omitempty" json:"cvpEdNumero,omitempty" yaml:"cvpEdNumero,omitempty"`
	CvpNoCliente           *string   `xml:"cvpNoCliente,omitempty" json:"cvpNoCliente,omitempty" yaml:"cvpNoCliente,omitempty"`
	CvpNuCep               *int64    `xml:"cvpNuCep,omitempty" json:"cvpNuCep,omitempty" yaml:"cvpNuCep,omitempty"`
	DescricaoErro          *string   `xml:"descricaoErro,omitempty" json:"descricaoErro,omitempty" yaml:"descricaoErro,omitempty"`
	EstSgEstado            *string   `xml:"estSgEstado,omitempty" json:"estSgEstado,omitempty" yaml:"estSgEstado,omitempty"`
	MonVarTarifaAdicional  *int      `xml:"monVarTarifaAdicional,omitempty" json:"monVarTarifaAdicional,omitempty" yaml:"monVarTarifaAdicional,omitempty"`
	MonVarTarifaServico    *int      `xml:"monVarTarifaServico,omitempty" json:"monVarTarifaServico,omitempty" yaml:"monVarTarifaServico,omitempty"`
	MonVarValorDescontos   *int      `xml:"monVarValorDescontos,omitempty" json:"monVarValorDescontos,omitempty" yaml:"monVarValorDescontos,omitempty"`
	MonVarValorImposto     *int      `xml:"monVarValorImposto,omitempty" json:"monVarValorImposto,omitempty" yaml:"monVarValorImposto,omitempty"`
	PrsCoProdutoServico    *int64    `xml:"prsCoProdutoServico,omitempty" json:"prsCoProdutoServico,omitempty" yaml:"prsCoProdutoServico,omitempty"`
	PveNu                  *int64    `xml:"pveNu,omitempty" json:"pveNu,omitempty" yaml:"pveNu,omitempty"`
	PveOrgNuAgencia        *int64    `xml:"pveOrgNuAgencia,omitempty" json:"pveOrgNuAgencia,omitempty" yaml:"pveOrgNuAgencia,omitempty"`
	PveOrgNuAgenciaDes     *int64    `xml:"pveOrgNuAgenciaDes,omitempty" json:"pveOrgNuAgenciaDes,omitempty" yaml:"pveOrgNuAgenciaDes,omitempty"`
	PveOrgNuAgenciaOri     *int64    `xml:"pveOrgNuAgenciaOri,omitempty" json:"pveOrgNuAgenciaOri,omitempty" yaml:"pveOrgNuAgenciaOri,omitempty"`
	RetornaCodErro         *int      `xml:"retornaCodErro,omitempty" json:"retornaCodErro,omitempty" yaml:"retornaCodErro,omitempty"`
	SitNoSituacao          *string   `xml:"sitNoSituacao,omitempty" json:"sitNoSituacao,omitempty" yaml:"sitNoSituacao,omitempty"`
	TlgTxDescricao         *string   `xml:"tlgTxDescricao,omitempty" json:"tlgTxDescricao,omitempty" yaml:"tlgTxDescricao,omitempty"`
	VapDhTransacao         *DateTime `xml:"vapDhTransacao,omitempty" json:"vapDhTransacao,omitempty" yaml:"vapDhTransacao,omitempty"`
	VapNuEtiquetaEncomenda *string   `xml:"vapNuEtiquetaEncomenda,omitempty" json:"vapNuEtiquetaEncomenda,omitempty" yaml:"vapNuEtiquetaEncomenda,omitempty"`
	VapVrCobradoEct        *float64  `xml:"vapVrCobradoEct,omitempty" json:"vapVrCobradoEct,omitempty" yaml:"vapVrCobradoEct,omitempty"`
	VapVrNominal           *float64  `xml:"vapVrNominal,omitempty" json:"vapVrNominal,omitempty" yaml:"vapVrNominal,omitempty"`
}

// ValidaEtiquetaPLP was auto-generated from WSDL.
type ValidaEtiquetaPLP struct {
	NumeroEtiqueta *string `xml:"numeroEtiqueta,omitempty" json:"numeroEtiqueta,omitempty" yaml:"numeroEtiqueta,omitempty"`
	IdPlp          *int64  `xml:"idPlp,omitempty" json:"idPlp,omitempty" yaml:"idPlp,omitempty"`
	Usuario        *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha          *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// ValidaEtiquetaPLPResponse was auto-generated from WSDL.
type ValidaEtiquetaPLPResponse struct {
	Return *bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ValidaPlp was auto-generated from WSDL.
type ValidaPlp struct {
	Cliente            *int64    `xml:"cliente,omitempty" json:"cliente,omitempty" yaml:"cliente,omitempty"`
	Numero             *string   `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Diretoria          *int64    `xml:"diretoria,omitempty" json:"diretoria,omitempty" yaml:"diretoria,omitempty"`
	Cartao             *string   `xml:"cartao,omitempty" json:"cartao,omitempty" yaml:"cartao,omitempty"`
	UnidadePostagem    *string   `xml:"unidadePostagem,omitempty" json:"unidadePostagem,omitempty" yaml:"unidadePostagem,omitempty"`
	Servico            *int64    `xml:"servico,omitempty" json:"servico,omitempty" yaml:"servico,omitempty"`
	ServicosAdicionais []*string `xml:"servicosAdicionais,omitempty" json:"servicosAdicionais,omitempty" yaml:"servicosAdicionais,omitempty"`
	Usuario            *string   `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha              *string   `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// ValidaPlpResponse was auto-generated from WSDL.
type ValidaPlpResponse struct {
	Return *bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ValidarPostagemReversa was auto-generated from WSDL.
type ValidarPostagemReversa struct {
	CodAdministrativo *string        `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	CodigoServico     *string        `xml:"codigoServico,omitempty" json:"codigoServico,omitempty" yaml:"codigoServico,omitempty"`
	CepDestinatario   *string        `xml:"cepDestinatario,omitempty" json:"cepDestinatario,omitempty" yaml:"cepDestinatario,omitempty"`
	IdCartaoPostagem  *string        `xml:"idCartaoPostagem,omitempty" json:"idCartaoPostagem,omitempty" yaml:"idCartaoPostagem,omitempty"`
	Coleta            *ColetaReversa `xml:"coleta,omitempty" json:"coleta,omitempty" yaml:"coleta,omitempty"`
	Usuario           *string        `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha             *string        `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// ValidarPostagemReversaResponse was auto-generated from WSDL.
type ValidarPostagemReversaResponse struct {
	Return *bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ValidarPostagemSimultanea was auto-generated from WSDL.
type ValidarPostagemSimultanea struct {
	CodAdministrativo *int              `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	CodigoServico     *int              `xml:"codigoServico,omitempty" json:"codigoServico,omitempty" yaml:"codigoServico,omitempty"`
	IdCartaoPostagem  *string           `xml:"idCartaoPostagem,omitempty" json:"idCartaoPostagem,omitempty" yaml:"idCartaoPostagem,omitempty"`
	CepDestinatario   *string           `xml:"cepDestinatario,omitempty" json:"cepDestinatario,omitempty" yaml:"cepDestinatario,omitempty"`
	Coleta            *ColetaSimultanea `xml:"coleta,omitempty" json:"coleta,omitempty" yaml:"coleta,omitempty"`
	Usuario           *string           `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha             *string           `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// ValidarPostagemSimultaneaResponse was auto-generated from WSDL.
type ValidarPostagemSimultaneaResponse struct {
	Return *bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ValorDeclarado was auto-generated from WSDL.
type ValorDeclarado struct {
	Maximo *float64 `xml:"maximo,omitempty" json:"maximo,omitempty" yaml:"maximo,omitempty"`
	Minimo *float64 `xml:"minimo,omitempty" json:"minimo,omitempty" yaml:"minimo,omitempty"`
}

// VerificaDisponibilidadeServico was auto-generated from WSDL.
type VerificaDisponibilidadeServico struct {
	CodAdministrativo *int    `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	NumeroServico     *string `xml:"numeroServico,omitempty" json:"numeroServico,omitempty" yaml:"numeroServico,omitempty"`
	CepOrigem         *string `xml:"cepOrigem,omitempty" json:"cepOrigem,omitempty" yaml:"cepOrigem,omitempty"`
	CepDestino        *string `xml:"cepDestino,omitempty" json:"cepDestino,omitempty" yaml:"cepDestino,omitempty"`
	Usuario           *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha             *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// VerificaDisponibilidadeServicoResponse was auto-generated from
// WSDL.
type VerificaDisponibilidadeServicoResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// VerificaModalTransporte was auto-generated from WSDL.
type VerificaModalTransporte struct {
	CodigoServico *string `xml:"codigoServico,omitempty" json:"codigoServico,omitempty" yaml:"codigoServico,omitempty"`
	CepOrigem     *string `xml:"cepOrigem,omitempty" json:"cepOrigem,omitempty" yaml:"cepOrigem,omitempty"`
	CepDestino    *string `xml:"cepDestino,omitempty" json:"cepDestino,omitempty" yaml:"cepDestino,omitempty"`
	Usuario       *string `xml:"usuario,omitempty" json:"usuario,omitempty" yaml:"usuario,omitempty"`
	Senha         *string `xml:"senha,omitempty" json:"senha,omitempty" yaml:"senha,omitempty"`
}

// VerificaModalTransporteResponse was auto-generated from WSDL.
type VerificaModalTransporteResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// VigenciaERP was auto-generated from WSDL.
type VigenciaERP struct {
	DataFinal   *DateTime `xml:"dataFinal,omitempty" json:"dataFinal,omitempty" yaml:"dataFinal,omitempty"`
	DataInicial *DateTime `xml:"dataInicial,omitempty" json:"dataInicial,omitempty" yaml:"dataInicial,omitempty"`
	DatajFim    *int      `xml:"datajFim,omitempty" json:"datajFim,omitempty" yaml:"datajFim,omitempty"`
	DatajIni    *int      `xml:"datajIni,omitempty" json:"datajIni,omitempty" yaml:"datajIni,omitempty"`
	Id          *int64    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
}

// Operation wrapper for VerificaSeTodosObjetosCancelados.
// OperationVerificaSeTodosObjetosCancelados was auto-generated
// from WSDL.
type OperationVerificaSeTodosObjetosCancelados struct {
	VerificaSeTodosObjetosCancelados *VerificaSeTodosObjetosCancelados `xml:"VerificaSeTodosObjetosCancelados,omitempty" json:"VerificaSeTodosObjetosCancelados,omitempty" yaml:"VerificaSeTodosObjetosCancelados,omitempty"`
}

// Operation wrapper for VerificaSeTodosObjetosCancelados.
// OperationVerificaSeTodosObjetosCanceladosResponse was auto-generated
// from WSDL.
type OperationVerificaSeTodosObjetosCanceladosResponse struct {
	VerificaSeTodosObjetosCanceladosResponse *VerificaSeTodosObjetosCanceladosResponse `xml:"VerificaSeTodosObjetosCanceladosResponse,omitempty" json:"VerificaSeTodosObjetosCanceladosResponse,omitempty" yaml:"VerificaSeTodosObjetosCanceladosResponse,omitempty"`
}

// Operation wrapper for AtualizaPagamentoNaEntrega.
// OperationAtualizaPagamentoNaEntrega was auto-generated from
// WSDL.
type OperationAtualizaPagamentoNaEntrega struct {
	AtualizaPagamentoNaEntrega *AtualizaPagamentoNaEntrega `xml:"atualizaPagamentoNaEntrega,omitempty" json:"atualizaPagamentoNaEntrega,omitempty" yaml:"atualizaPagamentoNaEntrega,omitempty"`
}

// Operation wrapper for AtualizaPagamentoNaEntrega.
// OperationAtualizaPagamentoNaEntregaResponse was auto-generated
// from WSDL.
type OperationAtualizaPagamentoNaEntregaResponse struct {
	AtualizaPagamentoNaEntregaResponse *AtualizaPagamentoNaEntregaResponse `xml:"atualizaPagamentoNaEntregaResponse,omitempty" json:"atualizaPagamentoNaEntregaResponse,omitempty" yaml:"atualizaPagamentoNaEntregaResponse,omitempty"`
}

// Operation wrapper for AtualizaRemessaAgrupada.
// OperationAtualizaRemessaAgrupada was auto-generated from WSDL.
type OperationAtualizaRemessaAgrupada struct {
	AtualizaRemessaAgrupada *AtualizaRemessaAgrupada `xml:"atualizaRemessaAgrupada,omitempty" json:"atualizaRemessaAgrupada,omitempty" yaml:"atualizaRemessaAgrupada,omitempty"`
}

// Operation wrapper for AtualizaRemessaAgrupada.
// OperationAtualizaRemessaAgrupadaResponse was auto-generated
// from WSDL.
type OperationAtualizaRemessaAgrupadaResponse struct {
	AtualizaRemessaAgrupadaResponse *AtualizaRemessaAgrupadaResponse `xml:"atualizaRemessaAgrupadaResponse,omitempty" json:"atualizaRemessaAgrupadaResponse,omitempty" yaml:"atualizaRemessaAgrupadaResponse,omitempty"`
}

// Operation wrapper for BloquearObjeto.
// OperationBloquearObjeto was auto-generated from WSDL.
type OperationBloquearObjeto struct {
	BloquearObjeto *BloquearObjeto `xml:"bloquearObjeto,omitempty" json:"bloquearObjeto,omitempty" yaml:"bloquearObjeto,omitempty"`
}

// Operation wrapper for BloquearObjeto.
// OperationBloquearObjetoResponse was auto-generated from WSDL.
type OperationBloquearObjetoResponse struct {
	BloquearObjetoResponse *BloquearObjetoResponse `xml:"bloquearObjetoResponse,omitempty" json:"bloquearObjetoResponse,omitempty" yaml:"bloquearObjetoResponse,omitempty"`
}

// Operation wrapper for BuscaCliente.
// OperationBuscaCliente was auto-generated from WSDL.
type OperationBuscaCliente struct {
	BuscaCliente *BuscaCliente `xml:"buscaCliente,omitempty" json:"buscaCliente,omitempty" yaml:"buscaCliente,omitempty"`
}

// Operation wrapper for BuscaCliente.
// OperationBuscaClienteResponse was auto-generated from WSDL.
type OperationBuscaClienteResponse struct {
	BuscaClienteResponse *BuscaClienteResponse `xml:"buscaClienteResponse,omitempty" json:"buscaClienteResponse,omitempty" yaml:"buscaClienteResponse,omitempty"`
}

// Operation wrapper for BuscaContrato.
// OperationBuscaContrato was auto-generated from WSDL.
type OperationBuscaContrato struct {
	BuscaContrato *BuscaContrato `xml:"buscaContrato,omitempty" json:"buscaContrato,omitempty" yaml:"buscaContrato,omitempty"`
}

// Operation wrapper for BuscaContrato.
// OperationBuscaContratoResponse was auto-generated from WSDL.
type OperationBuscaContratoResponse struct {
	BuscaContratoResponse *BuscaContratoResponse `xml:"buscaContratoResponse,omitempty" json:"buscaContratoResponse,omitempty" yaml:"buscaContratoResponse,omitempty"`
}

// Operation wrapper for BuscaDataAtual.
// OperationBuscaDataAtual was auto-generated from WSDL.
type OperationBuscaDataAtual struct {
	BuscaDataAtual *BuscaDataAtual `xml:"buscaDataAtual,omitempty" json:"buscaDataAtual,omitempty" yaml:"buscaDataAtual,omitempty"`
}

// Operation wrapper for BuscaDataAtual.
// OperationBuscaDataAtualResponse was auto-generated from WSDL.
type OperationBuscaDataAtualResponse struct {
	BuscaDataAtualResponse *BuscaDataAtualResponse `xml:"buscaDataAtualResponse,omitempty" json:"buscaDataAtualResponse,omitempty" yaml:"buscaDataAtualResponse,omitempty"`
}

// Operation wrapper for BuscaOpcoes.
// OperationBuscaOpcoes was auto-generated from WSDL.
type OperationBuscaOpcoes struct {
	BuscaOpcoes *BuscaOpcoes `xml:"buscaOpcoes,omitempty" json:"buscaOpcoes,omitempty" yaml:"buscaOpcoes,omitempty"`
}

// Operation wrapper for BuscaOpcoes.
// OperationBuscaOpcoesResponse was auto-generated from WSDL.
type OperationBuscaOpcoesResponse struct {
	BuscaOpcoesResponse *BuscaOpcoesResponse `xml:"buscaOpcoesResponse,omitempty" json:"buscaOpcoesResponse,omitempty" yaml:"buscaOpcoesResponse,omitempty"`
}

// Operation wrapper for BuscaPagamentoEntrega.
// OperationBuscaPagamentoEntrega was auto-generated from WSDL.
type OperationBuscaPagamentoEntrega struct {
	BuscaPagamentoEntrega *BuscaPagamentoEntrega `xml:"buscaPagamentoEntrega,omitempty" json:"buscaPagamentoEntrega,omitempty" yaml:"buscaPagamentoEntrega,omitempty"`
}

// Operation wrapper for BuscaPagamentoEntrega.
// OperationBuscaPagamentoEntregaResponse was auto-generated from
// WSDL.
type OperationBuscaPagamentoEntregaResponse struct {
	BuscaPagamentoEntregaResponse *BuscaPagamentoEntregaResponse `xml:"buscaPagamentoEntregaResponse,omitempty" json:"buscaPagamentoEntregaResponse,omitempty" yaml:"buscaPagamentoEntregaResponse,omitempty"`
}

// Operation wrapper for BuscaServicos.
// OperationBuscaServicos was auto-generated from WSDL.
type OperationBuscaServicos struct {
	BuscaServicos *BuscaServicos `xml:"buscaServicos,omitempty" json:"buscaServicos,omitempty" yaml:"buscaServicos,omitempty"`
}

// Operation wrapper for BuscaServicos.
// OperationBuscaServicosResponse was auto-generated from WSDL.
type OperationBuscaServicosResponse struct {
	BuscaServicosResponse *BuscaServicosResponse `xml:"buscaServicosResponse,omitempty" json:"buscaServicosResponse,omitempty" yaml:"buscaServicosResponse,omitempty"`
}

// Operation wrapper for BuscaServicosAdicionaisAtivos.
// OperationBuscaServicosAdicionaisAtivos was auto-generated from
// WSDL.
type OperationBuscaServicosAdicionaisAtivos struct {
	BuscaServicosAdicionaisAtivos *BuscaServicosAdicionaisAtivos `xml:"buscaServicosAdicionaisAtivos,omitempty" json:"buscaServicosAdicionaisAtivos,omitempty" yaml:"buscaServicosAdicionaisAtivos,omitempty"`
}

// Operation wrapper for BuscaServicosAdicionaisAtivos.
// OperationBuscaServicosAdicionaisAtivosResponse was auto-generated
// from WSDL.
type OperationBuscaServicosAdicionaisAtivosResponse struct {
	BuscaServicosAdicionaisAtivosResponse *BuscaServicosAdicionaisAtivosResponse `xml:"buscaServicosAdicionaisAtivosResponse,omitempty" json:"buscaServicosAdicionaisAtivosResponse,omitempty" yaml:"buscaServicosAdicionaisAtivosResponse,omitempty"`
}

// Operation wrapper for BuscaServicosValorDeclarado.
// OperationBuscaServicosValorDeclarado was auto-generated from
// WSDL.
type OperationBuscaServicosValorDeclarado struct {
	BuscaServicosValorDeclarado *BuscaServicosValorDeclarado `xml:"buscaServicosValorDeclarado,omitempty" json:"buscaServicosValorDeclarado,omitempty" yaml:"buscaServicosValorDeclarado,omitempty"`
}

// Operation wrapper for BuscaServicosValorDeclarado.
// OperationBuscaServicosValorDeclaradoResponse was auto-generated
// from WSDL.
type OperationBuscaServicosValorDeclaradoResponse struct {
	BuscaServicosValorDeclaradoResponse *BuscaServicosValorDeclaradoResponse `xml:"buscaServicosValorDeclaradoResponse,omitempty" json:"buscaServicosValorDeclaradoResponse,omitempty" yaml:"buscaServicosValorDeclaradoResponse,omitempty"`
}

// Operation wrapper for BuscaServicosXServicosAdicionais.
// OperationBuscaServicosXServicosAdicionais was auto-generated
// from WSDL.
type OperationBuscaServicosXServicosAdicionais struct {
	BuscaServicosXServicosAdicionais *BuscaServicosXServicosAdicionais `xml:"buscaServicosXServicosAdicionais,omitempty" json:"buscaServicosXServicosAdicionais,omitempty" yaml:"buscaServicosXServicosAdicionais,omitempty"`
}

// Operation wrapper for BuscaServicosXServicosAdicionais.
// OperationBuscaServicosXServicosAdicionaisResponse was auto-generated
// from WSDL.
type OperationBuscaServicosXServicosAdicionaisResponse struct {
	BuscaServicosXServicosAdicionaisResponse *BuscaServicosXServicosAdicionaisResponse `xml:"buscaServicosXServicosAdicionaisResponse,omitempty" json:"buscaServicosXServicosAdicionaisResponse,omitempty" yaml:"buscaServicosXServicosAdicionaisResponse,omitempty"`
}

// Operation wrapper for BuscaTarifaVale.
// OperationBuscaTarifaVale was auto-generated from WSDL.
type OperationBuscaTarifaVale struct {
	BuscaTarifaVale *BuscaTarifaVale `xml:"buscaTarifaVale,omitempty" json:"buscaTarifaVale,omitempty" yaml:"buscaTarifaVale,omitempty"`
}

// Operation wrapper for BuscaTarifaVale.
// OperationBuscaTarifaValeResponse was auto-generated from WSDL.
type OperationBuscaTarifaValeResponse struct {
	BuscaTarifaValeResponse *BuscaTarifaValeResponse `xml:"buscaTarifaValeResponse,omitempty" json:"buscaTarifaValeResponse,omitempty" yaml:"buscaTarifaValeResponse,omitempty"`
}

// Operation wrapper for CalculaTarifaServico.
// OperationCalculaTarifaServico was auto-generated from WSDL.
type OperationCalculaTarifaServico struct {
	CalculaTarifaServico *CalculaTarifaServico `xml:"calculaTarifaServico,omitempty" json:"calculaTarifaServico,omitempty" yaml:"calculaTarifaServico,omitempty"`
}

// Operation wrapper for CalculaTarifaServico.
// OperationCalculaTarifaServicoResponse was auto-generated from
// WSDL.
type OperationCalculaTarifaServicoResponse struct {
	CalculaTarifaServicoResponse *CalculaTarifaServicoResponse `xml:"calculaTarifaServicoResponse,omitempty" json:"calculaTarifaServicoResponse,omitempty" yaml:"calculaTarifaServicoResponse,omitempty"`
}

// Operation wrapper for CancelarObjeto.
// OperationCancelarObjeto was auto-generated from WSDL.
type OperationCancelarObjeto struct {
	CancelarObjeto *CancelarObjeto `xml:"cancelarObjeto,omitempty" json:"cancelarObjeto,omitempty" yaml:"cancelarObjeto,omitempty"`
}

// Operation wrapper for CancelarObjeto.
// OperationCancelarObjetoResponse was auto-generated from WSDL.
type OperationCancelarObjetoResponse struct {
	CancelarObjetoResponse *CancelarObjetoResponse `xml:"cancelarObjetoResponse,omitempty" json:"cancelarObjetoResponse,omitempty" yaml:"cancelarObjetoResponse,omitempty"`
}

// Operation wrapper for CancelarPedidoScol.
// OperationCancelarPedidoScol was auto-generated from WSDL.
type OperationCancelarPedidoScol struct {
	CancelarPedidoScol *CancelarPedidoScol `xml:"cancelarPedidoScol,omitempty" json:"cancelarPedidoScol,omitempty" yaml:"cancelarPedidoScol,omitempty"`
}

// Operation wrapper for CancelarPedidoScol.
// OperationCancelarPedidoScolResponse was auto-generated from
// WSDL.
type OperationCancelarPedidoScolResponse struct {
	CancelarPedidoScolResponse *CancelarPedidoScolResponse `xml:"cancelarPedidoScolResponse,omitempty" json:"cancelarPedidoScolResponse,omitempty" yaml:"cancelarPedidoScolResponse,omitempty"`
}

// Operation wrapper for ConsultaCEP.
// OperationConsultaCEP was auto-generated from WSDL.
type OperationConsultaCEP struct {
	ConsultaCEP *ConsultaCEP `xml:"consultaCEP,omitempty" json:"consultaCEP,omitempty" yaml:"consultaCEP,omitempty"`
}

// Operation wrapper for ConsultaCEP.
// OperationConsultaCEPResponse was auto-generated from WSDL.
type OperationConsultaCEPResponse struct {
	ConsultaCEPResponse *ConsultaCEPResponse `xml:"consultaCEPResponse,omitempty" json:"consultaCEPResponse,omitempty" yaml:"consultaCEPResponse,omitempty"`
}

// Operation wrapper for FechaPlp.
// OperationFechaPlp was auto-generated from WSDL.
type OperationFechaPlp struct {
	FechaPlp *FechaPlp `xml:"fechaPlp,omitempty" json:"fechaPlp,omitempty" yaml:"fechaPlp,omitempty"`
}

// Operation wrapper for FechaPlp.
// OperationFechaPlpResponse was auto-generated from WSDL.
type OperationFechaPlpResponse struct {
	FechaPlpResponse *FechaPlpResponse `xml:"fechaPlpResponse,omitempty" json:"fechaPlpResponse,omitempty" yaml:"fechaPlpResponse,omitempty"`
}

// Operation wrapper for FechaPlpVariosServicos.
// OperationFechaPlpVariosServicos was auto-generated from WSDL.
type OperationFechaPlpVariosServicos struct {
	FechaPlpVariosServicos *FechaPlpVariosServicos `xml:"fechaPlpVariosServicos,omitempty" json:"fechaPlpVariosServicos,omitempty" yaml:"fechaPlpVariosServicos,omitempty"`
}

// Operation wrapper for FechaPlpVariosServicos.
// OperationFechaPlpVariosServicosResponse was auto-generated from
// WSDL.
type OperationFechaPlpVariosServicosResponse struct {
	FechaPlpVariosServicosResponse *FechaPlpVariosServicosResponse `xml:"fechaPlpVariosServicosResponse,omitempty" json:"fechaPlpVariosServicosResponse,omitempty" yaml:"fechaPlpVariosServicosResponse,omitempty"`
}

// Operation wrapper for GeraDigitoVerificadorEtiquetas.
// OperationGeraDigitoVerificadorEtiquetas was auto-generated from
// WSDL.
type OperationGeraDigitoVerificadorEtiquetas struct {
	GeraDigitoVerificadorEtiquetas *GeraDigitoVerificadorEtiquetas `xml:"geraDigitoVerificadorEtiquetas,omitempty" json:"geraDigitoVerificadorEtiquetas,omitempty" yaml:"geraDigitoVerificadorEtiquetas,omitempty"`
}

// Operation wrapper for GeraDigitoVerificadorEtiquetas.
// OperationGeraDigitoVerificadorEtiquetasResponse was auto-generated
// from WSDL.
type OperationGeraDigitoVerificadorEtiquetasResponse struct {
	GeraDigitoVerificadorEtiquetasResponse *GeraDigitoVerificadorEtiquetasResponse `xml:"geraDigitoVerificadorEtiquetasResponse,omitempty" json:"geraDigitoVerificadorEtiquetasResponse,omitempty" yaml:"geraDigitoVerificadorEtiquetasResponse,omitempty"`
}

// Operation wrapper for GetStatusCartaoPostagem.
// OperationGetStatusCartaoPostagem was auto-generated from WSDL.
type OperationGetStatusCartaoPostagem struct {
	GetStatusCartaoPostagem *GetStatusCartaoPostagem `xml:"getStatusCartaoPostagem,omitempty" json:"getStatusCartaoPostagem,omitempty" yaml:"getStatusCartaoPostagem,omitempty"`
}

// Operation wrapper for GetStatusCartaoPostagem.
// OperationGetStatusCartaoPostagemResponse was auto-generated
// from WSDL.
type OperationGetStatusCartaoPostagemResponse struct {
	GetStatusCartaoPostagemResponse *GetStatusCartaoPostagemResponse `xml:"getStatusCartaoPostagemResponse,omitempty" json:"getStatusCartaoPostagemResponse,omitempty" yaml:"getStatusCartaoPostagemResponse,omitempty"`
}

// Operation wrapper for GetStatusPLP.
// OperationGetStatusPLP was auto-generated from WSDL.
type OperationGetStatusPLP struct {
	GetStatusPLP *GetStatusPLP `xml:"getStatusPLP,omitempty" json:"getStatusPLP,omitempty" yaml:"getStatusPLP,omitempty"`
}

// Operation wrapper for GetStatusPLP.
// OperationGetStatusPLPResponse was auto-generated from WSDL.
type OperationGetStatusPLPResponse struct {
	GetStatusPLPResponse *GetStatusPLPResponse `xml:"getStatusPLPResponse,omitempty" json:"getStatusPLPResponse,omitempty" yaml:"getStatusPLPResponse,omitempty"`
}

// Operation wrapper for IntegrarUsuarioScol.
// OperationIntegrarUsuarioScol was auto-generated from WSDL.
type OperationIntegrarUsuarioScol struct {
	IntegrarUsuarioScol *IntegrarUsuarioScol `xml:"integrarUsuarioScol,omitempty" json:"integrarUsuarioScol,omitempty" yaml:"integrarUsuarioScol,omitempty"`
}

// Operation wrapper for IntegrarUsuarioScol.
// OperationIntegrarUsuarioScolResponse was auto-generated from
// WSDL.
type OperationIntegrarUsuarioScolResponse struct {
	IntegrarUsuarioScolResponse *IntegrarUsuarioScolResponse `xml:"integrarUsuarioScolResponse,omitempty" json:"integrarUsuarioScolResponse,omitempty" yaml:"integrarUsuarioScolResponse,omitempty"`
}

// Operation wrapper for ObterClienteAtualizacao.
// OperationObterClienteAtualizacao was auto-generated from WSDL.
type OperationObterClienteAtualizacao struct {
	ObterClienteAtualizacao *ObterClienteAtualizacao `xml:"obterClienteAtualizacao,omitempty" json:"obterClienteAtualizacao,omitempty" yaml:"obterClienteAtualizacao,omitempty"`
}

// Operation wrapper for ObterClienteAtualizacao.
// OperationObterClienteAtualizacaoResponse was auto-generated
// from WSDL.
type OperationObterClienteAtualizacaoResponse struct {
	ObterClienteAtualizacaoResponse *ObterClienteAtualizacaoResponse `xml:"obterClienteAtualizacaoResponse,omitempty" json:"obterClienteAtualizacaoResponse,omitempty" yaml:"obterClienteAtualizacaoResponse,omitempty"`
}

// Operation wrapper for ObterEmbalagemLRS.
// OperationObterEmbalagemLRS was auto-generated from WSDL.
type OperationObterEmbalagemLRS struct {
	ObterEmbalagemLRS *ObterEmbalagemLRS `xml:"obterEmbalagemLRS,omitempty" json:"obterEmbalagemLRS,omitempty" yaml:"obterEmbalagemLRS,omitempty"`
}

// Operation wrapper for ObterEmbalagemLRS.
// OperationObterEmbalagemLRSResponse was auto-generated from WSDL.
type OperationObterEmbalagemLRSResponse struct {
	ObterEmbalagemLRSResponse *ObterEmbalagemLRSResponse `xml:"obterEmbalagemLRSResponse,omitempty" json:"obterEmbalagemLRSResponse,omitempty" yaml:"obterEmbalagemLRSResponse,omitempty"`
}

// Operation wrapper for ObterMensagemParametrizada.
// OperationObterMensagemParametrizada was auto-generated from
// WSDL.
type OperationObterMensagemParametrizada struct {
	ObterMensagemParametrizada *ObterMensagemParametrizada `xml:"obterMensagemParametrizada,omitempty" json:"obterMensagemParametrizada,omitempty" yaml:"obterMensagemParametrizada,omitempty"`
}

// Operation wrapper for ObterMensagemParametrizada.
// OperationObterMensagemParametrizadaResponse was auto-generated
// from WSDL.
type OperationObterMensagemParametrizadaResponse struct {
	ObterMensagemParametrizadaResponse *ObterMensagemParametrizadaResponse `xml:"obterMensagemParametrizadaResponse,omitempty" json:"obterMensagemParametrizadaResponse,omitempty" yaml:"obterMensagemParametrizadaResponse,omitempty"`
}

// Operation wrapper for PesquisarDimensoesServico.
// OperationPesquisarDimensoesServico was auto-generated from WSDL.
type OperationPesquisarDimensoesServico struct {
	PesquisarDimensoesServico *PesquisarDimensoesServico `xml:"pesquisarDimensoesServico,omitempty" json:"pesquisarDimensoesServico,omitempty" yaml:"pesquisarDimensoesServico,omitempty"`
}

// Operation wrapper for PesquisarDimensoesServico.
// OperationPesquisarDimensoesServicoResponse was auto-generated
// from WSDL.
type OperationPesquisarDimensoesServicoResponse struct {
	PesquisarDimensoesServicoResponse *PesquisarDimensoesServicoResponse `xml:"pesquisarDimensoesServicoResponse,omitempty" json:"pesquisarDimensoesServicoResponse,omitempty" yaml:"pesquisarDimensoesServicoResponse,omitempty"`
}

// Operation wrapper for PesquisarEmbalagensPorServico.
// OperationPesquisarEmbalagensPorServico was auto-generated from
// WSDL.
type OperationPesquisarEmbalagensPorServico struct {
	PesquisarEmbalagensPorServico *PesquisarEmbalagensPorServico `xml:"pesquisarEmbalagensPorServico,omitempty" json:"pesquisarEmbalagensPorServico,omitempty" yaml:"pesquisarEmbalagensPorServico,omitempty"`
}

// Operation wrapper for PesquisarEmbalagensPorServico.
// OperationPesquisarEmbalagensPorServicoResponse was auto-generated
// from WSDL.
type OperationPesquisarEmbalagensPorServicoResponse struct {
	PesquisarEmbalagensPorServicoResponse *PesquisarEmbalagensPorServicoResponse `xml:"pesquisarEmbalagensPorServicoResponse,omitempty" json:"pesquisarEmbalagensPorServicoResponse,omitempty" yaml:"pesquisarEmbalagensPorServicoResponse,omitempty"`
}

// Operation wrapper for PesquisarParametrosPorDescricao.
// OperationPesquisarParametrosPorDescricao was auto-generated
// from WSDL.
type OperationPesquisarParametrosPorDescricao struct {
	PesquisarParametrosPorDescricao *PesquisarParametrosPorDescricao `xml:"pesquisarParametrosPorDescricao,omitempty" json:"pesquisarParametrosPorDescricao,omitempty" yaml:"pesquisarParametrosPorDescricao,omitempty"`
}

// Operation wrapper for PesquisarParametrosPorDescricao.
// OperationPesquisarParametrosPorDescricaoResponse was auto-generated
// from WSDL.
type OperationPesquisarParametrosPorDescricaoResponse struct {
	PesquisarParametrosPorDescricaoResponse *PesquisarParametrosPorDescricaoResponse `xml:"pesquisarParametrosPorDescricaoResponse,omitempty" json:"pesquisarParametrosPorDescricaoResponse,omitempty" yaml:"pesquisarParametrosPorDescricaoResponse,omitempty"`
}

// Operation wrapper for PesquisarServicosAdicionais.
// OperationPesquisarServicosAdicionais was auto-generated from
// WSDL.
type OperationPesquisarServicosAdicionais struct {
	PesquisarServicosAdicionais *PesquisarServicosAdicionais `xml:"pesquisarServicosAdicionais,omitempty" json:"pesquisarServicosAdicionais,omitempty" yaml:"pesquisarServicosAdicionais,omitempty"`
}

// Operation wrapper for PesquisarServicosAdicionais.
// OperationPesquisarServicosAdicionaisResponse was auto-generated
// from WSDL.
type OperationPesquisarServicosAdicionaisResponse struct {
	PesquisarServicosAdicionaisResponse *PesquisarServicosAdicionaisResponse `xml:"pesquisarServicosAdicionaisResponse,omitempty" json:"pesquisarServicosAdicionaisResponse,omitempty" yaml:"pesquisarServicosAdicionaisResponse,omitempty"`
}

// Operation wrapper for SolicitaEtiquetas.
// OperationSolicitaEtiquetas was auto-generated from WSDL.
type OperationSolicitaEtiquetas struct {
	SolicitaEtiquetas *SolicitaEtiquetas `xml:"tns:solicitaEtiquetas,omitempty" json:"solicitaEtiquetas,omitempty" yaml:"solicitaEtiquetas,omitempty"`
}

// Operation wrapper for SolicitaEtiquetas.
// OperationSolicitaEtiquetasResponse was auto-generated from WSDL.
type OperationSolicitaEtiquetasResponse struct {
	SolicitaEtiquetasResponse *SolicitaEtiquetasResponse `xml:"solicitaEtiquetasResponse,omitempty" json:"solicitaEtiquetasResponse,omitempty" yaml:"solicitaEtiquetasResponse,omitempty"`
}

// Operation wrapper for SolicitaPLP.
// OperationSolicitaPLP was auto-generated from WSDL.
type OperationSolicitaPLP struct {
	SolicitaPLP *SolicitaPLP `xml:"solicitaPLP,omitempty" json:"solicitaPLP,omitempty" yaml:"solicitaPLP,omitempty"`
}

// Operation wrapper for SolicitaPLP.
// OperationSolicitaPLPResponse was auto-generated from WSDL.
type OperationSolicitaPLPResponse struct {
	SolicitaPLPResponse *SolicitaPLPResponse `xml:"solicitaPLPResponse,omitempty" json:"solicitaPLPResponse,omitempty" yaml:"solicitaPLPResponse,omitempty"`
}

// Operation wrapper for SolicitaXmlPlp.
// OperationSolicitaXmlPlp was auto-generated from WSDL.
type OperationSolicitaXmlPlp struct {
	SolicitaXmlPlp *SolicitaXmlPlp `xml:"solicitaXmlPlp,omitempty" json:"solicitaXmlPlp,omitempty" yaml:"solicitaXmlPlp,omitempty"`
}

// Operation wrapper for SolicitaXmlPlp.
// OperationSolicitaXmlPlpResponse was auto-generated from WSDL.
type OperationSolicitaXmlPlpResponse struct {
	SolicitaXmlPlpResponse *SolicitaXmlPlpResponse `xml:"solicitaXmlPlpResponse,omitempty" json:"solicitaXmlPlpResponse,omitempty" yaml:"solicitaXmlPlpResponse,omitempty"`
}

// Operation wrapper for SolicitarPostagemScol.
// OperationSolicitarPostagemScol was auto-generated from WSDL.
type OperationSolicitarPostagemScol struct {
	SolicitarPostagemScol *SolicitarPostagemScol `xml:"solicitarPostagemScol,omitempty" json:"solicitarPostagemScol,omitempty" yaml:"solicitarPostagemScol,omitempty"`
}

// Operation wrapper for SolicitarPostagemScol.
// OperationSolicitarPostagemScolResponse was auto-generated from
// WSDL.
type OperationSolicitarPostagemScolResponse struct {
	SolicitarPostagemScolResponse *SolicitarPostagemScolResponse `xml:"solicitarPostagemScolResponse,omitempty" json:"solicitarPostagemScolResponse,omitempty" yaml:"solicitarPostagemScolResponse,omitempty"`
}

// Operation wrapper for ValidaEtiquetaPLP.
// OperationValidaEtiquetaPLP was auto-generated from WSDL.
type OperationValidaEtiquetaPLP struct {
	ValidaEtiquetaPLP *ValidaEtiquetaPLP `xml:"validaEtiquetaPLP,omitempty" json:"validaEtiquetaPLP,omitempty" yaml:"validaEtiquetaPLP,omitempty"`
}

// Operation wrapper for ValidaEtiquetaPLP.
// OperationValidaEtiquetaPLPResponse was auto-generated from WSDL.
type OperationValidaEtiquetaPLPResponse struct {
	ValidaEtiquetaPLPResponse *ValidaEtiquetaPLPResponse `xml:"validaEtiquetaPLPResponse,omitempty" json:"validaEtiquetaPLPResponse,omitempty" yaml:"validaEtiquetaPLPResponse,omitempty"`
}

// Operation wrapper for ValidaPlp.
// OperationValidaPlp was auto-generated from WSDL.
type OperationValidaPlp struct {
	ValidaPlp *ValidaPlp `xml:"validaPlp,omitempty" json:"validaPlp,omitempty" yaml:"validaPlp,omitempty"`
}

// Operation wrapper for ValidaPlp.
// OperationValidaPlpResponse was auto-generated from WSDL.
type OperationValidaPlpResponse struct {
	ValidaPlpResponse *ValidaPlpResponse `xml:"validaPlpResponse,omitempty" json:"validaPlpResponse,omitempty" yaml:"validaPlpResponse,omitempty"`
}

// Operation wrapper for ValidarPostagemReversa.
// OperationValidarPostagemReversa was auto-generated from WSDL.
type OperationValidarPostagemReversa struct {
	ValidarPostagemReversa *ValidarPostagemReversa `xml:"validarPostagemReversa,omitempty" json:"validarPostagemReversa,omitempty" yaml:"validarPostagemReversa,omitempty"`
}

// Operation wrapper for ValidarPostagemReversa.
// OperationValidarPostagemReversaResponse was auto-generated from
// WSDL.
type OperationValidarPostagemReversaResponse struct {
	ValidarPostagemReversaResponse *ValidarPostagemReversaResponse `xml:"validarPostagemReversaResponse,omitempty" json:"validarPostagemReversaResponse,omitempty" yaml:"validarPostagemReversaResponse,omitempty"`
}

// Operation wrapper for ValidarPostagemSimultanea.
// OperationValidarPostagemSimultanea was auto-generated from WSDL.
type OperationValidarPostagemSimultanea struct {
	ValidarPostagemSimultanea *ValidarPostagemSimultanea `xml:"validarPostagemSimultanea,omitempty" json:"validarPostagemSimultanea,omitempty" yaml:"validarPostagemSimultanea,omitempty"`
}

// Operation wrapper for ValidarPostagemSimultanea.
// OperationValidarPostagemSimultaneaResponse was auto-generated
// from WSDL.
type OperationValidarPostagemSimultaneaResponse struct {
	ValidarPostagemSimultaneaResponse *ValidarPostagemSimultaneaResponse `xml:"validarPostagemSimultaneaResponse,omitempty" json:"validarPostagemSimultaneaResponse,omitempty" yaml:"validarPostagemSimultaneaResponse,omitempty"`
}

// Operation wrapper for VerificaDisponibilidadeServico.
// OperationVerificaDisponibilidadeServico was auto-generated from
// WSDL.
type OperationVerificaDisponibilidadeServico struct {
	VerificaDisponibilidadeServico *VerificaDisponibilidadeServico `xml:"verificaDisponibilidadeServico,omitempty" json:"verificaDisponibilidadeServico,omitempty" yaml:"verificaDisponibilidadeServico,omitempty"`
}

// Operation wrapper for VerificaDisponibilidadeServico.
// OperationVerificaDisponibilidadeServicoResponse was auto-generated
// from WSDL.
type OperationVerificaDisponibilidadeServicoResponse struct {
	VerificaDisponibilidadeServicoResponse *VerificaDisponibilidadeServicoResponse `xml:"verificaDisponibilidadeServicoResponse,omitempty" json:"verificaDisponibilidadeServicoResponse,omitempty" yaml:"verificaDisponibilidadeServicoResponse,omitempty"`
}

// Operation wrapper for VerificaModalTransporte.
// OperationVerificaModalTransporte was auto-generated from WSDL.
type OperationVerificaModalTransporte struct {
	VerificaModalTransporte *VerificaModalTransporte `xml:"verificaModalTransporte,omitempty" json:"verificaModalTransporte,omitempty" yaml:"verificaModalTransporte,omitempty"`
}

// Operation wrapper for VerificaModalTransporte.
// OperationVerificaModalTransporteResponse was auto-generated
// from WSDL.
type OperationVerificaModalTransporteResponse struct {
	VerificaModalTransporteResponse *VerificaModalTransporteResponse `xml:"verificaModalTransporteResponse,omitempty" json:"verificaModalTransporteResponse,omitempty" yaml:"verificaModalTransporteResponse,omitempty"`
}

// atendeCliente implements the AtendeCliente interface.
type atendeCliente struct {
	cli *soap.Client
}

// VerificaSeTodosObjetosCancelados was auto-generated from WSDL.
func (p *atendeCliente) VerificaSeTodosObjetosCancelados(VerificaSeTodosObjetosCancelados *VerificaSeTodosObjetosCancelados) (*VerificaSeTodosObjetosCanceladosResponse, error) {
	α := struct {
		OperationVerificaSeTodosObjetosCancelados `xml:"tns:VerificaSeTodosObjetosCancelados"`
	}{
		OperationVerificaSeTodosObjetosCancelados{
			VerificaSeTodosObjetosCancelados,
		},
	}

	γ := struct {
		OperationVerificaSeTodosObjetosCanceladosResponse `xml:"VerificaSeTodosObjetosCanceladosResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("VerificaSeTodosObjetosCancelados", α, &γ); err != nil {
		return nil, err
	}
	return γ.VerificaSeTodosObjetosCanceladosResponse, nil
}

// AtualizaPagamentoNaEntrega was auto-generated from WSDL.
func (p *atendeCliente) AtualizaPagamentoNaEntrega(AtualizaPagamentoNaEntrega *AtualizaPagamentoNaEntrega) (*AtualizaPagamentoNaEntregaResponse, error) {
	α := struct {
		OperationAtualizaPagamentoNaEntrega `xml:"tns:atualizaPagamentoNaEntrega"`
	}{
		OperationAtualizaPagamentoNaEntrega{
			AtualizaPagamentoNaEntrega,
		},
	}

	γ := struct {
		OperationAtualizaPagamentoNaEntregaResponse `xml:"atualizaPagamentoNaEntregaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("AtualizaPagamentoNaEntrega", α, &γ); err != nil {
		return nil, err
	}
	return γ.AtualizaPagamentoNaEntregaResponse, nil
}

// AtualizaRemessaAgrupada was auto-generated from WSDL.
func (p *atendeCliente) AtualizaRemessaAgrupada(AtualizaRemessaAgrupada *AtualizaRemessaAgrupada) (*AtualizaRemessaAgrupadaResponse, error) {
	α := struct {
		OperationAtualizaRemessaAgrupada `xml:"tns:atualizaRemessaAgrupada"`
	}{
		OperationAtualizaRemessaAgrupada{
			AtualizaRemessaAgrupada,
		},
	}

	γ := struct {
		OperationAtualizaRemessaAgrupadaResponse `xml:"atualizaRemessaAgrupadaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("AtualizaRemessaAgrupada", α, &γ); err != nil {
		return nil, err
	}
	return γ.AtualizaRemessaAgrupadaResponse, nil
}

// BloquearObjeto was auto-generated from WSDL.
func (p *atendeCliente) BloquearObjeto(BloquearObjeto *BloquearObjeto) (*BloquearObjetoResponse, error) {
	α := struct {
		OperationBloquearObjeto `xml:"tns:bloquearObjeto"`
	}{
		OperationBloquearObjeto{
			BloquearObjeto,
		},
	}

	γ := struct {
		OperationBloquearObjetoResponse `xml:"bloquearObjetoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BloquearObjeto", α, &γ); err != nil {
		return nil, err
	}
	return γ.BloquearObjetoResponse, nil
}

// BuscaCliente was auto-generated from WSDL.
func (p *atendeCliente) BuscaCliente(BuscaCliente *BuscaCliente) (*BuscaClienteResponse, error) {
	α := struct {
		OperationBuscaCliente `xml:"tns:buscaCliente"`
	}{
		OperationBuscaCliente{
			BuscaCliente,
		},
	}

	γ := struct {
		OperationBuscaClienteResponse `xml:"buscaClienteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaCliente", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaClienteResponse, nil
}

// BuscaContrato was auto-generated from WSDL.
func (p *atendeCliente) BuscaContrato(BuscaContrato *BuscaContrato) (*BuscaContratoResponse, error) {
	α := struct {
		OperationBuscaContrato `xml:"tns:buscaContrato"`
	}{
		OperationBuscaContrato{
			BuscaContrato,
		},
	}

	γ := struct {
		OperationBuscaContratoResponse `xml:"buscaContratoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaContrato", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaContratoResponse, nil
}

// BuscaDataAtual was auto-generated from WSDL.
func (p *atendeCliente) BuscaDataAtual(BuscaDataAtual *BuscaDataAtual) (*BuscaDataAtualResponse, error) {
	α := struct {
		OperationBuscaDataAtual `xml:"tns:buscaDataAtual"`
	}{
		OperationBuscaDataAtual{
			BuscaDataAtual,
		},
	}

	γ := struct {
		OperationBuscaDataAtualResponse `xml:"buscaDataAtualResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaDataAtual", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaDataAtualResponse, nil
}

// BuscaOpcoes was auto-generated from WSDL.
func (p *atendeCliente) BuscaOpcoes(BuscaOpcoes *BuscaOpcoes) (*BuscaOpcoesResponse, error) {
	α := struct {
		OperationBuscaOpcoes `xml:"tns:buscaOpcoes"`
	}{
		OperationBuscaOpcoes{
			BuscaOpcoes,
		},
	}

	γ := struct {
		OperationBuscaOpcoesResponse `xml:"buscaOpcoesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaOpcoes", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaOpcoesResponse, nil
}

// BuscaPagamentoEntrega was auto-generated from WSDL.
func (p *atendeCliente) BuscaPagamentoEntrega(BuscaPagamentoEntrega *BuscaPagamentoEntrega) (*BuscaPagamentoEntregaResponse, error) {
	α := struct {
		OperationBuscaPagamentoEntrega `xml:"tns:buscaPagamentoEntrega"`
	}{
		OperationBuscaPagamentoEntrega{
			BuscaPagamentoEntrega,
		},
	}

	γ := struct {
		OperationBuscaPagamentoEntregaResponse `xml:"buscaPagamentoEntregaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaPagamentoEntrega", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaPagamentoEntregaResponse, nil
}

// BuscaServicos was auto-generated from WSDL.
func (p *atendeCliente) BuscaServicos(BuscaServicos *BuscaServicos) (*BuscaServicosResponse, error) {
	α := struct {
		OperationBuscaServicos `xml:"tns:buscaServicos"`
	}{
		OperationBuscaServicos{
			BuscaServicos,
		},
	}

	γ := struct {
		OperationBuscaServicosResponse `xml:"buscaServicosResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaServicos", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaServicosResponse, nil
}

// BuscaServicosAdicionaisAtivos was auto-generated from WSDL.
func (p *atendeCliente) BuscaServicosAdicionaisAtivos(BuscaServicosAdicionaisAtivos *BuscaServicosAdicionaisAtivos) (*BuscaServicosAdicionaisAtivosResponse, error) {
	α := struct {
		OperationBuscaServicosAdicionaisAtivos `xml:"tns:buscaServicosAdicionaisAtivos"`
	}{
		OperationBuscaServicosAdicionaisAtivos{
			BuscaServicosAdicionaisAtivos,
		},
	}

	γ := struct {
		OperationBuscaServicosAdicionaisAtivosResponse `xml:"buscaServicosAdicionaisAtivosResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaServicosAdicionaisAtivos", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaServicosAdicionaisAtivosResponse, nil
}

// BuscaServicosValorDeclarado was auto-generated from WSDL.
func (p *atendeCliente) BuscaServicosValorDeclarado(BuscaServicosValorDeclarado *BuscaServicosValorDeclarado) (*BuscaServicosValorDeclaradoResponse, error) {
	α := struct {
		OperationBuscaServicosValorDeclarado `xml:"tns:buscaServicosValorDeclarado"`
	}{
		OperationBuscaServicosValorDeclarado{
			BuscaServicosValorDeclarado,
		},
	}

	γ := struct {
		OperationBuscaServicosValorDeclaradoResponse `xml:"buscaServicosValorDeclaradoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaServicosValorDeclarado", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaServicosValorDeclaradoResponse, nil
}

// BuscaServicosXServicosAdicionais was auto-generated from WSDL.
func (p *atendeCliente) BuscaServicosXServicosAdicionais(BuscaServicosXServicosAdicionais *BuscaServicosXServicosAdicionais) (*BuscaServicosXServicosAdicionaisResponse, error) {
	α := struct {
		OperationBuscaServicosXServicosAdicionais `xml:"tns:buscaServicosXServicosAdicionais"`
	}{
		OperationBuscaServicosXServicosAdicionais{
			BuscaServicosXServicosAdicionais,
		},
	}

	γ := struct {
		OperationBuscaServicosXServicosAdicionaisResponse `xml:"buscaServicosXServicosAdicionaisResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaServicosXServicosAdicionais", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaServicosXServicosAdicionaisResponse, nil
}

// BuscaTarifaVale was auto-generated from WSDL.
func (p *atendeCliente) BuscaTarifaVale(BuscaTarifaVale *BuscaTarifaVale) (*BuscaTarifaValeResponse, error) {
	α := struct {
		OperationBuscaTarifaVale `xml:"tns:buscaTarifaVale"`
	}{
		OperationBuscaTarifaVale{
			BuscaTarifaVale,
		},
	}

	γ := struct {
		OperationBuscaTarifaValeResponse `xml:"buscaTarifaValeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("BuscaTarifaVale", α, &γ); err != nil {
		return nil, err
	}
	return γ.BuscaTarifaValeResponse, nil
}

// CalculaTarifaServico was auto-generated from WSDL.
func (p *atendeCliente) CalculaTarifaServico(CalculaTarifaServico *CalculaTarifaServico) (*CalculaTarifaServicoResponse, error) {
	α := struct {
		OperationCalculaTarifaServico `xml:"tns:calculaTarifaServico"`
	}{
		OperationCalculaTarifaServico{
			CalculaTarifaServico,
		},
	}

	γ := struct {
		OperationCalculaTarifaServicoResponse `xml:"calculaTarifaServicoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CalculaTarifaServico", α, &γ); err != nil {
		return nil, err
	}
	return γ.CalculaTarifaServicoResponse, nil
}

// CancelarObjeto was auto-generated from WSDL.
func (p *atendeCliente) CancelarObjeto(CancelarObjeto *CancelarObjeto) (*CancelarObjetoResponse, error) {
	α := struct {
		OperationCancelarObjeto `xml:"tns:cancelarObjeto"`
	}{
		OperationCancelarObjeto{
			CancelarObjeto,
		},
	}

	γ := struct {
		OperationCancelarObjetoResponse `xml:"cancelarObjetoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CancelarObjeto", α, &γ); err != nil {
		return nil, err
	}
	return γ.CancelarObjetoResponse, nil
}

// CancelarPedidoScol was auto-generated from WSDL.
func (p *atendeCliente) CancelarPedidoScol(CancelarPedidoScol *CancelarPedidoScol) (*CancelarPedidoScolResponse, error) {
	α := struct {
		OperationCancelarPedidoScol `xml:"tns:cancelarPedidoScol"`
	}{
		OperationCancelarPedidoScol{
			CancelarPedidoScol,
		},
	}

	γ := struct {
		OperationCancelarPedidoScolResponse `xml:"cancelarPedidoScolResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CancelarPedidoScol", α, &γ); err != nil {
		return nil, err
	}
	return γ.CancelarPedidoScolResponse, nil
}

// ConsultaCEP was auto-generated from WSDL.
func (p *atendeCliente) ConsultaCEP(ConsultaCEP *ConsultaCEP) (*ConsultaCEPResponse, error) {
	α := struct {
		OperationConsultaCEP `xml:"tns:consultaCEP"`
	}{
		OperationConsultaCEP{
			ConsultaCEP,
		},
	}

	γ := struct {
		OperationConsultaCEPResponse `xml:"consultaCEPResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ConsultaCEP", α, &γ); err != nil {
		return nil, err
	}
	return γ.ConsultaCEPResponse, nil
}

// FechaPlp was auto-generated from WSDL.
func (p *atendeCliente) FechaPlp(FechaPlp *FechaPlp) (*FechaPlpResponse, error) {
	α := struct {
		OperationFechaPlp `xml:"tns:fechaPlp"`
	}{
		OperationFechaPlp{
			FechaPlp,
		},
	}

	γ := struct {
		OperationFechaPlpResponse `xml:"fechaPlpResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("FechaPlp", α, &γ); err != nil {
		return nil, err
	}
	return γ.FechaPlpResponse, nil
}

// FechaPlpVariosServicos was auto-generated from WSDL.
func (p *atendeCliente) FechaPlpVariosServicos(FechaPlpVariosServicos *FechaPlpVariosServicos) (*FechaPlpVariosServicosResponse, error) {
	α := struct {
		OperationFechaPlpVariosServicos `xml:"tns:fechaPlpVariosServicos"`
	}{
		OperationFechaPlpVariosServicos{
			FechaPlpVariosServicos,
		},
	}

	γ := struct {
		OperationFechaPlpVariosServicosResponse `xml:"fechaPlpVariosServicosResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("FechaPlpVariosServicos", α, &γ); err != nil {
		return nil, err
	}
	return γ.FechaPlpVariosServicosResponse, nil
}

// GeraDigitoVerificadorEtiquetas was auto-generated from WSDL.
func (p *atendeCliente) GeraDigitoVerificadorEtiquetas(GeraDigitoVerificadorEtiquetas *GeraDigitoVerificadorEtiquetas) (*GeraDigitoVerificadorEtiquetasResponse, error) {
	α := struct {
		OperationGeraDigitoVerificadorEtiquetas `xml:"tns:geraDigitoVerificadorEtiquetas"`
	}{
		OperationGeraDigitoVerificadorEtiquetas{
			GeraDigitoVerificadorEtiquetas,
		},
	}

	γ := struct {
		OperationGeraDigitoVerificadorEtiquetasResponse `xml:"geraDigitoVerificadorEtiquetasResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GeraDigitoVerificadorEtiquetas", α, &γ); err != nil {
		return nil, err
	}
	return γ.GeraDigitoVerificadorEtiquetasResponse, nil
}

// GetStatusCartaoPostagem was auto-generated from WSDL.
func (p *atendeCliente) GetStatusCartaoPostagem(GetStatusCartaoPostagem *GetStatusCartaoPostagem) (*GetStatusCartaoPostagemResponse, error) {
	α := struct {
		OperationGetStatusCartaoPostagem `xml:"tns:getStatusCartaoPostagem"`
	}{
		OperationGetStatusCartaoPostagem{
			GetStatusCartaoPostagem,
		},
	}

	γ := struct {
		OperationGetStatusCartaoPostagemResponse `xml:"getStatusCartaoPostagemResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetStatusCartaoPostagem", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStatusCartaoPostagemResponse, nil
}

// GetStatusPLP was auto-generated from WSDL.
func (p *atendeCliente) GetStatusPLP(GetStatusPLP *GetStatusPLP) (*GetStatusPLPResponse, error) {
	α := struct {
		OperationGetStatusPLP `xml:"tns:getStatusPLP"`
	}{
		OperationGetStatusPLP{
			GetStatusPLP,
		},
	}

	γ := struct {
		OperationGetStatusPLPResponse `xml:"getStatusPLPResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetStatusPLP", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStatusPLPResponse, nil
}

// IntegrarUsuarioScol was auto-generated from WSDL.
func (p *atendeCliente) IntegrarUsuarioScol(IntegrarUsuarioScol *IntegrarUsuarioScol) (*IntegrarUsuarioScolResponse, error) {
	α := struct {
		OperationIntegrarUsuarioScol `xml:"tns:integrarUsuarioScol"`
	}{
		OperationIntegrarUsuarioScol{
			IntegrarUsuarioScol,
		},
	}

	γ := struct {
		OperationIntegrarUsuarioScolResponse `xml:"integrarUsuarioScolResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("IntegrarUsuarioScol", α, &γ); err != nil {
		return nil, err
	}
	return γ.IntegrarUsuarioScolResponse, nil
}

// ObterClienteAtualizacao was auto-generated from WSDL.
func (p *atendeCliente) ObterClienteAtualizacao(ObterClienteAtualizacao *ObterClienteAtualizacao) (*ObterClienteAtualizacaoResponse, error) {
	α := struct {
		OperationObterClienteAtualizacao `xml:"tns:obterClienteAtualizacao"`
	}{
		OperationObterClienteAtualizacao{
			ObterClienteAtualizacao,
		},
	}

	γ := struct {
		OperationObterClienteAtualizacaoResponse `xml:"obterClienteAtualizacaoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ObterClienteAtualizacao", α, &γ); err != nil {
		return nil, err
	}
	return γ.ObterClienteAtualizacaoResponse, nil
}

// ObterEmbalagemLRS was auto-generated from WSDL.
func (p *atendeCliente) ObterEmbalagemLRS(ObterEmbalagemLRS *ObterEmbalagemLRS) (*ObterEmbalagemLRSResponse, error) {
	α := struct {
		OperationObterEmbalagemLRS `xml:"tns:obterEmbalagemLRS"`
	}{
		OperationObterEmbalagemLRS{
			ObterEmbalagemLRS,
		},
	}

	γ := struct {
		OperationObterEmbalagemLRSResponse `xml:"obterEmbalagemLRSResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ObterEmbalagemLRS", α, &γ); err != nil {
		return nil, err
	}
	return γ.ObterEmbalagemLRSResponse, nil
}

// ObterMensagemParametrizada was auto-generated from WSDL.
func (p *atendeCliente) ObterMensagemParametrizada(ObterMensagemParametrizada *ObterMensagemParametrizada) (*ObterMensagemParametrizadaResponse, error) {
	α := struct {
		OperationObterMensagemParametrizada `xml:"tns:obterMensagemParametrizada"`
	}{
		OperationObterMensagemParametrizada{
			ObterMensagemParametrizada,
		},
	}

	γ := struct {
		OperationObterMensagemParametrizadaResponse `xml:"obterMensagemParametrizadaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ObterMensagemParametrizada", α, &γ); err != nil {
		return nil, err
	}
	return γ.ObterMensagemParametrizadaResponse, nil
}

// PesquisarDimensoesServico was auto-generated from WSDL.
func (p *atendeCliente) PesquisarDimensoesServico(PesquisarDimensoesServico *PesquisarDimensoesServico) (*PesquisarDimensoesServicoResponse, error) {
	α := struct {
		OperationPesquisarDimensoesServico `xml:"tns:pesquisarDimensoesServico"`
	}{
		OperationPesquisarDimensoesServico{
			PesquisarDimensoesServico,
		},
	}

	γ := struct {
		OperationPesquisarDimensoesServicoResponse `xml:"pesquisarDimensoesServicoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("PesquisarDimensoesServico", α, &γ); err != nil {
		return nil, err
	}
	return γ.PesquisarDimensoesServicoResponse, nil
}

// PesquisarEmbalagensPorServico was auto-generated from WSDL.
func (p *atendeCliente) PesquisarEmbalagensPorServico(PesquisarEmbalagensPorServico *PesquisarEmbalagensPorServico) (*PesquisarEmbalagensPorServicoResponse, error) {
	α := struct {
		OperationPesquisarEmbalagensPorServico `xml:"tns:pesquisarEmbalagensPorServico"`
	}{
		OperationPesquisarEmbalagensPorServico{
			PesquisarEmbalagensPorServico,
		},
	}

	γ := struct {
		OperationPesquisarEmbalagensPorServicoResponse `xml:"pesquisarEmbalagensPorServicoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("PesquisarEmbalagensPorServico", α, &γ); err != nil {
		return nil, err
	}
	return γ.PesquisarEmbalagensPorServicoResponse, nil
}

// PesquisarParametrosPorDescricao was auto-generated from WSDL.
func (p *atendeCliente) PesquisarParametrosPorDescricao(PesquisarParametrosPorDescricao *PesquisarParametrosPorDescricao) (*PesquisarParametrosPorDescricaoResponse, error) {
	α := struct {
		OperationPesquisarParametrosPorDescricao `xml:"tns:pesquisarParametrosPorDescricao"`
	}{
		OperationPesquisarParametrosPorDescricao{
			PesquisarParametrosPorDescricao,
		},
	}

	γ := struct {
		OperationPesquisarParametrosPorDescricaoResponse `xml:"pesquisarParametrosPorDescricaoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("PesquisarParametrosPorDescricao", α, &γ); err != nil {
		return nil, err
	}
	return γ.PesquisarParametrosPorDescricaoResponse, nil
}

// PesquisarServicosAdicionais was auto-generated from WSDL.
func (p *atendeCliente) PesquisarServicosAdicionais(PesquisarServicosAdicionais *PesquisarServicosAdicionais) (*PesquisarServicosAdicionaisResponse, error) {
	α := struct {
		OperationPesquisarServicosAdicionais `xml:"tns:pesquisarServicosAdicionais"`
	}{
		OperationPesquisarServicosAdicionais{
			PesquisarServicosAdicionais,
		},
	}

	γ := struct {
		OperationPesquisarServicosAdicionaisResponse `xml:"pesquisarServicosAdicionaisResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("PesquisarServicosAdicionais", α, &γ); err != nil {
		return nil, err
	}
	return γ.PesquisarServicosAdicionaisResponse, nil
}

// SolicitaEtiquetas was auto-generated from WSDL.
func (p *atendeCliente) SolicitaEtiquetas(SolicitaEtiquetas *SolicitaEtiquetas) (*SolicitaEtiquetasResponse, error) {
	α := struct {
		OperationSolicitaEtiquetas `xml:"solicitaEtiquetas"`
	}{
		OperationSolicitaEtiquetas{
			SolicitaEtiquetas,
		},
	}

	γ := struct {
		OperationSolicitaEtiquetasResponse `xml:"solicitaEtiquetasResponse"`
	}{}

	if err := p.cli.RoundTripWithAction("", α, &γ); err != nil {
		return nil, err
	}
	return γ.SolicitaEtiquetasResponse, nil
}

// SolicitaPLP was auto-generated from WSDL.
func (p *atendeCliente) SolicitaPLP(SolicitaPLP *SolicitaPLP) (*SolicitaPLPResponse, error) {
	α := struct {
		OperationSolicitaPLP `xml:"tns:solicitaPLP"`
	}{
		OperationSolicitaPLP{
			SolicitaPLP,
		},
	}

	γ := struct {
		OperationSolicitaPLPResponse `xml:"solicitaPLPResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SolicitaPLP", α, &γ); err != nil {
		return nil, err
	}
	return γ.SolicitaPLPResponse, nil
}

// SolicitaXmlPlp was auto-generated from WSDL.
func (p *atendeCliente) SolicitaXmlPlp(SolicitaXmlPlp *SolicitaXmlPlp) (*SolicitaXmlPlpResponse, error) {
	α := struct {
		OperationSolicitaXmlPlp `xml:"tns:solicitaXmlPlp"`
	}{
		OperationSolicitaXmlPlp{
			SolicitaXmlPlp,
		},
	}

	γ := struct {
		OperationSolicitaXmlPlpResponse `xml:"solicitaXmlPlpResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SolicitaXmlPlp", α, &γ); err != nil {
		return nil, err
	}
	return γ.SolicitaXmlPlpResponse, nil
}

// SolicitarPostagemScol was auto-generated from WSDL.
func (p *atendeCliente) SolicitarPostagemScol(SolicitarPostagemScol *SolicitarPostagemScol) (*SolicitarPostagemScolResponse, error) {
	α := struct {
		OperationSolicitarPostagemScol `xml:"tns:solicitarPostagemScol"`
	}{
		OperationSolicitarPostagemScol{
			SolicitarPostagemScol,
		},
	}

	γ := struct {
		OperationSolicitarPostagemScolResponse `xml:"solicitarPostagemScolResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SolicitarPostagemScol", α, &γ); err != nil {
		return nil, err
	}
	return γ.SolicitarPostagemScolResponse, nil
}

// ValidaEtiquetaPLP was auto-generated from WSDL.
func (p *atendeCliente) ValidaEtiquetaPLP(ValidaEtiquetaPLP *ValidaEtiquetaPLP) (*ValidaEtiquetaPLPResponse, error) {
	α := struct {
		OperationValidaEtiquetaPLP `xml:"tns:validaEtiquetaPLP"`
	}{
		OperationValidaEtiquetaPLP{
			ValidaEtiquetaPLP,
		},
	}

	γ := struct {
		OperationValidaEtiquetaPLPResponse `xml:"validaEtiquetaPLPResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ValidaEtiquetaPLP", α, &γ); err != nil {
		return nil, err
	}
	return γ.ValidaEtiquetaPLPResponse, nil
}

// ValidaPlp was auto-generated from WSDL.
func (p *atendeCliente) ValidaPlp(ValidaPlp *ValidaPlp) (*ValidaPlpResponse, error) {
	α := struct {
		OperationValidaPlp `xml:"tns:validaPlp"`
	}{
		OperationValidaPlp{
			ValidaPlp,
		},
	}

	γ := struct {
		OperationValidaPlpResponse `xml:"validaPlpResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ValidaPlp", α, &γ); err != nil {
		return nil, err
	}
	return γ.ValidaPlpResponse, nil
}

// ValidarPostagemReversa was auto-generated from WSDL.
func (p *atendeCliente) ValidarPostagemReversa(ValidarPostagemReversa *ValidarPostagemReversa) (*ValidarPostagemReversaResponse, error) {
	α := struct {
		OperationValidarPostagemReversa `xml:"tns:validarPostagemReversa"`
	}{
		OperationValidarPostagemReversa{
			ValidarPostagemReversa,
		},
	}

	γ := struct {
		OperationValidarPostagemReversaResponse `xml:"validarPostagemReversaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ValidarPostagemReversa", α, &γ); err != nil {
		return nil, err
	}
	return γ.ValidarPostagemReversaResponse, nil
}

// ValidarPostagemSimultanea was auto-generated from WSDL.
func (p *atendeCliente) ValidarPostagemSimultanea(ValidarPostagemSimultanea *ValidarPostagemSimultanea) (*ValidarPostagemSimultaneaResponse, error) {
	α := struct {
		OperationValidarPostagemSimultanea `xml:"tns:validarPostagemSimultanea"`
	}{
		OperationValidarPostagemSimultanea{
			ValidarPostagemSimultanea,
		},
	}

	γ := struct {
		OperationValidarPostagemSimultaneaResponse `xml:"validarPostagemSimultaneaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ValidarPostagemSimultanea", α, &γ); err != nil {
		return nil, err
	}
	return γ.ValidarPostagemSimultaneaResponse, nil
}

// VerificaDisponibilidadeServico was auto-generated from WSDL.
func (p *atendeCliente) VerificaDisponibilidadeServico(VerificaDisponibilidadeServico *VerificaDisponibilidadeServico) (*VerificaDisponibilidadeServicoResponse, error) {
	α := struct {
		OperationVerificaDisponibilidadeServico `xml:"tns:verificaDisponibilidadeServico"`
	}{
		OperationVerificaDisponibilidadeServico{
			VerificaDisponibilidadeServico,
		},
	}

	γ := struct {
		OperationVerificaDisponibilidadeServicoResponse `xml:"verificaDisponibilidadeServicoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("VerificaDisponibilidadeServico", α, &γ); err != nil {
		return nil, err
	}
	return γ.VerificaDisponibilidadeServicoResponse, nil
}

// VerificaModalTransporte was auto-generated from WSDL.
func (p *atendeCliente) VerificaModalTransporte(VerificaModalTransporte *VerificaModalTransporte) (*VerificaModalTransporteResponse, error) {
	α := struct {
		OperationVerificaModalTransporte `xml:"tns:verificaModalTransporte"`
	}{
		OperationVerificaModalTransporte{
			VerificaModalTransporte,
		},
	}

	γ := struct {
		OperationVerificaModalTransporteResponse `xml:"verificaModalTransporteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("VerificaModalTransporte", α, &γ); err != nil {
		return nil, err
	}
	return γ.VerificaModalTransporteResponse, nil
}
