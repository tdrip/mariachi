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

// Table fields

// technologies
const technologiesTName = "technologies"

// Primay Key: id
const technologiesIdCName = "id"

// name
const technologiesNameCName = "name"

// type
const technologiesTypeCName = "type"

// icon
const technologiesIconCName = "icon"

// type_description
const technologiesType_descriptionCName = "type_description"

// HANDLER

type TechnologiesHandler struct {
	per.IStorageHandler
	Parent   *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewTechnologiesHandler(datastore *SQLL.SQLLiteDatastore) *TechnologiesHandler {
	ds := TechnologiesHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler
func (handler *TechnologiesHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *TechnologiesHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for Technologie
func (handler *TechnologiesHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures", "Executing Query")
	return handler.Executor.ExecuteQuery(`CREATE TABLE IF NOT EXISTS technologies (
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
)`)
}

// End Istorage

// This function Technologie removes all data for the table
func (handler *TechnologiesHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + technologiesTName))
}

// This adds Technologie to the database
func (handler *TechnologiesHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Technologie)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO "+technologiesTName+" ( "+"["+technologiesNameCName+"]"+",["+technologiesTypeCName+"]"+",["+technologiesIconCName+"]"+",["+technologiesType_descriptionCName+"]"+" ) VALUES (?,?,?,?)", data.Name, data.Type, data.Icon, data.Type_description))
}

func (handler *TechnologiesHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Technologie)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE "+technologiesTName+" SET "+"["+technologiesNameCName+"] = ? "+",["+technologiesTypeCName+"] = ? "+",["+technologiesIconCName+"] = ? "+",["+technologiesType_descriptionCName+"] = ? "+"  WHERE ["+technologiesIdCName+"] = ?", data.Name, data.Type, data.Icon, data.Type_description, data.Id))
}

func (handler *TechnologiesHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return SQLL.ResultToSQLLiteQueryResult(data)
}

func (handler *TechnologiesHandler) FindById(SearchData int64) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+technologiesIdCName+"],"+"["+technologiesNameCName+"]"+",["+technologiesTypeCName+"]"+",["+technologiesIconCName+"]"+",["+technologiesType_descriptionCName+"]"+"  FROM "+technologiesTName+" WHERE "+technologiesIdCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *TechnologiesHandler) FindByName(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+technologiesIdCName+"],"+"["+technologiesNameCName+"]"+",["+technologiesTypeCName+"]"+",["+technologiesIconCName+"]"+",["+technologiesType_descriptionCName+"]"+"  FROM "+technologiesTName+" WHERE "+technologiesNameCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *TechnologiesHandler) FindByType(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+technologiesIdCName+"],"+"["+technologiesNameCName+"]"+",["+technologiesTypeCName+"]"+",["+technologiesIconCName+"]"+",["+technologiesType_descriptionCName+"]"+"  FROM "+technologiesTName+" WHERE "+technologiesTypeCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *TechnologiesHandler) FindByIcon(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+technologiesIdCName+"],"+"["+technologiesNameCName+"]"+",["+technologiesTypeCName+"]"+",["+technologiesIconCName+"]"+",["+technologiesType_descriptionCName+"]"+"  FROM "+technologiesTName+" WHERE "+technologiesIconCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *TechnologiesHandler) FindByType_description(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+technologiesIdCName+"],"+"["+technologiesNameCName+"]"+",["+technologiesTypeCName+"]"+",["+technologiesIconCName+"]"+",["+technologiesType_descriptionCName+"]"+"  FROM "+technologiesTName+" WHERE "+technologiesType_descriptionCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *TechnologiesHandler) ReadAll() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+technologiesIdCName+"],"+"["+technologiesNameCName+"]"+",["+technologiesTypeCName+"]"+",["+technologiesIconCName+"]"+",["+technologiesType_descriptionCName+"]"+"  FROM "+technologiesTName, handler.ParseRows))
}

func (handler *TechnologiesHandler) ParseRows(rows *sql.Rows) per.IQueryResult {

	var Id *int64

	var Name *string

	var Type *string

	var Icon *string

	var Type_description *string

	results := []per.IDataItem{} //Technologie{}

	for rows.Next() {
		err := rows.Scan(&Id, &Name, &Type, &Icon, &Type_description)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)
		if err != nil {
			handler.Parent.LogErrorEf("ParseRows", "Row Scan error: %s", err)
		} else {
			res := data.Technologie{}

			if Id != nil {
				res.Id = *Id
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Id", *Id)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			if Name != nil {
				res.Name = *Name
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Name", *Name)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			if Type != nil {
				res.Type = *Type
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Type", *Type)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			if Icon != nil {
				res.Icon = *Icon
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Icon", *Icon)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			if Type_description != nil {
				res.Type_description = *Type_description
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Type_description", *Type_description)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			results = append(results, res)
		}

	}
	return SQLL.NewDataQueryResult(true, results)
}
