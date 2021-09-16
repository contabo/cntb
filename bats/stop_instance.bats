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

@test 'stop instance : ok' {
  run ./cntb start instance ${INSTANCE_ID}

  sleep 2

  run ./cntb stop instance ${INSTANCE_ID}
  assert_success
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'ACTION'
  assert_output --partial "${INSTANCE_ID}"
  assert_output --partial 'stop'

  sleep 2

  run ./cntb start instance ${INSTANCE_ID}
}