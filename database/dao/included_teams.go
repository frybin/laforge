package dao

import (
	"context"
	"time"

	"github.com/frybin/laforge/database/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllIncludedTeams is a function to get a slice of record(s) from included_teams table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllIncludedTeams(ctx context.Context, page, pagesize int64, order string) (includedteams []*model.IncludedTeams, totalRows int, err error) {

	includedteams = []*model.IncludedTeams{}

	includedteamsOrm := DB.Model(&model.IncludedTeams{})
	includedteamsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		includedteamsOrm = includedteamsOrm.Offset(offset).Limit(pagesize)
	} else {
		includedteamsOrm = includedteamsOrm.Limit(pagesize)
	}

	if order != "" {
		includedteamsOrm = includedteamsOrm.Order(order)
	}

	if err = includedteamsOrm.Find(&includedteams).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return includedteams, totalRows, nil
}

// GetIncludedTeams is a function to get a single record from the included_teams table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetIncludedTeams(ctx context.Context, argId string) (record *model.IncludedTeams, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddIncludedTeams is a function to add a single record to included_teams table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddIncludedTeams(ctx context.Context, record *model.IncludedTeams) (result *model.IncludedTeams, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateIncludedTeams is a function to update a single record from included_teams table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateIncludedTeams(ctx context.Context, argId string, updated *model.IncludedTeams) (result *model.IncludedTeams, RowsAffected int64, err error) {

	result = &model.IncludedTeams{}
	db := DB.First(result, argId)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteIncludedTeams is a function to delete a single record from included_teams table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteIncludedTeams(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.IncludedTeams{}
	db := DB.First(record, argId)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
