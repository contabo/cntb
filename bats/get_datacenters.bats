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

@test 'get datacenters : ok' {
  run ./cntb get datacenters
  assert_success
  assert_output --partial 'SLUG'
  assert_output --partial 'CAPABILITIES'
}

@test 'get datacenters : ok : format wide' {
  run ./cntb get datacenters -o wide
  assert_success
  assert_output --partial 'NAME'
  assert_output --partial 'SLUG'
  assert_output --partial 'CAPABILITIES'
  assert_output --partial 'S3URL'
}

@test 'get datacenters : ok : format json' {
  run ./cntb get datacenters -o json
  assert_success
  assert_output --partial '"slug"'
  assert_output --partial '"capabilities"'
}