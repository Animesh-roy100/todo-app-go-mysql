package seed

import (
	"log"
	"todolist/api/models"

	"github.com/jinzhu/gorm"
)

// Load is a function that loads the seed data into the database
// It takes a pointer to a gorm.DB object as an argument
// It creates the users and todos tables if they don't exist
// It also adds a foreign key to the todos table
// It then creates two users and two todos and associates them with each other
func Load(db *gorm.DB) {
 // Drop the users and todos tables if they exist
 err := db.Debug().DropTableIfExists(&models.ToDo{}, &models.User{}).Error
 if err != nil {
  // Log the error if it occurs
  log.Fatalf("cannot drop table: %v", err)
 }
 // Create the users and todos tables
 err = db.Debug().AutoMigrate(&models.User{}, &models.ToDo{}).Error
 if err != nil {
  // Log the error if it occurs
  log.Fatalf("cannot migrate table: %v", err)
 }

 // Add a foreign key to the todos table
 err = db.Debug().Model(&models.ToDo{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
 if err != nil {
  // Log the error if it occurs
  log.Fatalf("attaching foreign key error: %v", err)
 }

 // Create two users
 users := []models.User{
  models.User{
   Username: "stev",
   Email:    "stev@gmail.com",
   Password: "stev123",
  },
  models.User{
   Username: "martin",
   Email:    "martin@gmail.com",
   Password: "martin123",
  },
 }

 // Create two todos
 todos := []models.ToDo{
  models.ToDo{
   Title:   "Dance class",
   Content: "I have to find dance class near me",
  },
  models.ToDo{
   Title:   "Coding",
   Content: "I have to learn coding",
  },
 }

 // Loop through the users
 for i, _ := range users {
  // Create the user in the database
  err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
  if err != nil {
   // Log the error if it occurs
   log.Fatalf("cannot seed users table: %v", err)
  }
  // Set the author ID of the todo to the ID of the user
  todos[i].AuthorID = users[i].ID

  // Create the todo in the database
  err = db.Debug().Model(&models.ToDo{}).Create(&todos[i]).Error
  if err != nil {
   // Log the error if it occurs
   log.Fatalf("cannot seed todos table: %v", err)
  }
 }
}
