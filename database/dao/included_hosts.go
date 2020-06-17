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

// GetAllIncludedHosts is a function to get a slice of record(s) from included_hosts table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllIncludedHosts(ctx context.Context, page, pagesize int64, order string) (includedhosts []*model.IncludedHosts, totalRows int, err error) {

	includedhosts = []*model.IncludedHosts{}

	includedhostsOrm := DB.Model(&model.IncludedHosts{})
	includedhostsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		includedhostsOrm = includedhostsOrm.Offset(offset).Limit(pagesize)
	} else {
		includedhostsOrm = includedhostsOrm.Limit(pagesize)
	}

	if order != "" {
		includedhostsOrm = includedhostsOrm.Order(order)
	}

	if err = includedhostsOrm.Find(&includedhosts).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return includedhosts, totalRows, nil
}

// GetIncludedHosts is a function to get a single record from the included_hosts table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetIncludedHosts(ctx context.Context, argId string) (record *model.IncludedHosts, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddIncludedHosts is a function to add a single record to included_hosts table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddIncludedHosts(ctx context.Context, record *model.IncludedHosts) (result *model.IncludedHosts, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateIncludedHosts is a function to update a single record from included_hosts table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateIncludedHosts(ctx context.Context, argId string, updated *model.IncludedHosts) (result *model.IncludedHosts, RowsAffected int64, err error) {

	result = &model.IncludedHosts{}
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

// DeleteIncludedHosts is a function to delete a single record from included_hosts table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteIncludedHosts(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.IncludedHosts{}
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
