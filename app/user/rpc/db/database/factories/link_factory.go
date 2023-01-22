package factories

import (
	"github.com/bxcodec/faker/v3"
	"middle/app/user/rpc/model"
	"middle/pkg/helpers"
)

func MakeLinks(times int) []model.Links {

	var objs []model.Links

	for i := 0; i < times; i++ {
		m := model.Links{
			Name: faker.Username(),
			Url:  helpers.StringToNullString(faker.URL()),
		}
		objs = append(objs, m)
	}

	return objs
}
