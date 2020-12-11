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
// CREATE TABLE available_certs (
//    id INTEGER       PRIMARY KEY AUTOINCREMENT,
//    cn VARCHAR (250) NOT NULL,
//    dn TEXT          NOT NULL
//)
//

// Table fields

// available_certs
const available_certsTName = "available_certs"

// Primay Key: id
const available_certsIdCName = "id"

// cn
const available_certsCnCName = "cn"

// dn
const available_certsDnCName = "dn"

// HANDLER

type Available_certsHandler struct {
	per.IStorageHandler
	Parent   *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewAvailable_certsHandler(datastore *SQLL.SQLLiteDatastore) *Available_certsHandler {
	ds := Available_certsHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler
func (handler *Available_certsHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *Available_certsHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for Available_cert
func (handler *Available_certsHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures", "Executing Query")
	return handler.Executor.ExecuteQuery("CREATE TABLE IF NOT EXISTS available_certs (id INTEGER PRIMARY KEY AUTOINCREMENT, cn VARCHAR (250) NOT NULL, dn TEXT NOT NULL )")
}

// End Istorage

// This function Available_cert removes all data for the table
func (handler *Available_certsHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + available_certsTName))
}

// This adds Available_cert to the database
func (handler *Available_certsHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Available_cert)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO "+available_certsTName+" ( "+"["+available_certsCnCName+"]"+",["+available_certsDnCName+"]"+" ) VALUES (?,?)", data.Cn, data.Dn))
}

func (handler *Available_certsHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Available_cert)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE "+available_certsTName+" SET "+"["+available_certsCnCName+"] = ? "+",["+available_certsDnCName+"] = ? "+"  WHERE ["+available_certsIdCName+"] = ?", data.Cn, data.Dn, data.Id))
}

func (handler *Available_certsHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return SQLL.ResultToSQLLiteQueryResult(data)
}

func (handler *Available_certsHandler) FindById(SearchData int64) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+available_certsIdCName+"],"+"["+available_certsCnCName+"]"+",["+available_certsDnCName+"]"+"  FROM "+available_certsTName+" WHERE "+available_certsIdCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *Available_certsHandler) FindByCn(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+available_certsIdCName+"],"+"["+available_certsCnCName+"]"+",["+available_certsDnCName+"]"+"  FROM "+available_certsTName+" WHERE "+available_certsCnCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *Available_certsHandler) FindByDn(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+available_certsIdCName+"],"+"["+available_certsCnCName+"]"+",["+available_certsDnCName+"]"+"  FROM "+available_certsTName+" WHERE "+available_certsDnCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *Available_certsHandler) ReadAll() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+available_certsIdCName+"],"+"["+available_certsCnCName+"]"+",["+available_certsDnCName+"]"+"  FROM "+available_certsTName, handler.ParseRows))
}

func (handler *Available_certsHandler) ParseRows(rows *sql.Rows) per.IQueryResult {

	var Id int64

	var Cn string

	var Dn string

	results := []per.IDataItem{} //Available_cert{}

	for rows.Next() {
		rows.Scan(&Id, &Cn, &Dn)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)

		res := data.Available_cert{}

		res.Id = Id
		handler.Parent.LogDebugf("ParseRows", "Set '%v' for Id", Id)

		res.Cn = Cn
		handler.Parent.LogDebugf("ParseRows", "Set '%v' for Cn", Cn)

		res.Dn = Dn
		handler.Parent.LogDebugf("ParseRows", "Set '%v' for Dn", Dn)

		results = append(results, res)
	}
	return SQLL.NewDataQueryResult(true, results)
}
