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

// GetAllBuildConfigs is a function to get a slice of record(s) from build_configs table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBuildConfigs(ctx context.Context, page, pagesize int64, order string) (buildconfigs []*model.BuildConfigs, totalRows int, err error) {

	buildconfigs = []*model.BuildConfigs{}

	buildconfigsOrm := DB.Model(&model.BuildConfigs{})
	buildconfigsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		buildconfigsOrm = buildconfigsOrm.Offset(offset).Limit(pagesize)
	} else {
		buildconfigsOrm = buildconfigsOrm.Limit(pagesize)
	}

	if order != "" {
		buildconfigsOrm = buildconfigsOrm.Order(order)
	}

	if err = buildconfigsOrm.Find(&buildconfigs).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return buildconfigs, totalRows, nil
}

// GetBuildConfigs is a function to get a single record from the build_configs table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetBuildConfigs(ctx context.Context, argId string) (record *model.BuildConfigs, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddBuildConfigs is a function to add a single record to build_configs table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddBuildConfigs(ctx context.Context, record *model.BuildConfigs) (result *model.BuildConfigs, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateBuildConfigs is a function to update a single record from build_configs table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBuildConfigs(ctx context.Context, argId string, updated *model.BuildConfigs) (result *model.BuildConfigs, RowsAffected int64, err error) {

	result = &model.BuildConfigs{}
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

// DeleteBuildConfigs is a function to delete a single record from build_configs table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBuildConfigs(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.BuildConfigs{}
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
