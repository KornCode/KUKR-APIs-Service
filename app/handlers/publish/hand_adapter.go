package publishhand

import (
	handler "github.com/KornCode/KUKR-APIs-Service/app/handlers"
	publishsrv "github.com/KornCode/KUKR-APIs-Service/app/services/publish"
	"github.com/gofiber/fiber/v2"
)

type publishHandler struct {
	publishService publishsrv.PublishService
}

func NewPublishHandler(publishService publishsrv.PublishService) publishHandler {
	return publishHandler{publishService}
}

func (h publishHandler) CreateOne(c *fiber.Ctx) error {
	js := new(jsonPublishCreateOne)
	if err := c.BodyParser(js); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}
	if err := validateStruct(js); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}

	publish := publishsrv.Publish{
		Category:    js.Category,
		Text:        js.Text,
		Bibid:       js.Bibid,
		ContentType: js.ContentType,
		TitleTh:     js.TitleTh,
		TitleEn:     js.TitleEn,
		PubYear:     js.PubYear,
		AuthorTh:    js.AuthorTh,
		AuthorEn:    js.AuthorEn,
		Link:        js.Link,
	}

	serv_result_pk, err := h.publishService.CreateOne(&publish)
	if err != nil {
		return handler.HandleError(c, err)
	}

	return c.Status(201).JSON(fiber.Map{"data": fiber.Map{
		"primary_key": serv_result_pk,
	}})
}

func (h publishHandler) UpdateOneByPK(c *fiber.Ctx) error {
	js := new(jsonPublishUpdateOneByPK)
	if err := c.BodyParser(js); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}
	if err := validateStruct(js); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}

	publish := publishsrv.Publish{
		ID:          js.ID,
		Category:    js.Category,
		Text:        js.Text,
		Bibid:       js.Bibid,
		ContentType: js.ContentType,
		TitleTh:     js.TitleTh,
		TitleEn:     js.TitleEn,
		PubYear:     js.PubYear,
		AuthorTh:    js.AuthorTh,
		AuthorEn:    js.AuthorEn,
		Link:        js.Link,
	}

	if err := h.publishService.UpdateOneByPK(publish.ID, &publish); err != nil {
		return handler.HandleError(c, err)
	}

	return c.SendStatus(200)
}

func (h publishHandler) DeleteOneByPK(c *fiber.Ctx) error {
	js := new(jsonPublishDeleteOneByPK)
	if err := c.BodyParser(js); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}
	if err := validateStruct(js); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}

	if err := h.publishService.DeleteOneByPK(js.ID); err != nil {
		return handler.HandleError(c, err)
	}

	return c.SendStatus(200)
}

func (h publishHandler) GetByCategoryAndPubYear(c *fiber.Ctx) error {
	qs := new(queryPublishCategoryAndPubYear)
	if err := c.QueryParser(qs); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := validateStruct(qs); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}

	serv_results, err := h.publishService.GetByCategoryAndPubYear(qs.Category, qs.PubYear)
	if err != nil {
		return handler.HandleError(c, err)
	}

	return c.Status(200).JSON(fiber.Map{"data": serv_results})
}

func (h publishHandler) GetByBibid(c *fiber.Ctx) error {
	qs := new(queryPublishBibid)
	if err := c.QueryParser(qs); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := validateStruct(qs); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}

	serv_result, err := h.publishService.GetByBibid(qs.Bibid)
	if err != nil {
		return handler.HandleError(c, err)
	}

	return c.Status(200).JSON(fiber.Map{"data": serv_result})
}

func (h publishHandler) GetPaginateByOptions(c *fiber.Ctx) error {
	qs := new(queryPublishPaginateByOptions)
	if err := c.QueryParser(qs); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := validateStruct(qs); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}

	page := qs.Page
	if page == 0 {
		page = 1
	}

	limit := qs.Limit
	if limit == 0 {
		limit = 10
	}

	serv_result, err := h.publishService.GetPaginateByOptions(
		page, limit, map[string]interface{}{
			"category": qs.Category,
			"pub_year": qs.PubYear,
		},
	)
	if err != nil {
		return handler.HandleError(c, err)
	}

	return c.Status(200).JSON(fiber.Map{"data": serv_result})
}

func (h publishHandler) SyncDataSource(c *fiber.Ctx) error {
	js := new(jsonPublishPubYear)
	if err := c.BodyParser(js); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}
	if err := validateStruct(js); err != nil {
		return c.Status(406).JSON(fiber.Map{
			"error": err,
		})
	}

	err := h.publishService.SyncDataSource(js.PubYear)
	if err != nil {
		return handler.HandleError(c, err)

	}

	return c.SendStatus(201)
}
