curl -X PUT http://docker:5984/locations
curl -X PUT http://docker:5984/checkins
curl -X PUT http://docker:5984/checkins_needreview

curl -X POST localhost:4567/location/483691cff964a520ec4f1fe3/3176
curl -X POST localhost:4567/location/4ac6769df964a52075b420e3/8350
curl -X POST localhost:4567/location/49be98edf964a520c0541fe3/3538
curl -X POST localhost:4567/location/4a6274a3f964a5201ec41fe3/5320
curl -X POST localhost:4567/location/540dc3c1498e8ef1a0c2cf66/2100785
