package labs

type LabStore interface {
	CreateLab(lab *Lab) error
	DeleteLab(id uint) error
	GetLab(id uint) (*Lab, error)
}
