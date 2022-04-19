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

@test 'update instance display name: ok' {
  if [ ${INT_ENVIRONMENT} == 'prod' ]; then
    skip "Skip due to prod environment"
  fi

  run ./cntb create instance -p 6 --imageId "${STANDARD_IMAGE_ID}" --productId "V1" -r "EU"
  assert_success
  instanceId="$output"

  run ./cntb update instance "$instanceId" --displayName="testDisplayName"
  assert_success
}

@test "get instance wrong instance id type: nok" {
  run ./cntb update instance abc
  assert_failure
}

@test 'get instance inexistent instance id : nok' {
  run ./cntb update instance 123
  assert_failure
}