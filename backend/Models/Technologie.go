package pgumodel

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	per "github.com/tdrip/persist/pkg/interfaces"
)

//
// Built from:
// main - VTechsDatastores.Db
/*
 CREATE TABLE technologies (
    id               INTEGER      NOT NULL,
    name             VARCHAR (50) UNIQUE,
    type             VARCHAR (25),
    icon             VARCHAR (50),
    type_description VARCHAR (30),
    PRIMARY KEY (
        id
    ),
    UNIQUE (
        name
    )
)
*/
//

// Data storage IDataItem

// Built from: technologies
type Technologie struct {
	per.IDataItem `json:"-"`

	// id (SQL TYPE: INTEGER)
	Id int64 `json:"id"`

	// name (SQL TYPE: VARCHAR (50))
	Name string `json:"name,omitempty"`

	// type (SQL TYPE: VARCHAR (25))
	Type string `json:"type,omitempty"`

	// icon (SQL TYPE: VARCHAR (50))
	Icon string `json:"icon,omitempty"`

	// type_description (SQL TYPE: VARCHAR (30))
	Type_description string `json:"type_description,omitempty"`
}

func (data Technologie) ConvertFromIDataItem(input per.IDataItem) Technologie {
	res := input.(Technologie)
	return res
}

func (data Technologie) Print() string {
	return fmt.Sprintf(" %s ", data)
}

func (data *Technologie) String() string {
	str := ""

	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ", data.Id)

	// name (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Name)

	// type (SQL TYPE: VARCHAR (25))
	str = str + fmt.Sprintf(" %s ", data.Type)

	// icon (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Icon)

	// type_description (SQL TYPE: VARCHAR (30))
	str = str + fmt.Sprintf(" %s ", data.Type_description)

	return str //fmt.Sprintf(" %v, ",data)
}
