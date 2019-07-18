# Clusteradm user journeys 

Clusteradm primary goal is to provide a simplified consistent UX, so therefore outlining user journeys is essential.  In this document we will walk through some of the common user experiences. 

## Central themes 

1.  
2. 

## Getting started 

### (REPL/Fuse Mode)

clusteradm 
... Fuse mounts an FS and starts a repl (this is weird)
/> help
init version ls
/> init --providers=aws,vsphere 
... if no KUBECONFIG assume --local=true
... deploy clusteradm operator 
... clusteradm operator deploys aws and vsphere providers 
/> ls 
aws vsphere 
/> cd aws 
/aws/> ls 
. 
/aws/> help 
config version apply ls 
/aws/> config test
... default editor or have a conf? 
/aws/> save 
... applies the YAML
/aws/> ls 
test . 
/aws/> config test

TODO: Directory 
cat spec
cat status

Questions: 
1. Namespacing 
2. Using standard fs options and scripting `find . | grep foo`
3. 

### Subcommands
clusteradm init --providers=aws,vsphere 
TODO: finish subcommand standard workflow for getting started.

