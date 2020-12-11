package pgudatastore

import (
	hndlr "github.com/tdrip/backend/Handlers"

	SQLL "github.com/tdrip/persist/pkg/sqllite"
)

type VTechsDatastore struct {
	Datastore *SQLL.SQLLiteDatastore

	//
	//AttributesHandler *AttributesHandler
	//
	//Available_certsHandler *Available_certsHandler
	//
	//Ca_templatesHandler *Ca_templatesHandler
	//
	//ConfigurationsHandler *ConfigurationsHandler
	//
	//ObjectsHandler *ObjectsHandler
	//
	//ProductsHandler *ProductsHandler
	//
	//TechnologiesHandler *TechnologiesHandler
	//
}

func CreateDataStorage(filename string) *VTechsDatastore {

	res := VTechsDatastore{}
	ds := SQLL.CreateOpenSQLLiteDatastore(filename)

	// tests the example
	ds.SetStorageHander("Generic", SQLL.NewSQLLiteTableHandler(ds))

	ds.SetStorageHander("attributes", hndlr.NewAttributesHandler(ds))

	ds.SetStorageHander("available_certs", hndlr.NewAvailable_certsHandler(ds))

	ds.SetStorageHander("ca_templates", hndlr.NewCa_templatesHandler(ds))

	ds.SetStorageHander("configurations", hndlr.NewConfigurationsHandler(ds))

	ds.SetStorageHander("objects", hndlr.NewObjectsHandler(ds))

	ds.SetStorageHander("products", hndlr.NewProductsHandler(ds))

	ds.SetStorageHander("technologies", hndlr.NewTechnologiesHandler(ds))

	ds.CreateStructures()

	res.Datastore = ds

	return &res
}

func (fds *VTechsDatastore) GetAttributesHandler() *hndlr.AttributesHandler {

	data, ok := fds.Datastore.GetStorageHandler("attributes")
	if ok {
		res := data.(*hndlr.AttributesHandler)
		return res
	}
	return nils
}

func (fds *VTechsDatastore) GetAvailable_certsHandler() *hndlr.Available_certsHandler {

	data, ok := fds.Datastore.GetStorageHandler("available_certs")
	if ok {
		res := data.(*hndlr.Available_certsHandler)
		return res
	}
	return nil
}

func (fds *VTechsDatastore) GetCa_templatesHandler() *hndlr.Ca_templatesHandler {

	data, ok := fds.Datastore.GetStorageHandler("ca_templates")
	if ok {
		res := data.(*hndlr.Ca_templatesHandler)
		return res
	}
	return nil
}

func (fds *VTechsDatastore) GetConfigurationsHandler() *hndlr.ConfigurationsHandler {

	data, ok := fds.Datastore.GetStorageHandler("configurations")
	if ok {
		res := data.(*hndlr.ConfigurationsHandler)
		return res
	}
	return nil
}

func (fds *VTechsDatastore) GetObjectsHandler() *hndlr.ObjectsHandler {

	data, ok := fds.Datastore.GetStorageHandler("objects")
	if ok {
		res := data.(*hndlr.ObjectsHandler)
		return res
	}
	return nil
}

func (fds *VTechsDatastore) GetProductsHandler() *hndlr.ProductsHandler {

	data, ok := fds.Datastore.GetStorageHandler("products")
	if ok {
		res := data.(*hndlr.ProductsHandler)
		return res
	}
	return nil
}

func (fds *VTechsDatastore) GetTechnologiesHandler() *hndlr.TechnologiesHandler {

	data, ok := fds.Datastore.GetStorageHandler("technologies")
	if ok {
		res := data.(*hndlr.TechnologiesHandler)
		return res
	}
	return nil
}
