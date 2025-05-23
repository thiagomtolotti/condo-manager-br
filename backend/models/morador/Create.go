package moradorModel

import (
	"backend/db"
	"backend/schemas"
	"backend/utils"
	"context"
	"fmt"
)

func Create(data schemas.Morador) error {
	query, err := utils.LoadSQL("morador/create.sql")

	if err != nil {
		return fmt.Errorf("error reading create morador sql file: %w", err)
	}

	_, err = db.Connection.Exec(
		context.Background(),
		query,
		data.Cpf,
		data.Apartamento_id,
		data.Nome,
		data.Telefone,
		data.Responsavel,
		data.Proprietario,
	)

	if err != nil {
		return err
	}

	return nil
}
