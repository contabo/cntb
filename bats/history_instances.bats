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

@test 'get instance history : ok' {
  run ./cntb history instances
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'USERNAME'
}

@test 'get instances history : ok : format wide' {
  run ./cntb history instances -o wide
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'USERNAME'
  assert_output --partial 'CHANGEDBY'
  assert_output --partial 'REQUESTID'
  assert_output --partial 'TRACEID'
  assert_output --partial 'CHANGES'
}

@test 'get instances history : ok : format json' {
  run ./cntb history instances -o json
  assert_success
  assert_output --partial 'id'
  assert_output --partial 'instanceId'
  assert_output --partial 'action'
  assert_output --partial 'timestamp'
  assert_output --partial 'username'
  assert_output --partial 'changedBy'
  assert_output --partial 'requestId'
  assert_output --partial 'traceId'
  assert_output --partial 'changes'
}

@test 'get instances history : ok : format yaml' {
  run ./cntb history instances -o yaml
  assert_success
  assert_output --partial 'id:'
  assert_output --partial 'instanceId:'
  assert_output --partial 'action:'
  assert_output --partial 'timestamp:'
  assert_output --partial 'username:'
  assert_output --partial 'changedBy:'
  assert_output --partial 'requestId:'
  assert_output --partial 'traceId:'
  assert_output --partial 'changes:'
}


@test 'get image history : ok : filter by changedBy' {
  run ./cntb history instances --changedBy "${USER_ID}"
  assert_success
  # header
  assert_output --partial 'ID'
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'USERNAME'
}
