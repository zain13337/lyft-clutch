// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: bot/slackbot/v1/slackbot.proto

/*
Package slackbotv1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package slackbotv1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_SlackBotAPI_Event_0(ctx context.Context, marshaler runtime.Marshaler, client SlackBotAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EventRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Event(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SlackBotAPI_Event_0(ctx context.Context, marshaler runtime.Marshaler, server SlackBotAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EventRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Event(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterSlackBotAPIHandlerServer registers the http handlers for service SlackBotAPI to "mux".
// UnaryRPC     :call SlackBotAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterSlackBotAPIHandlerFromEndpoint instead.
func RegisterSlackBotAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server SlackBotAPIServer) error {

	mux.Handle("POST", pattern_SlackBotAPI_Event_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/clutch.bot.slackbot.v1.SlackBotAPI/Event", runtime.WithHTTPPathPattern("/v1/bot/slackbot/event"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SlackBotAPI_Event_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SlackBotAPI_Event_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterSlackBotAPIHandlerFromEndpoint is same as RegisterSlackBotAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSlackBotAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSlackBotAPIHandler(ctx, mux, conn)
}

// RegisterSlackBotAPIHandler registers the http handlers for service SlackBotAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSlackBotAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterSlackBotAPIHandlerClient(ctx, mux, NewSlackBotAPIClient(conn))
}

// RegisterSlackBotAPIHandlerClient registers the http handlers for service SlackBotAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "SlackBotAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "SlackBotAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "SlackBotAPIClient" to call the correct interceptors.
func RegisterSlackBotAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client SlackBotAPIClient) error {

	mux.Handle("POST", pattern_SlackBotAPI_Event_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/clutch.bot.slackbot.v1.SlackBotAPI/Event", runtime.WithHTTPPathPattern("/v1/bot/slackbot/event"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SlackBotAPI_Event_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SlackBotAPI_Event_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_SlackBotAPI_Event_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "bot", "slackbot", "event"}, ""))
)

var (
	forward_SlackBotAPI_Event_0 = runtime.ForwardResponseMessage
)
