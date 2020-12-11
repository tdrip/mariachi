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
// main - VenafiTechs.Db
/* CREATE TABLE attributes (
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

// Table fields

// attributes
const attributesTName = "attributes"

// Primay Key: id
const attributesIdCName = "id"

// name
const attributesNameCName = "name"

// description
const attributesDescriptionCName = "description"

// technologies_type
const attributesTechnologies_typeCName = "technologies_type"

// type
const attributesTypeCName = "type"

// technologies_invalid
const attributesTechnologies_invalidCName = "technologies_invalid"

// HANDLER

type AttributesHandler struct {
	per.IStorageHandler
	Parent   *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewAttributesHandler(datastore *SQLL.SQLLiteDatastore) *AttributesHandler {
	ds := AttributesHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler
func (handler *AttributesHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *AttributesHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for Attribute
func (handler *AttributesHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures", "Executing Query")
	return handler.Executor.ExecuteQuery("CREATE TABLE IF NOT EXISTS attributes (id INTEGER NOT NULL, name VARCHAR (50), description VARCHAR (250), technologies_type VARCHAR (50), type VARCHAR (10)  NOT NULL DEFAULT string, technologies_invalid TEXT (50), PRIMARY KEY (id))")
}

// End Istorage

// This function Attribute removes all data for the table
func (handler *AttributesHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + attributesTName))
}

// This adds Attribute to the database
func (handler *AttributesHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Attribute)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO "+attributesTName+" ( "+"["+attributesNameCName+"]"+",["+attributesDescriptionCName+"]"+",["+attributesTechnologies_typeCName+"]"+",["+attributesTypeCName+"]"+",["+attributesTechnologies_invalidCName+"]"+" ) VALUES (?,?,?,?,?)", data.Name, data.Description, data.Technologies_type, data.Type))
}

func (handler *AttributesHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Attribute)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE "+attributesTName+" SET "+"["+attributesNameCName+"] = ? "+",["+attributesDescriptionCName+"] = ? "+",["+attributesTechnologies_typeCName+"] = ? "+",["+attributesTypeCName+"] = ? "+",["+attributesTechnologies_invalidCName+"] = ? "+"  WHERE ["+attributesIdCName+"] = ?", data.Name, data.Description, data.Technologies_type, data.Type, data.Id))
}

func (handler *AttributesHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return SQLL.ResultToSQLLiteQueryResult(data)
}

func (handler *AttributesHandler) FindById(SearchData int64) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+attributesIdCName+"],"+"["+attributesNameCName+"]"+",["+attributesDescriptionCName+"]"+",["+attributesTechnologies_typeCName+"]"+",["+attributesTypeCName+"]"+",["+attributesTechnologies_invalidCName+"]"+"  FROM "+attributesTName+" WHERE "+attributesIdCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *AttributesHandler) FindByName(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+attributesIdCName+"],"+"["+attributesNameCName+"]"+",["+attributesDescriptionCName+"]"+",["+attributesTechnologies_typeCName+"]"+",["+attributesTypeCName+"]"+",["+attributesTechnologies_invalidCName+"]"+"  FROM "+attributesTName+" WHERE "+attributesNameCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *AttributesHandler) FindByDescription(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+attributesIdCName+"],"+"["+attributesNameCName+"]"+",["+attributesDescriptionCName+"]"+",["+attributesTechnologies_typeCName+"]"+",["+attributesTypeCName+"]"+",["+attributesTechnologies_invalidCName+"]"+"  FROM "+attributesTName+" WHERE "+attributesDescriptionCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *AttributesHandler) FindByTechnologies_type(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+attributesIdCName+"],"+"["+attributesNameCName+"]"+",["+attributesDescriptionCName+"]"+",["+attributesTechnologies_typeCName+"]"+",["+attributesTypeCName+"]"+",["+attributesTechnologies_invalidCName+"]"+"  FROM "+attributesTName+" WHERE "+attributesTechnologies_typeCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *AttributesHandler) FindByType(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+attributesIdCName+"],"+"["+attributesNameCName+"]"+",["+attributesDescriptionCName+"]"+",["+attributesTechnologies_typeCName+"]"+",["+attributesTypeCName+"]"+",["+attributesTechnologies_invalidCName+"]"+"  FROM "+attributesTName+" WHERE "+attributesTypeCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *AttributesHandler) ReadAll() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+attributesIdCName+"],"+"["+attributesNameCName+"]"+",["+attributesDescriptionCName+"]"+",["+attributesTechnologies_typeCName+"]"+",["+attributesTypeCName+"]"+",["+attributesTechnologies_invalidCName+"]"+"  FROM "+attributesTName, handler.ParseRows))
}

func (handler *AttributesHandler) ParseRows(rows *sql.Rows) per.IQueryResult {

	var Id int64

	var Name string

	var Description string

	var Technologies_type string

	var Type string

	results := []per.IDataItem{} //Attribute{}

	for rows.Next() {
		rows.Scan(&Id, &Name, &Description, &Technologies_type, &Type)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)

		res := data.Attribute{}

		res.Id = Id
		handler.Parent.LogDebugf("ParseRows", "Set '%v' for Id", Id)

		res.Name = Name
		handler.Parent.LogDebugf("ParseRows", "Set '%v' for Name", Name)

		res.Description = Description
		handler.Parent.LogDebugf("ParseRows", "Set '%v' for Description", Description)

		res.Technologies_type = Technologies_type
		handler.Parent.LogDebugf("ParseRows", "Set '%v' for Technologies_type", Technologies_type)

		res.Type = Type
		handler.Parent.LogDebugf("ParseRows", "Set '%v' for Type", Type)

		results = append(results, res)
	}
	return SQLL.NewDataQueryResult(true, results)
}
