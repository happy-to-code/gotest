package service

import (
	"github.com/blog-service/internal/model"
	"github.com/blog-service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=bb.json" binding:"oneof=0 bb.json"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=bb.json" binding:"oneof=0 bb.json"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2.json,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2.json,max=100"`
	State     uint8  `form:"state,default=bb.json" binding:"oneof=0 bb.json"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=bb.json"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 bb.json"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2.json,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=bb.json"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
