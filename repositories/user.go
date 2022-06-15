package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strconv"
	"telrobot/models"
	"telrobot/util"
	"telrobot/util/common"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository create a new UserRepository with db instance
func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&models.User{})

	return &UserRepository{db: db}
}

// List get all jobs by condition
func (r *UserRepository) List(name, email string, page, count int) ([]*models.User, uint, error) {
	db := r.db

	if name != "" {
		db = db.Where("name = ?", name)
	}
	if email != "" {
		db = db.Where("email = ?", email)
	}

	var total uint
	if err := db.Model(models.User{}).Count(&total).Error; err != nil {
		return nil, 0, util.CheckNotFoundError(err)
	}

	db = db.Offset(page-1).Count(count)

	var users []*models.User
	if err := db.Model(models.User{}).Find(&users).Error; err != nil {
		return nil, 0, util.CheckNotFoundError(err)
	}

	return users, total, nil
}

func (r *UserRepository) ListBySql(name, email string, limit, offset int) ([]*models.User, int, error) {
	var users []*models.User
	var total int
	var err error

	sql := "select uuid,name, email, phone, password from users where deleted_at is null "
	if name != "" {
		sql += " and name like '%" + name + "%'"
	}
	if email != "" {
		sql += " and email ='" + email + "'"
	}

	var resultCount []*models.ResultCount
	sqlCount := "select count(1) as Count from (" + sql + ") a "
	err = r.db.Raw(sqlCount).Scan(&resultCount).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}

	for _, value := range resultCount {
		total = value.Count
	}

	sql += "order by created_at "
	if limit != -1 {
		sql += " LIMIT " + strconv.Itoa(limit)
	}
	sql += " OFFSET " + strconv.Itoa(offset)

	err = r.db.Raw(sql).Scan(&users).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}

	return users, total, nil
}

// Get a job by id
func (r *UserRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, util.CheckNotFoundError(err)
	}

	return &user, nil
}

func (r *UserRepository) GetByUuid(uuid string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, util.CheckNotFoundError(err)
	}

	return &user, nil
}

// Update a user by id with data
func (r *UserRepository) UpdateByID(idStr string, data interface{}) (int64, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error("Update user id =", idStr)
	}

	query := r.db.Model(&models.User{Model: common.Model{ID: id}}).Updates(data)
	if err := query.Error; err != nil {
		return 0, err
	}

	return query.RowsAffected, nil
}

func (r *UserRepository) UpdateByUuid(uuid string, data interface{}) (int64, error) {

	query := r.db.Model(&models.User{Uuid: uuid}).Where("uuid = ?", uuid).Updates(data)
	if err := query.Error; err != nil {
		return 0, err
	}

	return query.RowsAffected, nil
}

// Create a job with data
func (r *UserRepository) Create(data *models.User) (*models.User, error) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}


// RemoveByID remove a user by id
func (r *UserRepository) RemoveByID(id string) (int64, error) {
	query := r.db.Delete(models.User{}, "id = ?", id)
	if err := query.Error; err != nil {
		return 0, err
	}

	return query.RowsAffected, nil
}

// RemoveByID remove a user by uuid
func (r *UserRepository) RemoveByUuid(uuid string) (int64, error) {
	query := r.db.Delete(models.User{}, "uuid = ?", uuid)
	if err := query.Error; err != nil {
		return 0, err
	}

	return query.RowsAffected, nil
}
