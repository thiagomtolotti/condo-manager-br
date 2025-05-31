package moradorModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
)

func ApartamentoHasMorador(apartamento_id uuid.UUID) (bool, *errs.AppError) {
	query, err := utils.LoadSQL("morador/apartamento_has_morador.sql")
	if err != nil {
		return false, errs.Unexpected(fmt.Errorf("Reading has morador in apartment sql file: %w", err))
	}

	var has_morador bool
	err = pgxscan.Get(context.Background(), db.Connection, &has_morador, query, apartamento_id.String())
	if err != nil {
		return false, errs.Unexpected(fmt.Errorf("Querying has morador in apartment: %w", err))
	}

	return has_morador, nil
}
