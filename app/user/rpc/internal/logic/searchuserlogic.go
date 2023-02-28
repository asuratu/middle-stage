package logic

import (
	"context"
	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/pkg/es"
	"middle/app/user/rpc/user"
	"middle/common/xerr"
	"strconv"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLogic) SearchUser(in *user.SearchUserReq) (*user.SearchUserResp, error) {
	// 先查询 es, 拿到 ids 再查询 mysql
	// 构造查询的DSL,
	// DSL的意思是Domain Specific Language，领域特定语言，是一种针对特定领域的语言，比如SQL就是一种针对数据库的DSL
	query := `
	{
		"_source":{
		  "excludes": ["introduction"]
		}, 
		"query": {
		  "bool": {
		    "should": [
		      {
		        "match": {
		          "name": "俊义"
		        }
		      },
		      {
		        "match": {
		          "name": "张青"
		        }
		      }
		    ]
		  }
		},
		"sort": [
		  {
		    "city": {
		      "order": "desc"
		    }
		  }
		]
	}
	`
	rsp, err := l.svcCtx.EsClient.Search(l.ctx, es.UserIndexName, query)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ElasticSearchError), "Failed to search user err : %v , in :%+v", err, in)
	}

	// TODO 再查询 mysql

	simpleUsers := make([]*user.SimpleUser, 0)
	for _, hit := range rsp {
		logx.Infof("======================= hit: %+v", hit)
		var simpleUser user.SimpleUser
		simpleUser.Id, _ = strconv.ParseInt(hit["id"].(string), 10, 64)
		source := hit["source"].(map[string]interface{})
		simpleUser.Name = source["name"].(string)
		simpleUser.City = source["city"].(string)
		simpleUsers = append(simpleUsers, &simpleUser)
	}

	return &user.SearchUserResp{
		User: simpleUsers,
	}, nil
}
