package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	proxy "github.com/shogo82148/go-sql-proxy"
)

const (
	logElasped = time.Second
	traceSQL   = false
)

func setMysqlEtc(db *sql.DB) {
	db.SetConnMaxLifetime(time.Second * 65) // poolするには65sとか
	db.SetMaxOpenConns(100)                 // mysqlのmax-connectionと揃える?
	db.SetMaxIdleConns(100)
}

// registerMysqlProxy("mysql-proxy")
// sql.Open("mysql-proxy", "dns")

func registerMysqlProxy(name string) {
	sql.Register(name, proxy.NewProxyContext(&mysql.MySQLDriver{}, &proxy.HooksContext{
		PreExec: func(ctx context.Context, stmt *proxy.Stmt, args []driver.NamedValue) (interface{}, error) {
			return time.Now(), nil
		},
		PostExec: func(ctx context.Context, start interface{}, stmt *proxy.Stmt, args []driver.NamedValue, result driver.Result, err error) error {
			elapsed := time.Since(start.(time.Time))
			switch {
			case err != nil || elapsed > time.Second:
				log.Printf("[SQL Trace] Err: %s, Query: %s; args: %v, elapsed: %s", err, stmt.QueryString, args, elapsed)
			case traceSQL:
				log.Printf("[SQL Trace] Query: %s; args: %v, elapsed: %s", stmt.QueryString, args, elapsed)
			}
			return nil
		},
		PreQuery: func(ctx context.Context, stmt *proxy.Stmt, args []driver.NamedValue) (interface{}, error) {
			return time.Now(), nil
		},
		PostQuery: func(ctx context.Context, start interface{}, stmt *proxy.Stmt, args []driver.NamedValue, rows driver.Rows, err error) error {
			elapsed := time.Since(start.(time.Time))
			switch {
			case err != nil || elapsed > time.Second:
				log.Printf("[SQL Trace] Err: %s, Query: %s; args: %v, elapsed: %s", err, stmt.QueryString, args, elapsed)
			case traceSQL:
				log.Printf("[SQL Trace] Query: %s; args: %v, elapsed: %s", stmt.QueryString, args, elapsed)
			}
			return nil
		},
	}))
}
