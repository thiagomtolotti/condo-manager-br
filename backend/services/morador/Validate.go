package moradorService

import (
	"backend/db"
	"backend/utils/cpf"
	"context"
)

func Validate(cpf cpf.CPF) (bool, error) {

	const query = `SELECT 1 FROM moradores WHERE cpf=$1`
	var exists int
	err := db.Connection.QueryRow(context.Background(), query, cpf).Scan(&exists)

	if exists == 0 {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	return false, nil
}
