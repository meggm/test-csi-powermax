/*
Copyright © 2021-2025 Dell Inc. or its subsidiaries. All Rights Reserved.

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
package service

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // #nosec G108
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

var (
	testStatus    int
	testStartTime time.Time
)

func TestMain(m *testing.M) {
	testStatus = 0
	testStartTime = time.Now()

	go http.ListenAndServe("localhost:6060", nil) // #nosec G114

	if st := m.Run(); st > testStatus {
		testStatus = st
	}

	fmt.Printf("status %d\n", testStatus)

	os.Exit(testStatus)
}

func TestGoDog(t *testing.T) {
	fmt.Printf("starting godog...\n")
	runOptions := godog.Options{
		Format: "pretty",
		Paths:  []string{"features"},
		Tags:   "v1.0.0, v1.1.0, v1.2.0, v1.3.0, v1.4.0, v1.5.0, v1.6.0, v2.2.0, v2.3.0, v2.4.0, v2.5.0, v2.6.0, v2.7.0, v2.8.0, v2.9.0, v2.11.0, v2.12.0, v2.13.0, v2.14.0",
		// Tags:   "wip",
		// Tags: "resiliency", // uncomment to run all node resiliency related tests,
	}
	testStatus = godog.TestSuite{
		Name:                "CSI Powermax Unit Test",
		ScenarioInitializer: FeatureContext,
		Options:             &runOptions,
	}.Run()

	fmt.Printf("godog finished\n")
	if testStatus != 0 {
		t.Error("Error encountered in godog testing")
	}
}
