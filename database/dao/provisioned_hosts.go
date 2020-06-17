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

// GetAllProvisionedHosts is a function to get a slice of record(s) from provisioned_hosts table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllProvisionedHosts(ctx context.Context, page, pagesize int64, order string) (provisionedhosts []*model.ProvisionedHosts, totalRows int, err error) {

	provisionedhosts = []*model.ProvisionedHosts{}

	provisionedhostsOrm := DB.Model(&model.ProvisionedHosts{})
	provisionedhostsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		provisionedhostsOrm = provisionedhostsOrm.Offset(offset).Limit(pagesize)
	} else {
		provisionedhostsOrm = provisionedhostsOrm.Limit(pagesize)
	}

	if order != "" {
		provisionedhostsOrm = provisionedhostsOrm.Order(order)
	}

	if err = provisionedhostsOrm.Find(&provisionedhosts).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return provisionedhosts, totalRows, nil
}

// GetProvisionedHosts is a function to get a single record from the provisioned_hosts table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetProvisionedHosts(ctx context.Context, argId string) (record *model.ProvisionedHosts, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddProvisionedHosts is a function to add a single record to provisioned_hosts table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddProvisionedHosts(ctx context.Context, record *model.ProvisionedHosts) (result *model.ProvisionedHosts, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateProvisionedHosts is a function to update a single record from provisioned_hosts table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateProvisionedHosts(ctx context.Context, argId string, updated *model.ProvisionedHosts) (result *model.ProvisionedHosts, RowsAffected int64, err error) {

	result = &model.ProvisionedHosts{}
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

// DeleteProvisionedHosts is a function to delete a single record from provisioned_hosts table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteProvisionedHosts(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.ProvisionedHosts{}
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
