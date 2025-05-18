package schemas

import "github.com/google/uuid"

type Apartamento struct {
	Numero int    `json:"numero" binding:"required"`
	Bloco  string `json:"bloco" binding:"required"`
}

type ApartamentoWithId struct {
	Id uuid.UUID `json:"id"`
	Apartamento
}
