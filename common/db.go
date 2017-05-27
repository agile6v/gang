package common

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type DB struct {
    Dsn string  // dsn
}

func (db *DB) GetAllTasks() (map[string][]*Task, error) {
	database, err := sql.Open("mysql", db.Dsn)
	if err != nil {
        fmt.Printf("Failed to open mysql(%s): %s",
            db.Dsn, err.Error())
		return nil, err
	}
	defer database.Close()

    task_map := make(map[string][]*Task, 0)
    sql := "SELECT h.host,h.group_id,h.task_id,t.name,t.runner," +
        "t.command,t.args,t.status FROM host as h LEFT JOIN task " +
        "as t ON (h.task_id = t.id)"
    rows, err := database.Query(sql)
	if err != nil {
		return nil, err
	}

    for rows.Next() {
        var task_id int64
        var status bool
        var group_id int
        var name, runner, command, args, host string

        if err := rows.Scan(&host, &task_id, &group_id,
            &name, &runner, &command, &args, &status); err != nil {
            fmt.Println("mysql scan error: ", err)
            continue
        }

        t := &Task{
            Id: task_id,
            Name: name,
            Status: status,
            Command: command,
            Args: args,
            Runner: runner,
        }

        task_map[host] = append(task_map[host], t)
    }

    return task_map, nil
}
