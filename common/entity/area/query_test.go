package area

import (
	"common/database"
	"testing"
)

func TestQuery(t *testing.T) {
	pool := database.NewDBPool()
	dsn := "guiwoo:guiwoo@tcp(127.0.0.1:3306)/workout_area?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := pool.ConnectArea(dsn, "debug", 10, 10, 10)
	if err != nil {
		t.Error(err)
	}
	if pool[database.DBAREA] == nil {
		t.Error("connect area failure")
	}
	var siggArea []SiggArea
	if err = db.Model(&SiggArea{}).Find(&siggArea).Error; err != nil {
		t.Error(err)
	}
	t.Log(siggArea)
}
