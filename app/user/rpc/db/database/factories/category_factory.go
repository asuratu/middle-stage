package factories

import (
	"github.com/bxcodec/faker/v3"
	"middle/app/user/rpc/model"
	"middle/pkg/helpers"
)

func MakeCategories(count int) []model.Categories {

	var objs []model.Categories

	// 设置唯一性，如 Category 模型的某个字段需要唯一，即可取消注释
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		categoryModel := model.Categories{
			Name:        faker.Username(),
			Description: helpers.StringToNullString(faker.Sentence()),
		}
		objs = append(objs, categoryModel)
	}

	return objs
}
