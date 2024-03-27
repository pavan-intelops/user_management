package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavan-intelops/user_management/client_service/pkg/rest/server/daos/clients/sqls"
	"github.com/pavan-intelops/user_management/client_service/pkg/rest/server/models"
	"github.com/pavan-intelops/user_management/client_service/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	"strconv"
)

type Client_methodsController struct {
	clientMethodsService *services.Client_methodsService
}

func NewClient_methodsController() (*Client_methodsController, error) {
	clientMethodsService, err := services.NewClient_methodsService()
	if err != nil {
		return nil, err
	}
	return &Client_methodsController{
		clientMethodsService: clientMethodsService,
	}, nil
}

func (clientMethodsController *Client_methodsController) CreateClient_methods(context *gin.Context) {
	// validate input
	var input models.Client_methods
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// trigger clientMethods creation
	clientMethodsCreated, err := clientMethodsController.clientMethodsService.CreateClient_methods(&input)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, clientMethodsCreated)
}

func (clientMethodsController *Client_methodsController) ListClient_methods(context *gin.Context) {
	// trigger all clientMethods fetching
	clientMethods, err := clientMethodsController.clientMethodsService.ListClient_methods()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, clientMethods)
}

func (clientMethodsController *Client_methodsController) FetchClient_methods(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger clientMethods fetching
	clientMethods, err := clientMethodsController.clientMethodsService.GetClient_methods(id)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	serviceName := os.Getenv("SERVICE_NAME")
	collectorURL := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if len(serviceName) > 0 && len(collectorURL) > 0 {
		// get the current span by the request context
		currentSpan := trace.SpanFromContext(context.Request.Context())
		currentSpan.SetAttributes(attribute.String("clientMethods.id", strconv.FormatInt(clientMethods.Id, 10)))
	}

	context.JSON(http.StatusOK, clientMethods)
}

func (clientMethodsController *Client_methodsController) UpdateClient_methods(context *gin.Context) {
	// validate input
	var input models.Client_methods
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger clientMethods update
	if _, err := clientMethodsController.clientMethodsService.UpdateClient_methods(id, &input); err != nil {
		log.Error(err)
		if errors.Is(err, sqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func (clientMethodsController *Client_methodsController) DeleteClient_methods(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger clientMethods deletion
	if err := clientMethodsController.clientMethodsService.DeleteClient_methods(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}
