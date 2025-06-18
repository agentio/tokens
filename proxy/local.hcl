
call "tokens" {
  name     = "Tokens"
  port     = 9000
  endpoint = "localhost"
  insecure = true
}

host "localhost" {
  name    = "Tokens"
  backend = "localhost:3000"
  atproto {
    name   = "Tokens"
    scopes = ["atproto"]
    local  = 9000
  }
}
