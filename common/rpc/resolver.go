package rpc

import (
	"google.golang.org/grpc/resolver"
	"sync"
)

type myResolverBuilder struct{}

func (b *myResolverBuilder) Scheme() string {
	return "my_resolver"
}

func (*myResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &myResolver{
		target: target,
		cc:     cc,
	}
	// 在这里根据 target.ServiceName 获取服务器地址
	// 并将服务器地址更新到 cc 中
	r.updateAddresses()
	return r, nil
}

type myResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	mu     sync.Mutex
}

func (r *myResolver) ResolveNow(resolver.ResolveNowOptions) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 在这里实现手动触发服务地址解析的逻辑
	r.updateAddresses()
}

func (r *myResolver) Close() {
	// 在这里实现资源释放的逻辑
}

func (r *myResolver) updateAddresses() {
	// 在这里根据 target.ServiceName 获取服务器地址
	// 然后将地址更新到 r.cc 中
	addresses := []resolver.Address{
		{Addr: "localhost:50051"},
		// ...
	}
	r.cc.UpdateState(resolver.State{Addresses: addresses})
}
