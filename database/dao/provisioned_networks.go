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

// GetAllProvisionedNetworks is a function to get a slice of record(s) from provisioned_networks table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllProvisionedNetworks(ctx context.Context, page, pagesize int64, order string) (provisionednetworks []*model.ProvisionedNetworks, totalRows int, err error) {

	provisionednetworks = []*model.ProvisionedNetworks{}

	provisionednetworksOrm := DB.Model(&model.ProvisionedNetworks{})
	provisionednetworksOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		provisionednetworksOrm = provisionednetworksOrm.Offset(offset).Limit(pagesize)
	} else {
		provisionednetworksOrm = provisionednetworksOrm.Limit(pagesize)
	}

	if order != "" {
		provisionednetworksOrm = provisionednetworksOrm.Order(order)
	}

	if err = provisionednetworksOrm.Find(&provisionednetworks).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return provisionednetworks, totalRows, nil
}

// GetProvisionedNetworks is a function to get a single record from the provisioned_networks table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetProvisionedNetworks(ctx context.Context, argId string) (record *model.ProvisionedNetworks, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddProvisionedNetworks is a function to add a single record to provisioned_networks table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddProvisionedNetworks(ctx context.Context, record *model.ProvisionedNetworks) (result *model.ProvisionedNetworks, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateProvisionedNetworks is a function to update a single record from provisioned_networks table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateProvisionedNetworks(ctx context.Context, argId string, updated *model.ProvisionedNetworks) (result *model.ProvisionedNetworks, RowsAffected int64, err error) {

	result = &model.ProvisionedNetworks{}
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

// DeleteProvisionedNetworks is a function to delete a single record from provisioned_networks table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteProvisionedNetworks(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.ProvisionedNetworks{}
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
