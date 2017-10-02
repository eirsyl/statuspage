package src

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func IncidentStatus(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if status, ok := field.Interface().(string); ok {
		switch status {
		case
			"Investigating",
			"Identified",
			"Watching",
			"Fixed":
			return true
		}
		return false
	}
	return true
}

func ServiceStatus(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if status, ok := field.Interface().(string); ok {
		switch status {
		case
			"Operational",
			"Performance Issues",
			"Partial Outage",
			"Major Outage":
			return true
		}
		return false
	}
	return true
}
