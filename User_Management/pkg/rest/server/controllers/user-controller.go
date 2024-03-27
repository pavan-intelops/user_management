package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavan-intelops/user_management/user_management/pkg/rest/server/daos/clients/sqls"
	"github.com/pavan-intelops/user_management/user_management/pkg/rest/server/models"
	"github.com/pavan-intelops/user_management/user_management/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	"strconv"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() (*UserController, error) {
	userService, err := services.NewUserService()
	if err != nil {
		return nil, err
	}
	return &UserController{
		userService: userService,
	}, nil
}
