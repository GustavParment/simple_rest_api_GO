package services

import (
	"errors"
	"simplerestapi/api/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return new(UserService)
}

/*TODO- IMPLEMENT USER SERVICE LOGIK
 */
func (us *UserService) GetUserByUsername(username string) (user *models.User, err error) {
	user = new(models.User)

}
