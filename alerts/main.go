package alerts

import (
	"strconv"
)

// CheckAlert func
func CheckAlert(operator string, valueStr string, compareValueStr string) (bool, error) {
	value, err := strconv.ParseFloat(valueStr, 32)
	if err != nil {
		return false, err
	}
	compareValue, err := strconv.ParseFloat(compareValueStr, 32)
	if err != nil {
		return false, err
	}
	switch operator {
	case ">":
		return value > compareValue, nil
	case ">=":
		return value >= compareValue, nil
	case "<":
		return value < compareValue, nil
	case "<=":
		return value <= compareValue, nil
	case "==":
		return value == compareValue, nil
	case "!=":
		return value != compareValue, nil
	default:
		return false, nil
	}
}

// CheckAlertOther for string
func CheckAlertOther(operator string, value string, compareValue string) (bool, error) {
	switch operator {
	case ">":
		return value > compareValue, nil
	case ">=":
		return value >= compareValue, nil
	case "<":
		return value < compareValue, nil
	case "<=":
		return value <= compareValue, nil
	case "==":
		return value == compareValue, nil
	case "!=":
		return value != compareValue, nil
	default:
		return false, nil
	}
}
