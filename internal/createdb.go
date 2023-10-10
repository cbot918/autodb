package internal

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func CreateDB(cfg *Config) error {
	var err error
	var cmd string
	determineMySQL := "MySQL Community Server - GPL."
	determinePostgres := "database system is ready to accept connections"

	if cfg.DB_DRIVER == "mysql" {
		cmd = fmt.Sprintf(`docker run -dit -p %s:3306 --name %s -e MYSQL_ROOT_PASSWORD=%s -e MYSQL_DATABASE=%s mysql:latest`,
			cfg.DB_PORT,
			cfg.CONTAINER,
			cfg.DB_PASSWORD,
			cfg.DB_NAME,
		)
	} else if cfg.DB_DRIVER == "postgres" {
		fmt.Println("in createdb postgres")
		cmd = fmt.Sprintf(`docker run -dit -p %s:5432 --name %s -e POSTGRES_PASSWORD=%s -e POSTGRES_DB=%s postgres`,
			cfg.DB_PORT,
			cfg.CONTAINER,
			cfg.DB_PASSWORD,
			cfg.DB_NAME,
		)
	} else {
		return fmt.Errorf("driver " + cfg.DB_DRIVER + " not support, check .env")
	}
	err = exec.Command("/bin/sh", "-c", cmd).Run()
	if err != nil {
		return err
	}

	logFlag := true
	for logFlag {
		fmt.Printf("waiting %s container start ...\n", cfg.CONTAINER)
		time.Sleep(2 * time.Second)
		cmd1 := fmt.Sprintf("docker logs %s", cfg.CONTAINER)
		result, err1 := exec.Command("/bin/sh", "-c", cmd1).Output()
		if err != nil {
			log.Fatal(err1)
		}
		if cfg.DB_DRIVER == "mysql" && strings.Contains(string(result), determineMySQL) {
			logFlag = false
		}
		if cfg.DB_DRIVER == "postgres" && strings.Contains(string(result), determinePostgres) {
			logFlag = false
		}

	}

	fmt.Printf("%s create success\n", cfg.CONTAINER)

	return nil
}
