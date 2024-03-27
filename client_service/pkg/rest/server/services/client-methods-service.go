package services

import (
	"github.com/pavan-intelops/user_management/client_service/pkg/rest/server/daos"
	"github.com/pavan-intelops/user_management/client_service/pkg/rest/server/models"
)

type Client_methodsService struct {
	clientMethodsDao *daos.Client_methodsDao
}

func NewClient_methodsService() (*Client_methodsService, error) {
	clientMethodsDao, err := daos.NewClient_methodsDao()
	if err != nil {
		return nil, err
	}
	return &Client_methodsService{
		clientMethodsDao: clientMethodsDao,
	}, nil
}

func (clientMethodsService *Client_methodsService) CreateClient_methods(clientMethods *models.Client_methods) (*models.Client_methods, error) {
	return clientMethodsService.clientMethodsDao.CreateClient_methods(clientMethods)
}

func (clientMethodsService *Client_methodsService) ListClient_methods() ([]*models.Client_methods, error) {
	return clientMethodsService.clientMethodsDao.ListClient_methods()
}

func (clientMethodsService *Client_methodsService) GetClient_methods(id int64) (*models.Client_methods, error) {
	return clientMethodsService.clientMethodsDao.GetClient_methods(id)
}

func (clientMethodsService *Client_methodsService) UpdateClient_methods(id int64, clientMethods *models.Client_methods) (*models.Client_methods, error) {
	return clientMethodsService.clientMethodsDao.UpdateClient_methods(id, clientMethods)
}

func (clientMethodsService *Client_methodsService) DeleteClient_methods(id int64) error {
	return clientMethodsService.clientMethodsDao.DeleteClient_methods(id)
}
