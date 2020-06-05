package util

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	//当对接方接收到订单时，可通过此接口进行物流单的创建。发单时需保证商家订单号在同一家店铺保持唯一。
	//
	//测试时请额外注意系统自动绑定的测试店铺，因为测试开发者共用，所以商家订单号最好与当前时间相关来避免重复。
	//
	//注意：除特殊说明外对接方均为非平台型接入即不需传入shop相关信息。
	CreateOrderUrl string = SERVER_HOST + "/open/api/external/createorder"
	//并非真正发单；用来验证是否可以发单并在成功时返回时效、计价等信息，也可用来验证地址以及时间是否在顺丰的配送范围内。
	PreCreateOrderUrl string = SERVER_HOST + "/open/api/external/precreateorder"
	//当商家处发生异常需要取消配送时，可调用此接口对订单进行取消操作，同步返回结果。
	CancelOrderUrl string = SERVER_HOST + "/open/api/external/cancelorder"
	//订单创建后，骑士未接单的情况下通过该接口对订单进行加小费，促进订单接单
	//
	//注意：小费单位都是分
	AddOrderGratuityFeeUrl string = SERVER_HOST + "/open/api/external/addordergratuityfee"
	//订单加小费后，可通过该接口获取加小费的详细信息
	GetOrderGratuityFeeUrl string = SERVER_HOST + "/open/api/external/getordergratuityfee"
	//此接口可获取到指定订单的状态信息，用来进行订单状态的查询。可通过此对订单进行状态补齐操作，当接收顺丰状态回调失败时，主动查询订单状态。
	ListOrderFeedUrl string = SERVER_HOST + "/open/api/external/listorderfeed"
	//当订单为配送状态中，可通过该接口发起催单
	ReminderOrderUrl string = SERVER_HOST + "/open/api/external/reminderorder"
	//当订单被接单后，可通过该接口改收件人信息
	ChangeOrderUrl string = SERVER_HOST + "/open/api/external/changeorder"
	//此接口用于获取订单配送员的实时经纬度坐标，一般情况下骑士经纬度30s更新一次。
	RiderLatestPositionUrl string = SERVER_HOST + "/open/api/external/riderlatestposition"
	//此接口可获取一个订单的骑士位置H5链接，可进行内嵌或发送给用户（内嵌时无法保证界面的兼容性，如发现兼容性问题可使用获取配送员坐标接口自行开发轨迹H5）。
	RiderViewV2Url string = SERVER_HOST + "/open/api/external/riderviewv2"
	//顺丰订单回调详细查看接口。可以订单维度查询所有的回调信息，并在回调信息接收出现问题的时候主动查询此接口进行订单状态同步。
	GetCallbackInfoUrl string = SERVER_HOST + "/open/api/external/getcallbackinfo"
)

func Sign(param []byte) (string, error) {
	//bts,err:= json.Marshal(param)
	//if err!=nil{
	//	fmt.Println(err)
	//	return "",err
	//}
	str := string(param) + "&" + fmt.Sprintf("%d", DevID) + "&" + DevKey
	//fmt.Println(str)
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	fmt.Println(md5str)
	base := base64.StdEncoding.EncodeToString([]byte(md5str))

	return base, nil
}
func SendRequest(bts []byte, url string) ([]byte, error) {
	sign, err := Sign(bts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(sign)
	reader := bytes.NewReader(bts)
	resp, err := http.Post(url+"?sign="+sign, "application/json", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
