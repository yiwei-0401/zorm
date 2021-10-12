package session

import (
	"database/sql"
	"strings"
	"zorm/log"
)

type Session struct {
	db *sql.DB
	sql strings.Builder
	sqlVals []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db :db}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVals = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVals = append(s.sqlVals, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVals)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVals...); err != nil {
		log.Error(err)
	}
	return
}

//QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVals)
	return s.DB().QueryRow(s.sql.String(), s.sqlVals...)
}

//QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVals)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVals...); err != nil {
		log.Error(err)
	}
	return
}

