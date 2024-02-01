terraform {
  required_providers {
    jwb = {
      source  = "JulienQNN/just-webhook"
      version = "0.1.5"
    }
  }
}

provider "jwb" {
}
