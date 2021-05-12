package postgresql

import (
	"go/internal/config"

	"github.com/labstack/echo/v4"
)

type IPostgresqlService interface {
	GetbyId(id int64, ctx echo.Context) PostgresqlResponse
	GetAll(ctx echo.Context) []PostgresqlResponse
	Create(req PostgresqlRequest, ctx echo.Context) string
	Update(id int64, req PostgresqlRequest, ctx echo.Context) string
	Delete(id int64, ctx echo.Context) string
}

type PostgresqlService struct {
	shared *config.GlobalShared
	repo   IPostgresqlRepository
}

func NewPostgresqlService(s *config.GlobalShared) IPostgresqlService {
	return PostgresqlService{
		shared: s,
		repo:   NewPostgresqlRepository(s),
	}
}

func (s PostgresqlService) GetbyId(id int64, ctx echo.Context) PostgresqlResponse {
	row := s.repo.FindById(id)
	resp := s.convertDAOtoDTO(row)
	return resp

}

func (s PostgresqlService) GetAll(ctx echo.Context) []PostgresqlResponse {

	var resp []PostgresqlResponse
	rows := s.repo.FindAll()
	for _, v := range rows {
		temp := s.convertDAOtoDTO(v)
		resp = append(resp, temp)
	}

	return resp

}

func (s PostgresqlService) Create(req PostgresqlRequest, ctx echo.Context) string {
	ent := s.convertDTOtoDAO(req)
	rows, tx := s.repo.Create(ent, ctx)

	err := tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	count, err := rows.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if count != 1 {
		panic("Insert Failed")
	}

	return "Insert Success"

}

func (s PostgresqlService) Update(id int64, req PostgresqlRequest, ctx echo.Context) string {

	ent := s.convertDTOtoDAO(req)
	row, tx := s.repo.Update(id, ent, ctx)

	err := tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	count, err := row.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if count != 1 {
		panic("Update Failed")
	}

	return "Update Success"

}

func (s PostgresqlService) Delete(id int64, ctx echo.Context) string {
	row, tx := s.repo.Delete(id, ctx)

	err := tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	count, err := row.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if count != 1 {
		panic("Delete Failed")
	}

	return "Delete Success"
}

func (s PostgresqlService) convertDTOtoDAO(req PostgresqlRequest) PostgresqlEntity {

	ent := PostgresqlEntity{
		Name:    req.Name,
		Avatar:  req.Avatar,
		Address: req.Address,
	}

	return ent
}

func (s PostgresqlService) convertDAOtoDTO(ent PostgresqlEntity) PostgresqlResponse {
	resp := PostgresqlResponse{
		Avatar:  ent.Avatar,
		Name:    ent.Name,
		Address: ent.Address,
	}

	return resp
}
