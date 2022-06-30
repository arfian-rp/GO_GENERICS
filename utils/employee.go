package utils

//Employee Interface

type Employyee interface {
	GetName() string
}

func GetName[T Employyee](p T) string {
	return p.GetName()
}

//Manager Interface

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

//VicePresident Interface

type VicePresident interface {
	GetName() string
	GetVPName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVPName() string {
	return m.Name
}
