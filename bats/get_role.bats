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


@test "get api permission role with different outputs" {
    run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb get role apiPermission "$roleId"
    assert_success
    assert_output --partial 'ROLEID'
    assert_output --partial 'NAME'
    assert_output --partial "${roleId}"

    run ./cntb get role apiPermission "$roleId" -o wide
    assert_success
    assert_output --partial 'ROLEID'
    assert_output --partial 'NAME'
    assert_output --partial "${roleId}"


    run ./cntb get role apiPermission "$roleId" -o json
    assert_success
    assert_output --partial 'roleId'
    assert_output --partial 'name'
    assert_output --partial 'apiPermissions'
    assert_output --partial 'apiName'

    run ./cntb get role apiPermission "$roleId" -o yaml
    assert_success
    assert_output --partial 'roleId:'
    assert_output --partial 'name:'
    assert_output --partial 'apiPermissions:'
    assert_output --partial 'apiName:'
    # clean up
    run ./cntb delete role apiPermission "${roleId}"
    assert_success
}

@test "get role wrong permission type: nok" {
 run ./cntb get role aa 1
 assert_failure
}

@test "create role wrong number of inputs: nok" {
 run ./cntb get role
 assert_failure

 run .run ./cntb create role apiPermission apiPermission
 assert_failure

 run .run ./cntb create role apiPermission apiPermission
 assert_failure
}