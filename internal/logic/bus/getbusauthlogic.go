package bus

import (
	"context"
	"encoding/json"
	"net/url"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBusAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBusAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBusAuthLogic {
	return &GetBusAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBusAuthLogic) GetBusAuth(req *types.GetBusAuthReq) (resp *types.GetBusAuthResp, err error) {
	_, yxyHeaders := yxyClient.GetYxyBaseReqParam("")
	yxyReq := map[string]string{
		"authType":    "2",
		"ymAppId":     "2011112043190345310",
		"authAppid":   "10337",
		"callbackUrl": "https://api.pinbayun.com/api/v1/zjgd_interface/?schoolCode=10337",
		"unionid":     req.UID,
		"schoolCode":  consts.SCHOOL_CODE,
	}

	client := yxyClient.GetClient()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20))
	r, err := client.R().
		SetHeaders(yxyHeaders).
		SetQueryParams(yxyReq).
		Get(consts.GET_BUS_AUTH_CODE_URL)

	if err != nil {
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	redirectURL := r.RawResponse.Request.Response.Header.Get("Location")
	if redirectURL == "" {
		return nil, xerr.WithCode(xerr.ErrUnknown, "Login failed")
	}

	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	queryParams := parsedURL.Query()
	openid := queryParams.Get("openid")
	corpcode := queryParams.Get("corpcode")

	if openid == "" || corpcode == "" {
		return nil, xerr.WithCode(xerr.ErrHttpClient, "Auth code not found")
	}

	r, err = client.R().
		SetBody(map[string]string{
			"openid":   openid,
			"corpcode": corpcode,
		}).
		Post(consts.GET_BUS_AUTH_TOKEN_URL)

	if err != nil {
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	var result map[string]interface{}
	if err := json.Unmarshal(r.Body(), &result); err != nil {
		return nil, xerr.WithCode(xerr.ErrUnknown, err.Error())
	}

	token, ok := result["token"].(string)
	if !ok || token == "" {
		return nil, xerr.WithCode(xerr.ErrBusToeknInvalid, "Token is empty")
	}

	client.SetRedirectPolicy(resty.NoRedirectPolicy()) // 把重定向策略改回去避免影响其他功能

	return &types.GetBusAuthResp{
		Token: token,
	}, nil
}
