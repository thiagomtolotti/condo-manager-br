package moradorModel

import (
	"backend/db"
	"backend/utils/cpf"
	"context"
)

func Delete(cpf cpf.CPF) (bool, error) {
	const query = `DELETE FROM moradores WHERE cpf=$1`

	result, err := db.Connection.Exec(context.Background(), query, cpf)

	if err != nil {
		return false, err
	}

	return result.RowsAffected() != 0, nil
}
