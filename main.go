package main

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"shunfengtongchengApi-Demo/util"
	"time"
)

func main() {
	go func(){
		//testcreateorder()
	//testprecreateorder()
	//testCancelOrder()
	//testAddOrderGratuityFee()
	//testGetOrderGratuityFee()
	//testListOrderFeed()
	//testReminderOrder()
	//testChangeOrder()
	//testRiderLatestPosition()
	//testRiderViewV2()
	testGetCallbackInfo()
	}()

	{
		http.HandleFunc("/change", chenge)
		http.HandleFunc("/ok", ok)
		http.HandleFunc("/cancel", cancel)
		http.HandleFunc("/error", error)
		err := http.ListenAndServe(":9999", nil)
		if err != nil {
			fmt.Println(err)
		}
	}

}

//注意  回调请求中，如果往响应中写入东西，顺丰就不会再次回调了，默认回调成功，如果不往响应中写入数据，顺丰会继续回调，达到20次就不会再回调了
//回调仅仅是通知，你的响应不会改变订单状态
//状态异常时请先调用回调信息接口GetCallbackInfo

func chenge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("change")
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	str, _ := util.Sign(bts)
	fmt.Println(r.URL.Query()["sign"][0] == str)
	if !(r.URL.Query()["sign"][0] == str) {
		//w.Write([]byte(`{"error_code": 0,"error_msg": "fail"}`))
		return
	}
	statuschange := util.StatusChangeCallBack{}
	json.Unmarshal(bts, &statuschange)
	fmt.Println(statuschange)
	//校验订单  查数据库，校验id等信息，错误就返回失败
	if statuschange.SfOrderId != "3337055525789311010" || statuschange.UrlIndex != "rider_status" {
		//w.Write([]byte(`{"error_code": 0,"error_msg": "fail"}`))
		fmt.Println("业务校验失败")
		return
	}
	//更新数据库订单状态

	w.Write([]byte(`{"error_code": 0,"error_msg": "success"}`))
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok")
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	str, _ := util.Sign(bts)
	fmt.Println(r.URL.Query()["sign"][0] == str)
	if !(r.URL.Query()["sign"][0] == str) {
		w.Write([]byte(`{"error_code": 0,"error_msg": "fail"}`))
		return
	}
	orderok := util.OrderOkCallBack{}
	json.Unmarshal(bts, &orderok)
	fmt.Println(orderok)
	//校验订单  查数据库，校验id等信息，错误就返回失败
	if orderok.SfOrderId != "3337055525789311010" || orderok.UrlIndex != "order_complete" {
		//w.Write([]byte(`{"error_code": 0,"error_msg": "fail"}`))
		return
	}
	//更新数据库订单状态

	w.Write([]byte(`{"error_code": 0,"error_msg": "success"}`))
}
func cancel(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cancel")
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	str, _ := util.Sign(bts)
	fmt.Println(r.URL.Query()["sign"][0] == str)
	if !(r.URL.Query()["sign"][0] == str) {
		w.Write([]byte(`{"error_code": 0,"error_msg": "fail"}`))
		return
	}
	orderok := util.OrderCancelCallBack{}
	json.Unmarshal(bts, &orderok)
	fmt.Println(orderok)
	//校验订单  查数据库，校验id等信息，错误就返回失败
	if orderok.SfOrderId != "3336950196759752729" || orderok.UrlIndex != "sf_cancel " {
		w.Write([]byte(`{"error_code": 0,"error_msg": "fail"}`))
		return
	}
	//记录订单取消原因并更新订单状态
	w.Write([]byte(`{"error_code": 0,"error_msg": "success"}`))
}
func error(w http.ResponseWriter, r *http.Request) {
	fmt.Println("error")
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	str, _ := util.Sign(bts)
	fmt.Println(r.URL.Query()["sign"][0] == str)
	if !(r.URL.Query()["sign"][0] == str) {
		w.Write([]byte(`{"error_code": 0,"error_msg": "fail"}`))
		return
	}
	orderok := util.OrderErrorCallBack{}
	json.Unmarshal(bts, &orderok)
	fmt.Println(orderok)
	//校验订单  查数据库，校验id等信息，错误就返回失败
	if orderok.SfOrderId != "3336950196759752729" || orderok.UrlIndex != "rider_exception" {
		w.Write([]byte(`{"error_code": 0,"error_msg": "fail"}`))
		return
	}
	//更新数据库订单状态
	w.Write([]byte(`{"error_code": 0,"error_msg": "success"}`))
}
func testcreateorder() {
	uid, _ := uuid.NewV4()
	order := util.CreateOrder{
		DevID:       util.DevID,
		ShopID:      util.ShopID,
		ShopType:    1,
		ShopOrderID: uid.String(),
		OrderSource: "1",
		PayType:     1,
		OrderTime:   time.Now().Unix(),
		IsAppoint:   0,
		IsInsured:   0,
		ReturnFlag:  511,
		PushTime:    time.Now().Unix(),
		Version:     17,
		Receive: util.Receive{
			UserName:    "yth",
			UserPhone:   "18571625199",
			UserAddress: "北京市海淀区学清嘉创大厦A座15层",
			UserLng:     "116.334424",
			UserLat:     "40.030177",
		},
		OrderDetail: util.OrderDetail{
			TotalPrice:     10000,
			ProductType:    1,
			WeightGram:     500,
			ProductNum:     1,
			ProductTypeNum: 1,
			ProductDetail: []util.ProductDetail{
				util.ProductDetail{
					ProductName: "黄焖鸡",
					ProductNum:  1,
				},
			},
		},
	}
	//sign,err:= util.Sign(order)
	//fmt.Println(sign,err)
	bts, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.CreateOrderUrl)
	if err != nil {
		fmt.Println(err)
	}
	returnparam := util.CreateOrderReturn{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
	//fmt.Println(string(body),err)
}
func testprecreateorder() {
	precreateorder := util.PreCreateOrder{
		DevID:       util.DevID,
		ShopID:      util.ShopID,
		UserAddress: "北京市海淀区学清嘉创大厦A座15层",
		Weight:      1000,
		ProductType: 1,
		IsAppoint:   0,
		PayType:     1,
		IsInsured:   0,
		ReturnFlag:  511,
		PushTime:    time.Now().Unix(),
	}
	bts, err := json.Marshal(precreateorder)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.PreCreateOrderUrl)
	if err != nil {
		fmt.Println(err)
	}
	returnparam := util.PreCreateOrderReturn{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}
func testCancelOrder() {
	cancelorder := util.CancelOrder{
		DevID:        util.DevID,
		OrderID:      "3337110344144271400",
		OrderType:    1, //这里选2则shopid必填
		ShopID:       "0",
		ShopType:     1,
		CancelReason: "不想要了",
		PushTime:     time.Now().Unix(),
	}
	bts, err := json.Marshal(cancelorder)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.CancelOrderUrl)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	returnparam := util.CanceOrderReturn{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}

func testAddOrderGratuityFee() {
	uuid, _ := uuid.NewV4()
	param := util.AddOrderGratuityFee{
		DevID:        util.DevID,
		OrderID:      "3336911806725256196",
		OrderType:    1,
		ShopID:       util.ShopID,
		ShopType:     1,
		GratuityFee:  1000,
		SerialNumber: uuid.String(),
		PushTime:     time.Now().Unix(),
	}
	bts, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.AddOrderGratuityFeeUrl)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	returnparam := util.AddOrderGratuityFeeReturn{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}
func testGetOrderGratuityFee() {
	param := util.GetOrderGratuityFee{
		DevID:     util.DevID,
		OrderID:   "3336911806725256196",
		OrderType: 1,
		ShopID:    util.ShopID,
		ShopType:  1,
		PushTime:  time.Now().Unix(),
	}
	bts, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.GetOrderGratuityFeeUrl)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	returnparam := util.GetOrderGratuityFeeReturn{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}
func testListOrderFeed() {
	param := util.ListOrderFeed{
		DevID:     util.DevID,
		OrderID:   "3337110344144271400",
		OrderType: 1,
		ShopID:    util.ShopID,
		ShopType:  1,
		PushTime:  time.Now().Unix(),
	}
	bts, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.ListOrderFeedUrl)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	returnparam := util.ListOrderFeedReturn{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}
func testReminderOrder() {
	param := util.ReminderOrder{
		DevID:     util.DevID,
		OrderID:   "3336925940269489168",
		OrderType: 1,
		ShopID:    util.ShopID,
		ShopType:  1,
		PushTime:  time.Now().Unix(),
	}
	bts, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.ReminderOrderUrl)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	returnparam := util.ReminderOrderReturn{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}
func testChangeOrder() {
	param := util.ChangeOrder{
		DevID:       util.DevID,
		OrderID:     "3336925940269489168",
		OrderType:   1,
		ShopID:      util.ShopID,
		ShopType:    1,
		UserName:    "yitenghao",
		UserPhone:   "13888888888",
		UserAddress: "",
		LbsType:     2,
		UserLng:     "116.334424",
		UserLat:     "40.030177",
		PushTime:    time.Now().Unix(),
	}
	bts, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.ChangeOrderUrl)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	returnparam := util.ChangeOrderReturn{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}
func testRiderLatestPosition() {
	param := util.RiderLatestPosition{
		DevID:     util.DevID,
		OrderID:   "3336932261919744021",
		OrderType: 1,
		ShopID:    util.ShopID,
		ShopType:  1,
		PushTime:  time.Now().Unix(),
	}
	bts, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.RiderLatestPositionUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	returnparam := util.RiderLatestPositionReturn{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}
func testRiderViewV2() {
	param := util.RiderViewV2{
		DevID:     util.DevID,
		OrderID:   "3337109380461382689",
		OrderType: 1,
		ShopID:    util.ShopID,
		ShopType:  1,
		PushTime:  time.Now().Unix(),
	}
	bts, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.RiderViewV2Url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	returnparam := util.RiderViewV2Return{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}

//订单状态在没有回调成功或回调20次前是不会更新的
func testGetCallbackInfo() {
	param := util.GetCallbackInfo{
		DevID:     util.DevID,
		OrderID:   "3337058897423780876",
		OrderType: 1,
		ShopID:    util.ShopID,
		ShopType:  1,
		PushTime:  time.Now().Unix(),
	}
	bts, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	body, err := util.SendRequest(bts, util.GetCallbackInfoUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	returnparam := util.GetCallbackInfoRetuen{}
	json.Unmarshal(body, &returnparam)
	fmt.Printf("%+v\n", returnparam)
}
