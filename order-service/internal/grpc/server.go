package grpc

import (
	"context"
	"errors"
	"order-service/internal/domain"
	"order-service/internal/usecase"
	"order-service/order-service/proto/orderpb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	usecase *usecase.OrderUsecase
}

func NewOrderServer(u *usecase.OrderUsecase) *OrderServer {
	return &OrderServer{usecase: u}
}
func convertDomainOrderToPB(o *domain.Order) *orderpb.Order {
	var pbItems []*orderpb.ProductItem
	for _, item := range o.Items {
		pbItems = append(pbItems, &orderpb.ProductItem{
			ProductId: item.ProductID.Hex(),
			Quantity:  int32(item.Quantity),
		})
	}

	return &orderpb.Order{
		Id:        o.ID.Hex(),
		UserId:    o.UserID,
		Items:     pbItems,
		Status:    string(o.Status),
		CreatedAt: o.CreatedAt,
	}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.OrderResponse, error) {
	pbOrder := req.Order

	var domainItems []domain.ProductItem
	for _, item := range pbOrder.Items {
		objID, err := primitive.ObjectIDFromHex(item.ProductId)
		if err != nil {
			return nil, errors.New("invalid product_id format")
		}
		domainItems = append(domainItems, domain.ProductItem{
			ProductID: objID,
			Quantity:  int(item.Quantity),
		})
	}

	domainOrder := &domain.Order{
		UserID: pbOrder.UserId,
		Items:  domainItems,
		Status: domain.OrderStatus(pbOrder.Status),
	}

	if err := s.usecase.Create(domainOrder); err != nil {
		return nil, err
	}

	return &orderpb.OrderResponse{
		Order: convertDomainOrderToPB(domainOrder),
	}, nil
}

func (s *OrderServer) GetOrderByID(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.OrderResponse, error) {
	domainOrder, err := s.usecase.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &orderpb.OrderResponse{
		Order: convertDomainOrderToPB(domainOrder),
	}, nil
}

func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *orderpb.UpdateOrderStatusRequest) (*orderpb.OrderResponse, error) {

	err := s.usecase.UpdateStatus(req.Id, domain.OrderStatus(req.Status))
	if err != nil {
		return nil, err
	}

	domainOrder, err := s.usecase.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &orderpb.OrderResponse{
		Order: convertDomainOrderToPB(domainOrder),
	}, nil
}

func (s *OrderServer) ListUserOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
	domainOrders, err := s.usecase.GetByUser(req.UserId)
	if err != nil {
		return nil, err
	}

	var pbOrders []*orderpb.Order
	for _, order := range domainOrders {
		pbOrders = append(pbOrders, convertDomainOrderToPB(&order))
	}
	return &orderpb.ListOrdersResponse{
		Orders: pbOrders,
	}, nil
}
