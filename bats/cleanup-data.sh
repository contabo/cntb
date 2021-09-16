#!/bin/bash
source $(dirname $0)/globals.bash

set -eu -o pipefail

export IDM_USERNAME="${IDM_USERNAME:-admin}"
IDM_PASSWORD="${IDM_PASSWORD:-aitie:kaij#eghai3Thi}"
IDM_BASE_URL="${IDM_BASE_URL:-https://idm-test-int.contabo.intra}"
PG_CLUSTER="${PG_CLUSTER:-test-postgres}"


# 1. clean up database
cat <<EOF > /tmp/cleanup-cntb-testdata.sql
\c server_compute
delete from custom_image where customer_id = '98546';
refresh materialized view mv_image;
\c idm_user_management
delete from role where customer_id = '98546';
delete from "user" where customer_id = '98546';
\c idm_secret
delete from secret where customer_id = '98546';
\c organizing_tags
delete from assignment where customer_id = '98546';
delete from tag where customer_id = '98546';
EOF

scp /tmp/cleanup-cntb-testdata.sql ${PG_CLUSTER}.contabo.intra:/tmp/cleanup-cntb-testdata.sql
ssh ${PG_CLUSTER}.contabo.intra psql -h ${PG_CLUSTER}.contabo.intra -U postgres < /tmp/cleanup-cntb-testdata.sql


# 2. remove keycloak users if there are leftovers
getToken() {
    echo "... getting new token"
    TOKEN=$(curl --silent --fail -d 'client_id=admin-cli' -d "username=${IDM_USERNAME}" -d "password=${IDM_PASSWORD}" -d 'grant_type=password' "${IDM_BASE_URL}/auth/realms/master/protocol/openid-connect/token" | jq -r '.access_token')
}

deleteUser() {
    KC_USER_ID=$(curl --fail --silent -X GET -H "Content-Type:application/json" -H "Authorization: bearer ${TOKEN}" "${IDM_BASE_URL}/auth/admin/realms/contabo/users/?username=${KC_USER_NAME}&exact=true" | jq -r '.[0].id')
    STATUS_CODE=$(curl --write-out %{http_code} --silent --output /dev/null -X DELETE -H "Content-Type:application/json" -H "Authorization: bearer ${TOKEN}" "${IDM_BASE_URL}/auth/admin/realms/contabo/users/${KC_USER_ID}")
    echo "... deleted ${KC_USER_ID} with ${STATUS_CODE}"
}

getToken
KC_USER_NAME="foo.bar${TEST_SUFFIX}@contabo.com" deleteUser

