package vagaModel

import (
	"backend/db"
	"backend/errs"
	"backend/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func ApartamentoHasVaga(apartamento_id uuid.UUID) (bool, *errs.AppError) {
	query, err := utils.LoadSQL("vaga/apartamento_has_vaga.sql")
	if err != nil {
		return false, errs.Unexpected(fmt.Errorf("reading apartamento_has_vaga.sql: %w", err))
	}

	var has_morador bool
	err = db.Connection.QueryRow(context.Background(), query, apartamento_id.String()).Scan(&has_morador)
	if err != nil {
		return false, errs.Unexpected(fmt.Errorf("querying apartamento has vaga: %w", err))
	}

	return has_morador, nil
}
