package auth

import (
	"errors"
	"fmt"

	casbin "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type ACTION string
type AUTHZ_ROLE string

const (
	DEFAULT AUTHZ_ROLE = "default"
	VIEWER  AUTHZ_ROLE = "viewer"
	EDITOR  AUTHZ_ROLE = "editor"
	MANAGER AUTHZ_ROLE = "manager"
	READ    ACTION     = "read"
	WRITE   ACTION     = "write"
	MANAGE  ACTION     = "manage"
)

var enforcer *casbin.Enforcer

var action2number = map[ACTION]int{
	READ:   1,
	WRITE:  2,
	MANAGE: 3,
}

func customFunction(key1 string, key2 string) bool {
	number1 := action2number[ACTION(key1)]
	number2 := action2number[ACTION(key2)]
	return number1 <= number2
}

func customFunctionWrapper(args ...interface{}) (interface{}, error) {
	key1 := args[0].(string)
	key2 := args[1].(string)

	return bool(customFunction(key1, key2)), nil
}

func InitCasbin(gormInstance *gorm.DB) {
	adapter, _ := gormadapter.NewAdapterByDB(gormInstance)
	enforcer, _ = casbin.NewEnforcer("./rbac_model.conf", adapter)
	enforcer.AddFunction("keyMatchCustom", customFunctionWrapper)
	enforcer.LoadPolicy()
}

func CheckAuthz(who string, resource string, action ACTION) (bool, error) {
	result, err := enforcer.Enforce(who, resource, string(action))
	return result, err
}

func CreateCommunity(resourceName string) error {
	var policies [][]string = [][]string{
		{fmt.Sprintf("%s:%s", VIEWER, resourceName), resourceName, string(READ)},
		{fmt.Sprintf("%s:%s", EDITOR, resourceName), resourceName, string(WRITE)},
		{fmt.Sprintf("%s:%s", MANAGER, resourceName), resourceName, string(MANAGE)},
	}
	result, err := enforcer.AddPolicies(policies)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("create error: data conflict")
	}
	return nil
}

func AddRole(resourceName string, who string, level AUTHZ_ROLE) error {
	result, err := enforcer.AddRoleForUser(who,
		fmt.Sprintf("%s:%s", level, resourceName))
	if err != nil {
		return err
	}
	if !result {
		return errors.New("add error: data conflict")
	}
	return nil
}

func RemoveRole(resourceName string, who string, level AUTHZ_ROLE) error {
	result, err := enforcer.DeleteRoleForUser(who,
		fmt.Sprintf("%s:%s", level, resourceName))
	if err != nil {
		return err
	}
	if !result {
		return errors.New("remove error: data conflict")
	}
	return nil
}

func EditRole(resourceName string, who string, oldRole AUTHZ_ROLE, newRole AUTHZ_ROLE) error {
	err := enforcer.GetAdapter().(*gormadapter.Adapter).Transaction(enforcer, func(e casbin.IEnforcer) error {
		result, err := e.DeleteRoleForUser(who,
			fmt.Sprintf("%s:%s", oldRole, resourceName))
		if err != nil {
			return err
		}
		if !result {
			return errors.New("edit error 1: data conflict")
		}
		result, err = e.AddRoleForUser(who,
			fmt.Sprintf("%s:%s", newRole, resourceName))
		if err != nil {
			return err
		}
		if !result {
			return errors.New("edit error 2: data conflict")
		}
		return nil
	})
	return err
}
