package moradorModel

import (
	"backend/db"
	"backend/schemas"
	"backend/utils"
	"backend/utils/cpf"
	"context"
	"fmt"
)

func Patch(cpf cpf.CPF, data schemas.MoradorWithoutCPF) error {
	query, err := utils.LoadSQL("morador/patch.sql")
	if err != nil {
		return fmt.Errorf("error reading patch morador sql: %w", err)
	}

	_, err = db.Connection.Exec(
		context.Background(),
		query,
		data.Apartamento_id,
		data.Nome,
		data.Telefone,
		data.Responsavel,
		data.Proprietario,
		cpf.Value,
	)

	if err != nil {
		return err
	}

	return err
}
