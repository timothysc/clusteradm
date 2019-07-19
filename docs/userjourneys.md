# Clusteradm user journeys 

Clusteradm primary goal is to provide a simplified consistent UX, so therefore outlining user journeys is essential.  In this document we will walk through some of the common user experiences. 
 
## Preconditions 
Some initial seed cluster which can be local needed.  If you do not have one we recommend using [KIND](https://github.com/kubernetes-sigs/kind)

## Getting started 

`go get https://github.com/timothysc/clusteradm` 

TODO: ACL - mgmt 

### Existing Managment Cluster

`clusteradm init --providers=aws,vsphere` 

### Bootstrap a new Management Cluster 

`clusteradm init --boostrap=aws --providers=aws,vsphere`

### Common workflow 

`clusteradm config --provider=aws -o test.yaml (default --falvor=dev)`

... edit test.yaml ... 

`clusteradm apply -f test.yaml`

... Should do some initial linting and type validation ... 

`clusteradm status` 

## Upgrade 

`clusteradm upgrade` 

## Remove 

`clusteradm reset` 

... removes cluster api based objects but doesn't touch the clusters ... 
