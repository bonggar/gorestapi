package database

import (
	"testing"

	"github.com/bonggar/gorestapi/model"
	"github.com/gin-gonic/gin"
)

func TestDBInitialization(t *testing.T) {
	gin.SetMode(gin.TestMode)
	SQLiteDBConnect()
	if db != nil {
		t.Logf("SQLite DB Connection Success")
		if db.HasTable(model.User{}) {
			t.Logf("DB Migration Success")
		} else {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}
