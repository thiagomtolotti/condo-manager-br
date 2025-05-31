package vagaModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func GetCount() (int, *errs.AppError) {
	query, err := utils.LoadSQL("vaga/get_count.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading get vagas count SQL file: %w", err))
		return 0, err
	}

	var total int
	err = pgxscan.DefaultAPI.Get(context.Background(), db.Connection, &total, query)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("fetching vagas count: %w", err))
		return 0, err
	}

	return total, nil
}
