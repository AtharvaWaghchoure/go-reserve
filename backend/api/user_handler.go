package api

import (
	"errors"

	"github.com/AtharvaWaghchoure/goreserve/db"
	"github.com/AtharvaWaghchoure/goreserve/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore,
	}
}

func (handler *UserHandler) HandlePutUser(ctx *fiber.Ctx) error {
	var (
		// values bson.M
		params types.UpdateUserParams
		userID = ctx.Params("id")
	)
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil
	}
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	if err := handler.userStore.UpdateUser(ctx.Context(), filter, params); err != nil {
		return err
	}
	return ctx.JSON(map[string]string{"updated": userID})
}

func (handler *UserHandler) HandleDeleteUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	if err := handler.userStore.DeleteUser(ctx.Context(), userID); err != nil {
		return err
	}

	return ctx.JSON(map[string]string{"deleted": userID})
}

func (handler *UserHandler) HandlePostUser(ctx *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := handler.userStore.InsertUser(ctx.Context(), user)
	if err != nil {
		return err
	}
	return ctx.JSON(insertedUser)
}

func (handler *UserHandler) HandleGetUser(ctx *fiber.Ctx) error {
	var (
		id = ctx.Params("id")
	)
	user, err := handler.userStore.GetUserByID(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.JSON(map[string]string{"error": "not found"})
		}
		return err
	}
	return ctx.JSON(user)
}

func (handler *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {
	users, err := handler.userStore.GetUsers(ctx.Context())
	if err != nil {
		return nil
	}
	return ctx.JSON(users)
}
