package entity_test

import (
	"b6606244/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeValidate(t *testing.T) {
	t.Run("Valid Employees", func(t *testing.T) {
		g := NewGomegaWithT(t)

		employees := entity.Employees{
			Name:         "John Doe",
			Salary:       18000.00,
			EmployeeCode: "HR-1024",
		}
		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
