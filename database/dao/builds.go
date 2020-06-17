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

// GetAllBuilds is a function to get a slice of record(s) from builds table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBuilds(ctx context.Context, page, pagesize int64, order string) (builds []*model.Builds, totalRows int, err error) {

	builds = []*model.Builds{}

	buildsOrm := DB.Model(&model.Builds{})
	buildsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		buildsOrm = buildsOrm.Offset(offset).Limit(pagesize)
	} else {
		buildsOrm = buildsOrm.Limit(pagesize)
	}

	if order != "" {
		buildsOrm = buildsOrm.Order(order)
	}

	if err = buildsOrm.Find(&builds).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return builds, totalRows, nil
}

// GetBuilds is a function to get a single record from the builds table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetBuilds(ctx context.Context, argId string) (record *model.Builds, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddBuilds is a function to add a single record to builds table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddBuilds(ctx context.Context, record *model.Builds) (result *model.Builds, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateBuilds is a function to update a single record from builds table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBuilds(ctx context.Context, argId string, updated *model.Builds) (result *model.Builds, RowsAffected int64, err error) {

	result = &model.Builds{}
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

// DeleteBuilds is a function to delete a single record from builds table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBuilds(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.Builds{}
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
