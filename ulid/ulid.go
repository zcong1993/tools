package ulid

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

var (
	entropy = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func genUlid() string {
	return ulid.MustNew(ulid.Now(), entropy).String()
}

// GenUlids gen nums of uuid
func GenUlids(nums int64) []string {
	var resp []string
	for i := int64(0); i < nums; i++ {
		resp = append(resp, genUlid())
	}
	return resp
}
