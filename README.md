barfight
========

Go application to poll twitter data from bars and populate a db

Initiall designed to work with couchdb running in a docker container. 


Setup
* Setup Docker on whatever OS you're running in
* Download the couchdb containier `docker pull klaemo/couchdb`
* Run container with ports exposed `docker run -d -p 5984:5984 --name couchdb klaemo/couchdb`
* (If on OSX) Add your boot2docker vm to the hosts file as 'docker' ie. `192.168.59.103  docker` 
