package system

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupMockDB() (sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	directory := postgres.New(postgres.Config{
		Conn: db,
	})
	global.TREND_DB, err = gorm.Open(directory, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return mock, nil
}

func TestSystemAreasServices(t *testing.T) {
	mock, err := SetupMockDB()
	assert.NoError(t, err)

	defer func() {
		sqlDB, err := global.TREND_DB.DB()
		assert.NoError(t, err)
		sqlDB.Close()
	}()

	rows := sqlmock.NewRows([]string{"id", "title", "code", "ancestry", "created_at", "updated_at"}).
		AddRow(1, "Root Area", 100, "", "2021-01-01 00:00:00", "2021-01-01 00:00:00").
		AddRow(2, "Child Area", 101, "100", "2021-01-01 00:00:00", "2021-01-01 00:00:00")

	mock.ExpectQuery(`SELECT \* FROM "areas"`).WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	baseSystemAreas := BaseSystemAreas{}
	data, err := baseSystemAreas.SystemAreasServices(c)
	assert.NoError(t, err)
	assert.Len(t, data, 1)
	assert.Equal(t, "Root Area", data[0].Title)
	assert.Len(t, data[0].Children, 1)
	assert.Equal(t, "Child Area", data[0].Children[0].Title)
	assert.NoError(t, mock.ExpectationsWereMet())
}
