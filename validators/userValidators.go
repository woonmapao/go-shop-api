package validators

import (
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/models"
)

// Check if a user with the given ID exists in the database
func userExists(userID int) bool {
	var user models.User
	result := initializer.DB.First(&user, userID)
	return result.RowsAffected > 0
}
