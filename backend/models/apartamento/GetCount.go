package apartamentoModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func GetCount() (int, *errs.AppError) {
	var total int
	query, err := utils.LoadSQL("apartamento/get_count.sql")

	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading get apartamento count SQL: %w", err))
		return 0, err
	}

	err = pgxscan.Select(context.Background(), db.Connection, &total, query)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("querying apartamento count: %w", err))
		return 0, err
	}

	return total, nil
}
