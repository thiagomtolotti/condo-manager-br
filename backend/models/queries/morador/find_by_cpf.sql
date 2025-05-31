SELECT (
	cpf, apartamento_id, nome, 
	telefone, responsavel, proprietario
) FROM moradores WHERE cpf=$1;