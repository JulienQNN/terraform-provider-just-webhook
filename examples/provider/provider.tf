terraform {
  required_providers {
    jwb = {
      source  = "JulienQNN/just-webhook"
      version = "0.2.1"
    }
  }
}

provider "jwb" {
}
