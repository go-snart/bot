package db

import (
	"fmt"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// DBBuilder ensures a RethinkDB DB exists before returning its Term.
//nolint:golint
type DBBuilder struct {
	Name interface{}
	Term *r.Term
}

// BuildDB makes a DBBuilder with a given name.
func BuildDB(name interface{}) *DBBuilder {
	return &DBBuilder{
		Name: name,
		Term: nil,
	}
}

// Build attempts to create the DB, and then returns its Term.
func (d *DBBuilder) Build(qe r.QueryExecutor) r.Term {
	_f := "(*DBBuilder).Build"

	if d.Term != nil {
		return *d.Term
	}

	dbList := []interface{}(nil)

	err := r.DBList().ReadAll(&dbList, qe)
	if err != nil {
		err = fmt.Errorf("readall dblist: %w", err)
		Log.Warn(_f, err)

		return r.Error(err)
	}

	found := false

	for _, dbName := range dbList {
		if dbName == d.Name {
			found = true
			break
		}
	}

	if !found {
		_, err := r.DBCreate(d.Name).RunWrite(qe)
		if err != nil {
			err = fmt.Errorf("runwrite dbcreate: %w", err)
			Log.Warn(_f, err)

			return r.Error(err)
		}
	}

	term := r.DB(d.Name)
	d.Term = &term

	return term
}

// TableBuilder ensures a RethinkDB Table exists before returning its Term.
type TableBuilder struct {
	DB         *DBBuilder
	Name       interface{}
	CreateOpts []r.TableCreateOpts
	Term       *r.Term
}

// BuildTable makes a TableBuilder with a given db, name, and optional creation options.
func BuildTable(db *DBBuilder, name interface{}, co ...r.TableCreateOpts) *TableBuilder {
	return &TableBuilder{
		DB:         db,
		Name:       name,
		CreateOpts: co,
		Term:       nil,
	}
}

// Build attempts to create the Table, and then returns its Term.
func (t *TableBuilder) Build(qe r.QueryExecutor) r.Term {
	_f := "(*TableBuilder).Build"

	if t.Term != nil {
		return *t.Term
	}

	db := t.DB.Build(qe)

	tableList := []interface{}(nil)

	err := db.TableList().ReadAll(&tableList, qe)
	if err != nil {
		err = fmt.Errorf("readall tablelist: %w", err)
		Log.Warn(_f, err)

		return r.Error(err)
	}

	found := false

	for _, tableName := range tableList {
		if tableName == t.Name {
			found = true
			break
		}
	}

	if !found {
		_, err := db.TableCreate(t.Name, t.CreateOpts...).RunWrite(qe)
		if err != nil {
			err = fmt.Errorf("runwrite tablecreate: %w", err)
			Log.Warn(_f, err)

			return r.Error(err)
		}
	}

	term := db.Table(t.Name)
	t.Term = &term

	return term
}
