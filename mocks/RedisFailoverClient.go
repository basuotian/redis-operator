// Code generated by mockery v1.0.0
package mocks

import failover "github.com/spotahome/redis-operator/pkg/failover"
import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/client-go/pkg/api/v1"
import v1beta1 "k8s.io/client-go/pkg/apis/apps/v1beta1"

// RedisFailoverClient is an autogenerated mock type for the RedisFailoverClient type
type RedisFailoverClient struct {
	mock.Mock
}

// CreateBootstrapPod provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) CreateBootstrapPod(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRedisService provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) CreateRedisService(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRedisStatefulset provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) CreateRedisStatefulset(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSentinelDeployment provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) CreateSentinelDeployment(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSentinelService provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) CreateSentinelService(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBootstrapPod provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) DeleteBootstrapPod(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRedisService provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) DeleteRedisService(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRedisStatefulset provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) DeleteRedisStatefulset(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSentinelDeployment provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) DeleteSentinelDeployment(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSentinelService provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) DeleteSentinelService(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllRedisfailovers provides a mock function with given fields:
func (_m *RedisFailoverClient) GetAllRedisfailovers() (*failover.RedisFailoverList, error) {
	ret := _m.Called()

	var r0 *failover.RedisFailoverList
	if rf, ok := ret.Get(0).(func() *failover.RedisFailoverList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*failover.RedisFailoverList)
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

// GetBootstrapName provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetBootstrapName(rFailover *failover.RedisFailover) string {
	ret := _m.Called(rFailover)

	var r0 string
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) string); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetBootstrapPod provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetBootstrapPod(rFailover *failover.RedisFailover) (*v1.Pod, error) {
	ret := _m.Called(rFailover)

	var r0 *v1.Pod
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) *v1.Pod); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failover.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRedisName provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetRedisName(rFailover *failover.RedisFailover) string {
	ret := _m.Called(rFailover)

	var r0 string
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) string); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetRedisPodsIPs provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetRedisPodsIPs(rFailover *failover.RedisFailover) ([]string, error) {
	ret := _m.Called(rFailover)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) []string); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failover.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRedisService provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetRedisService(rFailover *failover.RedisFailover) (*v1.Service, error) {
	ret := _m.Called(rFailover)

	var r0 *v1.Service
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) *v1.Service); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failover.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRedisStatefulset provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetRedisStatefulset(rFailover *failover.RedisFailover) (*v1beta1.StatefulSet, error) {
	ret := _m.Called(rFailover)

	var r0 *v1beta1.StatefulSet
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) *v1beta1.StatefulSet); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.StatefulSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failover.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSentinelDeployment provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetSentinelDeployment(rFailover *failover.RedisFailover) (*v1beta1.Deployment, error) {
	ret := _m.Called(rFailover)

	var r0 *v1beta1.Deployment
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) *v1beta1.Deployment); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failover.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSentinelName provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetSentinelName(rFailover *failover.RedisFailover) string {
	ret := _m.Called(rFailover)

	var r0 string
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) string); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetSentinelPodsIPs provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetSentinelPodsIPs(rFailover *failover.RedisFailover) ([]string, error) {
	ret := _m.Called(rFailover)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) []string); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failover.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSentinelService provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) GetSentinelService(rFailover *failover.RedisFailover) (*v1.Service, error) {
	ret := _m.Called(rFailover)

	var r0 *v1.Service
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) *v1.Service); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failover.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRedisStatefulset provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) UpdateRedisStatefulset(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSentinelDeployment provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) UpdateSentinelDeployment(rFailover *failover.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: rFailover
func (_m *RedisFailoverClient) UpdateStatus(rFailover *failover.RedisFailover) (*failover.RedisFailover, error) {
	ret := _m.Called(rFailover)

	var r0 *failover.RedisFailover
	if rf, ok := ret.Get(0).(func(*failover.RedisFailover) *failover.RedisFailover); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*failover.RedisFailover)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failover.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}