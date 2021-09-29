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

  run ./cntb reinstall instance "${REINSTALL_INSTANCE_ID}" -b true --imageId="${STANDARD_IMAGE_ID2}" --addOns '[{ "id": 1424, "quantity": 1}]'
  assert_success
}

@test "reinstall instance wrong instance id type: nok" {
  run ./cntb reinstall instance abc -b true --imageId="${STANDARD_IMAGE_ID2}"
  assert_failure
}

@test 'reinstall instance inexistent instance id : nok' {
  run ./cntb reinstall instance 123
  assert_failure
}

@test 'reinstall instance without image id: nok' {
  run ./cntb reinstall instance "${REINSTALL_INSTANCE_ID}" -b true
  assert_failure
}