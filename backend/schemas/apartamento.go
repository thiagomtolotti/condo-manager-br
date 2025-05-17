package schemas

import "github.com/google/uuid"

type Apartamento struct {
	Numero int    `json:"numero"`
	Bloco  string `json:"bloco"`
}

type ApartamentoWithId struct {
	Id uuid.UUID `json:"id"`
	Apartamento
}
