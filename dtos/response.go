package dtos

type Resposta struct{
	Mensagem string`json:"mensagem"`
	Status int`json:"status"`
}