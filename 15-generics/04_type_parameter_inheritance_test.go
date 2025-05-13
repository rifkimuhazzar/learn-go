package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

type String string

func (m *String) GetName() string {
	return string(*m)
}

func TestTypeParameterInheritance(t *testing.T) {
	assert.Equal(t, "Monkey D. Luffy", GetName(&MyManager{Name: "Monkey D. Luffy"}))
	fmt.Println(GetName(&MyManager{Name: "Monkey D. Luffy"}))

	assert.Equal(t, "Roronoa Zoro", GetName[VicePresident](&MyVicePresident{Name: "Roronoa Zoro"}))
	fmt.Println(GetName[VicePresident](&MyVicePresident{Name: "Roronoa Zoro"}))

	str := String("Hello World")
	assert.Equal(t, "Hello World", GetName(&str))
	fmt.Println(GetName(&str))
}
