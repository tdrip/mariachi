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
 CREATE TABLE objects (
    id                INTEGER       NOT NULL,
    name              VARCHAR (250),
    keyvaultname      VARCHAR (50),
    ip_address        VARCHAR (120),
    host              VARCHAR (50),
    port              VARCHAR (50),
    issync            BOOLEAN,
    technologies_type VARCHAR (25),
    technologies_name VARCHAR (50),
    comments          TEXT,
    certificate_dns   TEXT,
    ca_template       TEXT,
    ca_common_name    VARCHAR (50),
    tech              VARCHAR (35),
    region            VARCHAR (50),
    loadbalancer      VARCHAR (50),
    PRIMARY KEY (
        id
    )
)
*/
//

// Data storage IDataItem

// Built from: objects
type Object struct {
	per.IDataItem `json:"-"`

	// id (SQL TYPE: INTEGER)
	Id int64 `json:"id"`

	// name (SQL TYPE: VARCHAR (250))
	Name string `json:"name,omitempty"`

	// keyvaultname (SQL TYPE: VARCHAR (50))
	Keyvaultname string `json:"keyvaultname,omitempty"`

	// ip_address (SQL TYPE: VARCHAR (120))
	Ip_address string `json:"ip_address,omitempty"`

	// host (SQL TYPE: VARCHAR (50))
	Host string `json:"host,omitempty"`

	// port (SQL TYPE: VARCHAR (50))
	Port string `json:"port,omitempty"`

	// technologies_type (SQL TYPE: VARCHAR (25))
	Technologies_type string `json:"technologies_type,omitempty"`

	// technologies_name (SQL TYPE: VARCHAR (50))
	Technologies_name string `json:"technologies_name,omitempty"`

	// comments (SQL TYPE: TEXT)
	Comments string `json:"comments,omitempty"`

	// certificate_dns (SQL TYPE: TEXT)
	Certificate_dns string `json:"certificate_dns,omitempty"`

	// ca_template (SQL TYPE: TEXT)
	Ca_template string `json:"ca_template,omitempty"`

	// ca_common_name (SQL TYPE: VARCHAR (50))
	Ca_common_name string `json:"ca_common_name,omitempty"`

	// tech (SQL TYPE: VARCHAR (35))
	Tech string `json:"tech,omitempty"`

	// region (SQL TYPE: VARCHAR (50))
	Region string `json:"region,omitempty"`

	// loadbalancer (SQL TYPE: VARCHAR (50))
	Loadbalancer string `json:"loadbalancer,omitempty"`
}

func (data Object) ConvertFromIDataItem(input per.IDataItem) Object {
	res := input.(Object)
	return res
}

func (data Object) Print() string {
	return fmt.Sprintf(" %s ", data)
}

func (data *Object) String() string {
	str := ""

	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ", data.Id)

	// name (SQL TYPE: VARCHAR (250))
	str = str + fmt.Sprintf(" %s ", data.Name)

	// keyvaultname (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Keyvaultname)

	// ip_address (SQL TYPE: VARCHAR (120))
	str = str + fmt.Sprintf(" %s ", data.Ip_address)

	// host (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Host)

	// port (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Port)

	// technologies_type (SQL TYPE: VARCHAR (25))
	str = str + fmt.Sprintf(" %s ", data.Technologies_type)

	// technologies_name (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Technologies_name)

	// comments (SQL TYPE: TEXT)
	str = str + fmt.Sprintf(" %s ", data.Comments)

	// certificate_dns (SQL TYPE: TEXT)
	str = str + fmt.Sprintf(" %s ", data.Certificate_dns)

	// ca_template (SQL TYPE: TEXT)
	str = str + fmt.Sprintf(" %s ", data.Ca_template)

	// ca_common_name (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Ca_common_name)

	// tech (SQL TYPE: VARCHAR (35))
	str = str + fmt.Sprintf(" %s ", data.Tech)

	// region (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Region)

	// loadbalancer (SQL TYPE: VARCHAR (50))
	str = str + fmt.Sprintf(" %s ", data.Loadbalancer)

	return str //fmt.Sprintf(" %v, ",data)
}
