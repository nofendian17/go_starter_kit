package healthcheck

import (
	"context"
	"fmt"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/model/response"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

// Health returns the health status of the system.
//
// Context ctx is used to carry deadlines, cancellation signals, and other request-scoped values.
// Returns a pointer to response.HealthResponse and an error.
func (u *useCase) Health(ctx context.Context) (*response.HealthResponse, error) {

	c, err := cpu.PercentWithContext(ctx, 0, false)
	if err != nil {
		return nil, err
	}

	m, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return &response.HealthResponse{
		Version: u.config.Application.Version,
		Uptime:  time.Since(u.startAt).String(),
		CPU:     fmt.Sprintf("%.2f%%", c[0]),
		Memory:  fmt.Sprintf("%f%%", m.UsedPercent),
	}, nil
}
