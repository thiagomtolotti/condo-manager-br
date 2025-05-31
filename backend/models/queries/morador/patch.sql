UPDATE moradores SET (
	apartamento_id = $1,
	nome = $2,
	telefone = $3,
	responsavel = $4,
	proprietario = $5
 )WHERE cpf = $6;