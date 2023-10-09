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

	determineString := "MySQL Community Server - GPL."

	cmd := fmt.Sprintf(`docker run -dit -p %s:3306 --name %s -e MYSQL_ROOT_PASSWORD=%s -e MYSQL_DATABASE=%s mysql:latest`,
		cfg.DB_PORT,
		cfg.CONTAINER,
		cfg.DB_PASSWORD,
		cfg.DB_NAME,
	)

	err = exec.Command("/bin/sh", "-c", cmd).Run()
	if err != nil {
		log.Fatal(err)
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

		if strings.Contains(string(result), determineString) {
			logFlag = false
		}
	}

	fmt.Printf("%s create success\n", cfg.CONTAINER)

	return nil
}
