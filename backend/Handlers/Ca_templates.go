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
 CREATE TABLE ca_templates (
    id          INTEGER      PRIMARY KEY AUTOINCREMENT,
    description VARCHAR (50) UNIQUE
                            NOT NULL,
    dn VARCHAR  (250) UNIQUE
                            NOT NULL

)
*/
//

// Table fields

// ca_templates
const ca_templatesTName = "ca_templates"

// Primay Key: id
const ca_templatesIdCName = "id"

// description
const ca_templatesDescriptionCName = "description"

// dn
const ca_templatesDnCName = "dn"

// HANDLER

type Ca_templatesHandler struct {
	per.IStorageHandler
	Parent   *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewCa_templatesHandler(datastore *SQLL.SQLLiteDatastore) *Ca_templatesHandler {
	ds := Ca_templatesHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler
func (handler *Ca_templatesHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *Ca_templatesHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for Ca_template
func (handler *Ca_templatesHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures", "Executing Query")
	return handler.Executor.ExecuteQuery(`CREATE TABLE IF NOT EXISTS ca_templates (
    id          INTEGER      PRIMARY KEY AUTOINCREMENT,
    description VARCHAR (50) UNIQUE
                            NOT NULL,
    dn VARCHAR  (250) UNIQUE
                            NOT NULL
                                                         
)`)
}

// End Istorage

// This function Ca_template removes all data for the table
func (handler *Ca_templatesHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + ca_templatesTName))
}

// This adds Ca_template to the database
func (handler *Ca_templatesHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Ca_template)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO "+ca_templatesTName+" ( "+"["+ca_templatesDescriptionCName+"]"+",["+ca_templatesDnCName+"]"+" ) VALUES (?,?)", data.Description, data.Dn))
}

func (handler *Ca_templatesHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Ca_template)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE "+ca_templatesTName+" SET "+"["+ca_templatesDescriptionCName+"] = ? "+",["+ca_templatesDnCName+"] = ? "+"  WHERE ["+ca_templatesIdCName+"] = ?", data.Description, data.Dn, data.Id))
}

func (handler *Ca_templatesHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return SQLL.ResultToSQLLiteQueryResult(data)
}

func (handler *Ca_templatesHandler) FindById(SearchData int64) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+ca_templatesIdCName+"],"+"["+ca_templatesDescriptionCName+"]"+",["+ca_templatesDnCName+"]"+"  FROM "+ca_templatesTName+" WHERE "+ca_templatesIdCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *Ca_templatesHandler) FindByDescription(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+ca_templatesIdCName+"],"+"["+ca_templatesDescriptionCName+"]"+",["+ca_templatesDnCName+"]"+"  FROM "+ca_templatesTName+" WHERE "+ca_templatesDescriptionCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *Ca_templatesHandler) FindByDn(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+ca_templatesIdCName+"],"+"["+ca_templatesDescriptionCName+"]"+",["+ca_templatesDnCName+"]"+"  FROM "+ca_templatesTName+" WHERE "+ca_templatesDnCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *Ca_templatesHandler) ReadAll() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+ca_templatesIdCName+"],"+"["+ca_templatesDescriptionCName+"]"+",["+ca_templatesDnCName+"]"+"  FROM "+ca_templatesTName, handler.ParseRows))
}

func (handler *Ca_templatesHandler) ParseRows(rows *sql.Rows) per.IQueryResult {

	var Id *int64

	var Description *string

	var Dn *string

	results := []per.IDataItem{} //Ca_template{}

	for rows.Next() {
		rows.Scan(&Id, &Description, &Dn)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)

		res := data.Ca_template{}

		if Id != nil {
			res.Id = *Id
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Id", Id)
		} else {
			handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
		}

		if Description != nil {
			res.Description = *Description
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Description", Description)
		} else {
			handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
		}

		if Dn != nil {
			res.Dn = *Dn
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Dn", Dn)
		} else {
			handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
		}

		results = append(results, res)
	}
	return SQLL.NewDataQueryResult(true, results)
}
