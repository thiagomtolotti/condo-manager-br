package vagaModel

import (
	"backend/db"
	"backend/utils"
	"context"
	"fmt"
)

func GetCount() (int, error) {
	var total int

	query, err := utils.LoadSQL("vaga/get_count.sql")
	if err != nil {
		return 0, fmt.Errorf("error reading get vaga count sql: %w", err)
	}

	err = db.Connection.QueryRow(context.Background(), query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("error fetching vagas count: %w", err)
	}

	return total, nil
}
