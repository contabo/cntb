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
  run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success
#   instanceid="$output"
}

@test 'create instance : nok : missing arguments' {
  run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' -r "EU"
  assert_failure
}

@test 'create : nok : invalid arguments' {
  run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":aba,"quantity":1}]' --productId "V1" -r "EU"
  assert_failure
}
