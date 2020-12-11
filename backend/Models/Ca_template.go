package pgumodel

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	per "github.com/tdrip/persist/pkg/interfaces"
)

//
// Built from:
// main - VTechsDatastore.Db
/*
 CREATE TABLE ca_templates (
    id          INTEGER      PRIMARY KEY AUTOINCREMENT,
    description VARCHAR (50) UNIQUE
                            NOT NULL,
    dn VARCHAR  (250) UNIQUE
                            NOT NULL

)
*/
//

// Data storage IDataItem

// Built from: ca_templates
type Ca_template struct {
	per.IDataItem `json:"-"`

	// id (SQL TYPE: INTEGER)
	Id int64 `json:"id"`

	// description (SQL TYPE: VARCHAR (50))
	Description string `json:"description"`

	// dn (SQL TYPE: VARCHAR  (250))
	Dn string `json:"dn"`
}

func (data Ca_template) ConvertFromIDataItem(input per.IDataItem) Ca_template {
	res := input.(Ca_template)
	return res
}

func (data Ca_template) Print() string {
	return fmt.Sprintf(" %s ", data)
}

func (data *Ca_template) String() string {
	str := ""

	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ", data.Id)

	// description (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Description)

	// dn (SQL TYPE: VARCHAR  (250))
	str = str + fmt.Sprintf(" %s ", data.Dn)

	return str //fmt.Sprintf(" %v, ",data)
}
