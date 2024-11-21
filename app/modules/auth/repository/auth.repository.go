package repository

import (
	"golang/app/modules/auth/model"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	var user model.User

    r.db.First(&user, id)

	return &user, nil
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindAll(offset, limit int) ([]model.User, int64, error) {
    var users []model.User
    var total int64

    // Count total records
    if err := r.db.Model(&model.User{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // Fetch paginated data
    if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
        return nil, 0, err
    }

    return users, total, nil
}

func (r *userRepository) Update(user *model.User) (*model.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	
	return user, nil
}