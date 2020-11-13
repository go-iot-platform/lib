package authen

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
	"github.com/micro/go-micro/metadata"
	"github.com/mitchellh/mapstructure"
	"github.com/go-iot-platform/errors"
	"github.com/go-iot-platform/go-redis"
	"github.com/go-iot-platform/lib/rbac/model"
	"github.com/go-iot-platform/rbac/utl"
	resourcehelper "github.com/go-iot-platform/resource-helper"
)

// ServiceName Seperate Char
const ServiceName = "rbac"

// ResourceAll resource key
const ResourceAll = "*"

// ResourceList resource key
const ResourceList = "List"

// CallbackType type of callback func
type CallbackType uint8

const (
	//DEFAULT callbacktype is default
	DEFAULT = iota
	//CUSTOMER callbacktype is customer
	CUSTOMER
	//POLICY callbacktype is policy
	POLICY
)

// Callback func
type Callback func(CallbackType, map[string]string) (interface{}, error)

// RBAC model representation
type RBAC struct {
	DB               *redis.Redis
	UserCacheKey     string
	ActionCacheKey   string
	CustomerCacheKey string
	Key              string
	Algo             string
	Callback         Callback
	Cfg              map[string]string
}

func emptyCallback(callbackType CallbackType, cfg map[string]string) (interface{}, error) {
	return nil, errors.BadRequest(ServiceName, "rbac:emptyCallback:%+v\n", cfg)
}

// Init init redis receiver
func Init(db *redis.Redis, userCacheKey string, actionCacheKey string, customerCacheKey string, key string, algo string, callback Callback) *RBAC {
	if callback == nil {
		callback = emptyCallback
	}
	return &RBAC{
		DB:               db,
		UserCacheKey:     userCacheKey,
		ActionCacheKey:   actionCacheKey,
		CustomerCacheKey: customerCacheKey,
		Key:              key,
		Algo:             algo,
		Callback:         callback,
	}
}

// UserFromContext instance
func (r *RBAC) UserFromContext(c echo.Context) *model.AuthUser {
	customer := new(model.Customer)

	id := c.Get("id").(int)
	customerID := c.Get("customer_id").(int)
	customerNumber := c.Get("customer_number").(string)
	user := c.Get("username").(string)
	role := c.Get("role").(uint8)
	userUUID := c.Get("user_uuid").(string)

	//check cache
	result := model.AuthUser{
		ID:             id,
		Username:       user,
		CustomerNumber: customerNumber,
		CustomerID:     customerID,
		UserUuid:       userUUID,
		Role:           model.AccessRole(role),
	}

	err := enforceRole(role, model.CustomerUserRole)
	if err != nil {
		return &result
	}

	err = r.GetDataFromCache(DEFAULT, r.CustomerCacheKey, customerNumber, customer)
	if err != nil || customer == nil || customer.Clients == nil {
		return &result
	}

	for i := 0; i < len(customer.Clients); i++ {
		result.Clients = append(result.Clients, customer.Clients[i].AccNumber)
	}

	return &result
}

// UserFromMetadata instance
func (r *RBAC) UserFromMetadata(ctx context.Context) (*model.AuthUser, error) {
	customer := new(model.Customer)

	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, errors.Forbidden(ServiceName, "rbac:authen:UserFromMetadata:invalidParams:meta:%+v\n", meta)
	}

	token, ok := meta["authorization"]
	// support UPERCASE for rpc
	if !ok {
		token, ok = meta["Authorization"]
	}
	if !ok {
		return nil, errors.Forbidden(ServiceName, "rbac:authen:UserFromMetadata:missingToken:meta:%+v\n", meta)
	}
	result, err := utl.ParseToken(r.Key, r.Algo, token)
	if err != nil {
		return nil, err
	}

	err = enforceRole(uint8(result.Role), model.CustomerUserRole)
	if err != nil {
		return result, nil
	}

	err = r.GetDataFromCache(DEFAULT, r.CustomerCacheKey, result.CustomerNumber, customer)
	if err != nil || customer == nil || customer.Clients == nil {
		return result, nil
	}

	for i := 0; i < len(customer.Clients); i++ {
		result.Clients = append(result.Clients, customer.Clients[i].AccNumber)
	}

	return result, nil
}

// EnforcePolicy only have permission when meet requirement policy and resouce
// params:
// - serviceName string
// - action string
// - resourceType string
// - resourceValue string
// - accountnumber string
// - patrition string
// - region string
// EnforcePolicy enfore policy
func (r *RBAC) EnforcePolicy(role uint8, customerNumber string, userUUID string, action string, compareResources ...string) ([]string, error) {
	user := &model.User{}
	actionCache := &resourcehelper.ItemActionCache{}

	if action == "" || r == nil || role == 0 || customerNumber == "" || userUUID == "" {
		return nil, errors.Forbidden(ServiceName, "enforcePolicy:invalidParams:action:%+v:role:%+v:cn:%+v:useruuid:%+v\n", action, role, customerNumber, userUUID)
	}

	err := r.GetDataFromCache(DEFAULT, r.ActionCacheKey, action, actionCache)
	if err != nil {
		return nil, err
	}
	if actionCache.Name == "" {
		return nil, errors.Forbidden(ServiceName, "enforcePolicy:invalidAction")
	}

	results, err := checkAdmin(role, compareResources, actionCache, customerNumber)
	if err != nil {
		return nil, err
	}
	//---if role is superadmin,admin , bypass check policy (fullacess)
	if results != nil {
		return results, nil
	}

	//check resource, if empty set value "*"
	if compareResources == nil || len(compareResources) == 0 {
		compareResources = []string{ResourceAll}
	}
	//-------------get policies of user-------------
	err = r.GetDataFromCache(POLICY, r.UserCacheKey, userUUID, user)
	if err != nil {
		return nil, err
	}
	if user.Policies == nil || len(user.Policies) == 0 {
		// try to get policies again
		err = r.GetDataFromCacheWorkaround(POLICY, r.UserCacheKey, userUUID, user)
		if user.Policies == nil || len(user.Policies) == 0 {
			return nil, errors.Forbidden(ServiceName, "enforcePolicy:policy:empty:action:%+v - actionCache:%+v - customerNumber:%+v - user:%+v\n", action, actionCache, customerNumber, user)
		}
	}

	assignedResources := getAssignedResources(action, actionCache, customerNumber, user.Policies)
	if len(assignedResources) == 0 {
		return nil, errors.Forbidden(ServiceName, "enforcePolicy:assignedResources:empty:action:%+v - actionCache:%+v - customerNumber:%+v - policies:%+v\n", action, actionCache, customerNumber, user.Policies)
	}

	err = checkInvalidResource(compareResources, actionCache, customerNumber)
	if err != nil {
		return nil, err
	}

	result := getResources(compareResources, assignedResources, actionCache, customerNumber)
	if result == nil || len(result) == 0 {
		return nil, errors.Forbidden(ServiceName, "enforcePolicy:restricted:compareResources:%+v - assignedResources:%+v - actionCache:%+v - cn:%+v\n", compareResources, assignedResources, actionCache, customerNumber)
	}
	return result, nil
}

// EnforceRole authorizes request by AccessRole
func enforceRole(role uint8, roleBase model.AccessRole) error {
	if role <= uint8(roleBase) {
		return nil
	}
	return errors.Forbidden(ServiceName, "enforcePolicy:forbidden:role:%+v - base:%+v\n", role, roleBase)
}
func getAssignedResources(action string, actionCache *resourcehelper.ItemActionCache, customerNumber string, policies []model.Policy) []string {
	result := []string{}
	resources := []string{}
	if policies == nil || len(policies) == 0 {
		return result
	}
	for i := 0; i < len(policies); i++ {
		for j := 0; j < len(policies[i].ResourceTypes); j++ {
			resourceType := policies[i].ResourceTypes[j]
			isMatch := resourcehelper.Contains(resourceType.Actions, action) || resourcehelper.Contains(resourceType.Actions, "*")
			if !isMatch {
				continue
			}
			//--refine to make sure that any item does not contain each other.
			for k := 0; k < len(resourceType.Resources); k++ {
				temp := ""
				for m := 0; m < len(result); m++ {
					ok, _, _ := resourcehelper.ParseResource(resourceType.Resources[k], result[m], actionCache, customerNumber)
					if !ok {
						continue
					}
					temp = result[m]
				}
				if temp != "" {
					continue
				}
				resources = append(resources, resourceType.Resources[k])
			}
			if len(resources) == 0 {
				continue
			}

			//filter items not match new resource item
			//make sure that any item does not contain each other.
			for k := 0; k < len(resources); k++ {
				assignedResources := []string{}
				for m := 0; m < len(result); m++ {
					ok, _, _ := resourcehelper.ParseResource(result[m], resources[k], actionCache, customerNumber)
					if ok {
						continue
					}
					assignedResources = append(assignedResources, result[m])
				}
				result = append(result, resources[k])
			}
		}
	}
	return result
}
func checkInvalidResource(compareResources []string, actionCache *resourcehelper.ItemActionCache, customerNumber string) error {
	for i := 0; i < len(compareResources); i++ {
		result := resourcehelper.GetResourceFormat(compareResources[i], actionCache, customerNumber)
		if result == nil {
			return errors.Forbidden(ServiceName, "enforcePolicy:invalidResource:compareResources:%+v - actionCache:%+v - cn:%+v\n", compareResources[i], actionCache, customerNumber)
		}
	}
	return nil
}
func getResources(compareResources []string, assignedResources []string, actionCache *resourcehelper.ItemActionCache, customerNumber string) []string {
	result := []string{}

	for i := 0; i < len(compareResources); i++ {
		for j := 0; j < len(assignedResources); j++ {
			ok, dataSourceObject, _ := resourcehelper.ParseResource(compareResources[i], assignedResources[j], actionCache, customerNumber)
			if ok {
				result = append(result, dataSourceObject.ToString())
			}
			if dataSourceObject.Resource == ResourceAll && actionCache.Type == ResourceList {
				result = append(result, assignedResources...)
			}
		}
	}
	return result
}

// GetDataFromCache get data from cache
func (r *RBAC) GetDataFromCache(mode CallbackType, key string, field string, result interface{}) error {
	var temp interface{}
	err := r.DB.GetObject(key, field, &temp)
	if err == nil {
		b, err := json.Marshal(temp)
		if err != nil {
			return errors.InternalServerError(ServiceName, "enforcePolicy:marshalProblem:%+v\n", temp)
		}
		json.Unmarshal(b, result)

		if result == nil {
			return errors.InternalServerError(ServiceName, "enforcePolicy:wrongData")
		}
		return nil
	}
	//  update cache by callback
	switch mode {
	case POLICY:
		result := model.User{}
		input := map[string]string{
			"user_uuid": field,
		}
		fmt.Printf("callback func: %+v\n", input)
		res, err := r.Callback(POLICY, input)
		if err == nil && res != nil {
			if err = mapstructure.Decode(&res, &result.Policies); err == nil {
				return nil
			}
		}
		if err != nil {
			fmt.Printf("Callback error:%+v\n", err)
		}
		// try to get from cache
		err = r.DB.GetObject(key, field, &temp)
		if err != nil && err.Error() == errors.RedisEmpty {
			return errors.NotFound(ServiceName, "enforcePolicy:policy:empty:%s\n", field)
		}
		break
	default:
		break
	}
	if err != nil {
		return err
	}
	b, err := json.Marshal(temp)
	if err != nil {
		return errors.InternalServerError(ServiceName, "enforcePolicy:marshalProblem:%+v\n", temp)
	}
	json.Unmarshal(b, result)

	if result == nil {
		return errors.InternalServerError(ServiceName, "enforcePolicy:wrongData")
	}
	return nil
}

// GetDataFromCacheWorkaround workaround
func (r *RBAC) GetDataFromCacheWorkaround(mode CallbackType, key, field string, result interface{}) error {
	fmt.Printf("trying to get %v for key %s and field %s", mode, key, field)
	var temp interface{}
	switch mode {
	case POLICY:
		input := map[string]string{
			"user_uuid": field,
		}
		fmt.Printf("callback func workaround: %+v\n", input)
		_, err := r.Callback(POLICY, input)
		if err != nil {
			return err
		}
		break
	default:
		return nil
	}
	// try to get from cache
	err := r.DB.GetObject(key, field, &temp)
	if err != nil {
		if err.Error() == errors.RedisEmpty {
			return errors.NotFound(ServiceName, "enforcePolicy:policy:GetDataFromCacheWorkaround:empty:%s\n", field)
		}
		return err
	}
	b, err := json.Marshal(temp)
	if err != nil {
		return errors.InternalServerError(ServiceName, "enforcePolicy:GetDataFromCacheWorkaround:marshalProblem:%+v\n", temp)
	}
	json.Unmarshal(b, result)
	if result == nil {
		return errors.InternalServerError(ServiceName, "enforcePolicy:GetDataFromCacheWorkaround:wrongData")
	}
	return nil
}
func checkAdmin(role uint8, compareResources []string, actionCache *resourcehelper.ItemActionCache, customerNumber string) ([]string, error) {
	var results []string
	err := enforceRole(role, model.AdminRole)
	if err != nil {
		return nil, nil
	}
	for i := 0; i < len(compareResources); i++ {
		resource := resourcehelper.GetResourceFormat(compareResources[i], actionCache, customerNumber)
		if resource == nil {
			return nil, errors.InternalServerError(ServiceName, "enforcePolicy:invalidCustomer:%+v - %+v - %+v\n", compareResources[i], actionCache, customerNumber)
		}
		results = append(results, resource.ToString())
	}

	return results, nil
}
