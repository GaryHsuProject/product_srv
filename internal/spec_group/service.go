package spec_group

import (
	. "shop/proto/proto"
)

type SpecGroupService struct {
	specGroupRepo Repository
	UnimplementedSpecGroupSrvServer
}

func NewCategoryService(specGroupRepo Repository) SpecGroupSrvServer {
	return &SpecGroupService{specGroupRepo: specGroupRepo}
}
