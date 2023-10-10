package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func Migrate(cfg *Config, con *sql.DB) error {

	var err error

	// copy sql file to container /
	cmd := fmt.Sprintf("docker cp %s %s:/", cfg.SQL_FILE, cfg.CONTAINER)
	err = exec.Command("/bin/sh", "-c", cmd).Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s copied\n", cfg.SQL_FILE)

	// note: `docker exec -i container_name sh -c "mysql -uroot -p12345 sec-kill < /babytun.sql"`
	cmd = fmt.Sprintf(`docker exec -i %s sh -c "mysql -u%s -p%s %s < /%s"`,
		cfg.CONTAINER,
		cfg.DB_USER,
		cfg.DB_PASSWORD,
		cfg.DB_NAME,
		cfg.SQL_FILE,
	)
	err = exec.Command("/bin/sh", "-c", cmd).Start()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	fmt.Printf("starting load %s to %s ...\n", cfg.SQL_FILE, cfg.DB_NAME)

	// 抓 table 數量來比對 load sql 的進度
	q := fmt.Sprintf(`SELECT COUNT(table_name)
	FROM information_schema.tables
	WHERE table_schema = '%s';`, cfg.DB_NAME)

	var count int64
	target, err := strconv.ParseInt(cfg.DB_TABLES, 10, 64)
	if err != nil {
		return err
	}

	for count != target {
		time.Sleep(time.Second * 2)
		row := con.QueryRow(q)
		if err := row.Scan(&count); err != nil {
			log.Fatal(err)
		}
		fmt.Println("tables to be ready: ", target-count)
	}

	fmt.Println("load table finish!")

	return nil
}
