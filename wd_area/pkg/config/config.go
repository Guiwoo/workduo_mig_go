package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Area struct {
	Dsn      string    `yaml:"Dsn"`
	Listen   string    `yaml:"Listen"`
	Log      LogConfig `yaml:"Log"`
	DataPath Data      `yaml:"Data"`
}
type LogConfig struct {
	Path     string        `yaml:"Path"`
	FileName string        `yaml:"FileName"`
	Level    zerolog.Level `yaml:"Level"`
}
type Data struct {
	Path string `yaml:"Path"`
}

func (a *Area) ParseConfig(filepath string) {
	fs, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fs, a)
	if err != nil {
		panic(err)
	}
}

func (a *Area) PrintConfig() {
	sb := strings.Builder{}

	sb.WriteString("\n---- Config ---\n")
	sb.WriteString(fmt.Sprintf("DSN : %s\n", a.Dsn))
	sb.WriteString(fmt.Sprintf("Listen : %s\n", a.Listen))
	sb.WriteString(fmt.Sprintf("Data :%+v\n", a.DataPath))
	sb.WriteString(fmt.Sprintf("Log : %+v\n", a.Log))
	sb.WriteString("---- Config ---\n")

	fmt.Println(sb.String())
}
