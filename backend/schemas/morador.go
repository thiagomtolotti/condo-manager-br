package schemas

import "github.com/google/uuid"

type MoradorWithoutCPF struct {
	Apartamento_id uuid.UUID `json:"apartamento_id" binding:"required"`
	Nome           string    `json:"nome" binding:"required"`
	Telefone       string    `json:"telefone" binding:"required"`
	Responsavel    bool      `json:"responsavel"`
	Proprietario   bool      `json:"proprietario"`
}

type Morador struct {
	Cpf string `json:"cpf" binding:"required"`
	MoradorWithoutCPF
}
