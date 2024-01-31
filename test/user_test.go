package test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/dreamkanathip/LabExamFinal/entity"
	. "github.com/onsi/gomega"
)

func TestUserValid(t *testing.T)  {
	//เคสถูกต้องทั้งหมด
	g := NewGomegaWithT(t)

	t.Run(userOK,func(t *testing.T) {
		user := entity.User {
			StudentID: "B6418045",
			Password: "12345678",
			Age: 22,
			FirstName: "Kanathip",
			Email: "dreamkanathip@gmail.com",
			Profile: "abcdefg",
			Phone: "0123456789",
			URL: "https://github.com/dreamkanathip/LabExamFinal",
		}
		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
	//เคส studentID ผิด
	t.Run(studentID is wrong,func(t *testing.T) {
		user := entity.User {
			StudentID: "A6418045", //ตัว A บังคับBMD เท่านั้น
			Password: "12345678",
			Age: 22,
			FirstName: "Kanathip",
			Email: "dreamkanathip@gmail.com",
			Profile: "abcdefg",
			Phone: "0123456789",
			URL: "https://github.com/dreamkanathip/LabExamFinal",
		}
		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("ขึ้นต้นด้วย BMD เท่านั้น"))
	})
}