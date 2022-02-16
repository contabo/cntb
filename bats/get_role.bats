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


@test "get api permission role with different outputs : ok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb get role "$roleId"
    assert_success
    assert_output --partial 'ROLEID'
    assert_output --partial 'NAME'
    assert_output --partial "${roleId}"

    run ./cntb get role "$roleId" -o wide
    assert_success
    assert_output --partial 'ROLEID'
    assert_output --partial 'NAME'
    assert_output --partial 'ADMIN'
    assert_output --partial 'ACCESSALLRESOURCES'
    assert_output --partial "${roleId}"


    run ./cntb get role "$roleId" -o json
    assert_success
    assert_output --partial 'roleId'
    assert_output --partial 'name'
    assert_output --partial 'admin'
    assert_output --partial 'accessAllResources'
    assert_output --partial 'permissions'
    assert_output --partial 'apiName'

    run ./cntb get role "$roleId" -o yaml
    assert_success
    assert_output --partial 'roleId:'
    assert_output --partial 'name:'
    assert_output --partial 'admin:'
    assert_output --partial 'accessAllResources:'
    assert_output --partial 'permissions:'
    assert_output --partial 'apiName:'

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}

@test "get role with wrong flags : nok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb get role
    assert_failure

    run ./cntb get role resourcePermission $roleId
    assert_failure

    run ./cntb get role $roleId apiPermission
    assert_failure

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}
