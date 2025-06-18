package clients

import (
	xrpc_anonymous "github.com/agentio/atiquette/pkg/xrpc/anonymous"
)

var AnonymousClient = &xrpc_anonymous.Client{Host: "https://public.api.bsky.app"}
