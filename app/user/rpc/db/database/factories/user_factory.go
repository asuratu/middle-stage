// Package factories 存放工厂方法
package factories

import (
	"middle/app/user/rpc/model"
	"middle/pkg/helpers"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []model.Users {

	var objs []model.Users

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		m := model.Users{
			Name: faker.Username(),
			// 将 string 转为 sql.NullString
			Email:    helpers.StringToNullString(faker.Email()),
			Phone:    helpers.StringToNullString(helpers.RandomMobile()),
			Password: helpers.StringToNullString("$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe"),
		}
		objs = append(objs, m)
	}

	return objs
}
