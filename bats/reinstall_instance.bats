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

###
### arugments
###

@test 'reinstall instance : ok' {
  if [ ${INT_ENVIRONMENT} == 'prod' ]; then
    skip "Skip due to prod environment"
  fi

  run ./cntb reinstall instance "${REINSTALL_INSTANCE_ID}" --imageId="${STANDARD_IMAGE_ID}" 
  assert_success
}

@test "reinstall instance wrong instance id type: nok" {
  run ./cntb reinstall instance abc --imageId="${STANDARD_IMAGE_ID}"
  assert_failure
}

@test 'reinstall instance inexistent instance id : nok' {
  run ./cntb reinstall instance 123
  assert_failure
}

@test 'reinstall instance without image id: nok' {
  run ./cntb reinstall instance "${REINSTALL_INSTANCE_ID}"
  assert_failure
}