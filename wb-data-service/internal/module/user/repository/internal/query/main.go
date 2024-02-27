package query

import (
	"wb-data-service-golang/wb-data-service/internal/module/user/repository/internal/model"
	"wb-data-service-golang/wb-data-service/pkg/goqu"
)

func GetInsert(model model.User) (string, []any, error) {
	record := make(goqu.Record)

	if model.Username.Valid {
		record["username"] = model.Username
	}

	record["email"] = model.Email
	record["password_hash"] = model.Password

	return goqu.Dialect.
		Insert(model.TableName()).
		Rows(record).
		Returning(model).
		Prepared(true).
		ToSQL()
}

func GetSelectById(model model.User) (string, []any, error) {
	return goqu.Dialect.
		Select(model).
		From(model.TableName()).
		Where(
			goqu.C("id").Eq(model.Id),
		).Prepared(true).ToSQL()
}

func GetSelectByCreds(model model.User) (string, []any, error) {
	return goqu.Dialect.
		Select(model).
		From(model.TableName()).
		Where(
			goqu.C("email").Eq(model.Email),
			goqu.C("password_hash").Eq(model.Password),
		).Prepared(true).ToSQL()
}

func GetDelete(model model.User) (string, []any, error) {
	return goqu.Dialect.
		Delete(model.TableName()).
		Where(
			goqu.C("id").Eq(model.Id),
		).Prepared(true).ToSQL()
}

func GetUpdate(model model.User) (string, []any, error) {
	record := make(goqu.Record)

	if model.Username.Valid {
		record["username"] = model.Username.String
	}
	if model.Email != "" {
		record["email"] = model.Email
	}

	return goqu.Dialect.
		Update(model.TableName()).
		Set(record).
		Where(
			goqu.C("id").Eq(model.Id),
		).Prepared(true).ToSQL()
}

func GetUpdatePassword(model model.User) (string, []any, error) {
	return goqu.Dialect.
		Update(model.TableName()).
		Set(goqu.Record{
			"password_hash": model.Password},
		).
		Where(
			goqu.C("id").Eq(model.Id),
		).Prepared(true).ToSQL()
}
