package pkg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPackages(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "State Package Suite")
}
