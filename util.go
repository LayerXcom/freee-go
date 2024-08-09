package freee

import (
	"fmt"
	"net/url"
)

func SetCompanyID(v *url.Values, companyID int64) {
	v.Set("company_id", fmt.Sprintf("%d", companyID))
}
