#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
    store_config_files
    ensureTestConfig
}

function teardown_file() {
    restore_config_files
}


@test "create role without resources : ok" {
    run ./cntb create role --name "foo${TEST_SUFFIX}" --permissions '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb get role $roleId -o yaml
    assert_success
    assert_output --partial 'admin: false'
    assert_output --partial 'accessAllResources: false'

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}

@test "create role with admin : ok" {
    run ./cntb create role --name "foo${TEST_SUFFIX}" --admin
    assert_success
    roleId="$output"

    run ./cntb get role $roleId -o yaml
    assert_success
    assert_output --partial 'admin: true'

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}

@test "create role with accessAllResources : ok" {
    run ./cntb create role --name "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]' --accessAllResources
    assert_success
    roleId="$output"

    run ./cntb get role $roleId -o yaml
    assert_success
    assert_output --partial 'accessAllResources: true'

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}


@test "create role without resources short flags : ok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}

@test "create role without resources multiple endpoints : ok" {
    run ./cntb create role --name "foo${TEST_SUFFIX}" --permissions '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}, {"apiName":"/v1/tags","actions":["READ"]}]'
    assert_success
    roleId="$output"

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}

@test "create role with resources : ok" {
    run ./cntb create tag --name "foo${TEST_SUFFIX}"
    assert_success
    tagid="$output"

    run ./cntb create role --name "foo${TEST_SUFFIX}" --permissions "[{\"apiName\" : \"/v1/users\", \"actions\": [\"READ\", \"CREATE\"], \"resources\": [$tagid]}]"
    assert_success
    roleId="$output"

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}

@test "create role with resources multiple endpoints : ok" {
    run ./cntb create tag --name "foo${TEST_SUFFIX}"
    assert_success
    tagid="$output"

    run ./cntb create role --name "foo${TEST_SUFFIX}" --permissions "[{\"apiName\" : \"/v1/users\", \"actions\": [\"READ\", \"CREATE\"], \"resources\": [$tagid]}, {\"apiName\":\"/v1/tags\",\"actions\":[\"READ\"], \"resources\": [$tagid]}]"
    assert_success
    roleId="$output"

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}


@test "create role with resources multiple endpoints only one with tag : ok" {
    run ./cntb create tag --name "foo${TEST_SUFFIX}"
    assert_success
    tagid="$output"

    run ./cntb create role --name "foo${TEST_SUFFIX}" --permissions "[{\"apiName\" : \"/v1/users\", \"actions\": [\"READ\", \"CREATE\"], \"resources\": [$tagid]},{\"apiName\":\"/v1/tags\",\"actions\":[\"READ\"]}]"
    assert_success
    roleId="$output"

    run ./cntb get role $roleId -o yaml
    assert_success
    assert_output --partial 'admin: false'
    assert_output --partial 'accessAllResources: false'

    # clean up
    run ./cntb delete role "${roleId}"
    assert_success
}

@test "create role with foreign tagId : nok : check error" {
    run ./cntb create --name "foo${TEST_SUFFIX}" role --permissions '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"], "resources": [0]}]'
    assert_failure
    assert_output --partial 'not found'
    assert_output --partial '404'
}

@test "create role without name : nok" {
    run ./cntb create role --permissions '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_failure
}

@test "create role without permissions : nok" {
    run ./cntb create role --name="foo${TEST_SUFFIX}"
    assert_failure
}

@test "create role without name and permissions: nok" {
    run ./cntb create role
    assert_failure
}

@test "create role with wrong action in permissions : nok" {
    run ./cntb create role -n foo${TEST_SUFFIX} --permissions '[{"apiName": "/v1/tags", "actions": ["UNKNOWN"]}]'
    assert_failure
}

@test "create role with unknown api endpoint : nok" {
    run ./cntb create role -n foo${TEST_SUFFIX} --permissions '[{"apiName": "/unkown/api/endpoint", "actions": ["CREATE"]}]'
    assert_failure
}

@test "create role with wrong JSON permissions : nok" {
    run ./cntb create role -n foo${TEST_SUFFIX} --permissions '["apiName": "/unkown/api/endpoint", "actions": ["CREATE"]}]'
    assert_failure
}


wrapperCreateApiPermissionFile() {
    echo "$1" | ./cntb create role -f -
}

@test "create from file api permission : ok" {
    run wrapperCreateApiPermissionFile '{"name": "foo", "permissions": [{"apiName": "/v1/users", "actions": ["READ", "CREATE"], "resources": []}]}'
    assert_success
    roleId="$output"

    # clean up
    run ./cntb delete role "$roleId"
    assert_success
}

@test "create from file missing name : nok" {
    run wrapperCreateApiPermissionFile '{"permissions": [{"apiName": "/v1/users", "actions": ["READ", "CREATE"]}]}'
    assert_failure
}
