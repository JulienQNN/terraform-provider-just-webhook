terraform {
  required_providers {
    jwb = {
      source  = "JulienQNN/just-webhook"
      version = "0.2.2"
    }
  }
}

provider "jwb" {
}
