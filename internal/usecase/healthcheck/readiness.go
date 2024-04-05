package healthcheck

import (
	"context"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/model/response"
	"time"
)

const (
	statusUP   = "UP"   // statusUP represents the status of a component as "UP"
	statusDown = "DOWN" // statusDown represents the status of a component as "DOWN"
)

// Readiness checks the readiness of the application by verifying different components like database and Redis connections.
func (u *useCase) Readiness(ctx context.Context) (*response.ReadinessResponse, error) {
	ready := response.ReadinessResponse{
		Status: statusUP,
		Checks: []response.Check{},
	}

	var checks []response.Check

	// Check database connection
	checkDB := response.Check{
		Name:   u.db.GormDB.Name(),
		Status: statusUP,
	}
	err := u.db.SqlDB.PingContext(ctx)
	if err != nil {
		checkDB.Status = statusDown
		checkDB.Error = err.Error()
		ready.Status = statusDown
	}
	checks = append(checks, checkDB)

	// Check Redis connection
	checkRedis := response.Check{
		Name:   "Redis",
		Status: statusUP,
	}
	err = u.cacheClient.Ping(ctx)
	if err != nil {
		checkRedis.Status = statusDown
		checkRedis.Error = err.Error()
		ready.Status = statusDown
	}
	checks = append(checks, checkRedis)

	ready.Checks = append(ready.Checks, checks...)

	return &ready, nil
}

// Health returns the health status of the system.
func (u *useCase) Health(ctx context.Context) (*response.HealthResponse, error) {
	return &response.HealthResponse{
		Version: u.config.Application.Version,
		Uptime:  time.Since(u.startAt).String(),
	}, nil
}
