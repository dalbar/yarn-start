package yarnstart_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitGoBuild(t *testing.T) {
	suite := spec.New("yarn-start", spec.Report(report.Terminal{}), spec.Sequential())
	suite("Build", testBuild)
	suite("Detect", testDetect)
	suite("ProjectPathParser", testProjectPathParser)
	suite.Run(t)
}
