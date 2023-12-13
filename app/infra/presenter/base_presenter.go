package presenter

import "github.com/gofiber/fiber/v2"

func Success(message string, data interface{}, meta interface{}) fiber.Map {
	if meta != nil {
		return fiber.Map{
			"data":    data,
			"meta":    meta,
			"message": message,
			"status":  1,
			"code":    200,
		}
	} else {
		return fiber.Map{
			"data":    data,
			"message": message,
			"status":  1,
			"code":    200,
		}
	}
}

func SuccessAuth(message string, token string, data interface{}) fiber.Map {
	result := fiber.Map{
		"data":    data,
		"token":   token,
		"message": message,
		"status":  1,
		"code":    200,
	}

	return result
}

func Unauthorize(message string, data interface{}) fiber.Map {
	result := fiber.Map{
		"data":    data,
		"message": message,
		"status":  0,
		"code":    401,
	}

	return result
}

func Error(message string, data interface{}, code int) fiber.Map {
	result := fiber.Map{
		"data":    data,
		"message": message,
		"status":  0,
		"code":    code,
	}

	return result
}
