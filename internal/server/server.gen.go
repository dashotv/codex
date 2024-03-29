// Code generated by oto; DO NOT EDIT.

package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FileService interface {
	Index(echo.Context, *IndexRequest) (*FilesResponse, error)
	Walk(echo.Context, *KeyRequest) (*EmptyResponse, error)
}

type JobService interface {
	Create(echo.Context, *Request) (*JobResponse, error)
	Index(echo.Context, *Request) (*JobsResponse, error)
}

type fileServiceServer struct {
	fileService FileService
}

// Register adds the FileService to the otohttp.Server.
func RegisterFileService(e *echo.Group, fileService FileService) {
	handler := &fileServiceServer{
		fileService: fileService,
	}
	e.POST("/FileService.Index", handler.handleIndex)
	e.POST("/FileService.Walk", handler.handleWalk)
}

func (s *fileServiceServer) handleIndex(c echo.Context) error {
	request := &IndexRequest{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.fileService.Index(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *fileServiceServer) handleWalk(c echo.Context) error {
	request := &KeyRequest{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.fileService.Walk(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

type jobServiceServer struct {
	jobService JobService
}

// Register adds the JobService to the otohttp.Server.
func RegisterJobService(e *echo.Group, jobService JobService) {
	handler := &jobServiceServer{
		jobService: jobService,
	}
	e.POST("/JobService.Create", handler.handleCreate)
	e.POST("/JobService.Index", handler.handleIndex)
}

func (s *jobServiceServer) handleCreate(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.jobService.Create(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *jobServiceServer) handleIndex(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.jobService.Index(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}
