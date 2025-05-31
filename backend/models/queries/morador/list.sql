SELECT 
	cpf, apartamento_id, nome,
	telefone, responsavel, proprietario
FROM moradores LIMIT $1 OFFSET $2;