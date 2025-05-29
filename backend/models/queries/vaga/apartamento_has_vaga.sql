SELECT EXISTS (
	SELECT 1 FROM vagas WHERE apartamento_id = $1
) AS has_vaga;