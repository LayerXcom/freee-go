package freee

import (
	"fmt"
	"testing"
)

var testJsonData = map[string]int{
	"{\"error\":\"invalid_grant\",\"error_description\":\"指定された認可グラントは不正か、有効期限切れか、無効か、リダイレクトURIが異なるか、もしくは別のクライアントに適用されています。\"}":                                                                                                                                                                          1,
	"{\"status_code\":404,\"errors\":[{\"type\":\"status\",\"messages\":[\"リソースが見つかりません。\"]},{\"type\":\"validation\",\"messages\":[\"既に削除された、あるいは存在しない取引先です。\"]}]}":                                                                                                                                     2,
	"{\"status_code\":404,\"errors\":[{\"type\":\"status\",\"messages\":[\"リソースが見つかりません。\"]},{\"type\":\"error\",\"messages\":[\"存在しないか既に削除された取引です。\"]}]}":                                                                                                                                               2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"指定された partner_id は存在しません。\"]}]}":                                                                                                                                     2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"名前（通称）「yyyyy」はすでに存在します。\"]}]}":                                                                                                                                       2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"取引先「xxxxx」は既に存在します\"]}]}":                                                                                                                                            2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"Codeを入力してください。\"]}]}":                                                                                                                                                2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"今月アップロードできるファイル数の上限を超えています。ミニマムプラン、もしくはベーシックプランに契約してファイルボックスの枚数制限を解除しましょう。\"]}]}":                                                                                    2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"品目「AAAAAAAAAA」と勘定科目「BBBBBBBBBBB」は勘定科目の設定で紐付けされていません。\"]}]}":                                                                                                          2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"月締めがされています。2021-06-01以降の日付を入力してください。\"]}]}":                                                                                                                          2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"現在選択している会計年度の期首日以前の取引を登録することができません。発生日を期首日以降にするか、年度締めを巻き戻してください。\"]}]}":                                                                                              2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"Details 存在しない tag_id　が含まれています。\"]}]}":                                                                                                                                2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"すでに同じ名前の項目 が存在しています。\"]}]}":                                                                                                                                          2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"取引の承認状態を確認してください。\"]}]}":                                                                                                                                             2,
	"{\"status_code\":400,\"errors\":[{\"type\":\"status\",\"messages\":[\"不正なリクエストです。\"]},{\"type\":\"validation\",\"messages\":[\"貸借が一致していません。\",\"貸借が一致していません。\",\"決済に明細が登録されている取引は金額を変更することができません。変更する場合は、該当の明細の取引登録を解除してください\",\"振込が成功している取引の金額、取引先は変更できません。新しい取引を作成してください。\",\"取引単位は不正な値です。\"]}]}": 6,
	"{\"status_code\":401,\"errors\":[{\"type\":\"status\",\"messages\":[\"アクセス権限がありません。\",\"company_admin\",\"api/v1/segment_tags\",\"index_division_3\",\"このAPIにアクセスしたい場合、事業所の管理者にご確認ください。\"],\"codes\":[\"user_do_not_have_permission\"]}]}":                                                          5,
	"{\"message\": \"ログインをして下さい\",\"code\": \"invalid_access_token\"}": 1,
	"{\"messages\":[\"部門名はすでに存在します。\"]}":                               1,
	"{\"aaaaaa\":12345}": 0,
	"{}":                 0,
}

var testInvalidJsonData = []string{
	"",                              // brank
	"aaaaaaaa",                      // jsonではない
	"{\"error_description\":12345}", // error_descriptionにエラー文字列が入っていない
	"{\"errors\":\"エラーです\"}",        // errorsが配列ではなくエラー文字列が入っている
}

func TestExtractFreeeErrorMessage(t *testing.T) {
	t.Parallel()
	i := 0
	for data, num := range testJsonData {
		i := i
		data := data
		num := num
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			messages, err := ExtractFreeeErrorMessage(data)

			if err != nil {
				t.Fatal(err)
			}
			if num != len(messages) {
				t.Fatalf("unmatch message nums : %d, %d", num, len(messages))
			}
			t.Logf("message len = %d\n", len(messages))
		})
		i++
	}
}

func TestExtractFreeeInvalidErrorMessage(t *testing.T) {
	t.Parallel()
	for i, data := range testInvalidJsonData {
		i := i
		data := data
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			messages, err := ExtractFreeeErrorMessage(data)

			if err == nil {
				t.Fatal("not error")
			}
			if messages != nil {
				t.Fatal("valid parsed message")
			}
			t.Logf("error messages: %s\n", err.Error())
		})
	}
}

func TestErrorMessages(t *testing.T) {
	i := 0
	t.Parallel()
	for data, num := range testJsonData {
		i := i
		data := data
		num := num
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			err := &Error{
				StatusCode:              400,
				RawError:                data,
				IsAuthorizationRequired: false,
			}
			messages := err.Messages()
			if num != len(messages) {
				t.Fatalf("unmatch message nums : %d, %d", num, len(messages))
			}
			t.Logf("message len = %d\n", len(messages))
		})
		i++
	}
}

func TestInvalidErrorMessages(t *testing.T) {
	t.Parallel()
	for i, data := range testInvalidJsonData {
		i := i
		data := data
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			err := &Error{
				StatusCode:              400,
				RawError:                data,
				IsAuthorizationRequired: false,
			}
			messages := err.Messages()
			if len(messages) > 0 {
				t.Fatalf("unmatch message nums : %d", len(messages))
			}
			t.Logf("message len = %d\n", len(messages))
		})
	}
}
