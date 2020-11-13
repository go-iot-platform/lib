package payment

import (
	"github.com/go-iot-platform/lib/rbac/model"
)

type MethodType uint
type HttpType int
type JobAgent int

const (
	// Payment gateway method config
	GetConfigMethod      = "getconfig"
	TopupMethod          = "topup"
	RenewMethod          = "renew"
	ChargeMethod         = "charge"
	CreateCustomerMethod = "createcustomer"
	GetCustomerInfo      = "getcustomerinfo"
	// Charge type config
	OnskyWallet = "onsky"
)

// payService type config
const (
	Renew = iota + 1
	Charge
	Topup
	JobRunner
)

// Payment gateway call method config
const (
	GRPC = iota + 1
	API
)

// HTTP method config
const (
	GET = iota + 1
	POST
	PUT
	PATCH
	DELETE
)

func (s MethodType) String() string {
	switch s {
	case API:
		return "API"
	case GRPC:
		return "GRPC"
	default:
		return "GRPC"
	}
}

func (s JobAgent) String() string {
	switch s {
	case Renew:
		return "Renew"
	case Charge:
		return "Charge"
	case Topup:
		return "Topup"
	case JobRunner:
		return "Job"
	default:
		return ""
	}
}

func (s HttpType) String() string {
	switch s {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case PATCH:
		return "PATCH"
	case DELETE:
		return "DELETE"
	default:
		return "POST"
	}
}

// PaymentGateway represents package table
type PaymentGateway struct {
	model.Base
	UUID        string                `pg:"uuid,type:uuid,unique" json:"uuid" validate:"omitempty,uuid4"`
	Name        string                `pg:"name,type:varchar(50),unique" json:"name" validate:"required,max=50,alphanum"`
	Description string                `pg:"description,type:varchar(255)" json:"description" validate:"omitempty,max=255"`
	MethodName  string                `pg:"method_name,type:varchar(255)" json:"methodName" validate:"max=255"`
	MethodType  MethodType            `pg:"method_type" json:"methodType"`
	MethodURL   string                `pg:"method_url,type:varchar(500)" json:"methodURL"`
	Active      bool                  `pg:",use_zero" json:"active" validate:"oneof=true false"`
	Param       []PaymentGatewayParam `pg:"param" json:"param"`
	HTTPType    HttpType              `pg:"http_type" json:"http_type"`
	//Config      map[string]interface{} `json:"config" validate:"omitempty"`
}

//GetConfigRequest represent the get config request.
type GetConfigRequest struct {
	Name string `json:"name" validate:"max=50"`
}

//TopupRequest represent topup request.
type TopupRequest struct {
	Amount             uint64 `json:"amount" validate:"required"`
	PaymentGatewayName string `json:"payment_gateway_name" validate:"required,max=50,alphanum"`
	PaymentMethodNonce string `json:"payment_method_nonce,required"`
}

//PaymentGatewayParam represent param jsonb field in model.
type PaymentGatewayParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

//ChargeRequest represent a charge request
type ChargeRequest struct {
	PackageUUID        string `json:"package_uuid" validate:"uuid4"`
	PaymentGateway     string `json:"payment_gateway" validate:"max=25,alpha"`
	ServiceName        string `json:"service_name" validate:"max=50"`
	Time               uint32 `json:"time" validate:"numeric,min=0"`
	PaymentMethodNonce string `json:"payment_method_nonce,omitempty"`
	PromoCode          string `json:"promo_code,omitempty"`
	IsAutoCharge       bool   `json:"is_auto_charge"`
}

//CountFeeRequest represent a countFee request
type CountFeeRequest struct {
	PackageUUID string `json:"package_uuid" validate:"uuid4"`
	ServiceName string `json:"service_name,omitempty" validate:"max=50"`
	PromoCode   string `json:"promo_code,omitempty"`
}

//CountFeeResponse represent a countFee response
type CountFeeResponse struct {
	PackagePrice   string `json:"package_price"`
	PromotionValue string `json:"promotion_value"`
	FinalPrice     string `json:"final_price"`
}

//JobRunnerReq represent job runner request
type JobRunnerReq struct {
	CustomerNumber string `json:"customer_number"`
	PackageUUID    string `json:"package_uuid" validate:"uuid4"`
	PaymentGateway string `json:"payment_gateway"`
	UserUUID       string `json:"user_uuid" validate:"uuid4"`
}
