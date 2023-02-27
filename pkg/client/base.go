package client

import (
	"context"
	"gateway/pkg/db"
	"gateway/pkg/keys"
	"strconv"
	"time"

	"github.com/richkeyu/gocommons/client"
)

const (
	uriGetMerchantAll = "/v1/merchant/all"
	uriGetProjects    = "/v1/projects/projects"
	uriGetUserById    = "/v1/user/get"
)

type BaseCli struct {
	client *client.Request
}

func NewBaseCli() *BaseCli {
	return &BaseCli{
		client: client.NewHttpClientWithConfig("service", "base"),
	}
}

func (b *BaseCli) GetProjectByIdWithCache(ctx context.Context, id int64) (*Project, error) {
	allProject, err := b.GetProjectsWithCache(ctx)
	if err != nil {
		return nil, err
	}
	p := allProject[id]
	return &p, nil
}

func (b *BaseCli) GetProjectsWithCache(ctx context.Context) (map[int64]Project, error) {
	var allProject map[int64]Project
	if projectCache, ok := db.ProjectsCache.Get(keys.ProjectCacheKey()); ok {
		allProject = projectCache.(map[int64]Project)
	} else {
		resp, err := b.client.WithContext(ctx).Get(uriGetProjects)
		if err != nil {
			return nil, err
		}
		var res GetProjects
		_, err = resp.MustParseBody(&res)
		if err != nil {
			return nil, err
		}

		allProject = make(map[int64]Project, len(res.Data))
		for _, project := range res.Data {
			allProject[project.ID] = project
		}
		db.ProjectsCache.Set(keys.ProjectCacheKey(), allProject, time.Minute*10)
	}
	return allProject, nil
}

func (b *BaseCli) GetMerchantAllWithCache(ctx context.Context) (map[int64]Merchant, error) {
	var allMerchant map[int64]Merchant
	if merchantCache, ok := db.MerchantCache.Get(keys.MerchantCacheKey()); ok {
		allMerchant = merchantCache.(map[int64]Merchant)
	} else {
		resp, err := b.client.WithContext(ctx).Get(uriGetMerchantAll)
		if err != nil {
			return nil, err
		}
		var res GetMerchantAllResponse
		_, err = resp.MustParseBody(&res)
		if err != nil {
			return nil, err
		}

		allMerchant = make(map[int64]Merchant, len(res.Data))
		for i := range res.Data {
			allMerchant[res.Data[i].Id] = res.Data[i]
		}
		db.MerchantCache.Set(keys.MerchantCacheKey(), allMerchant, time.Minute*10)
	}
	return allMerchant, nil
}

func (b *BaseCli) GetMerchantByIdWithCache(ctx context.Context, id int64) (*Merchant, error) {
	allMerchant, err := b.GetMerchantAllWithCache(ctx)
	if err != nil {
		return nil, err
	}
	m := allMerchant[id]
	return &m, nil
}

func (b *BaseCli) GetUserByIdWithCache(ctx context.Context, userId int64) (*User, error) {
	var user User
	if userCache, ok := db.UserCache.Get(keys.UserCacheKey(userId)); ok {
		user = userCache.(User)
	} else {
		resp, err := b.client.WithContext(ctx).Get(uriGetUserById, client.Options{
			Query: map[string]interface{}{
				"id": strconv.FormatInt(userId, 10),
			}})
		if err != nil {
			return nil, err
		}
		var res GetUser
		_, err = resp.MustParseBody(&res)
		if err != nil {
			return nil, err
		}
		user = res.User
		db.UserCache.Set(keys.UserCacheKey(userId), user, time.Minute*10)
	}
	return &user, nil
}
