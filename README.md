#tutum-machine

Experimental project.
Provision Tutum node-clusters with Stackfiles from a single YAML file.

##Intructions

(Will likely change soon)

To launch clusters from a YAML file (see example.yml for pattern):

`go run main.go up` will launch a local tutum-machine.yml file

or

`go run main.go --file /path/to/yaml up` will launch the targeted file.

##TODO

- Process Stackfiles
- Handle Taging
- Error handling
