package countthemonkeys

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCountTheMonkeys(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountTheMonkeys Suite")
}
