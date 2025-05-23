package apartmentoModel

import (
	"backend/db"
	"backend/utils"
	"context"
	"fmt"
)

func GetCount() (int, error) {
	var count int
	query, err := utils.LoadSQL("apartamento/get_count.sql")

	if err != nil {
		return 0, fmt.Errorf("error reading get apartamento count sql: %w", err)
	}

	err = db.Connection.QueryRow(context.Background(), query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
