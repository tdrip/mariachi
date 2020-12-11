package pguhandlers

import (
	"database/sql"

	data "github.com/tdrip/mariachi/backend/Models"

	_ "github.com/mattn/go-sqlite3"
	per "github.com/tdrip/persist/pkg/interfaces"
	SQLL "github.com/tdrip/persist/pkg/sqllite"
)

//
// Built from:
// main - VTechsDatastore.Db
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

// Table fields

// configurations
const configurationsTName = "configurations"

// Primay Key: id
const configurationsIdCName = "id"

// user
const configurationsUserCName = "user"

// password
const configurationsPasswordCName = "password"

// baseurl
const configurationsBaseurlCName = "baseurl"

// HANDLER

type ConfigurationsHandler struct {
	per.IStorageHandler
	Parent   *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewConfigurationsHandler(datastore *SQLL.SQLLiteDatastore) *ConfigurationsHandler {
	ds := ConfigurationsHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler
func (handler *ConfigurationsHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *ConfigurationsHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for Configuration
func (handler *ConfigurationsHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures", "Executing Query")
	return handler.Executor.ExecuteQuery(`CREATE TABLE IF NOT EXISTS configurations (
    id       INTEGER       NOT NULL,
    user     VARCHAR (120),
    password VARCHAR (25),
    baseurl  VARCHAR (250),
    PRIMARY KEY (
        id
    )
)`)
}

// End Istorage

// This function Configuration removes all data for the table
func (handler *ConfigurationsHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + configurationsTName))
}

// This adds Configuration to the database
func (handler *ConfigurationsHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Configuration)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO "+configurationsTName+" ( "+"["+configurationsUserCName+"]"+",["+configurationsPasswordCName+"]"+",["+configurationsBaseurlCName+"]"+" ) VALUES (?,?,?)", data.User, data.Password, data.Baseurl))
}

func (handler *ConfigurationsHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Configuration)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE "+configurationsTName+" SET "+"["+configurationsUserCName+"] = ? "+",["+configurationsPasswordCName+"] = ? "+",["+configurationsBaseurlCName+"] = ? "+"  WHERE ["+configurationsIdCName+"] = ?", data.User, data.Password, data.Baseurl, data.Id))
}

func (handler *ConfigurationsHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return SQLL.ResultToSQLLiteQueryResult(data)
}

func (handler *ConfigurationsHandler) FindById(SearchData int64) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+configurationsIdCName+"],"+"["+configurationsUserCName+"]"+",["+configurationsPasswordCName+"]"+",["+configurationsBaseurlCName+"]"+"  FROM "+configurationsTName+" WHERE "+configurationsIdCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ConfigurationsHandler) FindByUser(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+configurationsIdCName+"],"+"["+configurationsUserCName+"]"+",["+configurationsPasswordCName+"]"+",["+configurationsBaseurlCName+"]"+"  FROM "+configurationsTName+" WHERE "+configurationsUserCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ConfigurationsHandler) FindByPassword(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+configurationsIdCName+"],"+"["+configurationsUserCName+"]"+",["+configurationsPasswordCName+"]"+",["+configurationsBaseurlCName+"]"+"  FROM "+configurationsTName+" WHERE "+configurationsPasswordCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ConfigurationsHandler) FindByBaseurl(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+configurationsIdCName+"],"+"["+configurationsUserCName+"]"+",["+configurationsPasswordCName+"]"+",["+configurationsBaseurlCName+"]"+"  FROM "+configurationsTName+" WHERE "+configurationsBaseurlCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ConfigurationsHandler) ReadAll() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+configurationsIdCName+"],"+"["+configurationsUserCName+"]"+",["+configurationsPasswordCName+"]"+",["+configurationsBaseurlCName+"]"+"  FROM "+configurationsTName, handler.ParseRows))
}

func (handler *ConfigurationsHandler) ParseRows(rows *sql.Rows) per.IQueryResult {

	var Id *int64

	var User *string

	var Password *string

	var Baseurl *string

	results := []per.IDataItem{} //Configuration{}

	for rows.Next() {
		err := rows.Scan(&Id, &User, &Password, &Baseurl)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)
		if err != nil {
			handler.Parent.LogErrorEf("ParseRows", "Row Scan error: %s", err)
		} else {
			res := data.Configuration{}

			if Id != nil {
				res.Id = *Id
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Id", *Id)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			if User != nil {
				res.User = *User
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for User", *User)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			if Password != nil {
				res.Password = *Password
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Password", *Password)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			if Baseurl != nil {
				res.Baseurl = *Baseurl
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Baseurl", *Baseurl)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			results = append(results, res)
		}

	}
	return SQLL.NewDataQueryResult(true, results)
}
