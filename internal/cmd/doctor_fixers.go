package cmd

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

// FixSimulatorBinary wraps the cargo build execution in a timeout context
func FixSimulatorBinary(ctx context.Context) error {
	buildCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(buildCtx, "cargo", "build", "--release")
	if err := cmd.Run(); err != nil {
		if errors.Is(buildCtx.Err(), context.DeadlineExceeded) {
			return errors.New("Build timed out")
		}
		return fmt.Errorf("cargo build failed: %w", err)
	}
	return nil
}
