package network

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kubevirt.io/kubevirt/pkg/log"
)

func TestNetwork(t *testing.T) {
	log.Log.SetIOWriter(GinkgoWriter)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Network Suite")
}
