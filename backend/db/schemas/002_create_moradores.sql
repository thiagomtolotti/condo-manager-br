CREATE TABLE IF NOT EXISTS moradores (
	cpf VARCHAR(11) PRIMARY KEY,
	apartamento_id UUID NOT NULL,
	nome VARCHAR(100) NOT NULL,
	telefone VARCHAR(15) NOT NULL,
	responsavel BOOLEAN NOT NULL,
	proprietario BOOLEAN NOT NULL,
	FOREIGN KEY (apartamento_id) REFERENCES apartamentos(id)
)