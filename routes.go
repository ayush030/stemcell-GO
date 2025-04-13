package hornet

import "github.com/labstack/echo/v4"

func (service *Service) AddRoutes(server *echo.Echo, s *Service) {
	/*
		- POST /api/resource - Create a new resource
		- GET /api/resource/{id} - Get a specific resource by ID
		- PUT /api/resource/{id} - Update an existing resource by ID
		- DELETE /api/resource/{id} - Delete a resource by ID

		* GET /api/resource/ - List all resources
	*/
	server.GET("/api/resource/:id", s.GetResourceWrapper)
	server.GET("/api/resource/:id/", s.GetResourceWrapper)

	server.GET("/api/resource", s.ListResourceWrapper)
	server.GET("/api/resource/", s.ListResourceWrapper)

	server.POST("/api/resource", s.CreateResourceWrapper)
	server.POST("/api/resource/", s.CreateResourceWrapper)

	server.PUT("/api/resource/:id", s.UpdateResourceWrapper)
	server.PUT("/api/resource/:id/", s.UpdateResourceWrapper)

	server.DELETE("/api/resource/:id", s.DeleteResourceWrapper)
	server.DELETE("/api/resource/:id/", s.DeleteResourceWrapper)
}
