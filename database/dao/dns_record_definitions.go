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

// GetAllDNSRecordDefinitions is a function to get a slice of record(s) from dns_record_definitions table in the laforge-2 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllDNSRecordDefinitions(ctx context.Context, page, pagesize int64, order string) (dnsrecorddefinitions []*model.DNSRecordDefinitions, totalRows int, err error) {

	dnsrecorddefinitions = []*model.DNSRecordDefinitions{}

	dnsrecorddefinitionsOrm := DB.Model(&model.DNSRecordDefinitions{})
	dnsrecorddefinitionsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		dnsrecorddefinitionsOrm = dnsrecorddefinitionsOrm.Offset(offset).Limit(pagesize)
	} else {
		dnsrecorddefinitionsOrm = dnsrecorddefinitionsOrm.Limit(pagesize)
	}

	if order != "" {
		dnsrecorddefinitionsOrm = dnsrecorddefinitionsOrm.Order(order)
	}

	if err = dnsrecorddefinitionsOrm.Find(&dnsrecorddefinitions).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return dnsrecorddefinitions, totalRows, nil
}

// GetDNSRecordDefinitions is a function to get a single record from the dns_record_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
func GetDNSRecordDefinitions(ctx context.Context, argId string) (record *model.DNSRecordDefinitions, err error) {
	if err = DB.First(&record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddDNSRecordDefinitions is a function to add a single record to dns_record_definitions table in the laforge-2 database
// error - ErrInsertFailed, db save call failed
func AddDNSRecordDefinitions(ctx context.Context, record *model.DNSRecordDefinitions) (result *model.DNSRecordDefinitions, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateDNSRecordDefinitions is a function to update a single record from dns_record_definitions table in the laforge-2 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateDNSRecordDefinitions(ctx context.Context, argId string, updated *model.DNSRecordDefinitions) (result *model.DNSRecordDefinitions, RowsAffected int64, err error) {

	result = &model.DNSRecordDefinitions{}
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

// DeleteDNSRecordDefinitions is a function to delete a single record from dns_record_definitions table in the laforge-2 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteDNSRecordDefinitions(ctx context.Context, argId string) (rowsAffected int64, err error) {

	record := &model.DNSRecordDefinitions{}
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
