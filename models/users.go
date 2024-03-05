package models

import (
	"time"
	// "gorm.io/gorm"
)

type TPersonal struct {
	Begda     time.Time `gorm:"column:BEGDA;default:CURRENT_DATE"`
	Endda     time.Time `gorm:"column:ENDDA;default:2999-01-01"`
	Prsnid    string    `gorm:"primaryKey;column:PRSNID;default:uuid_generate_v4()"`
	Nam       string    `gorm:"column:NAM"`
	Nik       string    `gorm:"column:NIK;uniqueIndex"`
	Eml       string    `gorm:"column:EML"`
	Pnum      string    `gorm:"column:PNUM"`
	Bucd      string    `gorm:"column:BUCD"`
	Divcd     string    `gorm:"column:DIVCD"`
	Poscd     string    `gorm:"column:POSCD"`
	Ropoid    string    `gorm:"column:ROPOID"`
	Chgda     time.Time `gorm:"column:CHGDA;default:CURRENT_DATE"`
	Chgby     string    `gorm:"column:CHGBY"`
	Rlcd      string    `gorm:"column:RLCD"`
	Lvl       string    `gorm:"column:LVL"`
	Isact     bool      `gorm:"column:ISACT;default:true"`
	Stat      string    `gorm:"column:STAT"`
	Bunm      string    `gorm:"column:BUNM"`
	Divnm     string    `gorm:"column:DIVNM"`
	Dirspv    string    `gorm:"column:DIRSPV"`
	Dirspvnik string    `gorm:"column:DIRSPVNIK"`
	Direk     string    `gorm:"column:DIREK"`
	Subdirek  string    `gorm:"column:SUBDIREK"`
	X5        string    `gorm:"column:X5"`
	X6        string    `gorm:"column:X6"`
	X7        string    `gorm:"column:X7"`
}

func (t *TPersonal) TableName() string {
	return "t_personal"
}

func (t *TPersonal) SoftDelete(deletedBy string) {
	t.Endda = time.Now().AddDate(0, 0, -1)
	t.Isact = false
	t.Chgby = deletedBy
}

func (t *TPersonal) Undelete(restoredBy string) {
	t.Endda = time.Date(2999, time.January, 1, 0, 0, 0, 0, time.Now().Local().Location())
	t.Isact = true
	t.Chgby = restoredBy
	t.Stat = ""
}

type TUsers struct {
	Begda     time.Time `gorm:"column:BEGDA;default:CURRENT_DATE"`
	Endda     time.Time `gorm:"column:ENDDA;default:'2999-01-01'"`
	Usid      string    `gorm:"column:USID;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Usrn      string    `gorm:"column:USRN"`
	Eml       string    `gorm:"column:EML"`
	Pswd      string    `gorm:"column:PSWD"`
	Rlcd      string    `gorm:"column:RLCD"`
	Logat     time.Time `gorm:"column:LOGAT;default:current_timestamp"`
	Isact     bool      `gorm:"column:ISACT"`
	Chgda     time.Time `gorm:"column:CHGDA;default:current_date"`
	Chgby     string    `gorm:"column:CHGBY"`
	Nam       string    `gorm:"column:NAM"`
	Nik       string    `gorm:"column:NIK"`
	Pnum      string    `gorm:"column:PNUM"`
	Bucd      string    `gorm:"column:BUCD"`
	Bunm      string    `gorm:"column:BUNM"`
	Divcd     string    `gorm:"column:DIVCD"`
	Divnm     string    `gorm:"column:DIVNM"`
	Poscd     string    `gorm:"column:POSCD"`
	Lvl       string    `gorm:"column:LVL"`
	Stat      string    `gorm:"column:STAT"`
	Spvid     string    `gorm:"column:SPVID"`
	Spvnm     string    `gorm:"column:SPVNM"`
	Cmpcd     string    `gorm:"column:CMPCD"`
	Ropoid    string    `gorm:"column:ROPOID"`
	Dirspv    string    `gorm:"column:DIRSPV"`
	Dirspnvik string    `gorm:"column:DIRSPVNIK"`
	X1        string    `gorm:"column:X1"`
	X2        string    `gorm:"column:X2"`
	X3        string    `gorm:"column:X3"`
	X4        string    `gorm:"column:X4"`
	X5        string    `gorm:"column:X5"`
	X6        string    `gorm:"column:X6"`
	X7        string    `gorm:"column:X7"`
}

func (TUsers) TableName() string {
	return "t_users"
}

// SoftDelete soft deletes the user by setting Isact to false and updating Endda
func (user *TUsers) SoftDelete() {
	user.Isact = false
	user.Endda = time.Now().Add(-24 * time.Hour) // Example: Set to yesterday
}

// Undelete undeletes the user by setting Endda back to the default value
func (user *TUsers) Undelete() {
	user.Endda = time.Date(2999, time.January, 1, 0, 0, 0, 0, time.Now().Local().Location())
	user.Isact = true
}

// Delete soft deletes the user (wrapper for SoftDelete for consistency)
func (user *TUsers) Delete() {
	user.SoftDelete()
}
