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

type GetBusRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBusRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBusRecordLogic {
	return &GetBusRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type GetBusRecordsYxyResp struct {
	Results []struct {
		DateInfo struct {
			Info struct {
				ID   string `json:"id"`
				Name string `json:"shuttle_name"`
			} `json:"shuttle_bus_vo"`
		} `json:"shuttle_bus_date_vo"`
		DepartureTime string `json:"departure_datetime"`
		PayTime       string `json:"pay_time"`
	} `json:"results"`
}

func (l *GetBusRecordLogic) GetBusRecord(req *types.GetBusRecordReq) (resp *types.GetBusRecordResp, err error) {
	var yxyResp GetBusRecordsYxyResp
	client := yxyClient.GetClient()
	r, err := client.R().
		SetQueryParams(map[string]string{
			"page":      strconv.Itoa(req.Page),
			"page_size": strconv.Itoa(req.PageSize),
			"status":    req.Status,
		}).
		SetHeader("Authorization", req.Token).
		SetResult(&yxyResp).
		Get(consts.GET_BUS_RECORD_URL)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", consts.GET_BUS_RECORD_URL, err)
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	if r.StatusCode() == 400 {
		return nil, xerr.WithCode(xerr.ErrHttpClient, fmt.Sprintf("yxy response: %v", r))
	} else if r.StatusCode() == 500 {
		return nil, xerr.WithCode(xerr.ErrHttpClient, fmt.Sprintf("yxy response: %v", r))
	}

	var records []types.BusRecord
	for _, row := range yxyResp.Results {
		record := types.BusRecord{
			ID:            row.DateInfo.Info.ID,
			Name:          row.DateInfo.Info.Name,
			DepartureTime: row.DepartureTime,
			PayTime:       row.PayTime,
		}
		records = append(records, record)
	}
	return &types.GetBusRecordResp{
		BusRecords: records,
	}, nil
}
