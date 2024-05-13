package usecase_test

import (
	"errors"
	"mini-project/repository"
	"mini-project/usecase"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetAllPinpoints(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	pinpointRepo := repository.NewPinpointRepository(*db)
	userRepo := repository.NewUserRepository(*db)
	tierRepo := repository.NewTierRepository(*db)
	missionRepo := repository.NewMissionRepository(*db)
	userMissionRepo := repository.NewUserMissionRepository(*db)
	userUseCase := usecase.NewUserUseCase(userRepo, tierRepo)
	missionUseCase := usecase.NewMissionUseCase(missionRepo)
	userMissionUseCase := usecase.NewUserMissionUseCase(userMissionRepo)
	pinpointUsecase := usecase.NewPinpointUseCase(*userMissionUseCase, pinpointRepo, userRepo, *missionUseCase, *userUseCase)

	type Testcase struct {
		Case              string
		ExpectedPinpoints []repository.Pinpoint
		ExpectedError     error
	}

	testcases := []Testcase{
		{
			Case: "positive - select return 1 data",
			ExpectedPinpoints: []repository.Pinpoint{
				{
					ID:          1,
					UserID:      1,
					Name:        "pinpoint",
					Description: "description",
					Latitude:    1.0,
					Longitude:   1.0,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			ExpectedError: nil,
		},
		{
			Case: "positive - select return more than 1 data",
			ExpectedPinpoints: []repository.Pinpoint{
				{
					ID:          1,
					UserID:      1,
					Name:        "pinpoint",
					Description: "description",
					Latitude:    1.0,
					Longitude:   1.0,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					ID:          2,
					UserID:      1,
					Name:        "pinpoint",
					Description: "description",
					Latitude:    1.0,
					Longitude:   1.0,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			ExpectedError: nil,
		},
		{
			Case:              "negative - error when query",
			ExpectedPinpoints: []repository.Pinpoint{},
			ExpectedError:     errors.New("error query"),
		},
	}

	for _, testcase := range testcases {
		if testcase.ExpectedError != nil {
			mock.ExpectQuery(`SELECT`).WillReturnError(testcase.ExpectedError)
		} else {
			rows := sqlmock.NewRows([]string{"ID", "UserID", "Name", "Description", "Latitude", "Longitude", "CreatedAt", "UpdatedAt"})
			for _, pinpoint := range testcase.ExpectedPinpoints {
				rows.AddRow(pinpoint.ID, pinpoint.UserID, pinpoint.Name, pinpoint.Description, pinpoint.Latitude, pinpoint.Longitude, pinpoint.CreatedAt, pinpoint.UpdatedAt)
			}
			mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
		}

		pintpoints, err := pinpointUsecase.GetAllPinpoints()
		assert.Equal(t, testcase.ExpectedPinpoints, pintpoints)
		assert.Equal(t, testcase.ExpectedError, err)
	}

}

func TestGetPinpoint(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	pinpointRepo := repository.NewPinpointRepository(*db)
	userRepo := repository.NewUserRepository(*db)
	tierRepo := repository.NewTierRepository(*db)
	missionRepo := repository.NewMissionRepository(*db)
	userMissionRepo := repository.NewUserMissionRepository(*db)
	userUseCase := usecase.NewUserUseCase(userRepo, tierRepo)
	missionUseCase := usecase.NewMissionUseCase(missionRepo)
	userMissionUseCase := usecase.NewUserMissionUseCase(userMissionRepo)
	pinpointUsecase := usecase.NewPinpointUseCase(*userMissionUseCase, pinpointRepo, userRepo, *missionUseCase, *userUseCase)

	type Testcase struct {
		Case             string
		PinpointID       uint
		ExpectedPinpoint repository.Pinpoint
		ExpectedError    error
	}

	testcases := []Testcase{
		{
			Case:       "positive - select pinpoint id 1",
			PinpointID: 1,
			ExpectedPinpoint: repository.Pinpoint{
				ID:          1,
				UserID:      1,
				Name:        "pinpoint",
				Description: "description",
				Latitude:    1.0,
				Longitude:   1.0,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			ExpectedError: nil,
		},
		{
			Case:       "positive - select pinpoint id 2",
			PinpointID: 1,
			ExpectedPinpoint: repository.Pinpoint{
				ID:          2,
				UserID:      1,
				Name:        "pinpoint #2",
				Description: "description #2",
				Latitude:    789.0,
				Longitude:   789.0,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			ExpectedError: nil,
		},
		{
			Case:             "negative - error when query",
			ExpectedPinpoint: repository.Pinpoint{},
			ExpectedError:    errors.New("error query"),
		},
	}

	// setup data
	rows := sqlmock.NewRows([]string{"ID", "UserID", "Name", "Description", "Latitude", "Longitude", "CreatedAt", "UpdatedAt"})
	for _, testcase := range testcases {
		pinpoint := testcase.ExpectedPinpoint
		rows.AddRow(pinpoint.ID, pinpoint.UserID, pinpoint.Name, pinpoint.Description, pinpoint.Latitude, pinpoint.Longitude, pinpoint.CreatedAt, pinpoint.UpdatedAt)
	}

	for _, testcase := range testcases {
		if testcase.ExpectedError != nil {
			mock.ExpectQuery(`SELECT`).WillReturnError(testcase.ExpectedError)
		} else {
			mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
		}

		pintpoint, err := pinpointUsecase.GetPinpoint(testcase.PinpointID)
		assert.Equal(t, testcase.ExpectedPinpoint, pintpoint)
		assert.Equal(t, testcase.ExpectedError, err)
	}

}

func TestCreatePinpoint(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	pinpointRepo := repository.NewPinpointRepository(*db)
	userRepo := repository.NewUserRepository(*db)
	tierRepo := repository.NewTierRepository(*db)
	missionRepo := repository.NewMissionRepository(*db)
	userMissionRepo := repository.NewUserMissionRepository(*db)
	userUseCase := usecase.NewUserUseCase(userRepo, tierRepo)
	missionUseCase := usecase.NewMissionUseCase(missionRepo)
	userMissionUseCase := usecase.NewUserMissionUseCase(userMissionRepo)
	pinpointUsecase := usecase.NewPinpointUseCase(*userMissionUseCase, pinpointRepo, userRepo, *missionUseCase, *userUseCase)

	type Testcase struct {
		Case          string
		Pinpoint      repository.Pinpoint
		ExpectedError error
	}

	now := time.Now()

	testcases := []Testcase{
		{
			Case: "positive - success insert",
			Pinpoint: repository.Pinpoint{
				ID:          1,
				UserID:      1,
				Name:        "pinpoint",
				Description: "description",
				Latitude:    1.0,
				Longitude:   1.0,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			ExpectedError: nil,
		},
		{
			Case: "positive - success insert #2",
			Pinpoint: repository.Pinpoint{
				ID:          2,
				UserID:      1,
				Name:        "pinpoint #2",
				Description: "description #2",
				Latitude:    789.0,
				Longitude:   789.0,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			ExpectedError: nil,
		},
		{
			Case:          "negative - error when query",
			Pinpoint:      repository.Pinpoint{},
			ExpectedError: errors.New("error query"),
		},
	}
	rows := sqlmock.NewRows([]string{"ID", "UserID", "Name", "Description", "Latitude", "Longitude", "CreatedAt", "UpdatedAt"})
	for _, testcase := range testcases {
		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		if testcase.ExpectedError != nil {
			mock.ExpectBegin()
			mock.ExpectExec(`INSERT`).WillReturnError(testcase.ExpectedError)
			mock.ExpectRollback()
		} else {
			mock.ExpectBegin()
			mock.ExpectExec(`INSERT`).WillReturnResult(sqlmock.NewResult(int64(testcase.Pinpoint.ID), 1))
			mock.ExpectCommit()
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_missions`")).WillReturnRows(rows)
		}

		pinpoint, err := pinpointUsecase.CreatePinpoint(testcase.Pinpoint.UserID, testcase.Pinpoint.Name, testcase.Pinpoint.Description, testcase.Pinpoint.Latitude, testcase.Pinpoint.Longitude)
		assert.Equal(t, testcase.ExpectedError, err)
		if err == nil {

			pinpoint.CreatedAt = now
			pinpoint.UpdatedAt = now
			assert.Equal(t, testcase.Pinpoint, pinpoint)
		}
	}

}

func TestUpdatePinpoint(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	pinpointRepo := repository.NewPinpointRepository(*db)
	userRepo := repository.NewUserRepository(*db)
	tierRepo := repository.NewTierRepository(*db)
	missionRepo := repository.NewMissionRepository(*db)
	userMissionRepo := repository.NewUserMissionRepository(*db)
	userUseCase := usecase.NewUserUseCase(userRepo, tierRepo)
	missionUseCase := usecase.NewMissionUseCase(missionRepo)
	userMissionUseCase := usecase.NewUserMissionUseCase(userMissionRepo)
	pinpointUsecase := usecase.NewPinpointUseCase(*userMissionUseCase, pinpointRepo, userRepo, *missionUseCase, *userUseCase)

	type Testcase struct {
		Case          string
		PinpointID    uint
		Pinpoint      repository.Pinpoint
		ExpectedError error
	}

	now := time.Now()

	testcases := []Testcase{
		{
			Case: "positive - update pinpoint id 1",
			Pinpoint: repository.Pinpoint{
				ID:          1,
				UserID:      1,
				Name:        "pinpoint",
				Description: "description",
				Latitude:    1.0,
				Longitude:   1.0,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			ExpectedError: nil,
		},
		{
			Case:          "negative - error when query",
			Pinpoint:      repository.Pinpoint{},
			ExpectedError: errors.New("error query"),
		},
	}

	for _, testcase := range testcases {
		if testcase.ExpectedError != nil {
			mock.ExpectBegin()
			mock.ExpectExec(`UPDATE`).WillReturnError(testcase.ExpectedError)
			mock.ExpectRollback()
		} else {

			mock.ExpectBegin()
			mock.ExpectExec(`UPDATE`).WillReturnResult(sqlmock.NewResult(int64(testcase.Pinpoint.ID), 1))
			mock.ExpectCommit()
		}

		_, err := pinpointUsecase.UpdatePinpoint(testcase.Pinpoint.ID, testcase.Pinpoint.Name, testcase.Pinpoint.Description, testcase.Pinpoint.Latitude, testcase.Pinpoint.Longitude)
		assert.Equal(t, testcase.ExpectedError, err)
	}

}

func TestDeletePinpoint(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	pinpointRepo := repository.NewPinpointRepository(*db)
	userRepo := repository.NewUserRepository(*db)
	tierRepo := repository.NewTierRepository(*db)
	missionRepo := repository.NewMissionRepository(*db)
	userMissionRepo := repository.NewUserMissionRepository(*db)
	userUseCase := usecase.NewUserUseCase(userRepo, tierRepo)
	missionUseCase := usecase.NewMissionUseCase(missionRepo)
	userMissionUseCase := usecase.NewUserMissionUseCase(userMissionRepo)
	pinpointUsecase := usecase.NewPinpointUseCase(*userMissionUseCase, pinpointRepo, userRepo, *missionUseCase, *userUseCase)

	type Testcase struct {
		Case          string
		PinpointID    uint
		Pinpoint      repository.Pinpoint
		ExpectedError error
	}

	now := time.Now()

	testcases := []Testcase{
		{
			Case: "positive - delete pinpoint id 1",
			Pinpoint: repository.Pinpoint{
				ID:          1,
				UserID:      1,
				Name:        "pinpoint",
				Description: "description",
				Latitude:    1.0,
				Longitude:   1.0,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			ExpectedError: nil,
		},
		{
			Case:          "negative - error when query",
			Pinpoint:      repository.Pinpoint{},
			ExpectedError: errors.New("error query"),
		},
	}

	for _, testcase := range testcases {
		if testcase.ExpectedError != nil {
			mock.ExpectBegin()
			mock.ExpectExec(`DELETE`).WillReturnError(testcase.ExpectedError)
			mock.ExpectRollback()
		} else {

			mock.ExpectBegin()
			mock.ExpectExec(`DELETE`).WillReturnResult(sqlmock.NewResult(int64(testcase.Pinpoint.ID), 1))
			mock.ExpectCommit()
		}

		err := pinpointUsecase.DeletePinpoint(testcase.Pinpoint.ID)
		assert.Equal(t, testcase.ExpectedError, err)
	}

}
