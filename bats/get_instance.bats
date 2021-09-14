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
 run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success

  run ./cntb get instance abc
  assert_failure
}

@test 'get instance inexistent instance id : nok' {
  run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success
  instanceId=$((output + 10))
  
  run ./cntb get instance "$instanceId"
  assert_failure
}