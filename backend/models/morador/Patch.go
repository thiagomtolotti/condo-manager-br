package moradorModel

import (
	"backend/db"
	"backend/schemas"
	"backend/utils/cpf"
	"context"
)

func Patch(cpf cpf.CPF, data schemas.MoradorWithoutCPF) error {
	const query = `
        UPDATE moradores SET 
            apartamento_id = $1,
            nome = $2,
            telefone = $3,
            responsavel = $4,
            proprietario = $5
        WHERE cpf = $6`

	_, err := db.Connection.Exec(
		context.Background(),
		query,
		data.Apartamento_id,
		data.Nome,
		data.Telefone,
		data.Responsavel,
		data.Proprietario,
		cpf,
	)

	if err != nil {
		return err
	}

	return err
}
