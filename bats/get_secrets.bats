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

@test "get secrets: ok" {
  run ./cntb create secret --name="secret${TEST_SUFFIX}" --value="Test12345!" --type="password"
  assert_success
  secretId="$output"

  run ./cntb get secrets
  assert_success
  assert_output --partial 'SECRETID'
  assert_output --partial 'NAME'
  assert_output --partial 'TYPE'

  run ./cntb get secrets -o wide
  assert_success
  assert_output --partial 'SECRETID'
  assert_output --partial 'NAME'
  assert_output --partial 'TYPE'
  assert_output --partial 'CUSTOMERID'
  assert_output --partial 'TENANTID'

  run ./cntb get secrets -o json
  assert_success
  assert_output --partial '"secretId"'
  assert_output --partial '"name"'
  assert_output --partial '"type"'
  assert_output --partial '"customerId"'
  assert_output --partial '"tenantId"'

  run ./cntb get secrets -o yaml
  assert_success
  assert_output --partial 'secretId:'
  assert_output --partial 'name:'
  assert_output --partial 'type:'
  assert_output --partial 'customerId:'
  assert_output --partial 'tenantId:'

  run ./cntb get secrets --name "secret${TEST_SUFFIX}"
  assert_success
  assert_output --partial 'SECRETID'
  assert_output --partial 'NAME'
  assert_output --partial 'TYPE'
  assert_output --partial "secret${TEST_SUFFIX}"
  assert_output --partial "$secretId"

  run ./cntb get secrets --type "password"
  assert_success
  assert_output --partial 'SECRETID'
  assert_output --partial 'NAME'
  assert_output --partial 'password'
  assert_output --partial "secret${TEST_SUFFIX}"
  assert_output --partial "$secretId"

  run ./cntb delete secret $secretId
  assert_success
}

@test "get secrets with wrong number of inputs: nok" {
  run .run ./cntb get secrets a b c
  assert_failure
}