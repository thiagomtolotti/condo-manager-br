CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS vagas (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	numero INTEGER NOT NULL CHECK (numero > 0),
	apartamento_id UUID NOT NULL,
	FOREIGN KEY (apartamento_id) REFERENCES apartamentos(id)
)