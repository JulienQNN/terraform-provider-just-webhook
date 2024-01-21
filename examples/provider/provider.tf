terraform {
  required_providers {
    just-webhook = {
      source  = "JulienQNN/just-webhook"
      version = "0.1.0"
    }
  }
}
provider "jwb" {
}
