provider "google" {
  version = "~> 2.11"

  credentials = "${file("./creds/terraform.json")}"
  project     = "ori-tsauvajon"
  region      = "us-east1"
}

resource "google_container_cluster" "gke-cluster" {
  name     = "twelvefa-gke-cluster"
  network  = "default"
  location = "us-east1-c"

  node_pool {
    name       = "twelvefa-node-pool"
    initial_node_count = 3

    autoscaling {
      min_node_count = 3
      max_node_count = 5
    }

    node_config {
      preemptible  = true
      machine_type = "n1-standard-1"
      disk_size_gb = 30

      oauth_scopes = [
        "https://www.googleapis.com/auth/logging.write",
        "https://www.googleapis.com/auth/monitoring",
        "https://www.googleapis.com/auth/devstorage.read_only", // read GCR images
      ]
    }
  }
}
