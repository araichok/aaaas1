package grpc

import (
	"context"
	"inventory-service/internal/domain"
	"inventory-service/internal/usecase"
	"inventory-service/inventory-service/proto/inventorypb"
)

type InventoryServer struct {
	inventorypb.UnimplementedInventoryServiceServer
	usecase usecase.ProductUsecase
}

func NewInventoryServer(u usecase.ProductUsecase) *InventoryServer {
	return &InventoryServer{usecase: u}
}

func (s *InventoryServer) CreateProduct(ctx context.Context, req *inventorypb.CreateProductRequest) (*inventorypb.ProductResponse, error) {
	p := req.Product
	domainProd := &domain.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       int(p.Stock),
		Category:    p.Category,
	}
	err := s.usecase.Create(domainProd)
	if err != nil {
		return nil, err
	}
	return &inventorypb.ProductResponse{Product: &inventorypb.Product{
		Id:          domainProd.ID,
		Name:        domainProd.Name,
		Description: domainProd.Description,
		Price:       domainProd.Price,
		Stock:       int32(domainProd.Stock),
		Category:    domainProd.Category,
	}}, nil
}

func (s *InventoryServer) GetProductByID(ctx context.Context, req *inventorypb.GetProductRequest) (*inventorypb.ProductResponse, error) {
	product, err := s.usecase.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &inventorypb.ProductResponse{Product: &inventorypb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
		Category:    product.Category,
	}}, nil
}

func (s *InventoryServer) UpdateProduct(ctx context.Context, req *inventorypb.UpdateProductRequest) (*inventorypb.ProductResponse, error) {
	p := req.Product
	domainProd := &domain.Product{
		ID:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       int(p.Stock),
		Category:    p.Category,
	}
	err := s.usecase.Update(domainProd)
	if err != nil {
		return nil, err
	}
	return &inventorypb.ProductResponse{Product: &inventorypb.Product{
		Id:          domainProd.ID,
		Name:        domainProd.Name,
		Description: domainProd.Description,
		Price:       domainProd.Price,
		Stock:       int32(domainProd.Stock),
		Category:    domainProd.Category,
	}}, nil
}

func (s *InventoryServer) DeleteProduct(ctx context.Context, req *inventorypb.DeleteProductRequest) (*inventorypb.Empty, error) {
	err := s.usecase.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &inventorypb.Empty{}, nil
}

func (s *InventoryServer) ListProducts(ctx context.Context, req *inventorypb.ListProductsRequest) (*inventorypb.ListProductsResponse, error) {
	products, err := s.usecase.List()
	if err != nil {
		return nil, err
	}
	var pbProducts []*inventorypb.Product
	for _, p := range products {
		pbProducts = append(pbProducts, &inventorypb.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       int32(p.Stock),
			Category:    p.Category,
		})
	}
	return &inventorypb.ListProductsResponse{Products: pbProducts}, nil
}
