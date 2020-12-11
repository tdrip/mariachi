package pgumodel

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	per "github.com/tdrip/persist/pkg/interfaces"
)

//
// Built from:
// main - VenafiTechs.Db
/*
 CREATE TABLE configurations (
    id       INTEGER       NOT NULL,
    user     VARCHAR (120),
    password VARCHAR (25),
    baseurl  VARCHAR (250),
    PRIMARY KEY (
        id
    )
)
*/
//

// Data storage IDataItem

// Built from: configurations
type Configuration struct {
	per.IDataItem `json:"-"`

	// id (SQL TYPE: INTEGER)
	Id int64 `json:"id"`

	// user (SQL TYPE: VARCHAR (120))
	User string `json:"user,omitempty"`

	// password (SQL TYPE: VARCHAR (25))
	Password string `json:"password,omitempty"`

	// baseurl (SQL TYPE: VARCHAR (250))
	Baseurl string `json:"baseurl,omitempty"`
}

func (data Configuration) ConvertFromIDataItem(input per.IDataItem) Configuration {
	res := input.(Configuration)
	return res
}

func (data Configuration) Print() string {
	return fmt.Sprintf(" %s ", data)
}

func (data *Configuration) String() string {
	str := ""

	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ", data.Id)

	// user (SQL TYPE: VARCHAR (120))
	str = str + fmt.Sprintf(" %s ", data.User)

	// password (SQL TYPE: VARCHAR (25))
	str = str + fmt.Sprintf(" %s ", data.Password)

	// baseurl (SQL TYPE: VARCHAR (250))
	str = str + fmt.Sprintf(" %s ", data.Baseurl)

	return str //fmt.Sprintf(" %v, ",data)
}
