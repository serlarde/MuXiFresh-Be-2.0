package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
    rest.RestConf
    ModifytheUseravatar     zrpc.RpcClientConf     
    ModifytheUsername       zrpc.RpcClientConf          
	ModifytheUsertype       zrpc.RpcClientConf    
	AdjustAdmissionstatus   zrpc.RpcClientConf  
}