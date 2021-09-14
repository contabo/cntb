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
  run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success
  instanceId="$output"

  run ./cntb reinstall instance "$instanceId" --imageId="ae423751-50fa-4bf6-9978-015673bf51c4"
  assert_success

}

@test "reinstall instance wrong instance id type: nok" {
 run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success

  run ./cntb reinstall instance abc --imageId="ae423751-50fa-4bf6-9978-015673bf51c4"
  assert_failure
}

@test 'reinstall instance inexistent instance id : nok' {
  run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success
  instanceId=$((output + 10))
  
  run ./cntb reinstall instance "$instanceId"
  assert_failure
}

@test 'reinstall instance without imageID: nok' {
  run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success
  instanceId="$output"

  run ./cntb reinstall instance "$instanceId"
  assert_failure
}