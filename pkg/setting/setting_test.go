package setting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetting(t *testing.T) {
	SetUp("../../conf/app.ini")
	t.Logf("%+v\n", DatabaseSetting)

	a := assert.New(t)
	a.Equal(123, 123, "they should be equal")
	a.Equal(DatabaseSetting.Type, "sqlite3", "sqlite3 testing")
	a.Equal(DatabaseSetting.Path, "./data.db", "sqlite3 testing")
}

