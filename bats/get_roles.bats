#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
    store_config_files
    ensureTestConfig
    deleteCache
}

function teardown_file() {
    restore_config_files
}


@test "get roles: normal ok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb get roles
    assert_success
    assert_output --partial 'ROLEID'
    assert_output --partial 'NAME'
    assert_output --partial "${roleId}"

    run ./cntb get roles -o wide
    assert_success
    assert_output --partial 'ROLEID'
    assert_output --partial 'NAME'
    assert_output --partial 'PERMISSIONS'
    assert_output --partial 'ADMIN'
    assert_output --partial 'ACCESSALLRESOURCES'

    run ./cntb get roles -o json
    assert_success
    assert_output --partial 'roleId'
    assert_output --partial 'name'
    assert_output --partial 'admin'
    assert_output --partial 'accessAllResources'
    assert_output --partial 'permissions'
    assert_output --partial 'apiName'

    run ./cntb get roles -o yaml
    assert_success
    assert_output --partial 'roleId:'
    assert_output --partial 'name:'
    assert_output --partial 'admin:'
    assert_output --partial 'accessAllResources:'
    assert_output --partial 'permissions:'
    assert_output --partial 'apiName:'

    # clean uo
    run ./cntb delete role "$roleId"
}
