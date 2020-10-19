#!/bin/sh
BASE_URL='https://blackduck.platform-services.services-platdev.x.gcpnp.anz'
AUTHENTICATE='api/tokens/authenticate'

TOKEN="token $1"
COMPONENT_JSON_URL="$2"

token=$(curl --request POST --url ${BASE_URL}/${AUTHENTICATE} --header "Authorization: $TOKEN" | jq -r '.bearerToken')
curl --request GET --url ${COMPONENT_JSON_URL} --header "Authorization: Bearer $token" | tee report.json

