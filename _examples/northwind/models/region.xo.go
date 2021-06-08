package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Region represents a row from 'public.region'.
type Region struct {
	RegionID          int16  `json:"region_id"`          // region_id
	RegionDescription string `json:"region_description"` // region_description
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Region exists in the database.
func (r *Region) Exists() bool {
	return r._exists
}

// Deleted returns true when the Region has been marked for deletion from
// the database.
func (r *Region) Deleted() bool {
	return r._deleted
}

// Insert inserts the Region to the database.
func (r *Region) Insert(ctx context.Context, db DB) error {
	switch {
	case r._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case r._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (basic)
	const sqlstr = `INSERT INTO public.region (` +
		`region_id, region_description` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	logf(sqlstr, r.RegionID, r.RegionDescription)
	if err := db.QueryRowContext(ctx, sqlstr, r.RegionID, r.RegionDescription).Scan(&r.RegionID); err != nil {
		return logerror(err)
	}
	// set exists
	r._exists = true
	return nil
}

// Update updates a Region in the database.
func (r *Region) Update(ctx context.Context, db DB) error {
	switch {
	case !r._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case r._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.region SET (` +
		`region_description` +
		`) = ( ` +
		`$1` +
		`) WHERE region_id = $2`
	// run
	logf(sqlstr, r.RegionDescription, r.RegionID)
	if _, err := db.ExecContext(ctx, sqlstr, r.RegionDescription, r.RegionID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Region to the database.
func (r *Region) Save(ctx context.Context, db DB) error {
	if r.Exists() {
		return r.Update(ctx, db)
	}
	return r.Insert(ctx, db)
}

// Upsert performs an upsert for Region.
//
// NOTE: PostgreSQL 9.5+ only
func (r *Region) Upsert(ctx context.Context, db DB) error {
	switch {
	case r._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.region (` +
		`region_id, region_description` +
		`) VALUES (` +
		`$1, $2` +
		`) ON CONFLICT (region_id) DO UPDATE SET (` +
		`region_id, region_description` +
		`) = (` +
		`EXCLUDED.region_id, EXCLUDED.region_description` +
		`)`
	// run
	logf(sqlstr, r.RegionID, r.RegionDescription)
	if _, err := db.ExecContext(ctx, sqlstr, r.RegionID, r.RegionDescription); err != nil {
		return err
	}
	// set exists
	r._exists = true
	return nil
}

// Delete deletes the Region from the database.
func (r *Region) Delete(ctx context.Context, db DB) error {
	switch {
	case !r._exists: // doesn't exist
		return nil
	case r._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.region WHERE region_id = $1`
	// run
	logf(sqlstr, r.RegionID)
	if _, err := db.ExecContext(ctx, sqlstr, r.RegionID); err != nil {
		return logerror(err)
	}
	// set deleted
	r._deleted = true
	return nil
}

// RegionByRegionID retrieves a row from 'public.region' as a Region.
//
// Generated from index 'region_pkey'.
func RegionByRegionID(ctx context.Context, db DB, regionID int16) (*Region, error) {
	// query
	const sqlstr = `SELECT ` +
		`region_id, region_description ` +
		`FROM public.region ` +
		`WHERE region_id = $1`
	// run
	logf(sqlstr, regionID)
	r := Region{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, regionID).Scan(&r.RegionID, &r.RegionDescription); err != nil {
		return nil, logerror(err)
	}
	return &r, nil
}
