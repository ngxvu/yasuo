package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"gitlab.com/merakilab9/yasuo/pkg/model"
	"gitlab.com/merakilab9/yasuo/pkg/service"
	"log"
)

type CategoryQueueHandlers struct {
	service service.CategoryInterface
}

func NewCategoryQueueHandlers(service service.CategoryInterface) *CategoryQueueHandlers {
	return &CategoryQueueHandlers{service: service}
}

func (h *CategoryQueueHandlers) HandleJsonCateDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p model.ShopeeCateData
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Json to Yasuo: data=%s\n", p.Data)

	fmt.Println(p.Data.CategoryList)

	catelist := p.Data.CategoryList
	if err := h.service.SaveCate(ctx, catelist); err != nil {
		return err
	}
	return nil

}
