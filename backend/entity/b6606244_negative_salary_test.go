package entity_test

import (
	"b6606244/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeSalaryNegativeValidate(t *testing.T) {
	t.Run("Salary < 15000", func(t *testing.T) {
		g := NewGomegaWithT(t)

		employees := entity.Employees{
			Name:         "John Doe",
			Salary:       14999.00,
			EmployeeCode: "HR-1024",
		}
		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})

	t.Run("Salary > 200,000", func(t *testing.T) {
		g := NewGomegaWithT(t)

		employees := entity.Employees{
			Name:         "John Doe",
			Salary:       210000.00,
			EmployeeCode: "HR-1024",
		}
		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})
}
