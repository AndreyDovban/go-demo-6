package request

import "github.com/go-playground/validator"

func IsValid[T any](payload T) error {
	validator := validator.New()
	err := validator.Struct(payload)
	return err
}

// if payload.Email == "" {
// 	resp.Json(w, "Email required", 402)
// 	return
// }
// reg, _ := regexp.Compile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`)
// if !reg.MatchString(payload.Email) {
// 	resp.Json(w, "Wrong email", 402)
// 	return
// }
// if payload.Password == "" {
// 	resp.Json(w, "Password required", 402)
// 	return
// }
