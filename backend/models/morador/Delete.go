package moradorModel

import (
	"backend/db"
	"context"
)

func Delete(cpf string) (bool, error) {
	const query = `DELETE FROM moradores WHERE cpf=$1`

	result, err := db.Connection.Exec(context.Background(), query, cpf)

	if err != nil {
		return false, err
	}

	return result.RowsAffected() != 0, nil
}
