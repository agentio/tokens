
call "@" {
  name     = "@"
  port     = 7000
  endpoint = "@"
  operation "com.atproto.repo.createRecord" {
    method      = "POST"
    path        = "/xrpc/com.atproto.repo.createRecord"
    description = "Create a record in a repo"
  }
  operation "com.atproto.repo.listRecords" {
    method      = "GET"
    path        = "/xrpc/com.atproto.repo.listRecords"
    description = "List records in a repo"
  }
  operation "app.bsky.actor.getProfile" {
    method      = "GET"
    path        = "/xrpc/app.bsky.actor.getProfile"
    description = "Get a user's profile"
  }
}

call "tokens" {
  name     = "tokens"
  port     = 9000
  endpoint = "localhost"
  insecure = true
}

ingress {
  http  = 80
  https = 443
}

host "localhost" {
  name    = "tokens"
  backend = "localhost:3000"
  atproto {
    name   = "Tokens"
    scopes = ["atproto"]
    local  = 9000
  }
}

