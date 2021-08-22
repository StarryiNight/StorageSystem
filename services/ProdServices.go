package services

import (
	"context"
	"storageSystem/dao"
	"storageSystem/pbfiles"
)

type ProdService struct {

}


func (p *ProdService)GetProdStock(ctx context.Context, u *pbfiles.ProdRequest) (*pbfiles.ProdResponse, error)  {
	value, err := dao.Query(u.Key)
	if err != nil {
		return nil, err
	}
	return &pbfiles.ProdResponse{Value: value},nil
}

func (p *ProdService)RegisterProdStock(ctx context.Context, u *pbfiles.ProdRegister) (*pbfiles.ProdResponse, error) {
	err := dao.Insert(*u)
	if err != nil {
		return nil, err
	}
	return &pbfiles.ProdResponse{Value: "添加成功"}, nil

}

func (p *ProdService)UpdateProdStock(ctx context.Context, u *pbfiles.ProdRegister) (*pbfiles.ProdResponse, error){
	err:=dao.Update(*u)
	if err != nil {
		return nil,err
	}
	return &pbfiles.ProdResponse{Value: "修改成功"},nil
}

func (p *ProdService)DeleteProdStock(ctx context.Context, u *pbfiles.ProdRequest) (*pbfiles.ProdResponse, error){
	err:=dao.Delete(u.Key)
	if err != nil {
		return nil,err
	}
	return &pbfiles.ProdResponse{Value: "删除成功"},nil
}