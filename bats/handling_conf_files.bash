#!/bin/bash

OAUTH2_CLIENT_ID="${OAUTH2_CLIENT_ID:-arcus-playground}"
OAUTH2_USER="${OAUTH2_USER:-de-98546@contabo.net}"
OAUTH2_PASS="${OAUTH2_PASS:-uog4ungoo%th0ieNg3oa}"
OAUTH2_CLIENT_SECRET="${OAUTH2_CLIENT_SECRET:-7271e905-239e-4d15-a8c5-527743a58390}"
OAUTH2_URL="${OAUTH2_URL:-https://idm-dev-int.contabo.intra/auth/realms/contabo/protocol/openid-connect/token}"
API_ENDPOINT="${API_ENDPOINT:-https://api-dev-ext.contabo.intra}"

function store_config_files() {
    echo "Storing existing conf files..." >&3
    tmp_store_dir=$(mktemp -d cntb-bat-XXXXXX)
    echo "${tmp_store_dir}" >&3
    if test -f "/etc/.cntb.yaml"; then
        mv /etc/.cntb.yaml ${tmp_store_dir}/etc.cntb.yaml
    fi
    if test -f "~/.cntb.yaml"; then
        mv ~/.cntb.yaml ${tmp_store_dir}/home.cntb.yaml
    fi
}

function restore_config_files() {
    echo "Restoring existing conf files..." >&3
    if test -f "${tmp_store_dir}/etc.cntb.yaml"; then
        mv -f ${tmp_store_dir}/etc.cntb.yaml /etc/.cntb.yaml
    fi
    if test -f "${tmp_store_dir}/home.cntb.yaml"; then
        mv -f ${tmp_store_dir}/home.cntb.yaml ~/.cntb.yaml
    fi
    rm -rf ${tmp_store_dir}
}

function deleteCache() {
    rm -f ~/.cache/cntb/token
}

function ensureTestConfig() {
    cat << EOF > ~/.cntb.yaml
debug: warn
oauth2-tokenurl: ${OAUTH2_URL}
oauth2-clientid: ${OAUTH2_CLIENT_ID}
oauth2-client-secret: ${OAUTH2_CLIENT_SECRET}
oauth2-user: ${OAUTH2_USER}
oauth2-password: ${OAUTH2_PASS}
api: ${API_ENDPOINT}
EOF
}
