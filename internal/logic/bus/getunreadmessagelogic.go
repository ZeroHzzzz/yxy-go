package bus

import (
	"context"
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
}

func (l *GetUnreadMessageLogic) GetUnreadMessage(req *types.GetUnreadMessageReq) (resp *types.GetUnreadMessageResp, err error) {
	var messagelist GetMessage
	client := yxyClient.GetClient()
	r, err := client.R().
		SetQueryParams(map[string]string{
			"page":      strconv.Itoa(req.Page),
			"page_size": strconv.Itoa(req.PageSize),
			"num_pages": "0", // TODO: what is this?
		}).
		SetHeader("Authorization", req.Token).
		SetResult(&messagelist).
		Get(consts.GET_BUS_MESSAGE_UUL)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", consts.GET_BUS_RECORD_URL, err)
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	if r.StatusCode() == 400 {
		return nil, xerr.WithCode(xerr.ErrHttpClient, fmt.Sprintf("yxy response: %v", r))
	} else if r.StatusCode() == 500 {
		return nil, xerr.WithCode(xerr.ErrHttpClient, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.GetUnreadMessageResp{
		Messages: messagelist.MessageList,
	}, nil
}
