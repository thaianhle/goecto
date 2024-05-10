package createuser

import (
	"context"
	"database/sql"
	"goecto/changeset"
	"goecto/repo"
)

type UserService struct {
	db   *sql.DB
	repo *repo.Repo
}

func (u *UserService) CreateUser(ctx context.Context, dto *UserCreateDTO) (*User, error) {
	userEntity := &User{}
	userChangeset := changeset.CastValues(userEntity, map[string]interface{}{
		// Id not need to cast, because it will have then save userChangeset into repo
		"Name":   dto.Name,
		"Age":    dto.Age,
		"Detail": dto.Detail,
	})

	if !userChangeset.ValidInsert() {
		return nil, userChangeset.NotNullErrors()
	}
	err := u.repo.Save(ctx, userChangeset)
	if err != nil {
		return nil, err
	}

	// then here, userEntity will have Id created from repo
	return userEntity, nil
}
