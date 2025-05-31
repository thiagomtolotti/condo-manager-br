package moradorModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func GetCount() (int, *errs.AppError) {
	query, err := utils.LoadSQL("morador/get_count.sql")
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("reading get moradores count SQL file: %w", err))
		return 0, err
	}

	var total int
	err = pgxscan.Get(context.Background(), db.Connection, &total, query)
	if err != nil {
		var err = errs.Unexpected(fmt.Errorf("querying moradores total count: %w", err))
		return 0, err
	}

	return total, nil

}
