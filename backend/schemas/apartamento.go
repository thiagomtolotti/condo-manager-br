package schemas

type Apartamento struct {
	Numero int    `json:"numero"`
	Bloco  string `json:"bloco"`
}

type ApartamentoWithId struct {
	Id string `json:id`
	Apartamento
}
