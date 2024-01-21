package config

import (
	"backend-ccff/internal/dbx"
	"backend-ccff/internal/models"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"testing"
	"time"
)

// TODO crear evento
func TestCreateEvent(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceConfig := NewServerConfig(db, &usr, txID)

	id := uuid.New().String()

	_, cod, err := serviceConfig.Event.CreateEvents(id, "event test", "event test", time.Now(), 0, 0, id)
	if err != nil {
		t.Fatalf("no se pudo validar la creacion de evento, cod: %d error: %v", cod, err)
	}
	t.Log("Create event was successful")
	fmt.Print("Create event was successful")
}

// TODO actualizar evento
func TestUpdateEvent(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceConfig := NewServerConfig(db, &usr, txID)

	id := strings.ToLower("18A20310-F7ED-41C5-9A87-0DFD5227DB4D")
	// id := uuid.New().String()
	_, cod, err := serviceConfig.Event.UpdateEvents(id, "event test", "event test", time.Now(), 0, 0, id)
	if err != nil {
		t.Fatalf("no se pudo validar la actualizacion de evento, cod: %d error: %v", cod, err)
	}
	t.Log("Update event was successful")
	fmt.Print("Update event was successful")
}

// TODO eliminar evento por id
/*func TestDeleteEvent(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceConfig := config.NewServerConfig(db, &usr, txID)

	id := strings.ToUpper("18A20310-F7ED-41C5-9A87-0DFD5227DB4DD")

	_, err := serviceConfig.Event.DeleteEvents(id)
	if err != nil {
		t.Fatalf("no se pudo eliminar el evento, error: %v", err)
	}
	t.Log("Delete event was successful")
	fmt.Print("Delete event was successful")
}*/

// TODO extraer evento por id
func TestGetEventsByID(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceConfig := NewServerConfig(db, &usr, txID)

	id := uuid.New().String()

	_, cod, err := serviceConfig.Event.GetEventsByID(id)
	if err != nil {
		t.Fatalf("no se pudo traer el evento por id, cod: %d, error: %v", cod, err)
	}
	t.Log("GetEventsByID event was successful")
	fmt.Print("GetEventsByID event was successful")
}

// TODO extraer todo los eventos
func TestGetAllEvents(t *testing.T) {
	db := dbx.GetConnection()
	txID := uuid.New().String()
	usr := models.User{}
	serviceConfig := NewServerConfig(db, &usr, txID)

	_, err := serviceConfig.Event.GetAllEvents()
	if err != nil {
		t.Fatalf("no se pudo traer todo el evento por id, error: %v", err)
	}
	t.Log("GetAllEvents event was successful")
	fmt.Print("GetAllEvents event was successful")
}
