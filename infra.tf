provider "google" {
  credentials = "${file("./creds/terraform.json")}"
  project     = "ori-tsauvajon"
  region      = "us-east1"
}

resource "google_container_cluster" "gke-cluster" {
  name               = "twelvefa-gke-cluster"
  network            = "default"
  location           = "us-east1-c"

  # can't create an empty cluster so we immediatly delete the node pool
  remove_default_node_pool = true
  initial_node_count = 1
}

resource "google_container_node_pool" "primary_nodes" {
  name       = "twelvefa-node-pool"
  location   = "us-east1"
  cluster    = "${google_container_cluster.gke-cluster.name}"
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = "f1-micro"

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}