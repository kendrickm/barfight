curl -X PUT http://docker:5984/config
curl -X PUT http://docker:5984/config/bbbd5b349b7aa42d7cf0c38dd7028812 -d '{"name":"The Hop and Vine", "twitter":"thehopandvine" }'
curl -X PUT http://docker:5984/config/bbbd5b349b7aa42d7cf0c38dd7028652 -d '{"name":"Hopworks", "twitter":"hopworksbeer" }'
curl -X PUT http://docker:5984/config/bbbd5b349b7aa42d7cf0c38dd70278f4 -d '{"name":"The Beermongers", "twitter":"thebeermongers" }'
curl -X PUT http://docker:5984/config/bbbd5b349b7aa42d7cf0c38dd702737d -d '{"name":"Saraveza", "twitter":"Saraveza" }'