package moradorModel

import (
	"backend/db"
	"context"
	"fmt"
)

func Delete(cpf string) error {
	const query = `DELETE FROM moradores WHERE cpf=$1`

	result, err := db.Connection.Exec(context.Background(), query, cpf)

	if err != nil {
		return err
	}

	// TODO: Return with status code to process the response better
	if result.RowsAffected() == 0 {
		return fmt.Errorf("morador not found")
	}

	return nil
}
