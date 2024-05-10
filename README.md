# GOECTO Expired from Elixir ecto changeset
## 1. Usage
Create User Entity implement changeset.Schema interface
- Example:
- In user.go
    ```go
    // user.go
    package user

    import "goecto/changeset"
    
    type UserDetail struct {
    	Address    string `json:"address"`
    	AvatarLink string `json:"avatar_link`
    }
    type User struct {
	    Id     uint32
	    Name   string
	    Age    uint8
	    Detail *UserDetail
    }
    
    // implement for changeset.Schema interface
    func (u *User) Validators() map[string]*changeset.Box {
    	return map[string]*changeset.Box{
    		"Id":     changeset.NewBox().Ops(changeset.AI, changeset.NotNullable),
    		"Name":   changeset.NewBox().Ops(changeset.Nullable).Size(40),
    		"Age":    changeset.NewBox().Ops(changeset.Nullable),
    		"Detail": changeset.NewBox().JSONField(),
    	}
    }
    
    // use for request and response on creating user
    type UserCreateDTO struct {
        Id uint32 `json:"id,omitempty"`
        Name string `json:"name"`
        Age uint8 `json:"age"`
    }
    ```
    
- In user_service.go
    ```go
    // user_service.go
    package userservice

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
    	
    	// because CastValues get first param is schema interface
    	// we implement before, then it can read all boxes within changeset
    	// to validate or require fields efficient
    	// one more things, on detail field use json
    	// then goecto cast automatically into mysql json type [string]
    	// when you use cast it will automatically cast struct in declare of implementing schema interface ("Detail": changeset.NewBox().JSONField())
    	// when we query otherwise, it will automatically cast mysql json type into struct above (here is Detail struct)
    	userChangeset := changeset.CastValues(userEntity, map[string]interface{}{
    		// Id => not need to cast, because it will have then save userChangeset into repo
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
    	// fmt.Println('User id saved: ', userEntity.Id)
    	return userEntity, nil
    }
    ```

## 2. Notes:
- Library not auto migrate
- It is one data mapper, you should use convention exactly as name in struct field
- Example: if struct field User is (Id, Name, Age), then you first should create column in names (Id, Name, Age)
    
## 3. Release
- This library is one beta version on implement stable version from my oragnization. I will release and test bugs as best as posibble


