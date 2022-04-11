package spu

import (
	. "shop/proto/proto"
)

type SpecParamService struct {
	spuRepo Repository
	UnimplementedSpuSrvServer
}

func NewSpuService(spuRepo Repository) SpuSrvServer {
	return &SpecParamService{spuRepo: spuRepo}
}
