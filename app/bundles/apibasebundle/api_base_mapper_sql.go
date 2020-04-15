package apibasebundle

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"github.com/mintance/go-uniqid"
	"os"
	"strings"
	"time"
)

// ParametrosBundleSQL define a SQL mapper
type ApiBaseBundleSQL struct {
	db *gorm.DB
}

// NewParametrosSQLMapper instance
func NewApiBaseSQLMapper(db *gorm.DB) *ApiBaseBundleSQL {
	return &ApiBaseBundleSQL{
		db: db,
	}
}

func (m *ApiBaseBundleSQL) Alive() float64 {
	var response float64
	var s []string
	start := time.Now()
	for i := 0; i < 1; i++ {
		id := uniqid.New(uniqid.Params{"test", true})
		//fmt.Printf("%x", id)
		data := []byte(id)
		h := md5.New()
		charid := strings.ToUpper(hex.EncodeToString(h.Sum(data)))
		s = append(s, charid)
	}
	ms := time.Since(start)
	response = ms.Seconds()
	return response
}

func (m *ApiBaseBundleSQL) Env() interface{} {
	x := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		x[pair[0]] = pair[1]
	}
	return x
}


