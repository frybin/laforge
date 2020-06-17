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

// GetAllIncludedSteps is a function to get a slice of record(s) from included_steps table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllIncludedSteps(ctx context.Context, page, pagesize int64, order string) (includedsteps []*model.IncludedSteps, totalRows int, err error) {

	includedsteps = []*model.IncludedSteps{}

	includedstepsOrm := DB.Model(&model.IncludedSteps{})
	includedstepsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		includedstepsOrm = includedstepsOrm.Offset(offset).Limit(pagesize)
	} else {
		includedstepsOrm = includedstepsOrm.Limit(pagesize)
	}

	if order != "" {
		includedstepsOrm = includedstepsOrm.Order(order)
	}

	if err = includedstepsOrm.Find(&includedsteps).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return includedsteps, totalRows, nil
}

// GetIncludedSteps is a function to get a single record from the included_steps table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetIncludedSteps(ctx context.Context, argId string) (record *model.IncludedSteps, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddIncludedSteps is a function to add a single record to included_steps table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddIncludedSteps(ctx context.Context, record *model.IncludedSteps) (result *model.IncludedSteps, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateIncludedSteps is a function to update a single record from included_steps table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateIncludedSteps(ctx context.Context, argId string, updated *model.IncludedSteps) (result *model.IncludedSteps, RowsAffected int64, err error) {

	result = &model.IncludedSteps{}
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

// DeleteIncludedSteps is a function to delete a single record from included_steps table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteIncludedSteps(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.IncludedSteps{}
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
