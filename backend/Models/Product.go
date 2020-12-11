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
 CREATE TABLE products (
    id          INTEGER      NOT NULL,
    description VARCHAR (50),
    img         VARCHAR (50),
    PRIMARY KEY (
        id
    ),
    UNIQUE (
        description
    )
)
*/
//

// Data storage IDataItem

// Built from: products
type Product struct {
	per.IDataItem `json:"-"`

	// id (SQL TYPE: INTEGER)
	Id int64 `json:"id"`

	// description (SQL TYPE: VARCHAR (50))
	Description string `json:"description,omitempty"`

	// img (SQL TYPE: VARCHAR (50))
	Img string `json:"img,omitempty"`
}

func (data Product) ConvertFromIDataItem(input per.IDataItem) Product {
	res := input.(Product)
	return res
}

func (data Product) Print() string {
	return fmt.Sprintf(" %s ", data)
}

func (data *Product) String() string {
	str := ""

	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ", data.Id)

	// description (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Description)

	// img (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Img)

	return str //fmt.Sprintf(" %v, ",data)
}
