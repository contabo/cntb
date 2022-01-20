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

@test 'create instance : ok' {
  if [ ${INT_ENVIRONMENT} == 'prod' ]; then
    skip "Skip due to prod environment"
  fi

  run ./cntb create instance -p 6 --imageId "${STANDARD_IMAGE_ID}" --productId "V1" -r "EU"
  assert_success
}

@test 'create instance without arguments : ok' {
  if [ ${INT_ENVIRONMENT} == 'prod' ]; then
    skip "Skip due to prod environment"
  fi

  run ./cntb create instance
  assert_success
}

@test 'create instance : nok : invalid arguments' {
  run ./cntb create instance -p 6 --imageId "${STANDARD_IMAGE_ID}" --productId 10987 -r "EU"
  assert_failure
}

@test 'create instance : nok : invalid arguments -> product id' {
  run ./cntb create instance -p 6 --imageId "${STANDARD_IMAGE_ID}" --productId "v1" -r "EU"
  assert_failure
  assert_output --partial '400'
  assert_output --partial 'product ID is wrong'
}

@test 'create instance : nok : invalid env var arguments' {
  export CNTB_IMAGEID='notauuid'
  run ./cntb create instance
  assert_failure
  assert_output --partial '400'
  assert_output --partial 'imageId must be a UUID'
  unset CNTB_IMAGEID
}
