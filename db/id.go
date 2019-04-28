package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
)

//Model ID string以实现多库ID唯一
type Model struct {
	ID        string     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func init() {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
}

type IDResolveInterface interface {
	ResolveID() interface{}
}

func (m Model) ResolveID() interface{} {
	return m.ID
}

type CreateTimeResolveInterface interface {
	ResolveCreateTime() interface{}
}

type UpdateTimeResolveInterface interface {
	ResolveUpdateTime() interface{}
}

func (m Model) ResolveCreateTime() interface{} {
	return m.CreatedAt.UTC()
}

func (m Model) ResolveUpdateTime() interface{} {
	return m.CreatedAt.UTC()
}

func newID() string {
	return ksuid.New().String()
}

//CreateCallback AutoMaintain ID
func CreateCallback(scope *gorm.Scope) {
	if scope.HasColumn("ID") {
		newid := newID()
		scope.SetColumn("ID", newid)
	}
}
