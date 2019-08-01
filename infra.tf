provider "google" {
  credentials = "${file("./creds/terraform.json")}"
  project     = "ori-tsauvajon"
  region      = "us-east1"
}

resource "google_container_cluster" "gke-cluster" {
  name               = "ori-tsauvajon-gke-cluster"
  network            = "default"
  zone               = "us-east1-c"
  initial_node_count = 3
}