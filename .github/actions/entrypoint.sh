#!/bin/sh -l

BASE_URL='https://blackduck.platform-services.services-platdev.x.gcpnp.anz'
AUTHENTICATE='api/tokens/authenticate'

TOKEN="token $1"

curl --request POST --url ${BASE_URL}/${AUTHENTICATE} --header "Authorization: $TOKEN" | jq -r '.bearerToken'