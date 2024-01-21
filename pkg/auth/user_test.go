package auth

import (
	"backend-ccff/internal/dbx"
	"backend-ccff/internal/models"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"testing"
)

// TODO crear usuario
func TestCreateUser(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceLogin := NewServerAuth(db, &usr, txID)

	id := uuid.New().String()

	_, cod, err := serviceLogin.Users.CreateUsers(id, "manager", "0020180656", "76797002",
		"manager", "manager", "lasname_mother", "root@test.com", "123456", 0, 0)
	if err != nil {
		t.Fatalf("no se pudo validar la creacion de usuario, cod: %d error: %v", cod, err)
	}
	t.Log("Create users was successful")
	fmt.Print("Create users was successful")
}

// TODO actualizar usuario
func TestUpdateUser(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceLogin := NewServerAuth(db, &usr, txID)

	id := strings.ToLower("5D815348-0DDD-42D9-8B36-FF9468C0DDC8")

	_, cod, err := serviceLogin.Users.UpdateUsers(id, "CHRISTIAN.BERNUY@UNAS.EDU.PE", "0020170414", "71434609",
		"CHRISTIAN RUBEN", "BERNUY", "OSORIO", "CHRISTIAN.BERNUY@UNAS.EDU.PE", "$2a$10$q7dSw8SBI.qind2c7gVH7OF4v0KWyk195iDD/p/Offr7LSQCgk5ky", 0, 0)
	if err != nil {
		t.Fatalf("no se pudo validar la actualizacion de usuario, cod: %d error: %v", cod, err)
	}
	t.Log("UpdateUsers users was successful")
	fmt.Print("UpdateUsers users was successful")
}

// TODO eliminar usuario por id
/*func TestDeleteUser(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceLogin := auth.NewServerAuth(db, &usr, txID)

	id := "5D815348-0DDD-42D9-8B36-FF9468C0DDC8"

	_, err := serviceLogin.Users.DeleteUsers(id)
	if err != nil {
		t.Fatalf("no se pudo eliminar usuario por id, error: %v", err)
	}
	t.Log("DeleteUsers users was successful")
	fmt.Print("DeleteUsers users was successful")
}*/

// TODO extraer usuario por id
func TestGetUsersByID(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceLogin := NewServerAuth(db, &usr, txID)

	id := strings.ToLower("5D815348-0DDD-42D9-8B36-FF9468C0DDC8")

	_, cod, err := serviceLogin.Users.GetUsersByID(id)
	if err != nil {
		t.Fatalf("no se pudo extraer usuario por id, cod: %d, error: %v", cod, err)
	}
	t.Log("GetUsersByID users was successful")
	fmt.Print("GetUsersByID users was successful")
}

// TODO extraer todos los usuarios
func TestGetAllUsers(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceLogin := NewServerAuth(db, &usr, txID)

	_, err := serviceLogin.Users.GetAllUsers()
	if err != nil {
		t.Fatalf("no se pudo extraer todo los usuarios, error: %v", err)
	}
	t.Log("GetAllUsers users was successful")
	fmt.Print("GetAllUsers users was successful")
}
