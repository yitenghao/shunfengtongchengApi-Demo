package util

//订单结构
type CreateOrder struct {
	DevID         int    `json:"dev_id"`         //开发者id
	ShopID        string `json:"shop_id"`        //店铺id
	ShopType      int    `json:"shop_type"`      //店铺类型 1：顺丰店铺ID ；2：接入方店铺ID
	ShopOrderID   string `json:"shop_order_id"`  //店铺订单id 不允许重复
	OrderSource   string `json:"order_source"`   //订单接入来源 1：美团；2：饿了么；3：百度；4：口碑； 其他请直接填写中文字符串值
	OrderSequence string `json:"order_sequence"` //取货序号 与order_source配合使用 如：饿了么10号单，表示如下： order_source=2;order_sequence=10。 用于骑士快速寻找配送物
	//LbsType int `json:"lbs_type"`  //坐标类型	1：百度坐标，2：高德坐标
	PayType       int         `json:"pay_type"`       //用户支付方式
	OrderTime     int64       `json:"order_time"`     //用户下单时间
	IsAppoint     int         `json:"is_appoint"`     //是否是预约单
	ExpectTime    int         `json:"expect_time"`    //用户期望送达时间	预约单需必传,秒级时间戳
	IsInsured     int         `json:"is_insured"`     //是否保价，0：非保价；1：保价
	DeclaredValue int         `json:"declared_value"` //保价金额	单位：分
	Remark        string      `json:"remark"`         //订单备注
	ReturnFlag    int         `json:"return_flag"`    //返回字段控制标志位（二进制） 1:商品总价格，2:配送距离，4:物品重量，8:起送时间，16:期望送达时间，32:支付费用，64:实际支持金额，128:优惠卷总金额，256:结算方式 例如全部返回为填入511
	PushTime      int64       `json:"push_time"`      //推单时间	秒级时间戳
	Version       int         `json:"version"`        //版本号	参照文档主版本号填写 如：文档版本号1.7,version=17
	Receive       Receive     `json:"receive"`        //收货人信息
	Shop          Shop        `json:"shop"`           //收货人信息
	OrderDetail   OrderDetail `json:"order_detail"`   //订单详情
}
type Shop struct {
	ShopName    string `json:"shop_name"`    //店铺名称
	PhopPhone   string `json:"shop_phone"`   //店铺电话
	ShopAddress string `json:"shop_address"` //店铺地址
	ShopLng     string `json:"shop_lng"`     //店铺经度
	ShopLat     string `json:"shop_lat"`     //店铺纬度
}

//收货人信息
type Receive struct {
	UserName    string `json:"user_name"`
	UserPhone   string `json:"user_phone"`
	UserAddress string `json:"user_address"`
	UserLng     string `json:"user_lng"`
	UserLat     string `json:"user_lat"`
}

//订单详情
type OrderDetail struct {
	TotalPrice     int             `json:"total_price"`
	ProductType    int             `json:"product_type"`
	WeightGram     int             `json:"weight_gram"`
	ProductNum     int             `json:"product_num"`
	ProductTypeNum int             `json:"product_type_num"`
	ProductDetail  []ProductDetail `json:"product_detail"`
}

//物品详情
type ProductDetail struct {
	ProductName string `json:"product_name"`
	ProductNum  int    `json:"product_num"`
}

//公用响应参数
type PubReturn struct {
	ErrorCode int         `json:"error_code"` //成功为0，失败非0
	ErrorMsg  string      `json:"error_msg"`
	ErrorData interface{} `json:"error_data"`
}
type CreateOrderReturn struct {
	PubReturn
	Result CreateOrderResult `json:"result"`
}

//创建订单结果详情
type CreateOrderResult struct {
	SfOrderID             string `json:"sf_order_id"`             //顺丰订单号
	ShopOrderID           string `json:"shop_order_id"`           //商家订单号
	PushTime              int64  `json:"push_time"`               //推送时间
	OverflowFee           int    `json:"overflow_fee"`            //超出金额
	OverflowExpectTime    int64  `json:"overflow_expect_time"`    //超出时间
	TotalPrice            int    `json:"total_price"`             //配送费总额
	DeliveryDistanceMeter int    `json:"delivery_distance_meter"` //配送距离
	WeightGram            int    `json:"weight_gram"`             //商品重量
	StartTime             int64  `json:"start_time"`              //起送时间
	ExpectTime            int64  `json:"expect_time"`             //期望送达时间
	TotalPayMoney         int    `json:"total_pay_money"`         //支付费用
	RealPayMoney          int    `json:"real_pay_money"`          //实际支付金额
	CouponsTotalFee       int    `json:"coupons_total_fee"`       //优惠券总金额
	SettlementType        int    `json:"settlement_type"`         //结算方式
	SfBillID              string `json:"sf_bill_id"`              //顺丰运单号（需要设置）
	CompleteCode          int    `json:"complete_code"`
	GratuityFee           int    `json:"gratuity_fee"` //小费
}

//预生成订单
type PreCreateOrder struct {
	DevID       int    `json:"dev_id"` //开发者id
	ShopID      string `json:"shop_id"`
	UserAddress string `json:"user_address"`
	Weight      int    `json:"weight"`
	ProductType int    `json:"product_type"`
	IsAppoint   int    `json:"is_appoint"`
	PayType     int    `json:"pay_type"`
	IsInsured   int    `json:"is_insured"`
	ReturnFlag  int    `json:"return_flag"`
	PushTime    int64  `json:"push_time"`
}

//预生成订单响应
type PreCreateOrderReturn struct {
	PubReturn
	Result PreCreateOrderResult `json:"result"`
}

//预生成订单结果详情
type PreCreateOrderResult struct {
	DeliveryType int   `json:"delivery_type"`
	ExpectTime   int64 `json:"expect_time"` //预计送达时间
	//StartTime int64 `json:"start_time"` //预计起送时间
	PromiseDeliveryTime int64 `json:"promise_delivery_time"` //预计配送时间
	//DeliveryDistanceMeter string `json:"delivery_distance_meter"` //配送距离
	EstimatePayMoney   int64                  `json:"estimate_pay_money"`   //预计支付金额
	ChargePriceList    map[string]interface{} `json:"charge_price_list"`    //费用价格列表
	GratuityFee        int                    `json:"gratuity_fee"`         //订单小费
	CouponsTotalFee    int                    `json:"coupons_total_fee"`    //优惠券金额
	OverflowExpectTime int64                  `json:"overflow_expect_time"` //超出时间
	OverflowFee        int                    `json:"overflow_fee"`         //超出金额
	PushTime           int64                  `json:"push_time"`
	RealPayMoney       int                    `json:"real_pay_money"`  //实际支付金额
	SettlementType     int                    `json:"settlement_type"` //结算方式
	TotalPayMoney      int                    `json:"total_pay_money"` //支付费用
	TotalPrice         int                    `json:"total_price"`     //配送费总额
	WeightGram         int                    `json:"weight_gram"`     //商品重量
}

//取消订单
type CancelOrder struct {
	DevID        int    `json:"dev_id"`        //开发者id
	OrderID      string `json:"order_id"`      //订单id
	OrderType    int    `json:"order_type"`    //订单类型 1、顺丰订单号 2、商家订单号
	ShopID       string `json:"shop_id"`       //店铺id 默认0; order_type=2时必传shop_id与shop_type
	ShopType     int    `json:"shop_type"`     //店铺ID类型	1、顺丰店铺ID 2、接入方店铺ID
	CancelReason string `json:"cancel_reason"` //取消原因
	PushTime     int64  `json:"push_time"`     //取消时间
}

//取消订单响应
type CanceOrderReturn struct {
	PubReturn
	Result CancelOrderResult `json:"result"`
}

//取消订单结果
type CancelOrderResult struct {
	SfOrderID   string `json:"sf_order_id"`   //顺丰订单号
	ShopOrderID string `json:"shop_order_id"` //商家订单号
	PushTime    int64  `json:"push_time"`     //推送时间
}

//添加订单小费
type AddOrderGratuityFee struct {
	DevID        int    `json:"dev_id"`        //开发者ID
	OrderID      string `json:"order_id"`      //订单ID
	OrderType    int    `json:"order_type"`    //订单ID类型	1、顺丰订单号 2、商家订单号
	ShopID       string `json:"shop_id"`       //店铺ID
	ShopType     int    `json:"shop_type"`     //店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
	GratuityFee  int    `json:"gratuity_fee"`  //订单小费，单位分，加小费最低不能少于100分
	SerialNumber string `json:"serial_number"` //加小费传入的唯一标识，用来幂等处理
	PushTime     int64  `json:"push_time"`     //取消时间；秒级时间戳
}

//添加订单消费响应
type AddOrderGratuityFeeReturn struct {
	PubReturn
	Result AddOrderGratuityFeeResult `json:"result"`
}

//添加订单消费结果
type AddOrderGratuityFeeResult struct {
	SfOrderID        string `json:"sf_order_id"`        //顺丰订单号
	ShopOrderID      string `json:"shop_order_id"`      //商家订单号
	GratuityFee      int    `json:"gratuity_fee"`       //本次加的小费
	TotalGratuityFee int    `json:"total_gratuity_fee"` //总小费
	PushTime         int64  `json:"push_time"`          //推送时间
}

//获取小费信息
type GetOrderGratuityFee struct {
	DevID     int    `json:"dev_id"`     //开发者ID
	OrderID   string `json:"order_id"`   //订单ID
	OrderType int    `json:"order_type"` //订单ID类型	1、顺丰订单号 2、商家订单号
	ShopID    string `json:"shop_id"`    //店铺ID
	ShopType  int    `json:"shop_type"`  //店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
	PushTime  int64  `json:"push_time"`  //取消时间；秒级时间戳
}

//获取小费响应
type GetOrderGratuityFeeReturn struct {
	PubReturn
	Result GetOrderGratuityFeeResult `json:"result"`
}

//获取小费结果
type GetOrderGratuityFeeResult struct {
	SfOrderID          string `json:"sf_order_id"`          //顺丰订单号
	ShopOrderID        string `json:"shop_order_id"`        //商家订单号
	TotalGratuityFee   int    `json:"total_gratuity_fee"`   //总小费
	TotalGratuityTimes int    `json:"total_gratuity_times"` //加小费总次数
	GratuityFeeList    []struct {
		GratuityFee  int `json:"gratuity_fee"`  //本次加的小费
		GratuityTime int `json:"gratuity_time"` //本次加的小费
	} `json:"gratuity_fee_list"`
}

//订单状态查询
type ListOrderFeed struct {
	DevID     int    `json:"dev_id"`     //开发者ID
	OrderID   string `json:"order_id"`   //订单ID
	OrderType int    `json:"order_type"` //订单ID类型	1、顺丰订单号 2、商家订单号
	ShopID    string `json:"shop_id"`    //店铺ID
	ShopType  int    `json:"shop_type"`  //店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
	PushTime  int64  `json:"push_time"`  //取消时间；秒级时间戳
}

//订单状态查询响应
type ListOrderFeedReturn struct {
	PubReturn
	Result ListOrderFeedResult `json:"result"`
}
type ListOrderFeedResult struct {
	SfOrderID   string `json:"sf_order_id"`   //顺丰订单号
	ShopOrderID string `json:"shop_order_id"` //商家订单号
	CreateTime  int64  `json:"create_time"`   //创建时间
	PushTime    int64  `json:"push_time"`     //推送时间
	Feed        []struct {
		SfOrderID     string `json:"sf_order_id"`    //顺丰订单号
		ShopOrderID   string `json:"shop_order_id"`  //商家订单号
		OrderStatus   int    `json:"order_status"`   //订单状态
		Operator      string `json:"operator"`       //骑手
		OperatorName  string `json:"operator_name"`  //骑手名
		OperatorPhone string `json:"operator_phone"` //骑手电话
		StatusDesc    string `json:"status_desc"`    //状态说明
		Content       string `json:"content"`        //正文
		CreateTime    string `json:"create_time"`    //记录时间
	} `json:"feed"`
}

//催单
type ReminderOrder struct {
	DevID     int    `json:"dev_id"`     //开发者ID
	OrderID   string `json:"order_id"`   //订单ID
	OrderType int    `json:"order_type"` //订单ID类型	1、顺丰订单号 2、商家订单号
	ShopID    string `json:"shop_id"`    //店铺ID
	ShopType  int    `json:"shop_type"`  //店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
	PushTime  int64  `json:"push_time"`  //取消时间；秒级时间戳
}
type ReminderOrderReturn struct {
	PubReturn
	Result ReminderOrderResult `json:"result"`
}
type ReminderOrderResult struct {
	SfOrderID   string `json:"sf_order_id"`   //顺丰订单号
	ShopOrderID string `json:"shop_order_id"` //商家订单号
	PushTime    int64  `json:"push_time"`     //推送时间
}

//修改收货人信息
type ChangeOrder struct {
	DevID       int    `json:"dev_id"`       //开发者ID
	OrderID     string `json:"order_id"`     //订单ID
	OrderType   int    `json:"order_type"`   //订单ID类型	1、顺丰订单号 2、商家订单号
	ShopID      string `json:"shop_id"`      //店铺ID
	ShopType    int    `json:"shop_type"`    //店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
	UserName    string `json:"user_name"`    //用户姓名
	UserPhone   string `json:"user_phone"`   //用户电话
	UserAddress string `json:"user_address"` //用户地址
	LbsType     int    `json:"lbs_type"`     //坐标类型	1：百度坐标，2：高德坐标（默认值为2，下面的经纬度依赖这个坐标系，不传默认高德）
	UserLng     string `json:"user_lng"`     //用户地址经度	传入用户地址经纬度顺丰侧则不根据用户地址解析
	UserLat     string `json:"user_lat"`     //用户地址纬度
	PushTime    int64  `json:"push_time"`    //取消时间；秒级时间戳
}
type ChangeOrderReturn struct {
	PubReturn
	Result ChangeOrderResult `json:"result"`
}
type ChangeOrderResult struct {
	SfOrderID   string `json:"sf_order_id"`   //顺丰订单号
	ShopOrderID string `json:"shop_order_id"` //商家订单号
	PushTime    int64  `json:"push_time"`     //推送时间
}

//获取骑手最新位置
type RiderLatestPosition struct {
	DevID     int    `json:"dev_id"`     //开发者ID
	OrderID   string `json:"order_id"`   //订单ID
	OrderType int    `json:"order_type"` //订单ID类型	1、顺丰订单号 2、商家订单号
	ShopID    string `json:"shop_id"`    //店铺ID
	ShopType  int    `json:"shop_type"`  //店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
	PushTime  int64  `json:"push_time"`  //取消时间；秒级时间戳
}

type RiderLatestPositionReturn struct {
	PubReturn
	Result RiderLatestPositionResult `json:"result"`
}
type RiderLatestPositionResult struct {
	SfOrderID   string `json:"sf_order_id"`   //顺丰订单号
	ShopOrderID string `json:"shop_order_id"` //商家订单号
	RiderName   string `json:"rider_name"`    //骑手姓名
	RiderPhone  string `json:"rider_phone"`   //骑手电话
	RiderLng    string `json:"rider_lng"`     //经度
	RiderLat    string `json:"rider_lat"`     //纬度
	UploadTime  string `json:"upload_time"`   //更新时间
}

//获取配送员轨迹H5
type RiderViewV2 struct {
	DevID     int    `json:"dev_id"`     //开发者ID
	OrderID   string `json:"order_id"`   //订单ID
	OrderType int    `json:"order_type"` //订单ID类型	1、顺丰订单号 2、商家订单号
	ShopID    string `json:"shop_id"`    //店铺ID
	ShopType  int    `json:"shop_type"`  //店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
	PushTime  int64  `json:"push_time"`  //取消时间；秒级时间戳
}
type RiderViewV2Return struct {
	PubReturn
	Result RiderViewV2Result `json:"result"`
}
type RiderViewV2Result struct {
	SfOrderID   string `json:"sf_order_id"`   //顺丰订单号
	ShopOrderID string `json:"shop_order_id"` //商家订单号
	Url         string `json:"url"`           //H5页面URL
}

//订单回调详情
type GetCallbackInfo struct {
	DevID     int    `json:"dev_id"`     //开发者ID
	OrderID   string `json:"order_id"`   //订单ID
	OrderType int    `json:"order_type"` //订单ID类型	1、顺丰订单号 2、商家订单号
	ShopID    string `json:"shop_id"`    //店铺ID
	ShopType  int    `json:"shop_type"`  //店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
	PushTime  int64  `json:"push_time"`  //取消时间；秒级时间戳
}
type GetCallbackInfoRetuen struct {
	PubReturn
	Result []GetCallbackInfoResult `json:"result"`
}
type GetCallbackInfoResult struct {
	ShopId        string `json:"shop_id"`        // 店铺id
	SfOrderId     string `json:"sf_order_id"`    // 顺丰订单id
	ShopOrderId   string `json:"shop_order_id"`  // 店铺订单id、
	UrlIndex      string `json:"url_index"`      // 骑手状态
	OperatorName  string `json:"operator_name"`  // 骑手姓名
	OperatorPhone string `json:"operator_phone"` //骑手手机
	RiderLng      string `json:"rider_lng"`      // 经度
	RiderLat      string `json:"rider_lat"`      // 纬度
	OrderStatus   int    `json:"order_status"`   // 订单状态
	StatusDesc    string `json:"status_desc"`    // 状态详情
	PushTime      int64  `json:"push_time"`      // 时间
}

//订单状态改变回调
type StatusChangeCallBack struct {
	SfOrderId     string `json:"sf_order_id"`    //顺丰订单ID
	ShopOrderIid  string `json:"shop_order_id"`  //商家订单ID
	UrlIndex      string `json:"url_index"`      //rider_status
	OperatorName  string `json:"operator_name"`  //配送员姓名
	OperatorPhone string `json:"operator_phone"` //配送员电话
	RiderLng      string `json:"rider_lng"`      //配送员位置经度
	RiderLat      string `json:"rider_lat"`      //配送员位置纬度
	OrderStatus   int    `json:"order_status"`   //订单状态	10-配送员确认;12:配送员到店;15:配送员配送中
	StatusDesc    string `json:"status_desc"`    //状态描述
	PushTime      int64  `json:"push_time"`      //状态变更时间
}

//订单完成回调
type OrderOkCallBack struct {
	SfOrderId    string `json:"sf_order_id"`   //顺丰订单ID
	ShopOrderIid string `json:"shop_order_id"` //商家订单ID
	UrlIndex     string `json:"url_index"`     //rider_status
	OperatorName string `json:"operator_name"` //配送员姓名
	RiderLng     string `json:"rider_lng"`     //配送员位置经度
	RiderLat     string `json:"rider_lat"`     //配送员位置纬度
	OrderStatus  int    `json:"order_status"`  //订单状态	10-配送员确认;12:配送员到店;15:配送员配送中
	StatusDesc   string `json:"status_desc"`   //状态描述
	PushTime     int64  `json:"push_time"`     //状态变更时间
}

//订单取消回调
type OrderCancelCallBack struct {
	SfOrderId    string `json:"sf_order_id"`   //顺丰订单ID
	ShopOrderIid string `json:"shop_order_id"` //商家订单ID
	UrlIndex     string `json:"url_index"`     //rider_status
	OperatorName string `json:"operator_name"` //配送员姓名
	OrderStatus  int    `json:"order_status"`  //订单状态	10-配送员确认;12:配送员到店;15:配送员配送中
	StatusDesc   string `json:"status_desc"`   //状态描述
	CancelReason string `json:"cancel_reason"` //状态描述
	PushTime     int64  `json:"push_time"`     //状态变更时间
}

//订单异常
//4003:托寄物丢失或损坏
//1001:商家出货慢
//2010:顾客拒绝实名认证
//3004:实名认证校验失败
//1007:更改取货地址
//2001:顾客电话无法接通
//2004:更改期望送达时间
//2005:顾客拒收
//2008:顾客不在家
//2009:更改送货地址
//4001:配送地址错误
//4002:其他
type OrderErrorCallBack struct {
	SfOrderId     string `json:"sf_order_id"`    //顺丰订单ID
	ShopOrderIid  string `json:"shop_order_id"`  //商家订单ID
	UrlIndex      string `json:"url_index"`      //rider_status
	OperatorName  string `json:"operator_name"`  //配送员姓名
	OperatorPhone string `json:"operator_phone"` //配送员电话
	ExId          int    `json:"ex_id"`          //异常id
	ExContent     int    `json:"ex_content"`     //异常详情
	ExpectTime    int    `json:"expect_time"`    //新的期望送达时间	如果发生期望送达时间的更新此字段有值
	PushTime      int64  `json:"push_time"`      //状态变更时间
}
