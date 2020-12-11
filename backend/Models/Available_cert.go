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
 CREATE TABLE available_certs (
    id INTEGER       PRIMARY KEY AUTOINCREMENT,
    cn VARCHAR (250) NOT NULL,
    dn TEXT          NOT NULL
)
*/
//

// Data storage IDataItem

// Built from: available_certs
type Available_cert struct {
	per.IDataItem `json:"-"`

	// id (SQL TYPE: INTEGER)
	Id int64 `json:"id"`

	// cn (SQL TYPE: VARCHAR (250))
	Cn string `json:"cn"`

	// dn (SQL TYPE: TEXT)
	Dn string `json:"dn"`
}

func (data Available_cert) ConvertFromIDataItem(input per.IDataItem) Available_cert {
	res := input.(Available_cert)
	return res
}

func (data Available_cert) Print() string {
	return fmt.Sprintf(" %s ", data)
}

func (data *Available_cert) String() string {
	str := ""

	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ", data.Id)

	// cn (SQL TYPE: VARCHAR (250))
	str = str + fmt.Sprintf(" %s ", data.Cn)

	// dn (SQL TYPE: TEXT)
	str = str + fmt.Sprintf(" %s ", data.Dn)

	return str //fmt.Sprintf(" %v, ",data)
}
