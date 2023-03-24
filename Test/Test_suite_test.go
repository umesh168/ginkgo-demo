package Test_test

import (
	"demo/rules"
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

var _ = Describe("Person", func() {
	Context("check Person structure validation", func() {
		It("can validate Person Struct", func() {
			employee := rules.Person{
				Name:                "Umesh Sonawane",
				Email:               "usonawane@flexera.com",
				Age:                 32,
				DriverLicenseNumber: "121212121212",
			}
			response := employee.ValidatePerson()
			Expect(response).To(BeNil())
		})
	})

})
