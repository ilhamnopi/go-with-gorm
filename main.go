package main

import (
	"errors"
	"fmt"
	"main/database"
	"main/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()
	// createUser("ilhamnopihendri@gmail.com")
	// getUserById(1)
	// updateUserById(1, "ilhamnopihendri@gmail.com")
	// createProduct(1, "YLO", "YYYY")
	getUserWithProducts()
}

func createUser(email string){
	db := database.GetDB()
	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil{
		fmt.Println("Error creating user data:", err)
		return
	}
	fmt.Println("New User Data:", User)
}

func getUserById(id uint)  {
	db := database.GetDB()
	user := models.User{}
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}
	fmt.Printf("User Data: %+v \n", user)
}

func updateUserById(id uint, email string)  {
	db := database.GetDB()
	user:= models.User{}
	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	// err := db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"email": "ilhamnopi@gmail.com"}).Error
	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update user's email %v+ \n", user.Email)
}


func createProduct(userId uint, brand string, name string){
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand: brand,
		Name: name,
	}
	err:= db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}
	fmt.Println("New Product Data", Product)
}

func getUserWithProducts(){
	db:= database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil{
		fmt.Println("Error getting user data with products:", err.Error())
		return
	}
	fmt.Println("User Datas with Products")
	fmt.Printf("%v", users)
}