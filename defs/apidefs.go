package defs

type WxappLoginVO struct {
	Token string `json:"token"`
}

type WxappUserInfoVO struct {
	Uid      int    `json:"uid"`
	Nickname string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
}

type WxappAuthUserInfoReq struct {
	NickName  string `json:"nickName" validate:"required"`
	AvatarUrl string `jsoN:"avatarUrl" validate:"required"`
	Gender    int    `json:"gender"`
	Country   string `json:"country"`
	Province  string `json:"province"`
	City      string `json:"city"`
}

type RequestVO struct {
	Url    string `json:"url"`    // 请求URL
	Method string `json:"method"` // 请求方法
	Body   string `json:"body"`   // body报文
}

// PortalBannerVO banner对象
type PortalBannerVO struct {
	Id           int    `json:"id"`           // ID
	Picture      string `json:"picture"`      // 图片地址
	BusinessType int    `json:"businessType"` // 类型
	BusinessId   int    `json:"businessId"`   // 业务ID
}

// WxappAuthPhoneReq 授权手机号请求对象
type WxappAuthPhoneReq struct {
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}

// PortalGridCategoryVO 宫格列表
type PortalGridCategoryVO struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`       // 宫格标题
	CategoryId int    `json:"categoryId"` // 关联的分类
	Picture    string `json:"picture"`    // 宫格图标
}

// PortalCategoryVO 分类对象
type PortalCategoryVO struct {
	Id   int    `json:"id"`   // 分类ID
	Name string `json:"name"` // 分类名称
}

// PortalGoodsListVO 商品列表
type PortalGoodsListVO struct {
	Id       int     `json:"id"`       // 商品ID
	Title    string  `json:"title"`    // 标题
	Price    float64 `json:"price"`    // 价格
	Picture  string  `json:"picture"`  // 图片
	HumanNum int     `json:"humanNum"` // 购买人数
}

// PortalGoodsInfo 商品详情
type PortalGoodsInfo struct {
	Id            int            `json:"id"`            // 商品ID
	Title         string         `json:"title"`         // 标题
	Price         float64        `json:"price"`         // 价格
	Picture       string         `json:"picture"`       // 主图
	BannerPicture string         `json:"bannerPicture"` // 详情图
	DetailPicture string         `json:"detailPicture"` // 轮播图
	Tags          string         `json:"tags"`          // 标签
	Description   string         `json:"description"`   // 详情
	SkuList       []PortalSkuVO  `json:"skuList"`       // sku列表
	SpecList      []PortalSpecVO `json:"specList"`      // 规格列表
}

type PortalSpecVO struct {
	SpecId   int                `json:"specId"`   // 规格ID
	Name     string             `json:"name"`     // 规格名称
	AttrList []PortalSpecAttrVO `json:"attrList"` // 规格属性
}

type PortalSpecAttrVO struct {
	AttrId int    `json:"attrId"` // 属性ID
	Value  string `json:"value"`  // 属性名称
}

type PortalSkuVO struct {
	Id      int     `json:"id"`      // skuId
	Picture string  `json:"picture"` // 图片
	Title   string  `json:"title"`   // 标题
	Price   float64 `json:"price"`   // 价格
	Code    string  `json:"code"`    // 编码
	Stock   int     `json:"stock"`   // 库存量
	Specs   string  `json:"specs"`   // 多规格属性
}

type PortalCartGoodsVO struct {
	Id      int     `json:"id"`      // 购物车ID
	GoodsId int     `json:"goodsId"` // 商品ID
	SkuId   int     `json:"skuId"`   // SKU ID
	Title   string  `json:"title"`   // 标题
	Price   float64 `json:"price"`   // SKU价格
	Picture string  `json:"picture"` // SKU图片
	Specs   string  `json:"specs"`   // specs值
	Num     int     `json:"num"`     // 数量
	Status  int     `json:"status"`  // 库存状态：0-正常 1-缺货 2-下架
}

// PortalCartGoodsReq 添加购物车
type PortalCartGoodsReq struct {
	GoodsId int `json:"goodsId"` // 商品ID
	SkuId   int `json:"skuId"`   // skuId
	Num     int `json:"num"`     // 数量
}

// PortalEditCartReq 编辑购物车
type PortalEditCartReq struct {
	Id  int `json:"id"`  // 主键
	Num int `json:"num"` // 数量：-1 减一件 0 删除 1 加一件
}

type PortalUserCouponVO struct {
	CLogId      int    `json:"CLogId"`      // 领取记录ID
	CouponId    int    `json:"couponId"`    // 优惠券ID
	Title       string `json:"title"`       // 标题
	FullMoney   string `json:"fullMoney"`   // 满减额
	Minus       string `json:"minus"`       // 优惠额
	Rate        string `json:"rate"`        // 折扣
	Type        int    `json:"type"`        // 券类型：1-满减券 2-折扣券 3-代金券 4-满金额折扣券
	StartTime   string `json:"startTime"`   // 开始时间
	EndTime     string `json:"endTime"`     // 结束时间
	Description string `json:"description"` // 描述
}

type PortalCouponVO struct {
	Id          int    `json:"id"`          // 优惠券ID
	Title       string `json:"title"`       // 标题
	FullMoney   string `json:"fullMoney"`   // 满减额
	Minus       string `json:"minus"`       // 优惠额
	Rate        string `json:"rate"`        // 折扣
	Type        int    `json:"type"`        // 券类型：1-满减券 2-折扣券 3-代金券 4-满金额折扣券
	StartTime   string `json:"startTime"`   // 开始时间
	EndTime     string `json:"endTime"`     // 结束时间
	Description string `json:"description"` // 描述
}

type PortalTakeCouponReq struct {
	CouponId int `json:"couponId" validate:"required"`
}

type PortalOrderListVO struct {
	Id        int                  `json:"id"`        // 订单ID
	OrderNo   string               `json:"orderNo"`   // 订单号
	PlaceTime string               `json:"placeTime"` // 下单时间
	PayAmount float64              `json:"payAmount"` // 支付金额
	Status    int                  `json:"status"`    // 订单状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成
	GoodsNum  int                  `json:"goodsNum"`  // 商品数量
	GoodsList []PortalOrderGoodsVO `json:"goodsList"`
}

type PortalOrderGoodsVO struct {
	GoodsId int     `json:"goodsId"` // 商品ID
	Title   string  `json:"title"`   // 标题
	Price   float64 `json:"price"`   // 价格
	Picture string  `json:"picture"` // 图片
	SkuId   int     `json:"skuId"`   // skuId
	Specs   string  `json:"specs"`   // specs值
	Num     int     `json:"num"`     // 数量
}

type PortalOrderDetailVO struct {
	Id             int                  `json:"id"`             // 订单ID
	OrderNo        string               `json:"orderNo"`        // 订单号
	GoodsAmount    float64              `json:"goodsAmount"`    // 商品小计
	DiscountAmount float64              `json:"discountAmount"` // 优惠金额
	DispatchAmount float64              `json:"dispatchAmount"` // 运费
	PayAmount      float64              `json:"payAmount"`      // 实付款
	Status         int                  `json:"status"`         // 订单状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成 4-（待发货）退款申请 5-已退款
	GoodsNum       int                  `json:"goodsNum"`       // 商品数量
	PlaceTime      string               `json:"placeTime"`      // 下单时间
	PayTime        string               `json:"payTime"`        // 付款时间
	DeliverTime    string               `json:"deliverTime"`    // 发货时间
	FinishTime     string               `json:"finishTime"`     // 完成时间
	GoodsList      []PortalOrderGoodsVO `json:"goodsList"`      // 订单商品
	Address        AddressSnapshot      `json:"address"`        // 收货地址
	RefundApply    OrderRefundApplyVO   `json:"refundApply"`    // 退款申请
}

type OrderRefundApplyVO struct {
	RefundNo string `json:"refundNo"` // 退款编号
}

type OrderRemindVO struct {
	WaitPay     int `json:"waitPay"`     // 待付款
	NotExpress  int `json:"notExpress"`  // 待发货
	WaitReceive int `json:"waitReceive"` // 待收货
}

type OrderRefundApplyReq struct {
	OrderNo string `json:"orderNo" validate:"required"` // 订单号
	Reason  string `json:"reason" validate:"required"`  // 退款原因
}

type OrderRefundDetailVO struct {
	RefundNo     string               `json:"refundNo"`     // 退款编号
	Reason       string               `json:"reason"`       // 退款原因
	RefundAmount float64              `json:"refundAmount"` // 退款金额
	Status       int                  `json:"status"`       // 状态：0-退款申请 1-商家处理申请 2-退款完成
	ApplyTime    string               `json:"applyTime"`    // 申请时间
	RefundTime   string               `json:"refundTime"`   // 退款时间
	GoodsList    []PortalOrderGoodsVO `json:"goodsList"`    // 订单商品
}

type PortalAddressVO struct {
	Id          int    `json:"id"`          // 地址ID
	Contacts    string `json:"contacts"`    // 联系人
	Mobile      string `json:"mobile"`      // 手机号
	ProvinceId  string `json:"provinceId"`  // 省份编码
	CityId      string `json:"cityId"`      // 城市编码
	AreaId      string `json:"areaId"`      // 地区编码
	ProvinceStr string `json:"provinceStr"` // 省份
	CityStr     string `json:"cityStr"`     // 城市
	AreaStr     string `json:"areaStr"`     // 地区
	Address     string `json:"address"`     // 详细地址
	IsDefault   int    `json:"isDefault"`   // 默认收货地址：0-否 1-是
}

type PortalAddressReq struct {
	Id          int    `json:"id"`          // 地址ID
	Contacts    string `json:"contacts"`    // 联系人
	Mobile      string `json:"mobile"`      // 手机号
	ProvinceId  string `json:"provinceId"`  // 省份编码
	CityId      string `json:"cityId"`      // 城市编码
	AreaId      string `json:"areaId"`      // 地区编码
	ProvinceStr string `json:"provinceStr"` // 省份
	CityStr     string `json:"cityStr"`     // 城市
	AreaStr     string `json:"areaStr"`     // 地区
	Address     string `json:"address"`     // 详细地址
	IsDefault   int    `json:"isDefault"`   // 默认收货地址：0-否 1-是
}

type PortalBrowseRecordVO struct {
	Id         int    `json:"id"`         // 记录ID
	GoodsId    int    `json:"goodsId"`    // 商品ID
	Picture    string `json:"picture"`    // 商品图片
	Title      string `json:"title"`      // 商品标题
	Price      string `json:"price"`      // 商品价格
	BrowseTime string `json:"browseTime"` // 浏览时间，格式：yyyy-MM-dd HH:mm:ss
}

type PortalCartPlaceOrderReq struct {
	AddressId      int               `json:"addressId"`      // 收货地址ID
	CouponLogId    int               `json:"couponLogId"`    // 优惠券记录ID
	DispatchAmount string            `json:"dispatchAmount"` // 运费
	ExpectAmount   string            `json:"expectAmount"`   // 预期支付金额
	GoodsList      []PortalCartGoods `json:"goodsList"`      // 下单商品
}

type PortalCartGoods struct {
	Num     int `json:"num"`     // 数量
	CartId  int `json:"cartId"`  // 购物车ID
	GoodsId int `json:"goodsId"` // 商品ID
	SkuId   int `json:"skuId"`   // skuID
}

type PortalPlaceOrderVO struct {
	OrderNo  string `json:"orderNo"`  // 订单号
	PrepayId string `json:"prepayId"` // 预支付ID
}
