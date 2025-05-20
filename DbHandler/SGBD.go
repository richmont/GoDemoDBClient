package dbhandler

import "fmt"

type SGBD uint8

const (
	UNDEFINED = iota
	MARIADB
)

func (s SGBD) String() string {
	switch s {
	case UNDEFINED:
		return "undefined"
	case MARIADB:
		return "MariaDB"
	default:
		panic(fmt.Errorf("unknown SGBD"))
	}
}

func (s SGBD) Array() []string {
	return []string{"MariaDB"}
}
