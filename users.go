package freee

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathUsers = "users"
)

type Me struct {
	User User `json:"user"`
}

type User struct {
	// ユーザーID
	ID int64 `json:"id"`
	// メールアドレス
	Email string `json:"email"`
	// 表示ユーザー名
	DisplayName *string `json:"display_name,omitempty"`
	// 名
	FirstName *string `json:"first_name,omitempty"`
	// 姓
	LastName *string `json:"last_name,omitempty"`
	// 名（カナ）
	FirstNameKana *string `json:"first_name_kana,omitempty"`
	// 姓（カナ）
	LastNameKana *string       `json:"last_name_kana,omitempty"`
	Companies    []UserCompany `json:"companies,omitempty"`
}

type UserCompany struct {
	// 事業所ID
	ID int64 `json:"id"`
	// 表示名
	DisplayName string `json:"display_name"`
	// ユーザーの権限
	Role string `json:"role"`
	// カスタム権限（true: 使用する、false: 使用しない）
	UseCustomRole bool `json:"use_custom_role"`
}

type GetUsersMeOpts struct {
	Companies bool `url:"companies,omitempty"`
}

func (c *Client) GetUsersMe(
	ctx context.Context, oauth2Token *oauth2.Token,
	opts GetUsersMeOpts,
) (*Me, *oauth2.Token, error) {
	var result Me

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	oauth2Token, err = c.call(ctx, path.Join(APIPathUsers, "me"), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}
