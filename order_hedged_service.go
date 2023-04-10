package phemex

import (
	"context"
	"encoding/json"

	"github.com/Krisa/go-phemex/common"
)

// CreateOrderHedgedService create order
type CreateOrderHedgedService struct {
	c                *Client
	hedged           bool
	symbol           string
	clOrdID          *string
	actionBy         *string
	side             SideType
	posSide          PosSideType
	orderQtyRq       *float64
	priceRp          *float64
	ordType          *OrderType
	stopPxEp         *int64
	timeInForce      *TimeInForceType
	reduceOnly       *bool
	closeOnTrigger   *bool
	takeProfitEp     *int64
	stopLossEp       *int64
	pegOffsetValueEp *int64
	triggerType      *TriggerType
	text             *string
	pegPriceType     *string
}

// Symbol set symbol
func (s *CreateOrderHedgedService) Symbol(symbol string) *CreateOrderHedgedService {
	s.symbol = symbol
	return s
}

// ClOrdID set clOrID
func (s *CreateOrderHedgedService) ClOrdID(clOrdID string) *CreateOrderHedgedService {
	s.clOrdID = &clOrdID
	return s
}

// ActionBy set actionBy
func (s *CreateOrderHedgedService) ActionBy(actionBy string) *CreateOrderHedgedService {
	s.actionBy = &actionBy
	return s
}

// Side set side
func (s *CreateOrderHedgedService) Side(side SideType) *CreateOrderHedgedService {
	s.side = side
	return s
}

// PosSide set posSide
func (s *CreateOrderHedgedService) PosSide(posSide PosSideType) *CreateOrderHedgedService {
	s.posSide = posSide
	return s
}

// OrderQtyRq set orderQtyRq
func (s *CreateOrderHedgedService) OrderQtyRq(orderQtyRq float64) *CreateOrderHedgedService {
	s.orderQtyRq = &orderQtyRq
	return s
}

// PriceRp set priceRp
func (s *CreateOrderHedgedService) PriceRp(priceRp float64) *CreateOrderHedgedService {
	s.priceRp = &priceRp
	return s
}

// OrdType set ordType
func (s *CreateOrderHedgedService) OrdType(ordType OrderType) *CreateOrderHedgedService {
	s.ordType = &ordType
	return s
}

// StopPxEp set stopPxEp
func (s *CreateOrderHedgedService) StopPxEp(stopPxEp int64) *CreateOrderHedgedService {
	s.stopPxEp = &stopPxEp
	return s
}

// TimeInForce set timeInForce
func (s *CreateOrderHedgedService) TimeInForce(timeInForce TimeInForceType) *CreateOrderHedgedService {
	s.timeInForce = &timeInForce
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateOrderHedgedService) ReduceOnly(reduceOnly bool) *CreateOrderHedgedService {
	s.reduceOnly = &reduceOnly
	return s
}

// CloseOnTrigger set closeOnTrigger
func (s *CreateOrderHedgedService) CloseOnTrigger(closeOnTrigger bool) *CreateOrderHedgedService {
	s.closeOnTrigger = &closeOnTrigger
	return s
}

// TakeProfitEp set takeProfitEp
func (s *CreateOrderHedgedService) TakeProfitEp(takeProfitEp int64) *CreateOrderHedgedService {
	s.takeProfitEp = &takeProfitEp
	return s
}

// StopLossEp set stopLossEp
func (s *CreateOrderHedgedService) StopLossEp(stopLossEp int64) *CreateOrderHedgedService {
	s.stopLossEp = &stopLossEp
	return s
}

// TriggerType set triggerType
func (s *CreateOrderHedgedService) TriggerType(triggerType TriggerType) *CreateOrderHedgedService {
	s.triggerType = &triggerType
	return s
}

// Text set text
func (s *CreateOrderHedgedService) Text(text string) *CreateOrderHedgedService {
	s.text = &text
	return s
}

// PegOffsetValueEp set pegOffsetValueEp
func (s *CreateOrderHedgedService) PegOffsetValueEp(pegOffsetValueEp int64) *CreateOrderHedgedService {
	s.pegOffsetValueEp = &pegOffsetValueEp
	return s
}

// PegPriceType set pegPriceType
func (s *CreateOrderHedgedService) PegPriceType(pegPriceType string) *CreateOrderHedgedService {
	s.pegPriceType = &pegPriceType
	return s
}

func (s *CreateOrderHedgedService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {

	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":  s.symbol,
		"side":    s.side,
		"posSide": s.posSide,
	}
	if s.clOrdID != nil {
		m["clOrdID"] = *s.clOrdID
	}
	if s.orderQtyRq != nil {
		m["orderQtyRq"] = *s.orderQtyRq
	}
	if s.actionBy != nil {
		m["actionBy"] = *s.actionBy
	}
	if s.priceRp != nil {
		m["priceRp"] = *s.priceRp
	}
	if s.ordType != nil {
		m["ordType"] = *s.ordType
	}
	if s.stopPxEp != nil {
		m["stopPxEp"] = *s.stopPxEp
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.closeOnTrigger != nil {
		m["closeOnTrigger"] = *s.closeOnTrigger
	}
	if s.takeProfitEp != nil {
		m["takeProfitEp"] = *s.takeProfitEp
	}
	if s.stopLossEp != nil {
		m["stopLossEp"] = *s.stopLossEp
	}
	if s.triggerType != nil {
		m["triggerType"] = *s.triggerType
	}
	if s.text != nil {
		m["text"] = *s.text
	}
	if s.pegPriceType != nil {
		m["pegPriceType"] = *s.pegPriceType
	}
	if s.pegOffsetValueEp != nil {
		m["pegOffsetValueEp"] = *s.pegOffsetValueEp
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateOrderHedgedService) Do(ctx context.Context, opts ...RequestOption) (res *OrderResponse, err error) {

	data, err := s.createOrder(ctx, "/g-orders", opts...)
	if err != nil {
		return nil, err
	}
	resp := new(BaseResponse)
	resp.Data = new(OrderResponse)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, &common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return resp.Data.(*OrderResponse), nil
}

// CreateReplaceOrderHedgedService create order
type CreateReplaceOrderHedgedService struct {
	c            *Client
	symbol       string
	posSide      PosSideType
	orderID      string
	origClOrdID  *string
	clOrdID      *string
	priceRp      *float64
	orderQtyRq   *float64
	stopPx       *float64
	stopPxEp     *int64
	takeProfit   *float64
	takeProfitEp *int64
	stopLoss     *float64
	stopLossEp   *int64
	pegOffset    *float64
	pegOffsetEp  *int64
}

// Symbol set symbol
func (s *CreateReplaceOrderHedgedService) Symbol(symbol string) *CreateReplaceOrderHedgedService {
	s.symbol = symbol
	return s
}

// PosSide set posSide
func (s *CreateReplaceOrderHedgedService) PosSide(posSide PosSideType) *CreateReplaceOrderHedgedService {
	s.posSide = posSide
	return s
}

// OrderID set orderID
func (s *CreateReplaceOrderHedgedService) OrderID(orderID string) *CreateReplaceOrderHedgedService {
	s.orderID = orderID
	return s
}

// OrigClOrdID set origClOrdID
func (s *CreateReplaceOrderHedgedService) OrigClOrdID(origClOrdID string) *CreateReplaceOrderHedgedService {
	s.origClOrdID = &origClOrdID
	return s
}

// ClOrdID set clOrID
func (s *CreateReplaceOrderHedgedService) ClOrdID(clOrdID string) *CreateReplaceOrderHedgedService {
	s.clOrdID = &clOrdID
	return s
}

// PriceRp set priceRp
func (s *CreateReplaceOrderHedgedService) PriceRp(priceRp float64) *CreateReplaceOrderHedgedService {
	s.priceRp = &priceRp
	return s
}

// OrderQtyRq set orderQtyRq
func (s *CreateReplaceOrderHedgedService) OrderQtyRq(orderQtyRq float64) *CreateReplaceOrderHedgedService {
	s.orderQtyRq = &orderQtyRq
	return s
}

// StopPx set stopPx
func (s *CreateReplaceOrderHedgedService) StopPx(stopPx float64) *CreateReplaceOrderHedgedService {
	s.stopPx = &stopPx
	return s
}

// StopPxEp set stopPxEp
func (s *CreateReplaceOrderHedgedService) StopPxEp(stopPxEp int64) *CreateReplaceOrderHedgedService {
	s.stopPxEp = &stopPxEp
	return s
}

// TakeProfit set takeProfit
func (s *CreateReplaceOrderHedgedService) TakeProfit(takeProfit float64) *CreateReplaceOrderHedgedService {
	s.takeProfit = &takeProfit
	return s
}

// TakeProfitEp set takeProfitEp
func (s *CreateReplaceOrderHedgedService) TakeProfitEp(takeProfitEp int64) *CreateReplaceOrderHedgedService {
	s.takeProfitEp = &takeProfitEp
	return s
}

// StopLoss set stopLoss
func (s *CreateReplaceOrderHedgedService) StopLoss(stopLoss float64) *CreateReplaceOrderHedgedService {
	s.stopLoss = &stopLoss
	return s
}

// StopLossEp set stopLossEp
func (s *CreateReplaceOrderHedgedService) StopLossEp(stopLossEp int64) *CreateReplaceOrderHedgedService {
	s.stopLossEp = &stopLossEp
	return s
}

// PegOffset set pegOffset
func (s *CreateReplaceOrderHedgedService) PegOffset(pegOffset float64) *CreateReplaceOrderHedgedService {
	s.pegOffset = &pegOffset
	return s
}

// PegOffsetEp set pegOffsetEp
func (s *CreateReplaceOrderHedgedService) PegOffsetEp(pegOffsetEp int64) *CreateReplaceOrderHedgedService {
	s.pegOffsetEp = &pegOffsetEp
	return s
}

func (s *CreateReplaceOrderHedgedService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "PUT",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("orderID", s.orderID)
	r.setParam("posSide", s.posSide)

	if s.origClOrdID != nil {
		r.setParam("origClOrdID", *s.origClOrdID)
	}
	if s.clOrdID != nil {
		r.setParam("clOrdID", *s.clOrdID)
	}
	if s.priceRp != nil {
		r.setParam("priceRp", *s.priceRp)
	}
	if s.orderQtyRq != nil {
		r.setParam("orderQtyRq", *s.orderQtyRq)
	}
	if s.stopPx != nil {
		r.setParam("stopPx", *s.stopPx)
	}
	if s.stopPxEp != nil {
		r.setParam("stopPxEp", *s.stopPxEp)
	}
	if s.takeProfit != nil {
		r.setParam("takeProfit", *s.takeProfit)
	}
	if s.takeProfitEp != nil {
		r.setParam("takeProfitEp", *s.takeProfitEp)
	}
	if s.stopLoss != nil {
		r.setParam("stopLoss", *s.stopLoss)
	}
	if s.stopLossEp != nil {
		r.setParam("stopLossEp", *s.stopLossEp)
	}
	if s.pegOffset != nil {
		r.setParam("pegOffset", *s.pegOffset)
	}
	if s.pegOffsetEp != nil {
		r.setParam("pegOffsetEp", *s.pegOffsetEp)
	}
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateReplaceOrderHedgedService) Do(ctx context.Context, opts ...RequestOption) (res *OrderResponse, err error) {
	data, err := s.createOrder(ctx, "/g-orders/replace", opts...)
	if err != nil {
		return nil, err
	}
	resp := new(BaseResponse)
	resp.Data = new(OrderResponse)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, &common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return resp.Data.(*OrderResponse), nil
}

// ListOpenOrdersHedgedService list opened orders
type ListOpenOrdersHedgedService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *ListOpenOrdersHedgedService) Symbol(symbol string) *ListOpenOrdersHedgedService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *ListOpenOrdersHedgedService) Do(ctx context.Context, opts ...RequestOption) (res []*OrderResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/orders/activeList",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*OrderResponse{}, err
	}

	resp := new(BaseResponse)
	resp.Data = new(RowsOrderResponse)

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, &common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	rows := resp.Data.(*RowsOrderResponse)
	return rows.Rows, nil
}

// CancelOrderHedgedService cancel an order
type CancelOrderHedgedService struct {
	c       *Client
	symbol  string
	orderID *string
	posSide PosSideType
}

// Symbol set symbol
func (s *CancelOrderHedgedService) Symbol(symbol string) *CancelOrderHedgedService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelOrderHedgedService) OrderID(orderID string) *CancelOrderHedgedService {
	s.orderID = &orderID
	return s
}

// PosSide set posSide
func (s *CancelOrderHedgedService) PosSide(posSide PosSideType) *CancelOrderHedgedService {
	s.posSide = posSide
	return s
}

// Do send request
func (s *CancelOrderHedgedService) Do(ctx context.Context, opts ...RequestOption) (res *OrderResponse, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/g-orders/cancel",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("posSide", s.posSide)
	if s.orderID != nil {
		r.setParam("orderID", *s.orderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	resp := new(BaseResponse)
	resp.Data = new(OrderResponse)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, &common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return resp.Data.(*OrderResponse), nil
}

// QueryOrderHedgedService cancel an order
type QueryOrderHedgedService struct {
	c         *Client
	symbol    string
	orderID   *string
	clOrderID *string
}

// Symbol set symbol
func (s *QueryOrderHedgedService) Symbol(symbol string) *QueryOrderHedgedService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *QueryOrderHedgedService) OrderID(orderID string) *QueryOrderHedgedService {
	s.orderID = &orderID
	return s
}

// ClOrderID set clOrderID
func (s *QueryOrderHedgedService) ClOrderID(clOrderID string) *QueryOrderHedgedService {
	s.clOrderID = &clOrderID
	return s
}

// Do send request
func (s *QueryOrderHedgedService) Do(ctx context.Context, opts ...RequestOption) (res []*OrderResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/exchange/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderID", *s.orderID)
	}
	if s.clOrderID != nil {
		r.setParam("clOrderID", *s.clOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	resp := new(BaseResponse)
	resp.Data = new([]*OrderResponse)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code > 0 {
		return nil, &common.APIError{
			Code:    resp.Code,
			Message: resp.Msg,
		}
	}
	return *resp.Data.(*[]*OrderResponse), nil
}
