package main

import (
	"errors"
	"fmt"
)

// custom type error
type ErrorValidation struct {
	Field string
	Msg   string
}

func (e ErrorValidation) Error() string {
	return fmt.Sprintf(" '%s': %s", e.Field, e.Msg)
}

func validasiUsia(umur int) error {
	if umur < 0 {
		return errors.New("Umur tidak boleh negatif")
	}
	if umur < 17 {
		return &ErrorValidation{"Umur", "Umur harus minimal 17 tahun"}
	}
	return nil
}

func main() {
	// error handling
	if err := validasiUsia(15); err != nil {
		fmt.Println(err)
	}
	if err := validasiUsia(-5); err != nil {
		fmt.Println(err)
	}

	// panic & recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi panic:", r)
		}
	}()

	fmt.Println("Sebelum panic")
	panic("Ini adalah panic!")

}
