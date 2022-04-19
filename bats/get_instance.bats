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
  instanceId="$output"

  run ./cntb get instance "$instanceId"
  assert_success
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'NAME'
  assert_output --partial 'DISPLAYNAME'
  assert_output --partial 'STATUS'
  assert_output --partial 'IMAGEID'
  assert_output --partial 'IPV4'
  assert_output --partial 'IPV6'

  run ./cntb get instance "$instanceId" -o wide
  assert_success
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'NAME'
  assert_output --partial 'DISPLAYNAME'
  assert_output --partial 'STATUS'
  assert_output --partial 'IMAGEID'
  assert_output --partial 'REGION'
  assert_output --partial 'PRODUCTID'
  assert_output --partial 'CUSTOMERID'
  assert_output --partial 'IPV4'
  assert_output --partial 'IPV6'

  run ./cntb get instance "$instanceId" -o json
  assert_success
  assert_output --partial 'instanceId'
  assert_output --partial 'name'
  assert_output --partial 'displayName'
  assert_output --partial 'status'
  assert_output --partial 'imageId'
  assert_output --partial 'region'
  assert_output --partial 'productId'
  assert_output --partial 'customerId'
  assert_output --partial 'ipv4'
  assert_output --partial 'ipv6'
  assert_output --partial 'ipConfig'


  run ./cntb get instance "$instanceId" -o yaml
  assert_success
  assert_output --partial 'instanceId:'
  assert_output --partial 'name:'
  assert_output --partial 'displayName:'
  assert_output --partial 'status:'
  assert_output --partial 'imageId:'
  assert_output --partial 'region:'
  assert_output --partial 'productId:'
  assert_output --partial 'customerId:'
  assert_output --partial 'ipv4:'
  assert_output --partial 'ipv6:'
  assert_output --partial 'ipConfig:'
}

@test "get instance wrong instance id type: nok" {
  run ./cntb get instance abc
  assert_failure
}

@test 'get instance inexistent instance id : nok' {
  run ./cntb get instance 123
  assert_failure
}