package api

import "gorm.io/gorm"

type KindName struct {
	Meta
}

type KindNameList []*KindName
type KindNameIndex map[string]*KindName

func (l KindNameList) Index() KindNameIndex {
	index := KindNameIndex{}
	for _, o := range l {
		index[o.ID] = o
	}
	return index
}

func (d *KindName) BeforeCreate(tx *gorm.DB) error {
	d.ID = NewID()
	return nil
}

type KindNamePatchRequest struct {

}
