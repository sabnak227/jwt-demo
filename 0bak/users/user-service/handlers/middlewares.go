package handlers

import (
	"context"
	"log"
	"time"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/sabnak227/jwt-demo/bak/users"
	"github.com/sabnak227/jwt-demo/bak/users/user-service/svc"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/metaverse/truss/_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	in.WrapAllLabeledExcept(timingMiddleware())
	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	return in
}

func WrapService(in pb.UserServer) pb.UserServer {
	return in
}

func timingMiddleware() svc.LabeledMiddleware {
	return func(name string, in endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
			defer func(begin time.Time) {
				log.Printf("Endpoint: %s", name)
				log.Printf("Input: %v", req)
				log.Printf("Output: %v", rsp)
				log.Printf("Elpased Time: %s", time.Since(begin).String())
			}(time.Now())
			return in(ctx, req)
		}
	}
}
