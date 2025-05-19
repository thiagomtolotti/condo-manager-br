package schemas

import "github.com/google/uuid"

type Vaga struct {
	Numero int `json:"numero" binding:"required"`
}

type VagaWithId struct {
	Id uuid.UUID `json:"id" binding:"required"`
	Vaga
}
