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

// GetAllCompetitions is a function to get a slice of record(s) from competitions table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllCompetitions(ctx context.Context, page, pagesize int64, order string) (competitions []*model.Competitions, totalRows int, err error) {

	competitions = []*model.Competitions{}

	competitionsOrm := DB.Model(&model.Competitions{})
	competitionsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		competitionsOrm = competitionsOrm.Offset(offset).Limit(pagesize)
	} else {
		competitionsOrm = competitionsOrm.Limit(pagesize)
	}

	if order != "" {
		competitionsOrm = competitionsOrm.Order(order)
	}

	if err = competitionsOrm.Find(&competitions).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return competitions, totalRows, nil
}

// GetCompetitions is a function to get a single record from the competitions table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetCompetitions(ctx context.Context, argId string) (record *model.Competitions, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddCompetitions is a function to add a single record to competitions table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddCompetitions(ctx context.Context, record *model.Competitions) (result *model.Competitions, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateCompetitions is a function to update a single record from competitions table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateCompetitions(ctx context.Context, argId string, updated *model.Competitions) (result *model.Competitions, RowsAffected int64, err error) {

	result = &model.Competitions{}
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

// DeleteCompetitions is a function to delete a single record from competitions table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteCompetitions(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.Competitions{}
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
