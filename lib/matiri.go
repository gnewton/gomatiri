package lib

import (
	"time"
)

type BackupEvent struct {
	ID        int64
	Host      string
	Port      int
	StartTime time.Time
	EndTime   time.Time
	User      string
	Bytes     int64
	File      string
	Sha256    string
	error     bool
}

type Database struct {
	ID        int64
	Completed int
	BackupId  int64
	Database  string
	File      string
	StartTime time.Time
	EndTime   time.Time
	Bytes     int64
	Sha256    string
	Error     bool
}

type Services struct {
}

func (s *Services) Init(dbfile string) error {
	return nil
}

//BackupEvents
func (s *Services) GetAllBackupEvents() {

}

func (s *Services) GetBackupEvent(backupId int64) *BackupEvent {
	return nil
}

func (s *Services) GetBackupEventYears() {

}

//Databases
func (s *Services) GetDatabases(backupId int64) *Database {
	return nil
}
