package moradorModel

import (
	"backend/db"
	"backend/schemas"
	"context"
)

func Create(data schemas.Morador) error {
	const query = `
        INSERT INTO moradores (
            cpf, apartamento_id, nome, 
            telefone, responsavel, proprietario
        ) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Connection.Exec(
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
