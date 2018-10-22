# OpenFaas Cloud Functions [![OpenFaaS](https://img.shields.io/badge/openfaas-cloud-blue.svg)](https://www.openfaas.com)
This repository contains all my functions that are published by OpenFaas cloud.

1. Point In Polygon
2. Google Maps URL

## Point In Polygon
OpenFaas Point In Polygon algorithm. 

This algorithm determines if a GPS location (`point`) is within a GPS polygon (`polygon`). To determine if the point is within the polygon, the [even-odd rule algorithm](https://en.wikipedia.org/wiki/Even%E2%80%93odd_rule) is used.

### Usage

This function can be build and run locally but is also published via [OpenFaas Cloud](https://github.com/openfaas/openfaas-cloud) and available at:

`https://michaelvlaar.o6s.io/point-in-polygon`

#### Payload

Example payload with GPS coordinate in Amsterdam and a (square) polygon around Amsterdam.

**Request**
```
{
   "point":{
      "longitude":4.895168,
      "latitude":52.370216
   },
   "polygon":[
      {
         "longitude":4.851994,
         "latitude":52.394340
      },
      {
         "longitude":4.971313,
         "latitude":52.392233
      },
      {
         "longitude":4.982312,
         "latitude":52.335609
      },
      {
         "longitude":4.831007,
         "latitude":52.339999
      }
   ]
}
```

**Response**
```
{
   "pointInPolygon":true
}
```

## Google Maps URL
OpenFaas Google Maps URL generator. 

This function creates a Google Maps URL that places a marker at a specified GPS location.

### Usage
This function can be build and run locally but is also published via [OpenFaas Cloud](https://github.com/openfaas/openfaas-cloud) and available at:

`https://michaelvlaar.o6s.io/google-maps-url`

### Payload

An example payload with the GPS coordinate of Amsterdam.

**Request**
```
{
    "longitude":4.895168,
    "latitude":52.370216
}
```

**Response**
```
https://maps.google.com/?q=52.370216,4.895168
```