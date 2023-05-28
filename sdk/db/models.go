package db

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

var URI string

type Record struct {
	ID           string
	Sex          string
	Birthdate    string
	PatientID    string
	ICD          string
	Diagnosis    string
	ServiceDate  string
	Title        string
	Prescription string
}

func NewRecord(c []string) *Record {
	r := &Record{
		ID:           uuid.NewString(),
		Sex:          c[0],
		Birthdate:    c[1],
		PatientID:    c[2],
		ICD:          c[3],
		Diagnosis:    c[4],
		ServiceDate:  c[5],
		Title:        c[6],
		Prescription: c[7],
	}

	return r
}

func newConnection() *sql.DB {
	dbb, _ := sql.Open("postgres", URI)
	dbb.SetMaxIdleConns(10)
	dbb.SetMaxOpenConns(10)
	dbb.SetConnMaxLifetime(0)

	return dbb
}

func SetURI(uri string) {
	URI = uri
}

func NewCopy(records []*Record) (bool, error) {
	dbb := newConnection()
	txn, err := dbb.Begin()
	if err != nil {
		return false, err
	}

	stmt, _ := txn.Prepare(pq.CopyIn(
		"messagedetailrecord",
		"_id",
		"sex",
		"birthdate",
		"patientid",
		"icd",
		"diagnosis",
		"servicedate",
		"title",
		"prescription"))

	for _, r := range records {
		_, err := stmt.Exec(r.ID, r.Sex, r.Birthdate, r.PatientID, r.ICD, r.Diagnosis, r.ServiceDate, r.Title, r.Prescription)
		if err != nil {
			return false, err
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		return false, err
	}
	err = stmt.Close()
	if err != nil {
		return false, err
	}
	err = txn.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Record) ToStrings() (string, string, string, string, string, string, string, string) {
	return r.ID, r.Sex, r.Birthdate, r.PatientID, r.ICD, r.ServiceDate, r.Title, r.Prescription
}
