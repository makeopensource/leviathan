package jobs

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
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	testBatchJobProcessor(t, 100)
}

func Test500Jobs(t *testing.T) {
	testBatchJobProcessor(t, 500)
}

func Test2000Jobs(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	testBatchJobProcessor(t, 2000)
}

func testBatchJobProcessor(t *testing.T, numJobs int) {
	SetupTest()

	testValues := maps.Values(testCases)

	for i := 0; i < numJobs; i++ {
		// Randomly choose from test cases
		testCaseIndex := rand.Intn(len(testValues))
		testCase := testValues[testCaseIndex]

		// Run the test for each job directly in the main test goroutine
		t.Run(fmt.Sprintf("Job_%d", i), func(t *testing.T) {
			// Create subtests for better reporting
			// Enable parallel execution for this subtest
			t.Parallel()
			testJobProcessor(t, testCase.studentFile, testCase.expectedOutput, defaultTimeout)
			fmt.Printf("Job %d finished\n", i)
		})
	}
}
