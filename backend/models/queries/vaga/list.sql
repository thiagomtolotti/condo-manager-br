SELECT 
	id, numero, apartamento_id
FROM vagas LIMIT $1 OFFSET $2;