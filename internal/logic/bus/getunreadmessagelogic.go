package bus

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnreadMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUnreadMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnreadMessageLogic {
	return &GetUnreadMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type GetMessage struct {
	MessageList []types.Message `json:"results"`
	Detail      struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	} `json:"detail"`
}

func (l *GetUnreadMessageLogic) GetUnreadMessage(req *types.GetUnreadMessageReq) (resp *types.GetUnreadMessageResp, err error) {
	var yxyResp GetMessage
	client := yxyClient.GetClient()
	r, err := client.R().
		SetQueryParams(map[string]string{
			"page":      strconv.Itoa(req.Page),
			"page_size": strconv.Itoa(req.PageSize),
			"num_pages": "0", // TODO: what is this?
		}).
		SetHeader("Authorization", req.Token).
		Get(consts.GET_BUS_MESSAGE_UUL)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", consts.GET_BUS_RECORD_URL, err)
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	err = json.Unmarshal(r.Body(), &yxyResp)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	if r.StatusCode() == 400 {
		if yxyResp.Detail.Code == "AUTH_FAIL" {
			return nil, xerr.WithCode(xerr.ErrTokenInvalid, "权限验证失败")
		} else {
			return nil, xerr.WithCode(xerr.ErrHttpClient, fmt.Sprintf("yxy response: %v", r))
		}
	} else if r.StatusCode() == 500 {
		return nil, xerr.WithCode(xerr.ErrHttpClient, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.GetUnreadMessageResp{
		List: yxyResp.MessageList,
	}, nil
}
