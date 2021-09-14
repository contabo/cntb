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

@test 'start instance : ok' {
  run ./cntb stop instance 100

  run ./cntb start instance 100
  assert_success
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'ACTION'
  assert_output --partial '100'
  assert_output --partial 'start'

  run ./cntb stop instance 100
}

@test 'start instance : ok : fromat wide' {
  run ./cntb stop instance 100

  run ./cntb start instance 100 -o wide
  assert_success
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'ACTION'
  assert_output --partial '100'
  assert_output --partial 'start'

  run ./cntb stop instance 100
}

@test 'start instance : ok : format json' {
  run ./cntb stop instance 100
  
  run ./cntb start instance 100 -o json
  assert_success
  assert_output --partial '"instanceId"'
  assert_output --partial '"action"'
  assert_output --partial '100'
  assert_output --partial 'start'

  run ./cntb stop instance 100
}