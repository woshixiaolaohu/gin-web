package system

type ServiceGroup struct {
	JwtService
	OperationRecordService
	ApiService
	MenuService
	BaseMenuService
	UserService
	CasbinService
	AuthorityService
	AuthorityBtnService
	DictionaryService
	DictionaryDetailService
	SysExportTemplateService
	InitDBService
}
