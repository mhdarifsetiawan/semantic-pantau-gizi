package typeconv

import (
	"fmt"
	"math"
	"strconv"
)

func ConvertToUint32(value interface{}) (uint32, error) {
	switch v := value.(type) {
	case uint32:
		return v, nil
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint64:
		if v > math.MaxUint32 {
			return 0, fmt.Errorf("value %d exceeds uint32 max value", v)
		}
		return uint32(v), nil
	case uint:
		if v > math.MaxUint32 {
			return 0, fmt.Errorf("value %d exceeds uint32 max value", v)
		}
		return uint32(v), nil
	case int8:
		if v < 0 {
			return 0, fmt.Errorf("negative value %d cannot be converted to uint32", v)
		}
		return uint32(v), nil
	case int16:
		if v < 0 {
			return 0, fmt.Errorf("negative value %d cannot be converted to uint32", v)
		}
		return uint32(v), nil
	case int32:
		if v < 0 {
			return 0, fmt.Errorf("negative value %d cannot be converted to uint32", v)
		}
		return uint32(v), nil
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("negative value %d cannot be converted to uint32", v)
		}
		if v > math.MaxUint32 {
			return 0, fmt.Errorf("value %d exceeds uint32 max value", v)
		}
		return uint32(v), nil
	case int:
		if v < 0 {
			return 0, fmt.Errorf("negative value %d cannot be converted to uint32", v)
		}
		if v > math.MaxUint32 {
			return 0, fmt.Errorf("value %d exceeds uint32 max value", v)
		}
		return uint32(v), nil
	case float32:
		if v < 0 {
			return 0, fmt.Errorf("negative value %f cannot be converted to uint32", v)
		}
		if v > math.MaxUint32 {
			return 0, fmt.Errorf("value %f exceeds uint32 max value", v)
		}
		return uint32(v), nil
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("negative value %f cannot be converted to uint32", v)
		}
		if v > math.MaxUint32 {
			return 0, fmt.Errorf("value %f exceeds uint32 max value", v)
		}
		return uint32(v), nil
	case string:
		if v == "" {
			return 0, fmt.Errorf("cannot convert empty string to uint32")
		}
		u, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0, fmt.Errorf("cannot convert string '%s' to uint32: %v", v, err)
		}
		return uint32(u), nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", value)
	}
}

func ConvertToFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot convert string '%s' to float64: %v", v, err)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", value)
	}
}
