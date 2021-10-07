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

  run ./cntb create instance -p 6 --imageId "${STANDARD_IMAGE_ID}" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success
  instanceId="$output"

  run ./cntb get instance "$instanceId"
  assert_success
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'NAME'
  assert_output --partial 'STATUS'

  run ./cntb get instance "$instanceId" -o wide
  assert_success
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'NAME'
  assert_output --partial 'STATUS'
  assert_output --partial 'IMAGEID'
  assert_output --partial 'REGION'
  assert_output --partial 'PRODUCTID'
  assert_output --partial 'CUSTOMERID'

  run ./cntb get instance "$instanceId" -o json
  assert_success
  assert_output --partial 'instanceId'
  assert_output --partial 'name'
  assert_output --partial 'status'
  assert_output --partial 'imageId'
  assert_output --partial 'region'
  assert_output --partial 'productId'
  assert_output --partial 'customerId'

  run ./cntb get instance "$instanceId" -o yaml
  assert_success
  assert_output --partial 'instanceId:'
  assert_output --partial 'name:'
  assert_output --partial 'status:'
  assert_output --partial 'imageId:'
  assert_output --partial 'region:'
  assert_output --partial 'productId:'
  assert_output --partial 'customerId:'

}

@test "get instance wrong instance id type: nok" {
  run ./cntb get instance abc
  assert_failure
}

@test 'get instance inexistent instance id : nok' {
  run ./cntb get instance 123
  assert_failure
}