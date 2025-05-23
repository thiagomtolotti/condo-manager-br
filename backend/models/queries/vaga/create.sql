INSERT INTO vagas (
	apartamento_id, numero
) VALUES ($1, $2) RETURNING id