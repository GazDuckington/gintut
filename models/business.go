package models

import (
	"gintut/helpers"
	"time"
)

type TBusinessunit struct {
	Begda time.Time `gorm:"column:BEGDA;default:CURRENT_DATE"`
	Endda time.Time `gorm:"column:ENDDA;default:2999-01-01"`
	Buid  string    `gorm:"column:BUID;primaryKey;default:uuid_generate_v4()"`
	Bucd  string    `gorm:"column:BUCD"`
	Bunm  string    `gorm:"column:BUNM"`
	Chgda time.Time `gorm:"column:CHGDA"`
	Chgby string    `gorm:"column:CHGBY"`
	X1    string    `gorm:"column:X1"`
	X2    string    `gorm:"column:X2"`
	X3    string    `gorm:"column:X3"`
	X4    string    `gorm:"column:X4"`
	X5    string    `gorm:"column:X5"`
	X6    string    `gorm:"column:X6"`
	X7    string    `gorm:"column:X7"`
}

func (TBusinessunit) TableName() string {
	return "t_businessunit"
}

func (bu *TBusinessunit) SoftDelete(deletedBy string) {
	bu.Chgby = deletedBy
	bu.Chgda = helpers.Today
	bu.Endda = helpers.Yesterday
}

func (bu *TBusinessunit) Undelete(restoredBy string) {
	bu.Endda = time.Date(2999, time.January, 1, 0, 0, 0, 0, time.Now().Local().Location())
	bu.Chgda = helpers.Today
	bu.Chgby = restoredBy
}

func (bu *TBusinessunit) Delete(deletedBy string) {
	bu.SoftDelete(deletedBy)
}
