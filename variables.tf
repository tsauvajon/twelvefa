variable "project_id" {
  type        = string
  description = "The ID of the GCP project"
  default     = "ori-tsauvajon"
}

variable "region" {
  type        = string
  description = "GCP region to deploy to"
  default     = "us-east1"
}

variable "location" {
  type        = string
  description = "GCP location to deploy to"
  default     = "us-east1-c"
}

variable "cluster_name" {
  type        = string
  description = "GKE main cluster name"
  default     = "twelvefa-gke-cluster"
}
