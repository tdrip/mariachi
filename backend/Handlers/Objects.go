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
/* CREATE TABLE objects (
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

// Table fields

// objects
const objectsTName = "objects"

// Primay Key: id
const objectsIdCName = "id"

// name
const objectsNameCName = "name"

// keyvaultname
const objectsKeyvaultnameCName = "keyvaultname"

// ip_address
const objectsIp_addressCName = "ip_address"

// host
const objectsHostCName = "host"

// port
const objectsPortCName = "port"

// issync
const objectsIssyncCName = "issync"

// technologies_type
const objectsTechnologies_typeCName = "technologies_type"

// technologies_name
const objectsTechnologies_nameCName = "technologies_name"

// comments
const objectsCommentsCName = "comments"

// certificate_dns
const objectsCertificate_dnsCName = "certificate_dns"

// ca_template
const objectsCa_templateCName = "ca_template"

// ca_common_name
const objectsCa_common_nameCName = "ca_common_name"

// tech
const objectsTechCName = "tech"

// region
const objectsRegionCName = "region"

// loadbalancer
const objectsLoadbalancerCName = "loadbalancer"

// HANDLER

type ObjectsHandler struct {
	per.IStorageHandler
	Parent   *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewObjectsHandler(datastore *SQLL.SQLLiteDatastore) *ObjectsHandler {
	ds := ObjectsHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler
func (handler *ObjectsHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *ObjectsHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for Object
func (handler *ObjectsHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures", "Executing Query")
	return handler.Executor.ExecuteQuery("CREATE TABLE IF NOT EXISTS objects ( id INTEGER NOT NULL, name VARCHAR (250),keyvaultname VARCHAR (50), ip_address VARCHAR (120),host VARCHAR (50),port VARCHAR (50), issync BOOLEAN, technologies_type VARCHAR (25), technologies_name VARCHAR (50),comments TEXT,certificate_dns TEXT,ca_template TEXT,ca_common_name VARCHAR (50),tech VARCHAR (35),region VARCHAR (50),loadbalancer VARCHAR (50),PRIMARY KEY (id)s)")
}

// End Istorage

// This function Object removes all data for the table
func (handler *ObjectsHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + objectsTName))
}

// This adds Object to the database
func (handler *ObjectsHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Object)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO "+objectsTName+" ( "+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+" ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", data.Name, data.Keyvaultname, data.Ip_address, data.Host, data.Port, data.Technologies_type, data.Technologies_name, data.Comments, data.Certificate_dns, data.Ca_template, data.Ca_common_name, data.Tech, data.Region, data.Loadbalancer))
}

func (handler *ObjectsHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Object)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE "+objectsTName+" SET "+"["+objectsNameCName+"] = ? "+",["+objectsKeyvaultnameCName+"] = ? "+",["+objectsIp_addressCName+"] = ? "+",["+objectsHostCName+"] = ? "+",["+objectsPortCName+"] = ? "+",["+objectsIssyncCName+"] = ? "+",["+objectsTechnologies_typeCName+"] = ? "+",["+objectsTechnologies_nameCName+"] = ? "+",["+objectsCommentsCName+"] = ? "+",["+objectsCertificate_dnsCName+"] = ? "+",["+objectsCa_templateCName+"] = ? "+",["+objectsCa_common_nameCName+"] = ? "+",["+objectsTechCName+"] = ? "+",["+objectsRegionCName+"] = ? "+",["+objectsLoadbalancerCName+"] = ? "+"  WHERE ["+objectsIdCName+"] = ?", data.Name, data.Keyvaultname, data.Ip_address, data.Host, data.Port, data.Technologies_type, data.Technologies_name, data.Comments, data.Certificate_dns, data.Ca_template, data.Ca_common_name, data.Tech, data.Region, data.Loadbalancer, data.Id))
}

func (handler *ObjectsHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return SQLL.ResultToSQLLiteQueryResult(data)
}

func (handler *ObjectsHandler) FindById(SearchData int64) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsIdCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByName(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsNameCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByKeyvaultname(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsKeyvaultnameCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByIp_address(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsIp_addressCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByHost(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsHostCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByPort(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsPortCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByTechnologies_type(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsTechnologies_typeCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByTechnologies_name(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsTechnologies_nameCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByComments(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsCommentsCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByCertificate_dns(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsCertificate_dnsCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByCa_template(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsCa_templateCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByCa_common_name(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsCa_common_nameCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByTech(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsTechCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByRegion(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsRegionCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) FindByLoadbalancer(SearchData string) SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName+" WHERE "+objectsLoadbalancerCName+" = ?", handler.ParseRows, SearchData))
}

func (handler *ObjectsHandler) ReadAll() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+"["+objectsIdCName+"],"+"["+objectsNameCName+"]"+",["+objectsKeyvaultnameCName+"]"+",["+objectsIp_addressCName+"]"+",["+objectsHostCName+"]"+",["+objectsPortCName+"]"+",["+objectsIssyncCName+"]"+",["+objectsTechnologies_typeCName+"]"+",["+objectsTechnologies_nameCName+"]"+",["+objectsCommentsCName+"]"+",["+objectsCertificate_dnsCName+"]"+",["+objectsCa_templateCName+"]"+",["+objectsCa_common_nameCName+"]"+",["+objectsTechCName+"]"+",["+objectsRegionCName+"]"+",["+objectsLoadbalancerCName+"]"+"  FROM "+objectsTName, handler.ParseRows))
}

func (handler *ObjectsHandler) ParseRows(rows *sql.Rows) per.IQueryResult {

	var Id int64

	var Name string

	var Keyvaultname string

	var Ip_address string

	var Host string

	var Port string

	var IsSync bool

	var Technologies_type string

	var Technologies_name string

	var Comments string

	var Certificate_dns string

	var Ca_template string

	var Ca_common_name string

	var Tech string

	var Region string

	var Loadbalancer string

	results := []per.IDataItem{} //Object{}

	for rows.Next() {
		err := rows.Scan(&Id, &Name, &Keyvaultname, &Ip_address, &Host, &Port, &IsSync, &Technologies_type, &Technologies_name, &Comments, &Certificate_dns, &Ca_template, &Ca_common_name, &Tech, &Region, &Loadbalancer)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)
		if err == nil {
			res := data.Object{}

			res.Id = Id
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Id", Id)

			res.Name = Name
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Name", Name)

			res.Keyvaultname = Keyvaultname
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Keyvaultname", Keyvaultname)

			res.Ip_address = Ip_address
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Ip_address", Ip_address)

			res.Host = Host
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Host", Host)

			res.Port = Port
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Port", Port)

			res.Technologies_type = Technologies_type
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Technologies_type", Technologies_type)

			res.Technologies_name = Technologies_name
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Technologies_name", Technologies_name)

			res.Comments = Comments
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Comments", Comments)

			res.Certificate_dns = Certificate_dns
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Certificate_dns", Certificate_dns)

			res.Ca_template = Ca_template
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Ca_template", Ca_template)

			res.Ca_common_name = Ca_common_name
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Ca_common_name", Ca_common_name)

			res.Tech = Tech
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Tech", Tech)

			res.Region = Region
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Region", Region)

			res.Loadbalancer = Loadbalancer
			handler.Parent.LogDebugf("ParseRows", "Set '%v' for Loadbalancer", Loadbalancer)

			results = append(results, res)
		} else {
			handler.Parent.LogErrorEf("ParseRows", "Got Error '%s'", err)
		}

	}
	return SQLL.NewDataQueryResult(true, results)
}
