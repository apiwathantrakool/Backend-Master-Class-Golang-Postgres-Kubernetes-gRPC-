package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIp  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	// Extract metadata from the incoming context
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if userAgents, exists := md["user-agent"]; exists {
			mtdt.UserAgent = userAgents[0]
		}
	}

	// Extract client IP from peer info
	p, ok := peer.FromContext(ctx)
	if ok {
		clientIP := p.Addr.String()
		mtdt.ClientIp = clientIP
	}

	return mtdt
}
