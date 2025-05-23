package moradorModel

import (
	"backend/db"
	"backend/utils"
	"context"
	"fmt"
)

func GetCount() (int, error) {
	var total int

	query, err := utils.LoadSQL("morador/get_count.sql")
	if err != nil {
		return 0, fmt.Errorf("error reading get morador count sql: %w", err)
	}

	err = db.Connection.QueryRow(context.Background(), query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("error querying total count: %w", err)
	}

	return total, nil

}
