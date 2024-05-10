package createuser

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

type UserCreateDTO struct {
	Id     uint32 `json:"id,omitempty"`
	Name   string `json:"name"`
	Age    uint8  `json:"age"`
	Detail struct {
		Address    string `json:"address"`
		AvatarLink string `json:"avatar_link"`
	} `json:"detail,omitempty"`
}

func (u *User) Validators() map[string]*changeset.Box {
	return map[string]*changeset.Box{
		"Id":     changeset.NewBox().Ops(changeset.AI, changeset.NotNullable),
		"Name":   changeset.NewBox().Ops(changeset.Nullable).Size(40),
		"Age":    changeset.NewBox().Ops(changeset.Nullable),
		"Detail": changeset.NewBox().JSONField(),
	}
}
