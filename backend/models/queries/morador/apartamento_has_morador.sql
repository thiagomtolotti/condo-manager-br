SELECT EXISTS (
	SELECT 1 FROM moradores WHERE apartamento_id = $1
) AS has_morador;