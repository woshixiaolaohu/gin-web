package system

type ServiceGroup struct {
	JwtService
	OperationRecordService
	ApiService
	UserService
	CasbinService
}
