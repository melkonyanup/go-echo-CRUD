package postgresql

type MainResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type PostgresqlResponse struct {
	Avatar  string `validate:"required"`
	Name    string `validate:"required"`
	Address string `validate:"required"`
}
