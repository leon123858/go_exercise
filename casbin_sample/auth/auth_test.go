package auth_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"testing"

	utils "github.com/leon123858/committee-meeting-assistan/data-api/utils"
	"github.com/leon123858/committee-meeting-assistan/data-api/utils/auth"
	"github.com/stretchr/testify/assert"
)

// 印出 struct
//
// res, err := PrettyStruct(agents)
//
//	if err != nil {
//			log.Fatal(err)
//	}
//
// fmt.Println(res)
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func truncateTables(tables []string) {
	for _, element := range tables {
		if element == "pgmigrations" {
			continue
		}
		utils.DB.Exec(fmt.Sprintf("TRUNCATE TABLE \"%s\" RESTART IDENTITY CASCADE;", element))
	}
}

func emptyTables() {
	var tables []string
	if err := utils.DB.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
		panic(err)
	}
	truncateTables(tables)
}

func setUpAll() func() {
	// set test path
	_, filename, _, _ := runtime.Caller(0)
	println(filename)
	dir := path.Join(path.Dir(filename), "../../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	// start test
	utils.GetConfig("debug")
	utils.InitDB()
	// SELECT * FROM information_schema.tables;
	auth.InitCasbin(utils.DB)
	return func() {
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
	emptyTables()
	log.Println("---")
}

func teardown() {
	log.Println("---")
}

func TestAuthz(t *testing.T) {
	setup()
	// create workspace
	err := auth.CreateCommunity("testWorkSpaceId1")
	assert.Equal(t, nil, err)
	err = auth.CreateCommunity("testWorkSpaceId2")
	assert.Equal(t, nil, err)
	// set role
	// can do role can do
	err = auth.AddRole("testWorkSpaceId1", "baby1", auth.EDITOR)
	assert.Equal(t, nil, err)
	result, err := auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.READ)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, result)
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.WRITE)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, result)
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.MANAGE)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, result)
	// can not access other workspace
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId2", auth.READ)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, result)
	// can edit role for user
	// reject when origin is wrong
	err = auth.EditRole("testWorkSpaceId1", "baby1", auth.MANAGER, auth.VIEWER)
	assert.NotEqual(t, nil, err)
	// reject should still same
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.READ)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, result)
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.WRITE)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, result)
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.MANAGE)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, result)
	// can edit success
	err = auth.EditRole("testWorkSpaceId1", "baby1", auth.EDITOR, auth.VIEWER)
	assert.Equal(t, nil, err)
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.READ)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, result)
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.WRITE)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, result)
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.MANAGE)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, result)
	// remove role
	// should wrong when not the role
	err = auth.RemoveRole("testWorkSpaceId1", "baby1", auth.EDITOR)
	assert.NotEqual(t, nil, err)
	err = auth.RemoveRole("testWorkSpaceId1", "baby1", auth.VIEWER)
	assert.Equal(t, nil, err)
	result, err = auth.CheckAuthz("baby1", "testWorkSpaceId1", auth.READ)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, result)
	teardown()
}
