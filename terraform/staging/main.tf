variable "bundle_url" {
  type    = string
  default = ""
}

resource "null_resource" "fake_deployment" {
  triggers = {
    bundle_url = var.bundle_url
  }

  provisioner "local-exec" {
    command = "echo \"pretending to deploy from ${var.bundle_url}\"..."
  }
}
