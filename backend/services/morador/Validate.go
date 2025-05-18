package moradorService

import (
	"backend/db"
	"backend/utils"
	"context"
)

func Validate(cpf string) (bool, error) {
	if !utils.ValidateCPF(cpf) {
		return false, nil
	}

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
