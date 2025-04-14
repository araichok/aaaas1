package grpc

import (
	"context"
	"errors"
	"order-service/internal/domain"
	"order-service/internal/usecase"
	"order-service/order-service/proto/orderpb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderServer реализует orderpb.OrderServiceServer
type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	usecase *usecase.OrderUsecase
}

func NewOrderServer(u *usecase.OrderUsecase) *OrderServer {
	return &OrderServer{usecase: u}
}

// convertDomainOrderToPB преобразует доменную модель Order в proto-модель
func convertDomainOrderToPB(o *domain.Order) *orderpb.Order {
	var pbItems []*orderpb.ProductItem
	for _, item := range o.Items {
		pbItems = append(pbItems, &orderpb.ProductItem{
			Product_id: item.ProductID.Hex(),
			Quantity:   int32(item.Quantity),
		})
	}

	return &orderpb.Order{
		Id:         o.ID.Hex(),
		User_id:    o.UserID,
		Items:      pbItems,
		Status:     string(o.Status),
		Created_at: o.CreatedAt,
	}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.OrderResponse, error) {
	pbOrder := req.Order

	// Преобразование входных данных в доменную модель
	var domainItems []domain.ProductItem
	for _, item := range pbOrder.Items {
		// Преобразуем строковое значение в ObjectID
		objID, err := primitive.ObjectIDFromHex(item.Product_id)
		if err != nil {
			return nil, errors.New("invalid product_id format")
		}
		domainItems = append(domainItems, domain.ProductItem{
			ProductID: objID,
			Quantity:  int(item.Quantity),
		})
	}

	domainOrder := &domain.Order{
		UserID: pbOrder.User_id,
		Items:  domainItems,
		Status: domain.OrderStatus(pbOrder.Status),
		// CreatedAt будет установлен в репозитории
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
	// Сначала пробуем обновить статус заказа
	err := s.usecase.UpdateStatus(req.Id, domain.OrderStatus(req.Status))
	if err != nil {
		return nil, err
	}

	// После обновления получим заказ для возврата
	domainOrder, err := s.usecase.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &orderpb.OrderResponse{
		Order: convertDomainOrderToPB(domainOrder),
	}, nil
}

func (s *OrderServer) ListUserOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
	domainOrders, err := s.usecase.GetByUser(req.User_id)
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
