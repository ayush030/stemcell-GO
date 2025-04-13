package hornet

import (
	"context"
	"fmt"
	"hornet/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (service *Service) GetResourceWrapper(c echo.Context) error {
	ctx := c.Request().Context()
	identifier := c.Param("id")

	resource, err := service.GetResource(ctx, identifier)
	if err != nil {
		fmt.Println("GetResourceWrapper: Unable to get resource. Error: " + err.Error())
		return err
	}

	return c.JSON(http.StatusOK, resource)
}

func (service *Service) ListResourceWrapper(c echo.Context) error {
	ctx := c.Request().Context()

	resources, err := service.ListResource(ctx)
	if err != nil {
		fmt.Println("ListResourceWrapper: Unable to get resource. Error: " + err.Error())
		return err
	}

	return c.JSON(http.StatusOK, resources)
}

func (service *Service) CreateResourceWrapper(c echo.Context) error {
	ctx := c.Request().Context()

	createPayload := &models.CreateUpdateResource{}

	err := c.Bind(&createPayload)
	if err != nil {
		fmt.Println("Invalid update request. Error: " + err.Error())
		return err
	}

	resource, err := service.CreateResource(ctx, createPayload)
	if err != nil {
		fmt.Println("Error in creating request. Error: " + err.Error())
		return err
	}

	return c.JSON(http.StatusCreated, resource)
}

func (service *Service) UpdateResourceWrapper(c echo.Context) error {
	ctx := c.Request().Context()
	identifier := c.Param("id")
	updatePayload := &models.CreateUpdateResource{}

	err := c.Bind(&updatePayload)
	if err != nil {
		fmt.Println("Invalid update request. Error: " + err.Error())
		return err
	}

	err = service.UpdateResource(ctx, identifier, updatePayload)
	if err != nil {
		fmt.Println("Error in updating request. Error: " + err.Error())
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (service *Service) DeleteResourceWrapper(c echo.Context) error {
	ctx := c.Request().Context()
	identifier := c.Param("id")

	err := service.DeleteResource(ctx, identifier)
	if err != nil {
		fmt.Println("Error in updating request. Error: " + err.Error())
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}

// ---------------------------------------------------------------------------------------------------

func (service *Service) GetResource(ctx context.Context, identifier string) (interface{}, error) {
	resource := models.Resource{}

	err := service.DB(ctx).Table("resources").Where("id=?", identifier).Take(&resource).Error
	if err != nil {
		fmt.Println("GetResource: Unable to get resource form DB. Error: " + err.Error())

		return nil, err
	}

	return resource, nil
}

func (service *Service) ListResource(ctx context.Context) (interface{}, error) {
	var resources []models.Resource

	err := service.DB(ctx).Table("resources").Select("*").Find(&resources).Error
	if err != nil {
		fmt.Println("ListResource: Unable to list resources from DB. Error: " + err.Error())
		return nil, err
	}

	return resources, nil
}

func (service *Service) CreateResource(ctx context.Context, payload *models.CreateUpdateResource) (interface{}, error) {
	resource := models.NewResource(payload.Payload)

	err := service.DB(ctx).Table("resources").Create(&resource).Error
	if err != nil {
		fmt.Println("CreateResource: Unable to create resource in DB. Error: " + err.Error())
		return nil, err
	}

	return resource, nil
}

func (service *Service) UpdateResource(ctx context.Context, identifier string, payload *models.CreateUpdateResource) error {
	err := service.DB(ctx).Table("resources").Where("id=?", identifier).Update("payload", payload.Payload).Error
	if err != nil {
		fmt.Println("UpdateResource: Unable to update resource in DB. Error: " + err.Error())
		return err
	}

	return nil
}

func (service *Service) DeleteResource(ctx context.Context, identifier string) error {
	err := service.DB(ctx).Table("resources").Delete(&models.Resource{}, "id="+"'"+identifier+"'").Error
	if err != nil {
		fmt.Println("DeleteResource: Unable to delete resource in DB. Error: " + err.Error())
		return err
	}

	return nil
}
