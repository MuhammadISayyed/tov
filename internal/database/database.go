/*
Using sqlite with go
1. Use a migration tool. Top options are goose and db mate.
2. If you're using goose, run the goose create command to create a new migration.
3. Run goose up to execute the migration.
4. To connect to the database, do what's in the Connect func in the database.go file.
5. Create a model directory (already created) and create a go file for each table.
6. To group queries create a sql directory within the model directory with a file for each table.
7. Create a struct for each table that requires a DB *sql.DB. This will be the receiver for all of our methods regarding the table.
8. All of the above steps you can do following this tutorial: https://www.youtube.com/watch?v=qmXrYZavQIc
*/

package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./tov.db")
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}