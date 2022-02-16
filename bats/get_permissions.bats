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

@test "get permissions with different outputs" {
    run ./cntb get permissions
    assert_success

    assert_output --partial 'APINAME'
    assert_output --partial 'ACTIONS'

    run ./cntb get permissions -o wide
    assert_success
    assert_output --partial 'APINAME'
    assert_output --partial 'ACTIONS'

    run ./cntb get permissions -o json
    assert_success
    assert_output --partial 'apiName'
    assert_output --partial 'actions'

    run ./cntb get permissions -o yaml
    assert_success
    assert_output --partial 'apiName:'
    assert_output --partial 'actions:'
}