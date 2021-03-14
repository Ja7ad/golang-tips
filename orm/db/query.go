package db

type Query interface {
	AddBook(map[string]interface{}) error
	GetBook(uint) (map[string]interface{}, error)
	GetAllBook() (interface{}, error)
	UpdateBook(uint, map[string]interface{}) error
	DeleteBook(uint) error
}

// Add book
func (b Books) AddBook(book map[string]interface{}) error {
	if err := DB.Model(&b).Create(book).Error; err != nil {
		return err
	}
	return nil
}

// Get one book
func (b Books) GetBook(id uint) (interface{}, error) {
	if err := DB.First(&b, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return b, nil
}

// Get All books
func (Books) GetAllBook() (interface{}, error) {
	var allBooks []Books
	if err := DB.Find(&allBooks).Error; err != nil {
		return nil, err
	}
	return allBooks, nil
}

// Update one book
func (Books) UpdateBook(id uint, book map[string]interface{}) error {
	if err := DB.Model(Books{}).Where("id = ?", id).Updates(book).Error; err != nil {
		return err
	}
	return nil
}

// Delete one book
func (b Books) DeleteBook(id uint) error {
	if err := DB.Model(&b).Where("id = ?", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}
