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

@test "cancel instance wrong instance id type: nok" {
  run ./cntb cancel instance abc
  assert_failure
}

@test 'cancel instance inexistent instance id : nok' {
  run ./cntb cancel instance 123
  assert_failure
}


@test "cancel instance : nok" {

  if [ ${INT_ENVIRONMENT} == 'prod' ]; then
    skip "Skip due to prod environment"
  fi

  run ./cntb create instance -p 6 --imageId "${STANDARD_IMAGE_ID}" --productId "V1" -r "EU"
  assert_success
  instanceId="$output"
  
  run ./cntb cancel instance "$instanceId"
  assert_failure
  assert_output --partial 'Error'
  assert_output --partial '500,'
}