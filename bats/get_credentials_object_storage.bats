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

# on hold until credentials are moved to users api

# @test 'get object storage: normal ok' {
#   run ./cntb get objectStorage credentials
#   assert_success
#   assert_output --partial 'ACCESSKEY'
#   assert_output --partial 'SECRETKEY'
# }