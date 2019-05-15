package api

type DocNameMap struct{
	DbName string
	CollName string
}

var MemberDocNameMap = DocNameMap{"test", "member"}
var AdmDocNameMap = DocNameMap{"test", "adm"}
var StoreDocNameMap = DocNameMap{"test", "store"}
var NewsDocNameMap = DocNameMap{"test", "news"}
var ProductDocNameMap = DocNameMap{"test", "product"}