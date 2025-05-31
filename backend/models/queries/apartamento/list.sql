SELECT 
	id, numero, bloco 
FROM apartamentos LIMIT $1 OFFSET $2;