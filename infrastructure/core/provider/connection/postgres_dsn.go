package connection

import (
	"fmt"
	"strings"
)

type PostgresDNS struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  bool
	Timezone string
}

func (i *PostgresDNS) ToString() string {
	var s []string

	if i.Host != "" {
		s = append(s, fmt.Sprintf("host=%s", i.Host))
	} else {
		s = append(s, fmt.Sprintf("host=%s", "localhost"))
	}

	if i.Port != "" {
		s = append(s, fmt.Sprintf("port=%s", i.Port))
	} else {
		s = append(s, fmt.Sprintf("port=%s", i.Port))
	}

	if i.User != "" {
		s = append(s, fmt.Sprintf("user=%s", i.User))
	}

	if i.Password != "" {
		s = append(s, fmt.Sprintf("password=%s", i.Password))
	}

	if i.DBName != "" {
		s = append(s, fmt.Sprintf("dbname=%s", i.DBName))
	}

	if i.SSLMode {
		s = append(s, fmt.Sprintf("sslmode=%s", "require"))
	} else {
		s = append(s, fmt.Sprintf("sslmode=%s", "disable"))
	}

	if i.Timezone != "" {
		s = append(s, fmt.Sprintf("timezone=%s", i.Timezone))
	}

	return strings.Join(s, " ")
}
