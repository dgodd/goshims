// Code generated by counterfeiter. DO NOT EDIT.
// with command: counterfeiter -p -o /Users/pivotal/workspace/diego-release/src/code.cloudfoundry.org/goshims/grpcshim google.golang.org/grpc
package grpcshim

import (
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/tap"
)

type GrpcShim struct{}

func (sh *GrpcShim) RoundRobin(r naming.Resolver) grpc.Balancer {
	return grpc.RoundRobin(r)
}

func (sh *GrpcShim) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
	return grpc.Invoke(ctx, method, args, reply, cc, opts...)
}

func (sh *GrpcShim) WithCodec(c grpc.Codec) grpc.DialOption {
	return grpc.WithCodec(c)
}

func (sh *GrpcShim) WithCompressor(cp grpc.Compressor) grpc.DialOption {
	return grpc.WithCompressor(cp)
}

func (sh *GrpcShim) WithDecompressor(dc grpc.Decompressor) grpc.DialOption {
	return grpc.WithDecompressor(dc)
}

func (sh *GrpcShim) WithBalancer(b grpc.Balancer) grpc.DialOption {
	return grpc.WithBalancer(b)
}

func (sh *GrpcShim) WithServiceConfig(c <-chan grpc.ServiceConfig) grpc.DialOption {
	return grpc.WithServiceConfig(c)
}

func (sh *GrpcShim) WithBackoffMaxDelay(md time.Duration) grpc.DialOption {
	return grpc.WithBackoffMaxDelay(md)
}

func (sh *GrpcShim) WithBackoffConfig(b grpc.BackoffConfig) grpc.DialOption {
	return grpc.WithBackoffConfig(b)
}

func (sh *GrpcShim) WithBlock() grpc.DialOption {
	return grpc.WithBlock()
}

func (sh *GrpcShim) WithInsecure() grpc.DialOption {
	return grpc.WithInsecure()
}

func (sh *GrpcShim) WithTransportCredentials(creds credentials.TransportCredentials) grpc.DialOption {
	return grpc.WithTransportCredentials(creds)
}

func (sh *GrpcShim) WithPerRPCCredentials(creds credentials.PerRPCCredentials) grpc.DialOption {
	return grpc.WithPerRPCCredentials(creds)
}

func (sh *GrpcShim) WithTimeout(d time.Duration) grpc.DialOption {
	return grpc.WithTimeout(d)
}

func (sh *GrpcShim) WithDialer(f func(string, time.Duration) (net.Conn, error)) grpc.DialOption {
	return grpc.WithDialer(f)
}

func (sh *GrpcShim) WithStatsHandler(h stats.Handler) grpc.DialOption {
	return grpc.WithStatsHandler(h)
}

func (sh *GrpcShim) FailOnNonTempDialError(f bool) grpc.DialOption {
	return grpc.FailOnNonTempDialError(f)
}

func (sh *GrpcShim) WithUserAgent(s string) grpc.DialOption {
	return grpc.WithUserAgent(s)
}

func (sh *GrpcShim) WithUnaryInterceptor(f grpc.UnaryClientInterceptor) grpc.DialOption {
	return grpc.WithUnaryInterceptor(f)
}

func (sh *GrpcShim) WithStreamInterceptor(f grpc.StreamClientInterceptor) grpc.DialOption {
	return grpc.WithStreamInterceptor(f)
}

func (sh *GrpcShim) Dial(target string, opts ...grpc.DialOption) (ClientConn, error) {
	conn, err := grpc.Dial(target, opts...)
	if err!= nil {
		return nil, err
	}
	return &ClientConnShim{
		ClientConn: conn,
	}, nil
}

func (sh *GrpcShim) DialContext(ctx context.Context, target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	return grpc.DialContext(ctx, target, opts...)
}

func (sh *GrpcShim) NewGZIPCompressor() grpc.Compressor {
	return grpc.NewGZIPCompressor()
}

func (sh *GrpcShim) NewGZIPDecompressor() grpc.Decompressor {
	return grpc.NewGZIPDecompressor()
}

func (sh *GrpcShim) Header(md *metadata.MD) grpc.CallOption {
	return grpc.Header(md)
}

func (sh *GrpcShim) Trailer(md *metadata.MD) grpc.CallOption {
	return grpc.Trailer(md)
}

func (sh *GrpcShim) FailFast(failFast bool) grpc.CallOption {
	return grpc.FailFast(failFast)
}

func (sh *GrpcShim) Code(err error) codes.Code {
	return grpc.Code(err)
}

func (sh *GrpcShim) ErrorDesc(err error) string {
	return grpc.ErrorDesc(err)
}

func (sh *GrpcShim) Errorf(c codes.Code, format string, a ...interface{}) error {
	return grpc.Errorf(c, format, a...)
}

func (sh *GrpcShim) CustomCodec(codec grpc.Codec) grpc.ServerOption {
	return grpc.CustomCodec(codec)
}

func (sh *GrpcShim) RPCCompressor(cp grpc.Compressor) grpc.ServerOption {
	return grpc.RPCCompressor(cp)
}

func (sh *GrpcShim) RPCDecompressor(dc grpc.Decompressor) grpc.ServerOption {
	return grpc.RPCDecompressor(dc)
}

func (sh *GrpcShim) MaxMsgSize(m int) grpc.ServerOption {
	return grpc.MaxMsgSize(m)
}

func (sh *GrpcShim) MaxConcurrentStreams(n uint32) grpc.ServerOption {
	return grpc.MaxConcurrentStreams(n)
}

func (sh *GrpcShim) Creds(c credentials.TransportCredentials) grpc.ServerOption {
	return grpc.Creds(c)
}

func (sh *GrpcShim) UnaryInterceptor(i grpc.UnaryServerInterceptor) grpc.ServerOption {
	return grpc.UnaryInterceptor(i)
}

func (sh *GrpcShim) StreamInterceptor(i grpc.StreamServerInterceptor) grpc.ServerOption {
	return grpc.StreamInterceptor(i)
}

func (sh *GrpcShim) InTapHandle(h tap.ServerInHandle) grpc.ServerOption {
	return grpc.InTapHandle(h)
}

func (sh *GrpcShim) StatsHandler(h stats.Handler) grpc.ServerOption {
	return grpc.StatsHandler(h)
}

func (sh *GrpcShim) NewServer(opt ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opt...)
}

func (sh *GrpcShim) SetHeader(ctx context.Context, md metadata.MD) error {
	return grpc.SetHeader(ctx, md)
}

func (sh *GrpcShim) SendHeader(ctx context.Context, md metadata.MD) error {
	return grpc.SendHeader(ctx, md)
}

func (sh *GrpcShim) SetTrailer(ctx context.Context, md metadata.MD) error {
	return grpc.SetTrailer(ctx, md)
}

func (sh *GrpcShim) NewClientStream(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (_ grpc.ClientStream, err error) {
	return grpc.NewClientStream(ctx, desc, cc, method, opts...)
}
