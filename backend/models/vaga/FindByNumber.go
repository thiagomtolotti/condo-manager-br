package vagaModel

import (
	"backend/db"
	"backend/errs"
	"backend/schemas"
	"backend/utils"
	"context"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func FindByNumber(number int) (*schemas.VagaWithApartment, *errs.AppError) {
	query, err := utils.LoadSQL("vaga/find_by_number.sql")
	if err != nil {
		return nil, errs.Unexpected(fmt.Errorf("reading find_by_number.sql: %w", err))
	}

	var vaga schemas.VagaWithApartment
	err = pgxscan.Get(context.Background(), db.Connection, &vaga, query, number)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errs.Unexpected(fmt.Errorf("querying vaga by number: %w", err))
	}

	return &vaga, nil
}
