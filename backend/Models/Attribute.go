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
 CREATE TABLE attributes (
    id                   INTEGER       NOT NULL,
    name                 VARCHAR (50),
    description          VARCHAR (250),
    technologies_type    VARCHAR (50),
    type                 VARCHAR (10)  NOT NULL
                                       DEFAULT string,
    technologies_invalid TEXT (50),
    PRIMARY KEY (
        id
    )
)
*/
//

// Data storage IDataItem

// Built from: attributes
type Attribute struct {
	per.IDataItem `json:"-"`

	// id (SQL TYPE: INTEGER)
	Id int64 `json:"id"`

	// name (SQL TYPE: VARCHAR (50))
	Name string `json:"name,omitempty"`

	// description (SQL TYPE: VARCHAR (250))
	Description string `json:"description,omitempty"`

	// technologies_type (SQL TYPE: VARCHAR (50))
	Technologies_type string `json:"technologies_type,omitempty"`

	// type (SQL TYPE: VARCHAR (10))
	Type string `json:"type"`
}

func (data Attribute) ConvertFromIDataItem(input per.IDataItem) Attribute {
	res := input.(Attribute)
	return res
}

func (data Attribute) Print() string {
	return fmt.Sprintf(" %s ", data)
}

func (data *Attribute) String() string {
	str := ""

	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ", data.Id)

	// name (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Name)

	// description (SQL TYPE: VARCHAR (250))
	str = str + fmt.Sprintf(" %s ", data.Description)

	// technologies_type (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Technologies_type)

	// type (SQL TYPE: VARCHAR (10))
	str = str + fmt.Sprintf(" %s ", data.Type)

	return str //fmt.Sprintf(" %v, ",data)
}
