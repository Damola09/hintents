package simulator

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"
)

// TestPerfRegression loops 3 times, records ns/op, takes the mean,
// and compares it against the baseline from benchmark_sim_baseline.txt.
func TestPerfRegression(t *testing.T) {
	baselineData, err := os.ReadFile("benchmark_sim_baseline.txt")
	if err != nil {
		t.Skipf("No baseline found: %v", err)
	}

	re := regexp.MustCompile(`([0-9.]+)s`)
	matches := re.FindStringSubmatch(string(baselineData))
	if len(matches) < 2 {
		t.Fatalf("Could not parse baseline: %v", string(baselineData))
	}
	
	baselineSecs, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		t.Fatalf("Failed to parse baseline float: %v", err)
	}
	baselineMeanNs := float64(baselineSecs * 1e9)

	const iters = 3
	var totalNs float64
	for i := 0; i < iters; i++ {
		start := time.Now()
		// Performance critical path goes here
		time.Sleep(1 * time.Millisecond)
		totalNs += float64(time.Since(start).Nanoseconds())
	}

	meanNs := totalNs / float64(iters)
	fmt.Printf("ns/op: %.2f\n", meanNs)
	fmt.Printf("baseline (ns): %.2f\n", baselineMeanNs)

	// Compare mean against baseline
	if meanNs > baselineMeanNs*1.2 { // allowing arbitrary 20% jitter
		t.Errorf("Performance regression: mean %v ns > baseline %v ns", meanNs, baselineMeanNs)
	}
}
