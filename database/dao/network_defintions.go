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

// GetAllNetworkDefintions is a function to get a slice of record(s) from network_defintions table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllNetworkDefintions(ctx context.Context, page, pagesize int64, order string) (networkdefintions []*model.NetworkDefintions, totalRows int, err error) {

	networkdefintions = []*model.NetworkDefintions{}

	networkdefintionsOrm := DB.Model(&model.NetworkDefintions{})
	networkdefintionsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		networkdefintionsOrm = networkdefintionsOrm.Offset(offset).Limit(pagesize)
	} else {
		networkdefintionsOrm = networkdefintionsOrm.Limit(pagesize)
	}

	if order != "" {
		networkdefintionsOrm = networkdefintionsOrm.Order(order)
	}

	if err = networkdefintionsOrm.Find(&networkdefintions).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return networkdefintions, totalRows, nil
}

// GetNetworkDefintions is a function to get a single record from the network_defintions table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetNetworkDefintions(ctx context.Context, argId string) (record *model.NetworkDefintions, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddNetworkDefintions is a function to add a single record to network_defintions table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddNetworkDefintions(ctx context.Context, record *model.NetworkDefintions) (result *model.NetworkDefintions, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateNetworkDefintions is a function to update a single record from network_defintions table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateNetworkDefintions(ctx context.Context, argId string, updated *model.NetworkDefintions) (result *model.NetworkDefintions, RowsAffected int64, err error) {

	result = &model.NetworkDefintions{}
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

// DeleteNetworkDefintions is a function to delete a single record from network_defintions table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteNetworkDefintions(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.NetworkDefintions{}
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
