package service

import (
	"fmt"
	"github.com/google/uuid"
)

func IDGenerator(target string) string {
	uuid, _ := uuid.NewUUID()
	return fmt.Sprintf("%s_%s", target, uuid.String())
}
