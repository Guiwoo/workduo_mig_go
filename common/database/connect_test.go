package database

import "testing"

func Test_ConnectOnArea(t *testing.T) {
	pool := NewDBPool()

	dsn := "guiwoo:guiwoo@tcp(127.0.0.1:3306)/workout_area?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := pool.ConnectArea(dsn, "debug", 10, 10, 10)
	if err != nil {
		t.Error(err)
	}
	if pool[DBAREA] == nil {
		t.Error("connect area failure")
	}
}
