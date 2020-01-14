package unit_test

import (
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestUnit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Suite")
}

func IsDockerAvailable() bool {
	cmd := exec.Command("docker", "system", "info")

	session, err := gexec.Start(cmd, nil, nil)
	Expect(err).ToNot(HaveOccurred())

	session.Wait(5 * time.Second)
	return session.ExitCode() == 0
}
