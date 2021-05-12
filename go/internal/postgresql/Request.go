package postgresql

type PostgresqlRequest struct {
	Avatar  string `validate:"required" json:"avatar"`
	Name    string `validate:"required" json:"name"`
	Address string `validate:"required" json:"address"`
}
