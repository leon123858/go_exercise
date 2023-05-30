package utils_test

import (
	"log"
	"os"
	"testing"

	"github.com/leon123858/committee-meeting-assistan/data-api/model"
	_ "github.com/leon123858/committee-meeting-assistan/data-api/testing_init"

	utils "github.com/leon123858/committee-meeting-assistan/data-api/utils"
	"github.com/stretchr/testify/assert"
)

func setUpAll() func() {
	utils.GetConfig("debug")
	utils.InitDB()
	return func() {
		utils.DB.Exec("TRUNCATE TABLE album;")
		utils.DisconnectDB()
	}
}

func TestMain(m *testing.M) {
	tearDownAll := setUpAll()

	code := m.Run()

	tearDownAll() // you cannot use defer tearDownAll()
	os.Exit(code)
}

func setup() {
	log.Println("---")
}

func teardown() {
	log.Println("---")
}

func TestSomething(t *testing.T) {
	setup()
	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")
	utils.DB.Create(model.Album{ID: "556888", Title: "SSS", Artist: "AAA", Price: 55})
	teardown()

}

func TestSomething2(t *testing.T) {
	setup()
	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")
	teardown()
}
