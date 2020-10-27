#!/bin/bash

BASE_URL='https://blackduck.platform-services.services-platdev.x.gcpnp.anz'
AUTHENTICATE='api/tokens/authenticate'
bearerToken=$(curl --request POST --url ${BASE_URL}/${AUTHENTICATE} --header "Authorization: token $1" | jq -r '.bearerToken')

echo $bearerToken
