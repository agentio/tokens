host "tokens.babu.dev" {
  name    = "Tokens"
  backend = "nomad:tokens"
  atproto {
    name   = "Tokens"
    scopes = ["atproto"]
  }
}
