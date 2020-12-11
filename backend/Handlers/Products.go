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

// Table fields

// products
const productsTName = "products"

// Primay Key: id
const productsIdCName = "id"

// description
const productsDescriptionCName = "description"

// img
const productsImgCName = "img"

// HANDLER

type ProductsHandler struct {
	per.IStorageHandler
	Parent   *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewProductsHandler(datastore *SQLL.SQLLiteDatastore) *ProductsHandler {
	ds := ProductsHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler
func (handler *ProductsHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *ProductsHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for Product
func (handler *ProductsHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures", "Executing Query")
	return handler.Executor.ExecuteQuery(`CREATE TABLE IF NOT EXISTS products (
    id          INTEGER      NOT NULL,
    description VARCHAR (50),
    img         VARCHAR (50),
    PRIMARY KEY (
        id
    ),
    UNIQUE (
        description
    )
)`)
}

// End Istorage

// This function Product removes all data for the table
func (handler *ProductsHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + productsTName))
}

// This adds Product to the database
func (handler *ProductsHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Product)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO "+productsTName+" ( "+"["+productsDescriptionCName+"]"+",["+productsImgCName+"]"+" ) VALUES (?,?)", data.Description, data.Img))
}

func (handler *ProductsHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Product)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE "+productsTName+" SET "+"["+productsDescriptionCName+"] = ? "+",["+productsImgCName+"] = ? "+"  WHERE ["+productsIdCName+"] = ?", data.Description, data.Img, data.Id))
}

func (handler *ProductsHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return SQLL.ResultToSQLLiteQueryResult(data)
}

func (handler *ProductsHandler) FindById(SearchData int64) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+productsIdCName+"],"+"["+productsDescriptionCName+"]"+",["+productsImgCName+"]"+"  FROM "+productsTName+" WHERE "+productsIdCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ProductsHandler) FindByDescription(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+productsIdCName+"],"+"["+productsDescriptionCName+"]"+",["+productsImgCName+"]"+"  FROM "+productsTName+" WHERE "+productsDescriptionCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ProductsHandler) FindByImg(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+productsIdCName+"],"+"["+productsDescriptionCName+"]"+",["+productsImgCName+"]"+"  FROM "+productsTName+" WHERE "+productsImgCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ProductsHandler) ReadAll() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+productsIdCName+"],"+"["+productsDescriptionCName+"]"+",["+productsImgCName+"]"+"  FROM "+productsTName, handler.ParseRows))
}

func (handler *ProductsHandler) ParseRows(rows *sql.Rows) per.IQueryResult {

	var Id *int64

	var Description *string

	var Img *string

	results := []per.IDataItem{} //Product{}

	for rows.Next() {
		err := rows.Scan(&Id, &Description, &Img)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)
		if err != nil {
			handler.Parent.LogErrorEf("ParseRows", "Row Scan error: %s", err)
		} else {
			res := data.Product{}

			if Id != nil {
				res.Id = *Id
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Id", *Id)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			if Description != nil {
				res.Description = *Description
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Description", *Description)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			if Img != nil {
				res.Img = *Img
				handler.Parent.LogDebugf("ParseRows", "Set '%v' for Img", *Img)
			} else {
				handler.Parent.LogDebugf("ParseRows", "{.Name}} was NULL")
			}

			results = append(results, res)
		}

	}
	return SQLL.NewDataQueryResult(true, results)
}
