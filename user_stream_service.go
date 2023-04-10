package phemex

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var (
	baseURL = "wss://phemex.com/ws"
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = 5 * time.Second
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
)

// WsAuthService create listen key for user stream service
type WsAuthService struct {
	c   *Client
	url *string
}

// URL set url. wss://testnet.phemex.com/ws for test net
func (s *WsAuthService) URL(url string) *WsAuthService {
	s.url = &url
	return s
}

// Do send request
func (s *WsAuthService) Do(ctx context.Context, opts ...RequestOption) (c *websocket.Conn, err error) {
	if s.url != nil {
		baseURL = *s.url
	}
	s.c.debug("dial URL: %s", baseURL)

	c, _, err = websocket.DefaultDialer.Dial(baseURL, nil)
	//c.SetReadLimit()
	if err != nil {
		return nil, err
	}

	expiry := currentTimestamp() + 60
	raw := fmt.Sprintf("%s%v", s.c.APIKey, expiry)
	signedString, err := s.c.signString(raw)
	if err != nil {
		return nil, err
	}

	err = c.WriteJSON(map[string]interface{}{
		"method": "user.auth",
		"params": []interface{}{
			"API",
			s.c.APIKey,
			signedString,
			expiry,
		},
		"id": 100,
	})

	if err != nil {
		return nil, err
	}

	_, _, err = c.ReadMessage()

	if err != nil {
		return nil, err
	}

	return c, nil
}

// StartWsAOPService create listen key for user stream service
type StartWsAOPService struct {
	c  *Client
	id *int64
}

// SetID set id
func (s *StartWsAOPService) SetID(id int64) *StartWsAOPService {
	s.id = &id
	return s
}

// Do send request
func (s *StartWsAOPService) Do(c *websocket.Conn, handler WsHandler, errHandler ErrHandler, opts ...RequestOption) (err error) {

	stop := make(chan struct{})

	if c == nil {
		return fmt.Errorf("the connection is not initialized (%v)", *s.id)
	}
	go func() {
		defer close(stop)

		if WebsocketKeepalive {
			keepAlive(c, *s.id, stop, errHandler)
		}

		err = c.WriteJSON(map[string]interface{}{
			"id":     *s.id,
			"method": "aop.subscribe",
			"params": []interface{}{},
		})

		if err != nil {
			errHandler(err)
			return
		}

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				errHandler(err)
				return
			}
			var resp interface{}
			switch {
			case strings.HasPrefix(string(msg), "{\"error\""):
				resp = new(WsError)
			case strings.HasPrefix(string(msg), "{\"position_info\""):
				resp = new(WsPositionInfo)
			default:
				resp = new(WsAOP)
			}
			err = json.Unmarshal(msg, &resp)

			if err != nil {
				errHandler(err)
				return
			}
			handler(resp)
		}
	}()
	return
}

// WsAccount ws account
type WsAccount struct {
	AccountBalanceEv   int64  `json:"accountBalanceEv"`
	AccountID          int64  `json:"accountID"`
	BonusBalanceEv     int64  `json:"bonusBalanceEv"`
	Currency           string `json:"currency"`
	TotalUsedBalanceEv int64  `json:"totalUsedBalanceEv"`
	UserID             int64  `json:"userID"`
}

// WsOrder ws order
type WsOrder struct {
	AccountID               int64       `json:"accountID"`
	Action                  string      `json:"action"`
	ActionBy                string      `json:"actionBy"`
	ActionTimeNs            int64       `json:"actionTimeNs"`
	AddedSeq                int64       `json:"addedSeq"`
	BonusChangedAmountEv    int64       `json:"bonusChangedAmountEv"`
	ClOrdID                 string      `json:"clOrdID"`
	ClosedPnlEv             int64       `json:"closedPnlEv"`
	ClosedSize              float64     `json:"closedSize"`
	Code                    int64       `json:"code"`
	CumQty                  float64     `json:"cumQty"`
	CurAccBalanceEv         int64       `json:"curAccBalanceEv"`
	CurAssignedPosBalanceEv int64       `json:"curAssignedPosBalanceEv"`
	CurLeverageEr           int64       `json:"curLeverageEr"`
	CurPosSide              string      `json:"curPosSide"`
	CurPosSize              float64     `json:"curPosSize"`
	CurPosTerm              int64       `json:"curPosTerm"`
	CurPosValueEv           int64       `json:"curPosValueEv"`
	CurRiskLimitEv          int64       `json:"curRiskLimitEv"`
	Currency                string      `json:"currency"`
	CxlRejReason            int64       `json:"cxlRejReason"`
	DisplayQty              float64     `json:"displayQty"`
	ExecFeeEv               int64       `json:"execFeeEv"`
	ExecID                  string      `json:"execID"`
	ExecInst                string      `json:"execInst"`
	ExecPriceEp             int64       `json:"execPriceEp"`
	ExecQty                 float64     `json:"execQty"`
	ExecSeq                 float64     `json:"execSeq"`
	ExecStatus              string      `json:"execStatus"`
	ExecValueEv             int64       `json:"execValueEv"`
	FeeRateEr               int64       `json:"feeRateEr"`
	LeavesQty               float64     `json:"leavesQty"`
	LeavesValueEv           int64       `json:"leavesValueEv"`
	Message                 string      `json:"message"`
	OrdStatus               string      `json:"ordStatus"`
	OrdType                 string      `json:"ordType"`
	OrderID                 string      `json:"orderID"`
	OrderQty                float64     `json:"orderQty"`
	PegOffsetValueEp        int64       `json:"pegOffsetValueEp"`
	Platform                string      `json:"platform"`
	PriceEp                 int64       `json:"priceEp"`
	RelatedPosTerm          int64       `json:"relatedPosTerm"`
	RelatedReqNum           int64       `json:"relatedReqNum"`
	Side                    string      `json:"side"`
	StopDirection           string      `json:"stopDirection"`
	StopLossEp              int64       `json:"stopLossEp"`
	StopPxEp                int64       `json:"stopPxEp"`
	Symbol                  string      `json:"symbol"`
	TakeProfitEp            int64       `json:"takeProfitEp"`
	TimeInForce             string      `json:"timeInForce"`
	TransactTimeNs          int64       `json:"transactTimeNs"`
	Trigger                 TriggerType `json:"trigger"`
	UserID                  int64       `json:"userID"`
}

// WsPosition ws position
type WsPosition struct {
	AccountID              int64   `json:"accountID"`
	ActionTimeNs           int64   `json:"actionTimeNs"`
	AssignedPosBalanceEv   int64   `json:"assignedPosBalanceEv"`
	AvgEntryPriceEp        int64   `json:"avgEntryPriceEp"`
	BankruptCommEv         int64   `json:"bankruptCommEv"`
	BankruptPriceEp        int64   `json:"bankruptPriceEp"`
	BuyLeavesQty           float64 `json:"buyLeavesQty"`
	BuyLeavesValueEv       int64   `json:"buyLeavesValueEv"`
	BuyValueToCostEr       int64   `json:"buyValueToCostEr"`
	CreatedAtNs            int64   `json:"createdAtNs"`
	CrossSharedBalanceEv   int64   `json:"crossSharedBalanceEv"`
	CumClosedPnlEv         int64   `json:"cumClosedPnlEv"`
	CumFundingFeeEv        int64   `json:"cumFundingFeeEv"`
	CumTransactFeeEv       int64   `json:"cumTransactFeeEv"`
	CurTermRealisedPnlEv   int64   `json:"curTermRealisedPnlEv"`
	Currency               string  `json:"currency"`
	DataVer                float64 `json:"dataVer"`
	DeleveragePercentileEr int64   `json:"deleveragePercentileEr"`
	DisplayLeverageEr      int64   `json:"displayLeverageEr"`
	EstimatedOrdLossEv     int64   `json:"estimatedOrdLossEv"`
	ExecSeq                int64   `json:"execSeq"`
	FreeCostEv             int64   `json:"freeCostEv"`
	FreeQty                float64 `json:"freeQty"`
	InitMarginReqEr        int64   `json:"initMarginReqEr"`
	LastFundingTime        int64   `json:"lastFundingTime"`
	LastTermEndTime        int64   `json:"lastTermEndTime"`
	LeverageEr             int64   `json:"leverageEr"`
	LiquidationPriceEp     int64   `json:"liquidationPriceEp"`
	MaintMarginReqEr       int64   `json:"maintMarginReqEr"`
	MakerFeeRateEr         int64   `json:"makerFeeRateEr"`
	MarkPriceEp            int64   `json:"markPriceEp"`
	OrderCostEv            int64   `json:"orderCostEv"`
	PosCostEv              int64   `json:"posCostEv"`
	PositionMarginEv       int64   `json:"positionMarginEv"`
	PositionStatus         string  `json:"positionStatus"`
	RiskLimitEv            int64   `json:"riskLimitEv"`
	SellLeavesQty          float64 `json:"sellLeavesQty"`
	SellLeavesValueEv      int64   `json:"sellLeavesValueEv"`
	SellValueToCostEr      int64   `json:"sellValueToCostEr"`
	Side                   string  `json:"side"`
	Size                   float64 `json:"size"`
	Symbol                 string  `json:"symbol"`
	TakerFeeRateEr         int64   `json:"takerFeeRateEr"`
	Term                   int64   `json:"term"`
	UnrealisedPnlEv        int64   `json:"unrealisedPnlEv"`
	UpdatedAtNs            int64   `json:"updatedAtNs"`
	UsedBalanceEv          int64   `json:"usedBalanceEv"`
	UserID                 int64   `json:"userID"`
	ValueEv                int64   `json:"valueEv"`
}

// Error ws error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Status status
type Status struct {
	Status string `json:"status"`
}

// WsAOP ws AOP
type WsAOP struct {
	Accounts  []*WsAccount  `json:"accounts"`
	Orders    []*WsOrder    `json:"orders"`
	Positions []*WsPosition `json:"positions"`
	Sequence  int64         `json:"sequence"`
	Timestamp int64         `json:"timestamp"`
	Type      string        `json:"type"`
}

// WsError ws error
type WsError struct {
	Error  *Error  `json:"error"`
	Result *Status `json:"result"`
	ID     int     `json:"id"`
}

// PositionInfo position info
type PositionInfo struct {
	AccountID int64   `json:"accountID"`
	Light     float64 `json:"light"`
	Symbol    string  `json:"symbol"`
	UserID    int64   `json:"userID"`
}

// WsPositionInfo ws position info
type WsPositionInfo struct {
	PositionInfo *PositionInfo `json:"position_info"`
	Sequence     int64         `json:"sequence"`
}

type WsAOPHedged struct {
	Accounts  []*WsAccountHedged  `json:"accounts_p"`
	Orders    []*WsOrderHedged    `json:"orders_p"`
	Positions []*WsPositionHedged `json:"positions_p"`
	Sequence  int64               `json:"sequence"`
	Timestamp int64               `json:"timestamp"`
	Type      string              `json:"type"`
}

type WsAccountHedged struct {
	AccountBalanceRv   string `json:"accountBalanceRv"`
	AccountID          int64  `json:"accountID"`
	BonusBalanceRv     string `json:"bonusBalanceRv"`
	Currency           string `json:"currency"`
	TotalUsedBalanceRv string `json:"totalUsedBalanceRv"`
	UserID             int    `json:"userID"`
}

type WsOrderHedged struct {
	AccountID               int64  `json:"accountID"`
	Action                  string `json:"action"`
	ActionBy                string `json:"actionBy"`
	ActionTimeNs            int64  `json:"actionTimeNs"`
	AddedSeq                int    `json:"addedSeq"`
	ApRp                    string `json:"apRp"`
	BonusChangedAmountRv    string `json:"bonusChangedAmountRv"`
	BpRp                    string `json:"bpRp"`
	ClOrdID                 string `json:"clOrdID"`
	ClosedPnlRv             string `json:"closedPnlRv"`
	ClosedSize              string `json:"closedSize"`
	Code                    int    `json:"code"`
	CumFeeRv                string `json:"cumFeeRv"`
	CumQty                  string `json:"cumQty"`
	CumValueRv              string `json:"cumValueRv"`
	CurAccBalanceRv         string `json:"curAccBalanceRv"`
	CurAssignedPosBalanceRv string `json:"curAssignedPosBalanceRv"`
	CurBonusBalanceRv       string `json:"curBonusBalanceRv"`
	CurLeverageRr           string `json:"curLeverageRr"`
	CurPosSide              string `json:"curPosSide"`
	CurPosSize              string `json:"curPosSize"`
	CurPosTerm              int    `json:"curPosTerm"`
	CurPosValueRv           string `json:"curPosValueRv"`
	CurRiskLimitRv          string `json:"curRiskLimitRv"`
	Currency                string `json:"currency"`
	CxlRejReason            int    `json:"cxlRejReason"`
	DisplayQty              string `json:"displayQty"`
	ExecFeeRv               string `json:"execFeeRv"`
	ExecID                  string `json:"execID"`
	ExecPriceRp             string `json:"execPriceRp"`
	ExecQty                 string `json:"execQty"`
	ExecSeq                 int    `json:"execSeq"`
	ExecStatus              string `json:"execStatus"`
	ExecValueRv             string `json:"execValueRv"`
	FeeRateRr               string `json:"feeRateRr"`
	LeavesQty               string `json:"leavesQty"`
	LeavesValueRv           string `json:"leavesValueRv"`
	Message                 string `json:"message"`
	OrdStatus               string `json:"ordStatus"`
	OrdType                 string `json:"ordType"`
	OrderID                 string `json:"orderID"`
	OrderQty                string `json:"orderQty"`
	PegOffsetValueRp        string `json:"pegOffsetValueRp"`
	PosSide                 string `json:"posSide"`
	PriceRp                 string `json:"priceRp"`
	RelatedPosTerm          int    `json:"relatedPosTerm"`
	RelatedReqNum           int    `json:"relatedReqNum"`
	Side                    string `json:"side"`
	SlTrigger               string `json:"slTrigger"`
	StopLossRp              string `json:"stopLossRp"`
	StopPxRp                string `json:"stopPxRp"`
	Symbol                  string `json:"symbol"`
	TakeProfitRp            string `json:"takeProfitRp"`
	TimeInForce             string `json:"timeInForce"`
	TpTrigger               string `json:"tpTrigger"`
	TradeType               string `json:"tradeType"`
	TransactTimeNs          int64  `json:"transactTimeNs"`
	UserID                  int    `json:"userID"`
}

type WsPositionHedged struct {
	AccountID              int64  `json:"accountID"`
	AssignedPosBalanceRv   string `json:"assignedPosBalanceRv"`
	AvgEntryPriceRp        string `json:"avgEntryPriceRp"`
	BankruptCommRv         string `json:"bankruptCommRv"`
	BankruptPriceRp        string `json:"bankruptPriceRp"`
	BuyLeavesQty           string `json:"buyLeavesQty"`
	BuyLeavesValueRv       string `json:"buyLeavesValueRv"`
	BuyValueToCostRr       string `json:"buyValueToCostRr"`
	CreatedAtNs            int    `json:"createdAtNs"`
	CrossSharedBalanceRv   string `json:"crossSharedBalanceRv"`
	CumClosedPnlRv         string `json:"cumClosedPnlRv"`
	CumFundingFeeRv        string `json:"cumFundingFeeRv"`
	CumTransactFeeRv       string `json:"cumTransactFeeRv"`
	CurTermRealisedPnlRv   string `json:"curTermRealisedPnlRv"`
	Currency               string `json:"currency"`
	DataVer                int    `json:"dataVer"`
	DeleveragePercentileRr string `json:"deleveragePercentileRr"`
	DisplayLeverageRr      string `json:"displayLeverageRr"`
	EstimatedOrdLossRv     string `json:"estimatedOrdLossRv"`
	ExecSeq                int    `json:"execSeq"`
	FreeCostRv             string `json:"freeCostRv"`
	FreeQty                string `json:"freeQty"`
	InitMarginReqRr        string `json:"initMarginReqRr"`
	LastFundingTime        int64  `json:"lastFundingTime"`
	LastTermEndTime        int    `json:"lastTermEndTime"`
	LeverageRr             string `json:"leverageRr"`
	LiquidationPriceRp     string `json:"liquidationPriceRp"`
	MaintMarginReqRr       string `json:"maintMarginReqRr"`
	MakerFeeRateRr         string `json:"makerFeeRateRr"`
	MarkPriceRp            string `json:"markPriceRp"`
	MinPosCostRv           string `json:"minPosCostRv"`
	OrderCostRv            string `json:"orderCostRv"`
	PosCostRv              string `json:"posCostRv"`
	PosMode                string `json:"posMode"`
	PosSide                string `json:"posSide"`
	PositionMarginRv       string `json:"positionMarginRv"`
	PositionStatus         string `json:"positionStatus"`
	RiskLimitRv            string `json:"riskLimitRv"`
	SellLeavesQty          string `json:"sellLeavesQty"`
	SellLeavesValueRv      string `json:"sellLeavesValueRv"`
	SellValueToCostRr      string `json:"sellValueToCostRr"`
	Side                   string `json:"side"`
	Size                   string `json:"size"`
	Symbol                 string `json:"symbol"`
	TakerFeeRateRr         string `json:"takerFeeRateRr"`
	Term                   int    `json:"term"`
	TransactTimeNs         int64  `json:"transactTimeNs"`
	UnrealisedPnlRv        string `json:"unrealisedPnlRv"`
	UpdatedAtNs            int    `json:"updatedAtNs"`
	UsedBalanceRv          string `json:"usedBalanceRv"`
	UserID                 int    `json:"userID"`
	ValueRv                string `json:"valueRv"`
}
