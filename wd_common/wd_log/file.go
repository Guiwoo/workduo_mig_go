package wd_log

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
	"wd_common/wd_datetime"
)

type FileWriter struct {
	log  chan []byte
	file *os.File
	cur  time.Time
	path string
}

func prettier(s string) []byte {
	if !strings.Contains(s, "REQUEST") {
		return []byte(s)
	}
	sb := strings.Builder{}
	data := strings.Split(s, "\\n")
	l := data[len(data)-1]
	data[len(data)-1] = l[:len(l)-2]
	data = append(data, "}")

	for idx, v := range data {
		if idx != 0 && idx != len(data)-1 {
			sb.WriteString(fmt.Sprintf("\t%s\n", v))
		} else {
			sb.WriteString(v + "\n")
		}
	}
	return []byte(strings.ReplaceAll(sb.String(), "\\", ""))
}

func (f *FileWriter) Write(p []byte) (n int, err error) {
	now := time.Now()
	if f.IsRotate(now) {
		if err := f.RotateFile(now); err != nil {
			log.Fatal().Err(err).Msgf("fail to rotate file %+v", f)
		}
	}
	f.log <- prettier(string(p))
	return len(p), nil
}

func (f *FileWriter) WriteLog(msg []byte) (int, error) {
	return f.file.Write(msg)
}

func (f *FileWriter) setFile() error {

	file, err := os.OpenFile(
		fmt.Sprintf("%s.log", f.path),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		return err
	}
	f.file = file
	return nil
}

func (f *FileWriter) IsRotate(now time.Time) bool {
	if f.cur.Add(time.Hour * 24).Before(now) {
		return true
	}
	return false
}

func (f *FileWriter) RotateFile(now time.Time) error {
	if f.file != nil {
		if err := f.file.Close(); err != nil {
			return err
		}
	}
	day := fmt.Sprintf("%d-%d-%d", f.cur.Year(), int(f.cur.Month()), f.cur.Day())
	if err := os.Rename(
		fmt.Sprintf("%s.log", f.path),
		fmt.Sprintf("%s-%s.log", f.path, day)); err != nil {
		return err
	}
	return f.setFile()
}

func (f *FileWriter) initFile() error {
	now := time.Now()

	// 파일을 열고 오류 처리
	file, err := os.OpenFile(fmt.Sprintf("%s.log", f.path), os.O_RDONLY, 0644)
	if err != nil {
		return f.setFile()
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Err(err).Msgf("fail to close file")
		}
	}()

	// 파일 정보 가져오기
	stat, err := file.Stat()
	if err != nil {
		return f.setFile()
	}

	fixTime := wd_datetime.ParseToDay(stat.ModTime())
	f.cur = fixTime

	// 파일 로테이션 필요한지 확인하고 수행
	if f.IsRotate(now) {
		return f.RotateFile(now)
	}

	return f.setFile()
}

func NewFileWriter(path, name string) *FileWriter {
	f := FileWriter{
		log:  make(chan []byte, 1024),
		cur:  wd_datetime.ParseToDay(time.Now()),
		path: fmt.Sprintf("%s/%s", path, name),
	}

	if err := f.initFile(); err != nil {
		log.Fatal().Err(err).Msgf("fail to create or open file %+v", f.path)
	}

	go func() {
		for v := range f.log {
			if _, err := f.WriteLog(v); err != nil {
				log.Fatal().Err(err).Msgf("file write log fail")
			}
		}
	}()

	return &f
}
