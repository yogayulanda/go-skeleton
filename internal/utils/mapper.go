package utils

import (
	"fmt"
	"reflect"
)

// ProtoToModel mengubah proto object ke model domain secara generic
func ProtoToModel(proto interface{}, model interface{}) error {
	protoValue := reflect.ValueOf(proto)
	modelValue := reflect.ValueOf(model).Elem()

	// Validasi input
	if protoValue.Kind() != reflect.Ptr || modelValue.Kind() != reflect.Ptr {
		return fmt.Errorf("both proto and model must be pointers")
	}

	// Looping field pada proto dan model
	for i := 0; i < protoValue.Elem().NumField(); i++ {
		protoField := protoValue.Elem().Field(i)
		modelField := modelValue.FieldByName(protoValue.Elem().Type().Field(i).Name)

		if modelField.IsValid() && modelField.CanSet() {
			// Set model field dengan nilai dari proto field
			if modelField.Kind() == protoField.Kind() {
				modelField.Set(protoField)
			}
		}
	}
	return nil
}

// ModelToProto mengubah model domain ke proto object secara generic
func ModelToProto(model interface{}, proto interface{}) error {
	modelValue := reflect.ValueOf(model)
	protoValue := reflect.ValueOf(proto).Elem()

	// Validasi input
	if modelValue.Kind() != reflect.Ptr || protoValue.Kind() != reflect.Ptr {
		return fmt.Errorf("both model and proto must be pointers")
	}

	// Looping field pada model dan proto
	for i := 0; i < modelValue.Elem().NumField(); i++ {
		modelField := modelValue.Elem().Field(i)
		protoField := protoValue.FieldByName(modelValue.Elem().Type().Field(i).Name)

		if protoField.IsValid() && protoField.CanSet() {
			// Set proto field dengan nilai dari model field
			if protoField.Kind() == modelField.Kind() {
				protoField.Set(modelField)
			}
		}
	}
	return nil
}
