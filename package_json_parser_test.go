package yarnstart_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
	yarnstart "github.com/paketo-buildpacks/yarn-start"
	"github.com/sclevine/spec"
)

func testPackageJsonParser(t *testing.T, context spec.G, it spec.S) {
	Expect := NewWithT(t).Expect

	context("when parsing a valid package.json with start scripts", func() {
		var packageLocation string
		var workingDir string

		it.Before(func() {
			var err error
			workingDir, err = os.MkdirTemp("", "working-dir")
			Expect(err).NotTo(HaveOccurred())

			content := `{
				"scripts": {
					"poststart": "echo \"poststart\"",
					"prestart": "echo \"prestart\"",
					"start": "echo \"start\" && node server.js"
				}
			  },			
			`

			packageLocation = filepath.Join(workingDir, "package.json")
			Expect(os.WriteFile(packageLocation, []byte(content), 0600)).To(Succeed())
		})

		it.After(func() {
			Expect(os.RemoveAll(workingDir)).To(Succeed())
		})

		it("successfully extracts the scripts information", func() {
			pkg, err := yarnstart.NewPackageJsonFromPath(packageLocation)
			Expect(err).ToNot(HaveOccurred())

			Expect(pkg.Scripts.Start).To(ContainSubstring(`echo "start" && node server.js`))
			Expect(pkg.Scripts.PreStart).To(Equal(`echo "prestart"`))
			Expect(pkg.Scripts.PostStart).To(Equal(`echo "poststart"`))
		})
	})

	context("when the package.json is not a valid json file", func() {
		var packageLocation string
		var workingDir string

		it.Before(func() {
			var err error
			workingDir, err = os.MkdirTemp("", "working-dir")
			Expect(err).NotTo(HaveOccurred())

			packageLocation = filepath.Join(workingDir, "package.json")
			Expect(os.WriteFile(packageLocation, nil, 0600)).To(Succeed())
		})

		it.After(func() {
			Expect(os.RemoveAll(workingDir)).To(Succeed())
		})

		it("fails parsing", func() {
			_, err := yarnstart.NewPackageJsonFromPath(packageLocation)
			Expect(err).To(HaveOccurred())
		})
	})

	context("when the path to package.json is invalid", func() {
		it("fails parsing", func() {
			_, err := yarnstart.NewPackageJsonFromPath("/tmp/non-existent")
			Expect(err).To(HaveOccurred())
		})
	})
}
