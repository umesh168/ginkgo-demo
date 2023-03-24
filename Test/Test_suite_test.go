package Test_test

import (
	"demo/structures"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = Describe("Employee", func() {
	Context("check employee first name", func() {
		It("can extract the author's last name", func() {
			employee := structures.Employee{
				FirstName: "Les Miserables",
				LastName:  "Victor Hugo",
				ID:        2783,
			}
			response := employee.FirstNameData()
			Expect(response).To(Equal("Les Miserables"))
		})
	})

})
