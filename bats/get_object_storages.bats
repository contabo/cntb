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

@test 'get object storages: normal ok' {
  run ./cntb get objectStorages -s 2
  assert_success
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'DATACENTER'
  assert_output --partial 'REGION'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'TOTALPURCHASEDSPACETB'
}

@test 'get object storages: wide ok' {
  run ./cntb get objectStorages -o wide -s 2
  assert_success
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'DATACENTER'
  assert_output --partial 'REGION'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'STATUS'
  assert_output --partial 'TOTALPURCHASEDSPACETB'
  assert_output --partial 'S3URL'
}
