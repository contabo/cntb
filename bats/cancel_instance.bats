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

@test 'cancel instance : ok' {
  run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success
  instanceId="$output"

  run ./cntb cancel instance "$instanceId"
  assert_success
  assert_output --partial 'TENANTID'
  assert_output --partial 'CUSTOMERID'
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'ORDERID'

  run ./cntb cancel instance "$instanceId" -o wide
  assert_success
  assert_output --partial 'TENANTID'
  assert_output --partial 'CUSTOMERID'
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'ORDERID'

  run ./cntb cancel instance "$instanceId" -o json
  assert_success
  assert_output --partial 'tenantId'
  assert_output --partial 'customerId'
  assert_output --partial 'instanceId'
  assert_output --partial 'createdDate'
  assert_output --partial 'orderId'

  run ./cntb cancel instance "$instanceId" -o yaml
  assert_success
assert_output --partial 'tenantId:'
  assert_output --partial 'customerId:'
  assert_output --partial 'instanceId:'
  assert_output --partial 'createdDate:'
  assert_output --partial 'orderId:'


}

@test "cancel instance wrong instance id type: nok" {
 run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success

  run ./cntb cancel instance abc
  assert_failure
}

@test 'cancel instance inexistent instance id : nok' {
  run ./cntb create instance -b true -p 12 --imageId "ae423751-50fa-4bf6-9978-015673bf51c4" --addOns '[{"id":1424,"quantity":1}]' --productId "V1" -r "EU"
  assert_success
  instanceId=$((output + 10))
  
  run ./cntb cancel instance "$instanceId"
  assert_failure
}