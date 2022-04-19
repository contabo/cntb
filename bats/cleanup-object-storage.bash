#!/bin/bash

OAUTH2_URL="${OAUTH2_URL:-https://idm-dev-int.contabo.intra/auth/realms/contabo/protocol/openid-connect/token}"
OAUTH2_CLIENT_ID="${OAUTH2_CLIENT_ID:-arcus-playground}"
OAUTH2_CLIENT_SECRET="${OAUTH2_CLIENT_SECRET:-secret}"
USER_TO_IMPERSONATE="${USER_TO_IMPERSONATE:-internal-user-admin-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce-integration}"
KEYCLOAK_USER="${KEYCLOAK_USER:-user}"
KEYCLOAK_PASSWORD="${KEYCLOAK_PASSWORD:-pass}"
KEYCLOACK_CLIENT_ID="${KEYCLOACK_CLIENT_ID:-admin-cli}"
DELETE_URL="${DELETE_URL:-https://storage-object-storage-dev-int.contabo.intra/internal/v1/object-storages/}"


getInternalToken() {
  TOKEN=$(curl -s \
    -d "client_id=${KEYCLOACK_CLIENT_ID}" \
    -d "username=${KEYCLOAK_USER}" \
    --data-urlencode "password=${KEYCLOAK_PASSWORD}" \
    -d 'grant_type=password' -k "${OAUTH2_URL}" | jq -r '.access_token')

  DELETE_TOKEN=$(curl -s -k -X POST "${OAUTH2_URL}" \
    -H "Content-Type: application/x-www-form-urlencoded" \
    --data-urlencode "grant_type=urn:ietf:params:oauth:grant-type:token-exchange" \
    -d "client_id=${OAUTH2_CLIENT_ID}" \
    -d "client_secret=${OAUTH2_CLIENT_SECRET}" \
    -d "requested_subject=${USER_TO_IMPERSONATE}" \
    --data-urlencode "requested_token_type=urn:ietf:params:oauth:token-type:access_token" \
    -d "subject_token=${TOKEN}" | jq  -r '.access_token')

  echo "$DELETE_TOKEN"
}

deleteObjectStorage() {

  local internal_token=$(getInternalToken)
  curl -s -k -X DELETE "${DELETE_URL}$1" \
    -H "x-request-id: 8246fe75-dc7d-4aeb-9381-750e22f9c2ba" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${internal_token}" \
    -H "x-contabo-access: true" || true

  echo "deleted $1"
}

deleteObjectsInBucketIfExisting(){
  ./cntb delete object --region "$1" --bucket "$2" --path "test"
}

deleteBucketsInRegionIfExisting(){
  local buckets=$(./cntb get buckets -r "$1" -o=json)

  if [[ "$buckets" != "null" ]]; then
    local bucketnames=$(echo "$buckets" | jq -r '.[].name')

    local buckets_array=(`echo ${bucketnames}`)

    local num_of_buckets=${#buckets_array[@]}

    for (( i=0; i<${num_of_buckets}; i++ )); do
      deleteObjectsInBucketIfExisting "$1" ${buckets_array[$i]}
      ./cntb delete bucket "$1" ${buckets_array[$i]}
    done
  fi
}

getObjectStorageInRegion(){
  OBJECTSTORAGE=$(./cntb get objectStorages --regionName "$1" -o=json)
  OBJECTSORAGEID=$(echo "$OBJECTSTORAGE" | jq -r '.[].objectStorageId')
  echo "$OBJECTSORAGEID"
}

deleteObjectStorageIfExisting() {
  existingObjectStorage=$(getObjectStorageInRegion "$1")

  if  [ ! -z "$existingObjectStorage" ] ; then
      deleteBucketsInRegionIfExisting "$1"
      deleteObjectStorage "$existingObjectStorage"

      sleep 4

      deleteObjectStorageIfExisting "$1"
  else
    sleep 4
  fi
}
