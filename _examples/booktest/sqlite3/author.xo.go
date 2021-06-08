package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Author represents a row from 'authors'.
type Author struct {
	AuthorID int    `json:"author_id"` // author_id
	Name     string `json:"name"`      // name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Author exists in the database.
func (a *Author) Exists() bool {
	return a._exists
}

// Deleted returns true when the Author has been marked for deletion from
// the database.
func (a *Author) Deleted() bool {
	return a._deleted
}

// Insert inserts the Author to the database.
func (a *Author) Insert(ctx context.Context, db DB) error {
	switch {
	case a._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case a._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO authors (` +
		`name` +
		`) VALUES (` +
		`?` +
		`)`
	// run
	logf(sqlstr, a.Name)
	res, err := db.ExecContext(ctx, sqlstr, a.Name)
	if err != nil {
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	} // set primary key
	a.AuthorID = int(id)
	// set exists
	a._exists = true
	return nil
}

// Update updates a Author in the database.
func (a *Author) Update(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case a._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE authors SET ` +
		`name = ?` +
		` WHERE author_id = ?`
	// run
	logf(sqlstr, a.Name, a.AuthorID)
	if _, err := db.ExecContext(ctx, sqlstr, a.Name, a.AuthorID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Author to the database.
func (a *Author) Save(ctx context.Context, db DB) error {
	if a.Exists() {
		return a.Update(ctx, db)
	}
	return a.Insert(ctx, db)
}

// Delete deletes the Author from the database.
func (a *Author) Delete(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return nil
	case a._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM authors WHERE author_id = ?`
	// run
	logf(sqlstr, a.AuthorID)
	if _, err := db.ExecContext(ctx, sqlstr, a.AuthorID); err != nil {
		return logerror(err)
	}
	// set deleted
	a._deleted = true
	return nil
}

// AuthorByAuthorID retrieves a row from 'authors' as a Author.
//
// Generated from index 'authors_author_id_pkey'.
func AuthorByAuthorID(ctx context.Context, db DB, authorID int) (*Author, error) {
	// query
	const sqlstr = `SELECT ` +
		`author_id, name ` +
		`FROM authors ` +
		`WHERE author_id = ?`
	// run
	logf(sqlstr, authorID)
	a := Author{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, authorID).Scan(&a.AuthorID, &a.Name); err != nil {
		return nil, logerror(err)
	}
	return &a, nil
}

// AuthorsByName retrieves a row from 'authors' as a Author.
//
// Generated from index 'authors_name_idx'.
func AuthorsByName(ctx context.Context, db DB, name string) ([]*Author, error) {
	// query
	const sqlstr = `SELECT ` +
		`author_id, name ` +
		`FROM authors ` +
		`WHERE name = ?`
	// run
	logf(sqlstr, name)
	rows, err := db.QueryContext(ctx, sqlstr, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Author
	for rows.Next() {
		a := Author{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&a.AuthorID, &a.Name); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &a)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
