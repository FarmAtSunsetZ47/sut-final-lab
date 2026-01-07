package entity_test

import (
	"b6606244/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeCodeNegativeValidate(t *testing.T) {
	t.Run("Employee Code Invalid 1", func(t *testing.T) {
		g := NewGomegaWithT(t)

		employees := entity.Employees{
			Name:         "John Doe",
			Salary:       21000.00,
			EmployeeCode: "HR-abcd",
		}
		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))
	})

	t.Run("Employee Code Invalid 2", func(t *testing.T) {
		g := NewGomegaWithT(t)

		employees := entity.Employees{
			Name:         "John Doe",
			Salary:       21000.00,
			EmployeeCode: "ab-1234",
		}
		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))
	})

	t.Run("Employee Code Invalid 3", func(t *testing.T) {
		g := NewGomegaWithT(t)

		employees := entity.Employees{
			Name:         "John Doe",
			Salary:       21000.00,
			EmployeeCode: "ab-cdef",
		}
		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))
	})
}
