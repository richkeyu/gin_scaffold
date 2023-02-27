package keys

import "fmt"

func ProjectCacheKey() string {
	return fmt.Sprint("pay:project:all")
}

func MerchantCacheKey() string {
	return fmt.Sprint("pay:merchant:all")
}

func UserCacheKey(id int64) string {
	return fmt.Sprintf("pay:user:%d", id)
}
