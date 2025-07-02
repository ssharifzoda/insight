package service

import (
	"github.com/lib/pq"
	"green/internal/database"
	"green/internal/models"
)

type Authorization interface {
	GetUserByUsername(username string) (*models.User, error)
	UpdateRefreshToken(username, accessToken, refreshToken string) error
	ChangeUserPassword(request *models.ChangePassword) error
	GetTokensByUsername(username string) (string, string, error)
	GetUserPermissionsByUsername(username string) ([]string, error)
}

type User interface {
	AddNewUser(params *models.UserInput) error
	ChangeUserParams(userParams *models.UserInput) error
	ChangeUserActivity(userId int, block bool) error
	GetAllUser(search string, page, count int) ([]*models.User, int, error)
	GetUserById(userId int) (*models.User, error)
	DeleteUser(userId int) error
	GetCities() ([]*models.City, error)
	GetUserRoles() ([]*models.Role, error)
	GetUsersExcel(roles, cities string) ([]*models.UserOut, error)
	CheckUserById(userId int) (bool, error)
	CheckUsernameAndPhone(username, phone string) (bool, bool, error)
}

type Application interface {
	CreatePasswordApplicationByPhone(phone string) error
	CheckPhoneNumber(phone string) (string, error)
	CreatePasswordApplicationByMail(email string) error
}

type Report interface {
	CreateNewReport(reports []*models.Report) error
	GetMyReports(username string, cropId, page, limit, reportId int) ([]*models.Report, error)
	UpdateMyReport(input *models.Report) error
	GetReportById(id int) (*models.Report, error)
	GetReports(search, dateFrom, dateTo string, crops, plantDiseases, states []string, page, count int) ([]*models.ReportList, int, error)
	ChangeReportState(id int, changeState bool) error
	GetReportsExcel(dateFrom, dateTo string, crops, plantDiseases, states []string) ([]*models.Report, error)
	GetReportStates() ([]*models.ReportState, error)
	GetCrops(id pq.Int64Array) ([]*models.Crops, error)
	GetDiseases() ([]*models.PlantDiseases, error)
}

type Notification interface {
	GetMyNotifications(userId int) (*models.NotifyResponse, error)
	ChangeNotificationState(userId int) error
}

type Data interface {
	GetMyData(username, search, dateFrom, dateTo string, filter []string, page, count int) ([]*models.Data, int, error)
	UploadDataFile(username, filename string) error
	ParsingDataFile(username string, file []byte) error
	GetDataList(userId int) ([]*models.UserDataFiles, error)
}

type Worker interface {
	GetEmailPasswordApplications() ([]*models.PasswordApplication, error)
	UpdateApplicationState(application *models.PasswordApplication) error
}

type Admin interface {
	GetAllPermissions() ([]*models.Permission, error)
	AddPermissionInRole(permissionId, roleId int) error
	DeleteRolePermission(permissionId, roleId int) error
	AddCrop(crop *models.Crops) error
	DeleteCrop(id int, show bool) error
	EditCrop(crop *models.Crops) error
	AddPermission(permission *models.Permission) error
	AddRoute(route *models.Route) error
}

type Service struct {
	Authorization
	User
	Application
	Report
	Notification
	Data
	Worker
	Admin
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization: NewAuthService(db),
		User:          NewUserService(db),
		Application:   NewApplicationService(db),
		Report:        NewReportService(db),
		Notification:  NewNotificationService(db),
		Data:          NewDataService(db),
		Worker:        NewWorkerService(db),
		Admin:         NewAdminService(db),
	}
}
