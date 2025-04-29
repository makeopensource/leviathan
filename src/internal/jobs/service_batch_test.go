package jobs_test

import (
	"fmt"
	"golang.org/x/exp/maps"
	"math/rand"
	"os"
	"testing"
)

func Test50Jobs(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	testBatchJobProcessor(t, 50)
}

func Test100Jobs(t *testing.T) {
	testBatchJobProcessor(t, 100)
}

func Test500Jobs(t *testing.T) {
	testBatchJobProcessor(t, 500)
}
func Test1000Jobs(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	testBatchJobProcessor(t, 1000)
}

func Test2000Jobs(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	testBatchJobProcessor(t, 2000)
}

func Test10000Jobs(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	testBatchJobProcessor(t, 10000)
}

func testBatchJobProcessor(t *testing.T, numJobs int) {
	setupTest()

	testValues := maps.Values(testFuncs)

	for i := 0; i < numJobs; i++ {
		// Randomly choose from test cases
		testCaseIndex := rand.Intn(len(testValues))
		testCase := testValues[testCaseIndex]

		// Run the test for each job directly in the main test goroutine
		t.Run(fmt.Sprintf("Job_%d", i), func(t *testing.T) {
			// Create subtests for better reporting
			// Enable parallel execution for this subtest
			t.Parallel()
			testCase(t)
		})
	}
}
