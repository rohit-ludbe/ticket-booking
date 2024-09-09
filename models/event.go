package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID                   uint      `json:"id" gorm:"primarykey"`
	Name                 string    `json:"name"`
	Location             string    `json:"location"`
	TotalTicketPurchased int64     `json:"totalTicketPurchased" gorm:"-"`
	TotalTicketEntered   int64     `json:"totalTicketEntered" gorm:"-"`
	Date                 time.Time `json:"date"`
	CreateAt             time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

type EventRepository interface {
	GetAllEvents(ctx context.Context) ([]*Event, error)
	GetOneEvent(ctx context.Context, eventId uint) (*Event, error)
	CreateOneEvent(ctx context.Context, event *Event) (*Event, error)
	UpdateOneEvent(ctx context.Context, eventId uint, updateData map[string]interface{}) (*Event, error)
	DeleteOneEvent(ctx context.Context, eventId uint) error
}

func (e *Event) AfterFind(db *gorm.DB) (err error) {
	baseQuery := db.Model(&Ticket{}).Where(&Ticket{EventID: e.ID})

	if res := baseQuery.Count(&e.TotalTicketPurchased); res.Error != nil {
		return res.Error
	}

	if res := baseQuery.Where("entered = ?", true).Count(&e.TotalTicketEntered); res.Error != nil {
		return res.Error
	}

	return nil
}
