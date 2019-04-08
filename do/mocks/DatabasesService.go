// Code generated by mockery v1.0.0. DO NOT EDIT.

// Generated: please do not edit by hand

package mocks

import do "github.com/digitalocean/doctl/do"
import godo "github.com/digitalocean/godo"
import mock "github.com/stretchr/testify/mock"

// DatabasesService is an autogenerated mock type for the DatabasesService type
type DatabasesService struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *DatabasesService) Create(_a0 *godo.DatabaseCreateRequest) (*do.Database, error) {
	ret := _m.Called(_a0)

	var r0 *do.Database
	if rf, ok := ret.Get(0).(func(*godo.DatabaseCreateRequest) *do.Database); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Database)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.DatabaseCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDB provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) CreateDB(_a0 string, _a1 *godo.DatabaseCreateDBRequest) (*do.DatabaseDB, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.DatabaseDB
	if rf, ok := ret.Get(0).(func(string, *godo.DatabaseCreateDBRequest) *do.DatabaseDB); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabaseDB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.DatabaseCreateDBRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePool provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) CreatePool(_a0 string, _a1 *godo.DatabaseCreatePoolRequest) (*do.DatabasePool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.DatabasePool
	if rf, ok := ret.Get(0).(func(string, *godo.DatabaseCreatePoolRequest) *do.DatabasePool); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabasePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.DatabaseCreatePoolRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReplica provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) CreateReplica(_a0 string, _a1 *godo.DatabaseCreateReplicaRequest) (*do.DatabaseReplica, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.DatabaseReplica
	if rf, ok := ret.Get(0).(func(string, *godo.DatabaseCreateReplicaRequest) *do.DatabaseReplica); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabaseReplica)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.DatabaseCreateReplicaRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) CreateUser(_a0 string, _a1 *godo.DatabaseCreateUserRequest) (*do.DatabaseUser, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.DatabaseUser
	if rf, ok := ret.Get(0).(func(string, *godo.DatabaseCreateUserRequest) *do.DatabaseUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabaseUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.DatabaseCreateUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0
func (_m *DatabasesService) Delete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDB provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) DeleteDB(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePool provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) DeletePool(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReplica provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) DeleteReplica(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) DeleteUser(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *DatabasesService) Get(_a0 string) (*do.Database, error) {
	ret := _m.Called(_a0)

	var r0 *do.Database
	if rf, ok := ret.Get(0).(func(string) *do.Database); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Database)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnection provides a mock function with given fields: _a0
func (_m *DatabasesService) GetConnection(_a0 string) (*do.DatabaseConnection, error) {
	ret := _m.Called(_a0)

	var r0 *do.DatabaseConnection
	if rf, ok := ret.Get(0).(func(string) *do.DatabaseConnection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabaseConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDB provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) GetDB(_a0 string, _a1 string) (*do.DatabaseDB, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.DatabaseDB
	if rf, ok := ret.Get(0).(func(string, string) *do.DatabaseDB); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabaseDB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaintenance provides a mock function with given fields: _a0
func (_m *DatabasesService) GetMaintenance(_a0 string) (*do.DatabaseMaintenanceWindow, error) {
	ret := _m.Called(_a0)

	var r0 *do.DatabaseMaintenanceWindow
	if rf, ok := ret.Get(0).(func(string) *do.DatabaseMaintenanceWindow); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabaseMaintenanceWindow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPool provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) GetPool(_a0 string, _a1 string) (*do.DatabasePool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.DatabasePool
	if rf, ok := ret.Get(0).(func(string, string) *do.DatabasePool); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabasePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplica provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) GetReplica(_a0 string, _a1 string) (*do.DatabaseReplica, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.DatabaseReplica
	if rf, ok := ret.Get(0).(func(string, string) *do.DatabaseReplica); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabaseReplica)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplicaConnection provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) GetReplicaConnection(_a0 string, _a1 string) (*do.DatabaseConnection, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.DatabaseConnection
	if rf, ok := ret.Get(0).(func(string, string) *do.DatabaseConnection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabaseConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) GetUser(_a0 string, _a1 string) (*do.DatabaseUser, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.DatabaseUser
	if rf, ok := ret.Get(0).(func(string, string) *do.DatabaseUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.DatabaseUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *DatabasesService) List() (do.Databases, error) {
	ret := _m.Called()

	var r0 do.Databases
	if rf, ok := ret.Get(0).(func() do.Databases); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.Databases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBackups provides a mock function with given fields: _a0
func (_m *DatabasesService) ListBackups(_a0 string) (do.DatabaseBackups, error) {
	ret := _m.Called(_a0)

	var r0 do.DatabaseBackups
	if rf, ok := ret.Get(0).(func(string) do.DatabaseBackups); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.DatabaseBackups)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDBs provides a mock function with given fields: _a0
func (_m *DatabasesService) ListDBs(_a0 string) (do.DatabaseDBs, error) {
	ret := _m.Called(_a0)

	var r0 do.DatabaseDBs
	if rf, ok := ret.Get(0).(func(string) do.DatabaseDBs); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.DatabaseDBs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPools provides a mock function with given fields: _a0
func (_m *DatabasesService) ListPools(_a0 string) (do.DatabasePools, error) {
	ret := _m.Called(_a0)

	var r0 do.DatabasePools
	if rf, ok := ret.Get(0).(func(string) do.DatabasePools); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.DatabasePools)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReplicas provides a mock function with given fields: _a0
func (_m *DatabasesService) ListReplicas(_a0 string) (do.DatabaseReplicas, error) {
	ret := _m.Called(_a0)

	var r0 do.DatabaseReplicas
	if rf, ok := ret.Get(0).(func(string) do.DatabaseReplicas); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.DatabaseReplicas)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: _a0
func (_m *DatabasesService) ListUsers(_a0 string) (do.DatabaseUsers, error) {
	ret := _m.Called(_a0)

	var r0 do.DatabaseUsers
	if rf, ok := ret.Get(0).(func(string) do.DatabaseUsers); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.DatabaseUsers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) Migrate(_a0 string, _a1 *godo.DatabaseMigrateRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *godo.DatabaseMigrateRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Resize provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) Resize(_a0 string, _a1 *godo.DatabaseResizeRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *godo.DatabaseResizeRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMaintenance provides a mock function with given fields: _a0, _a1
func (_m *DatabasesService) UpdateMaintenance(_a0 string, _a1 *godo.DatabaseUpdateMaintenanceRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *godo.DatabaseUpdateMaintenanceRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
