
package helpers
import (
	"e-commerce/models"
	"e-commerce/utilities"
)
var db = utilities.ConnecDB()
func ValidateUser(user_id string) string {
	var user models.User
	err := db.Where("id = ? ", user_id).First(&user).Error
	if err != nil {
		return "user not found"
	}
	return "success"
}