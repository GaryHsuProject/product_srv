package spec_param

import (
	"context"
	"shop/driver"
	"shop/models"
	. "shop/proto/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SpecParamService struct {
	specParamRepo Repository
}

func NewCategoryService(specParamRepo Repository) SpecParamSrvServer {
	return &SpecParamService{specParamRepo: specParamRepo}
}

func (s *SpecParamService) GetAll(context.Context, *emptypb.Empty) (*SpecParams, error) {
	specParams, err := s.specParamRepo.FindAll()
	if err != nil {
		driver.Logger.Error(err)
	}
	result := SpecParams{}

	for _, specParam := range specParams {
		spDetail := SpecParamDetail{
			Id:        int32(specParam.Id),
			SpgId:     int32(specParam.SpgId),
			Name:      specParam.Name,
			Numeric:   specParam.Numeric,
			Unit:      specParam.Unit,
			Searching: specParam.Searching,
			Generic:   specParam.Generic,
			Segements: specParam.Segements,
		}
		result.SpecParam = append(result.SpecParam, &spDetail)
	}
	return &result, nil
}

func (s *SpecParamService) Create(ctx context.Context, req *CreateSpecParam) (*emptypb.Empty, error) {
	specParam := &models.SpecParam{
		SpgId:     int(req.SpgId),
		Name:      req.Name,
		Numeric:   req.Numeric,
		Searching: req.Searching,
		Generic:   req.Generic,
		Segements: req.Segements,
		Unit:      req.Unit,
	}
	err := s.specParamRepo.Insert(specParam)
	if err != nil {
		driver.Logger.Error(err)
	}
	return &emptypb.Empty{}, err
}
