package sku

import (
	. "shop/proto/proto"
)

type SpecParamService struct {
	skuRepository Repository
	UnimplementedSkuSrvServer
}

func NewCategoryService(skuRepository Repository) SkuSrvServer {
	return &SpecParamService{skuRepository: skuRepository}
}
