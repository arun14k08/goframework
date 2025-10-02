package framework

import (
	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(APIResponse{
		Code:    "success",
		Message: "Request successful",
		Data:    data,
	})
}

func SuccessWithMsg(ctx *fiber.Ctx, msg string,data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(APIResponse{
		Code:    "success",
		Message: msg,
		Data:    data,
	})
}

func Error(ctx *fiber.Ctx, status int, code, message string) error {
	return ctx.Status(status).JSON(APIResponse{
		Code:    code,
		Message: message,
	})
}

// --- Presets ---

func Created(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusCreated).JSON(APIResponse{
		Code: "resource_created",
		Message: "Request successful",
		Data: data,
	})
}

func InternalError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(APIResponse{
		Code:    "internal_error",
		Message: "Something went wrong, please try again later",
	})
}

func BadRequest(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Code:    "bad_request",
		Message: msg,
	})
}

func NotFound(ctx *fiber.Ctx, resource string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(APIResponse{
		Code:    "not_found",
		Message: resource + " not found",
	})
}

func UnAuthorized(ctx  *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(APIResponse{
		Code: "not_authorized",
		Message: msg,
	})
}