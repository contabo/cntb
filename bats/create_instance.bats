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

  run ./cntb create instance -b true -p 12 --imageId "${STANDARD_IMAGE_ID}" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success
}

@test 'create instance : nok : missing arguments' {
  run ./cntb create instance -b true -p 12 --imageId "${STANDARD_IMAGE_ID}" --addOns '[{"id":1424,"quantity":1}]' -r "EU"
  assert_failure
}

@test 'create : nok : invalid arguments' {
  run ./cntb create instance -b true -p 12 --imageId "${STANDARD_IMAGE_ID}" --addOns '[{"id":aba,"quantity":1}]' --productId "V1" -r "EU"
  assert_failure
}
