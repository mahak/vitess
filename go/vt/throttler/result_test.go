/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package throttler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	resultIncreased = Result{
		Now:                          sinceZero(1234 * time.Millisecond),
		RateChange:                   increasedRate,
		lastRateChange:               sinceZero(1 * time.Millisecond),
		OldState:                     stateIncreaseRate,
		TestedState:                  stateIncreaseRate,
		NewState:                     stateIncreaseRate,
		OldRate:                      100,
		NewRate:                      100,
		Reason:                       "increased the rate",
		CurrentRate:                  99,
		GoodOrBad:                    goodRate,
		MemorySkipReason:             "",
		HighestGood:                  95,
		LowestBad:                    0,
		LagRecordNow:                 lagRecord(sinceZero(1234*time.Millisecond), 101, 1),
		LagRecordBefore:              replicationLagRecord{},
		PrimaryRate:                  99,
		GuessedReplicationRate:       0,
		GuessedReplicationBacklogOld: 0,
		GuessedReplicationBacklogNew: 0,
	}
	resultDecreased = Result{
		Now:                          sinceZero(5000 * time.Millisecond),
		RateChange:                   decreasedRate,
		lastRateChange:               sinceZero(1234 * time.Millisecond),
		OldState:                     stateIncreaseRate,
		TestedState:                  stateDecreaseAndGuessRate,
		NewState:                     stateDecreaseAndGuessRate,
		OldRate:                      200,
		NewRate:                      100,
		Reason:                       "decreased the rate",
		CurrentRate:                  200,
		GoodOrBad:                    badRate,
		MemorySkipReason:             "",
		HighestGood:                  95,
		LowestBad:                    200,
		LagRecordNow:                 lagRecord(sinceZero(5000*time.Millisecond), 101, 2),
		LagRecordBefore:              lagRecord(sinceZero(1234*time.Millisecond), 101, 1),
		PrimaryRate:                  200,
		GuessedReplicationRate:       150,
		GuessedReplicationBacklogOld: 10,
		GuessedReplicationBacklogNew: 20,
	}
	resultEmergency = Result{
		Now:                          sinceZero(10123 * time.Millisecond),
		RateChange:                   decreasedRate,
		lastRateChange:               sinceZero(5000 * time.Millisecond),
		OldState:                     stateDecreaseAndGuessRate,
		TestedState:                  stateEmergency,
		NewState:                     stateEmergency,
		OldRate:                      100,
		NewRate:                      50,
		Reason:                       "emergency state decreased the rate",
		CurrentRate:                  100,
		GoodOrBad:                    badRate,
		MemorySkipReason:             "",
		HighestGood:                  95,
		LowestBad:                    100,
		LagRecordNow:                 lagRecord(sinceZero(10123*time.Millisecond), 101, 23),
		LagRecordBefore:              lagRecord(sinceZero(5000*time.Millisecond), 101, 2),
		PrimaryRate:                  0,
		GuessedReplicationRate:       0,
		GuessedReplicationBacklogOld: 0,
		GuessedReplicationBacklogNew: 0,
	}
)

func TestResultString(t *testing.T) {
	testcases := []struct {
		r    Result
		want string
	}{
		{
			resultIncreased,
			`rate was: increased from: 100 to: 100
alias: cell1-0000000101 lag: 1s
last change: 1.2s rate: 99 good/bad? good skipped b/c:  good/bad: 95/0
state (old/tested/new): I/I/I 
lag before: n/a (n/a ago) rates (primary/replica): 99/0 backlog (old/new): 0/0
reason: increased the rate`,
		},
		{
			resultDecreased,
			`rate was: decreased from: 200 to: 100
alias: cell1-0000000101 lag: 2s
last change: 3.8s rate: 200 good/bad? bad skipped b/c:  good/bad: 95/200
state (old/tested/new): I/D/D 
lag before: 1s (3.8s ago) rates (primary/replica): 200/150 backlog (old/new): 10/20
reason: decreased the rate`,
		},
		{
			resultEmergency,
			`rate was: decreased from: 100 to: 50
alias: cell1-0000000101 lag: 23s
last change: 5.1s rate: 100 good/bad? bad skipped b/c:  good/bad: 95/100
state (old/tested/new): D/E/E 
lag before: 2s (5.1s ago) rates (primary/replica): 0/0 backlog (old/new): 0/0
reason: emergency state decreased the rate`,
		},
	}

	for _, tc := range testcases {
		got := tc.r.String()
		require.Equal(t, tc.want, got)
	}
}

func TestResultRing(t *testing.T) {
	// Test data.
	r1 := Result{Reason: "r1"}
	r2 := Result{Reason: "r2"}
	r3 := Result{Reason: "r3"}

	rr := newResultRing(2)

	// Use the ring partially.
	rr.add(r1)
	got, want := rr.latestValues(), []Result{r1}
	require.Equal(t, want, got, "items not correctly added to resultRing")

	// Use it fully.
	rr.add(r2)
	got, want = rr.latestValues(), []Result{r2, r1}
	require.Equal(t, want, got, "items not correctly added to resultRing")

	// Let it wrap.
	rr.add(r3)
	got, want = rr.latestValues(), []Result{r3, r2}
	require.Equal(t, want, got, "resultRing did not wrap correctly")
}
