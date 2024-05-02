package controllers

import (
	"repo_pattern/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{taskService}
}

// @Summary Create a New Task
// @Description Create a new task
// @Tags Task
// @Accept json
// @Produce json
// @Param task body entities.CreateTask true "Task object to be created"
// @Success 201 {object} models.Task
// @Router /tasks [post]
func (tc *TaskController) CreateTask(c *fiber.Ctx) error {
	var task services.TaskServiceStruct
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := tc.taskService.CreateTask(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

// @Summary Update Task by ID
// @Description Update a task by its ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path integer true "Task ID" Format(uint64)
// @Param task body entities.UpdateTask true "Updated Task object"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [put]
func (tc *TaskController) UpdateTask(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task ID"})
	}

	var task services.TaskServiceStruct
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := tc.taskService.UpdateTask(id, &task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

// @Summary Delete Task by ID
// @Description Delete a task by its ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path integer true "Task ID" Format(uint64)
// @Success 204 "No Content"
// @Router /tasks/{id} [delete]
func (tc *TaskController) DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task ID"})
	}

	if err := tc.taskService.DeleteTask(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// @Summary Get all tasks
// @Description Get all tasks
// @Tags Task
// @Router /tasks [get]
func (tc *TaskController) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := tc.taskService.GetAllTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(tasks)
}

// @Summary Get Task by ID
// @Description Get a task by its ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path integer true "Task ID" Format(uint64)
// @Success 200 {object} entities.Task
// @Router /tasks/{id} [get]
func (tc *TaskController) GetTaskById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task ID"})
	}

	task, err := tc.taskService.GetTaskById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(task)
}

// @Summary Change Task Status
// @Description Change status of a task by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path integer true "Task ID" Format(uint64)
// @Param task body entities.StatusChangeRequest true "Status object to be created"
// @Success 200 {object} entities.Task
// @Router /tasks/status/{id} [patch]
func (tc *TaskController) ChangeTaskStatus(c *fiber.Ctx) error {
	// Parse task ID from request path parameter
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task ID"})
	}

	// Parse new status from request body
	var status services.StatusChangeStruct
	if err := c.BodyParser(&status); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Call service method to change task status
	err = tc.taskService.ChangeTaskStatus(uint(id), &status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
