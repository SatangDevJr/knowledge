#! /bin/bash
## CHECK if service exist

RESPONSE_SERVICE=$(curl --write-out "%{http_code}\n" \
  --silent --output /dev/null \
  -X GET http://${KONG_EP}/services/service_${KONG_HOST_DOMAIN})

echo "\nRESPONSE_SERVICE http://${KONG_EP}/services/service_${KONG_HOST_DOMAIN} :" $RESPONSE_SERVICE
## Decide to POST or PATCH
if [ "$RESPONSE_SERVICE" == "404" ]; then
  echo "Kong Service Not found, Creating Service ..";
  curl  -X POST \
    http://${KONG_EP}/services \
    -H 'Content-Type: application/json' \
    -d '{
      "name": "service_${KONG_HOST_DOMAIN}",
      "host": "${KONG_HOST_PROXY}",
      "protocol": "http",
      "tags" : ["${KONG_HOST_DOMAIN}"],
      "port": 80
    }';
else
  echo "Kong Service found, Patching Service ..";
  curl  -X PATCH \
    http://${KONG_EP}/services/service_${KONG_HOST_DOMAIN} \
    -H 'Content-Type: application/json' \
    -d '{
      "host": "${KONG_HOST_PROXY}",
      "protocol": "http",
      "tags" : ["${KONG_HOST_DOMAIN}"],
      "port": 80
    }';
fi

# ## Route
RESPONSE_ROUTE=$(curl --write-out "%{http_code}\n" \
  --silent --output /dev/null \
  -X GET http://${KONG_EP}/routes/route_${KONG_HOST_DOMAIN})

echo "\nRESPONSE_ROUTE http://${KONG_EP}/routes/route_${KONG_HOST_DOMAIN} :" $RESPONSE_ROUTE

if [ "$RESPONSE_ROUTE" == "404" ]; then
  echo "Kong Route Not found, Creating Route ..";
  curl  -X POST \
    http://${KONG_EP}/routes \
    -H 'Content-Type: application/json' \
    -d '{
      "name": "route_${KONG_HOST_DOMAIN}",
      "paths": [
        "/"
      ],
      "protocols": [
        "https"
      ],
      "service": {
        "name": "service_${KONG_HOST_DOMAIN}"
      },
      "strip_path": false,
      "preserve_host": true,
      "tags" : ["${KONG_HOST_DOMAIN}"],
      "hosts": [
        "${KONG_HOST_DOMAIN}"
      ]
    }';
else
  echo "Kong Route found, Patching Route ..";
  curl  -X PATCH \
    http://${KONG_EP}/routes/route_${KONG_HOST_DOMAIN} \
    -H 'Content-Type: application/json' \
    -d '{
      "paths": [
        "/"
      ],
      "protocols": [
        "https"
      ],
      "service": {
        "name": "service_${KONG_HOST_DOMAIN}"
      },
      "strip_path": false,
      "preserve_host": true,
      "tags" : ["${KONG_HOST_DOMAIN}"],
      "hosts": [
        "${KONG_HOST_DOMAIN}"
      ]
    }';
fi
