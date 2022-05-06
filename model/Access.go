package model

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Access package initialized")
}

type Access struct {
	Ip       string
	Path     string
	AccessAt time.Time
}

func (a Access) String() string {
	return "hello world"
}
