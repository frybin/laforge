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

// GetAllIncludedNetworks is a function to get a slice of record(s) from included_networks table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllIncludedNetworks(ctx context.Context, page, pagesize int64, order string) (includednetworks []*model.IncludedNetworks, totalRows int, err error) {

	includednetworks = []*model.IncludedNetworks{}

	includednetworksOrm := DB.Model(&model.IncludedNetworks{})
	includednetworksOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		includednetworksOrm = includednetworksOrm.Offset(offset).Limit(pagesize)
	} else {
		includednetworksOrm = includednetworksOrm.Limit(pagesize)
	}

	if order != "" {
		includednetworksOrm = includednetworksOrm.Order(order)
	}

	if err = includednetworksOrm.Find(&includednetworks).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return includednetworks, totalRows, nil
}

// GetIncludedNetworks is a function to get a single record from the included_networks table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetIncludedNetworks(ctx context.Context, argId string) (record *model.IncludedNetworks, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddIncludedNetworks is a function to add a single record to included_networks table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddIncludedNetworks(ctx context.Context, record *model.IncludedNetworks) (result *model.IncludedNetworks, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateIncludedNetworks is a function to update a single record from included_networks table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateIncludedNetworks(ctx context.Context, argId string, updated *model.IncludedNetworks) (result *model.IncludedNetworks, RowsAffected int64, err error) {

	result = &model.IncludedNetworks{}
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

// DeleteIncludedNetworks is a function to delete a single record from included_networks table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteIncludedNetworks(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.IncludedNetworks{}
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
