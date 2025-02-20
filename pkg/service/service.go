package service

import "gorm.io/gorm"

type Service struct {
	Contact    *Contact
	Experience *Experience
	Project    *Project
	Technology *Technology
	User       *User
	Visit      *Visit
}

type Base[T any] struct {
	DB *gorm.DB
}

func (s *Base[T]) GetAll() ([]T, error) {
	var records []T
	err := s.DB.Model(new(T)).Find(&records).Error
	return records, err
}

// Obtener un registro por ID
func (s *Base[T]) GetByID(id uint) (*T, error) {
	var record T
	err := s.DB.Model(new(T)).First(&record, id).Error
	return &record, err
}

func (s *Base[T]) GetWithFilters(filters map[string]interface{}) ([]T, error) {
	var records []T
	query := s.DB.Model(new(T))

	// Aplicar filtros dinámicamente
	for field, value := range filters {
		query = query.Where(field+" = ?", value)
	}

	err := query.Find(&records).Error
	return records, err
}

// Crear un nuevo registro
func (s *Base[T]) Create(record *T) error {
	return s.DB.Model(new(T)).Create(record).Error
}

// Actualizar un registro
func (s *Base[T]) Update(record *T) error {
	return s.DB.Model(new(T)).Save(record).Error
}

// Eliminar un registro
func (s *Base[T]) Delete(id uint) error {
	return s.DB.Model(new(T)).Delete(new(T), id).Error
}
