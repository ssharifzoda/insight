package database

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"green/internal/models"
)

type Authorization interface {
	GetUserByUsername(username string) (*models.User, error)
	UpdateRefreshToken(authParams *models.UserAuth) error
	ChangeUserPassword(request *models.ChangePassword) error
	GetTokensByUsername(username string) (string, string, error)
	GetUserPermissionsByUsername(username string) ([]string, error)
}

type User interface {
	AddNewUser(params *models.UserInput, authPrams *models.UserAuth) error
	ChangeUserParams(userParams *models.UserInput, authParams *models.UserAuth) error
	ChangeUserActivity(userId int, block bool) error
	GetAllUser(limit, offset int) ([]*models.User, int, error)
	SearchUser(search string) ([]*models.User, int, error)
	GetUserById(userId int) (*models.User, error)
	DeleteUser(userId int) error
	GetCities() ([]*models.City, error)
	GetUserRoles() ([]*models.Role, error)
	GetUsersExcel(rawSQl string) ([]*models.UserOut, error)
	CheckUserById(userId int) (bool, error)
	CheckUsernameAndPhone(username, phone string) (bool, bool, error)
}

type Application interface {
	CreatePasswordApplicationByPhone(phone string) error
	CheckPhoneNumber(phone string) (string, error)
	CreatePasswordApplicationByMail(email string) error
}

type Report interface {
	CreateNewReport(report []*models.Report) error
	GetMyReports(username string, cropId, limit, offset int) ([]*models.Report, error)
	SearchMyReports(username string, reportId int) ([]*models.Report, error)
	UpdateMyReport(input *models.Report) error
	GetReportById(id int) (*models.Report, error)
	GetReports(filterPart string, count, offset int) ([]*models.ReportList, int, error)
	SearchReport(search string) ([]*models.ReportList, int, error)
	ChangeReportState(id int, changeState bool) error
	GetReportsExcel(sqlRaw string) ([]*models.Report, error)
	GetReportStates() ([]*models.ReportState, error)
	GetCrops(id pq.Int64Array) ([]*models.Crops, error)
	GetDiseases() ([]*models.PlantDiseases, error)
}

type Notification interface {
	GetMyNotifications(userId int) ([]*models.Notification, error)
	ChangeNotificationState(userId int) error
}

type Data interface {
	GetMyData(username, filterSQl, dateSQl string, limit, offset int) ([]*models.Data, int, error)
	SearchMyData(username, search string) ([]*models.Data, int, error)
	UploadDataFile(username, filename string) error
	UploadData(username string, items []*models.Data) error
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
	EditCrop(crop *models.Crops) error
	DeleteCrop(id int, show bool) error
	AddPermission(permission *models.Permission) error
	AddRoute(route *models.Route) error
}

type Database struct {
	Authorization
	User
	Application
	Report
	Notification
	Data
	Worker
	Admin
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{
		Authorization: NewAuthPostgres(conn),
		User:          NewUserPostgres(conn),
		Application:   NewApplicationPostgres(conn),
		Report:        NewReportPostgres(conn),
		Notification:  NewNotificationPostgres(conn),
		Data:          NewDataPostgres(conn),
		Worker:        NewWorkerPostgres(conn),
		Admin:         NewAdminPostgres(conn),
	}
}
