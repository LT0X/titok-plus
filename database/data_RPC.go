package database

import (
	"douyin/rpc/contact/contact_client"
	"douyin/rpc/user/user_client"
	"douyin/rpc/video/video_client"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

var RPC *RPCServiceContext

type RPCServiceContext struct {
	VideoRpc   video_client.Video
	UserRpc    user_client.User
	ContactRpc contact_client.Contact
}

func NewRPCServiceContext() {
	RPC = &RPCServiceContext{
		VideoRpc: video_client.NewVideo(zrpc.MustNewClient(zrpc.RpcClientConf{
			Etcd: discov.EtcdConf{
				Hosts: []string{"192.168.169.128:2379"},
				Key:   "video.rpc",
			},
			//Endpoints: []string{"127.0.0.1:8003"},
			NonBlock: true,
		})),
		UserRpc: user_client.NewUser(zrpc.MustNewClient(zrpc.RpcClientConf{
			Etcd: discov.EtcdConf{
				Hosts: []string{"192.168.169.128:2379"},
				Key:   "user.rpc",
			},
			//Endpoints: []string{"127.0.0.1:8002"},
			NonBlock: true,
		})),
		ContactRpc: contact_client.NewContact(zrpc.MustNewClient(zrpc.RpcClientConf{
			Etcd: discov.EtcdConf{
				Hosts: []string{"192.168.169.128:2379"},
				Key:   "contact.rpc",
			},
			//Endpoints: []string{"127.0.0.1:8001"},
			NonBlock: true,
		})),
	}
}
