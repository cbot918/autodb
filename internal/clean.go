package internal

import (
	"database/sql"
	"fmt"
	"os/exec"
)

func Clean(cfg *Config, con *sql.DB) error {
	var err error
	var cmd string
	defer con.Close()

	cmd = fmt.Sprintf("docker container stop %s ", cfg.CONTAINER)
	err = exec.Command("/bin/sh", "-c", cmd).Run()
	if err != nil {
		return err
	}

	cmd = fmt.Sprintf("docker container rm %s ", cfg.CONTAINER)
	err = exec.Command("/bin/sh", "-c", cmd).Run()
	if err != nil {
		return err
	}

	fmt.Println("clean finish")

	return nil
}
