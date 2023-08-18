package scheduler

import (
	"area/config"
	"area/interfaces/data"
	"common/database"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

/**
{
	"type": "Feature",
	"properties": {
	"adm_nm": "서울특별시 종로구 사직동",
"adm_cd": "1101053",
"adm_cd2": "1111053000",
"sgg": "11110",
"sido": "11",
"sidonm": "서울특별시",
"temp": "종로구 사직동",
"sggnm": "종로구",
"adm_cd8": "11010530" },
*/

func Test_Something01(t *testing.T) {
	t.Log("Test_Something01")

	file, err := os.ReadFile("../../interfaces/data/HangJeongDong_ver20230701.json")

	if err != nil {
		t.Errorf("file read error : %s", err)
	}

	var result data.HangJeongDong
	err = json.Unmarshal(file, &result)
	if err != nil {
		t.Errorf("json unmarshal error : %s", err)
	}

	t.Logf("result : %v ,total : %d", result.Type, len(result.Features))
}

func Test_Shceduler(t *testing.T) {
	pool := database.NewDBPool()
	dsn := "guiwoo:guiwoo@tcp(127.0.0.1:3306)/workout_area?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := pool.ConnectArea(dsn, "debug", 10, 10, 10)
	if err != nil {
		t.Error(err)
	}

	a := &AreaScheduler{
		path: "../../interfaces/data/HangJeongDong_ver20230701.json",
		db:   db,
	}

	err = a.Process()
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func Test_Process(t *testing.T) {
	pool := database.NewDBPool()
	dsn := "guiwoo:guiwoo@tcp(127.0.0.1:3306)/workout_area?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := pool.ConnectArea(dsn, "debug", 10, 10, 10)
	if err != nil {
		t.Error(err)
	}

	f, err := os.ReadFile("../../config/config.yaml")
	if err != nil {
		t.Error(err)
	}
	var cfg config.AreaConfig
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		t.Error(err)
	}

	a := NewAreaScheduler(db, &cfg, logrus.StandardLogger())

	err = <-a.Run(context.Background())

	if err != nil {
		t.Error(err)
	}

}
