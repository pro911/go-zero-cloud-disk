package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pro911/go-zero-cloud-disk/core/models"
	"log"

	"github.com/pro911/go-zero-cloud-disk/core/internal/svc"
	"github.com/pro911/go-zero-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	resp.Message = "Hello Cloud Disk!"
	//获取用户列表
	data := make([]*models.UserBasic, 0)
	err = models.Engine.Find(&data)
	if err != nil {
		log.Println("Get UserBasic Error:", err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Marshal Error:", err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		log.Println("Json Indent Error:", err)
	}
	fmt.Println(dst.String())
	resp = new(types.Response)
	resp.Message = dst.String()
	return
}
