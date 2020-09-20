package model

// SecurityType request or resposne
type SecurityType int8

// NotificationStatus type constant
type NotificationStatus int8

// Mode type constant
type Mode int8

// NotificationType request or resposne
type NotificationType int8

// CallStatus call event
type CallStatus string

// Type of call or message
type Type int8

// DataType of a property of thing
type DataType int8

// StorageType of a property of thing
type StorageType int8

const (
	Picture = iota + 1
	Video
	Text
	Others
)

var StorageTypes = []StorageType{
	Picture,
	Video,
	Text,
	Others,
}

const (
	Number = iota + 1
	String
	Json
	Image
	Boolean
	ThingObject
	XML
	DateTime
	Location
	Schedule
	HTML
)

// Message Type constant
const (
	Motion = iota
	Smoke
	Co
	SOS
	TempHumd
	DoorLock
	OSLocus
	BedSensor
)

// Mode Type constant
const (
	Security = iota
	Safety
	Basic
)

// Status type constant
const (
	Initial = iota
	Running
	Done
	Timeout
	Canceled
	Notified
)

// TemplateType type constant
type TemplateType int8

// MessageKey type constant
type MessageKey string

// Message key
const (
	MotionDetect                = "motion_detect"
	DoorOpen                    = "door_open"
	DoorClose                   = "door_close"
	SecurityBreachMessage       = "security_breach"
	SafetyBreachMessageCO       = "safety_breach_co"
	SafetyBreachMessageSOS      = "safety_breach_sos"
	SafetyBreachMessageSmoke    = "safety_breach_smoke"
	SecurityAlarmOff            = "security_alarm_off"
	SecurityAlarmHome           = "security_alarm_home"
	SecurityAlarmAway           = "security_alarm_away"
	SafetyAlarmEnable           = "safety_alarm_enable"
	SafetyAlarmDisable          = "safety_alarm_disable"
	SafetyBreachMessageTempHumd = "safety_breach_temp_humd"
	LowBatteryMessage           = "battery_low"
	FullBatteryMessage          = "battery_full"
	VibrationMessage            = "vibration"
)

// Template type constant
const (
	Gateway = iota
	LightRGB
	LightTemp
	LightDim
	MotionTemplate
	Door
	CO
	SmokeTemplate
	Siren
	Switch1C
	Switch3C
	FourPWM
	LAMP130
	LAMP110
	PlugS
	PlugA
	Turnable
	LightRGBOld
	LightDimOld
	TemperatureHumidity
	ZigbeeDoorLock
	SOSButton
)

func (s TemplateType) String() string {
	switch s {
	case LightRGB:
		return "LightRGB"
	case LightTemp:
		return "LightTemp"
	case LightDim:
		return "LightDim"
	case Motion:
		return "Motion"
	case Door:
		return "Door"
	case CO:
		return "CO"
	case SmokeTemplate:
		return "Smoke"
	case Siren:
		return "Siren"
	case Switch1C:
		return "Switch1C"
	case Switch3C:
		return "Switch3C"
	case FourPWM:
		return "FourPWM"
	case LAMP130:
		return "LAMP130"
	case LAMP110:
		return "LAMP110"
	case PlugS:
		return "PlugS"
	case PlugA:
		return "PlugA"
	case Turnable:
		return "Turnable"
	case LightRGBOld:
		return "LightRGBOld"
	case LightDimOld:
		return "LightDimOld"
	case TemperatureHumidity:
		return "TemperatureHumidity"
	case ZigbeeDoorLock:
		return "ZigbeeDoorLock"
	case SOSButton:
		return "SOSButton"
	// case Gateway:
	// 	return "Gateway"
	default:
		return "Gateway"
	}
}
func (s StorageType) String() string {
	switch s {
	case Picture:
		return "Image"
	case Video:
		return "Video"
	case Text:
		return "Text"
	default:
		return "Others"
	}
}
func (s DataType) String() string {
	switch s {
	case Number:
		return "Number"
	case String:
		return "String"
	case Json:
		return "Json"
	case Image:
		return "Image"
	case Boolean:
		return "Boolean"
	case ThingObject:
		return "ThingObject"
	case XML:
		return "XML"
	case DateTime:
		return "DateTime"
	case Location:
		return "Location"
	case Schedule:
		return "Schedule"
	case HTML:
		return "HTML"
	default:
		return "Unknown"
	}
}

// GetDataTypeList get data
func GetDataTypeList() []Configure {
	var result = make([]Configure, 11)
	for i := 1; i < 12; i++ {
		id := i
		dataType := DataType(i)
		result[i-1] = Configure{
			&id,
			dataType.String(),
		}
	}

	return result
}

func (s SecurityType) String() string {
	switch s {
	case Co:
		return "Co"
	case Smoke:
		return "Smoke"
	case SOS:
		return "SOS"
	case Motion:
		return "Motion"
	case TempHumd:
		return "TempHumd"
	case DoorLock:
		return "DoorLock"
	case OSLocus:
		return "OS Locus"
	case BedSensor:
		return "Bed Sensor"
	default:
		return "Motion"
	}
}
func (s Mode) String() string {
	switch s {
	case Safety:
		return "Safety"
	case Security:
		return "Security"
	case Basic:
		return "Basic"
	default:
		return "Basic"
	}
}

func (s NotificationStatus) String() string {
	switch s {
	case Running:
		return "Running"
	case Done:
		return "Done"
	case Timeout:
		return "Timeout"
	case Canceled:
		return "Canceled"
	case Notified:
		return "Notified"
	case Initial:
		return "Initial"
	default:
		return "Initial"
	}
}

// Message Type constant
const (
	Simple = iota
	DoorSensor
	MotionSensor
	SecurityBreach
	SafetyBreachCO
	SafetyBreachSOS
	SafetyBreachSmoke
	SecurityAlarm
	SafetyAlarm
	SafetyBreachTempHumd
	OSLocusTemp
	OSLocusSOS
	LowBattery
	Vibration
	BedSensorSOS
	BedSensorHeartStop
	BedSensorBreathStop
	BedSensorTachycardia
	BedSensorBradycardia
	BedSensorEpilepsy
	Complex
)

func (s NotificationType) String() string {
	switch s {
	case Complex:
		return "Complex"
	case DoorSensor:
		return "DoorSensor"
	case MotionSensor:
		return "MotionSensor"
	case SecurityBreach:
		return "SecurityBreach"
	case SafetyBreachCO:
		return "SafetyBreachCO"
	case SafetyBreachSOS:
		return "SafetyBreachSOS"
	case SafetyBreachSmoke:
		return "SafetyBreachSmoke"
	case SafetyBreachTempHumd:
		return "SafetyBreachTempHumd"
	case SecurityAlarm:
		return "SecurityAlarm"
	case SafetyAlarm:
		return "SafetyAlarm"
	case Simple:
		return "Simple"
	case OSLocusTemp:
		return "OSLocusTemp"
	case OSLocusSOS:
		return "OSLocusSOS"
	case LowBattery:
		return "LowBattery"
	case Vibration:
		return "Vibration"
	case BedSensorSOS:
		return "BedSensorSOS"
	case BedSensorHeartStop:
		return "BedSensorHeartStop"
	case BedSensorBreathStop:
		return "BedSensorBreathStop"
	case BedSensorTachycardia:
		return "BedSensorTachycardia"
	case BedSensorBradycardia:
		return "BedSensorBradycardia"
	case BedSensorEpilepsy:
		return "BedSensorEpilepsy"
	default:
		return "Simple"
	}
}

// Status type constant
const (
	Unknown = "unknown"
	Queued  = "queued"
	Failed  = "failed"
	// = "//" call status
	Initiated = "initiated"
	Ringing   = "ringing"
	Busy      = "busy"
	Answered  = "answered"
	Completed = "completed"
	// = "//" sms status
	Sent        = "sent"
	Delivered   = "delivered"
	Undelivered = "undelivered"
)

func (s CallStatus) String() string {
	switch s {
	case Queued:
		return "Queued"
	case Initiated:
		return "Initiated"
	case Ringing:
		return "Ringing"
	case Busy:
		return "Busy"
	case Answered:
		return "Answered"
	case Completed:
		return "Completed"
	case Sent:
		return "Sent"
	case Delivered:
		return "Delivered"
	case Undelivered:
		return "Undelivered"
	case Failed:
		return "Failed"
	case Unknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

// Call
const (
	Call = iota
	SMS
)

func (s Type) String() string {
	switch s {
	case SMS:
		return "SMS"
	case Call:
		return "Call"
	default:
		return "Call"
	}
}
