package validators

import (
	"github.com/woonmapao/go-shop-api/initializer"
	"github.com/woonmapao/go-shop-api/models"
)

// Check if a user with the given ID exists in the database
func userExists(userID int) bool {
	var user models.User
	result := initializer.DB.First(&user, userID)
	return result.RowsAffected > 0
}
