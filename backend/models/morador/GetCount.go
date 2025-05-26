package moradorModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"context"
	"fmt"
)

func GetCount() (int, *errs.AppError) {
	var total int

	query, err := utils.LoadSQL("morador/get_count.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading get moradores count SQL file: %w", err))
		return 0, err
	}

	err = db.Connection.QueryRow(context.Background(), query).Scan(&total)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("querying moradores total count: %w", err))
		return 0, err
	}

	return total, nil

}
