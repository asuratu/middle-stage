package factories

import (
	"github.com/bxcodec/faker/v3"
	"middle/app/user/rpc/model"
)

func MakeTopics(count int) []model.Topics {

	var objs []model.Topics

	for i := 0; i < count; i++ {
		topicModel := model.Topics{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryId: 3,
			UserId:     1,
		}
		objs = append(objs, topicModel)
	}

	return objs
}
