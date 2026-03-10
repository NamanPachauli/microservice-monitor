terraform {
  required_version = ">= 1.5.0"

  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "~> 2.4"
    }
  }
}

provider "local" {}

resource "local_file" "example" {
  filename        = "C:/Users/PIS/microservice-monitor/terraform/hello.txt"
  content         = <<EOT
Hello Terraform!
This file is created using Terraform.
You can expand this workflow to manage Docker, Kubernetes, and more.
Happy Learning! 
EOT
  file_permission = "0644"
}

output "file_path" {
  value       = local_file.example.filename
  description = "Path of the file created by Terraform"
}