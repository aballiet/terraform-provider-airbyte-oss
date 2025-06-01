terraform {
  required_providers {
    airbyte = {
      source  = "aballiet/airbyte-oss"
      version = "1.2.3"
    }
  }
}

provider "airbyte" {
  server_url = "http://localhost:8001/api/" # Must match your Airbyte server endpoint (airbyte-server-svc on k8s)
  username   = "YOUR_USERNAME"
  password   = "YOUR_PASSWORD"
  timeout    = 600 # 10 minutes - default is 60 seconds
}