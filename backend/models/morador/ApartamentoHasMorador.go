package moradorModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func ApartamentoHasMorador(apartamento_id uuid.UUID) (bool, *errs.AppError) {
	query, err := utils.LoadSQL("morador/apartamento_has_morador.sql")
	if err != nil {
		return false, errs.Unexpected(fmt.Errorf("Reading has morador in apartment sql file: %w", err))
	}

	var has_morador bool
	err = db.Connection.QueryRow(context.Background(), query, apartamento_id.String()).Scan(&has_morador)
	if err != nil {
		return false, errs.Unexpected(fmt.Errorf("Querying has morador in apartment: %w", err))
	}

	return has_morador, nil
}
