package service

import (
	"Go-000/Week02/dao"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func GetUser(id int) (*dao.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id 不能为: %v", id)
	}
	u, err := dao.SelectUser(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, fmt.Sprintf("找不到 id: %v 的用户", id))
		}
		return nil, errors.Wrap(err, "获取用户失败")
	}
	return u, nil
}
