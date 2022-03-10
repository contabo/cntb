#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load cleanup-object-storage.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
  store_config_files
  ensureTestConfig
}

function teardown_file() {
  restore_config_files
}

@test 'get object storage history : ok' {
  run ./cntb history objectStorages
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'USERNAME'
  assert_output --partial 'TIMESTAMP'
}

@test 'get object storage history : ok : format wide' {
  run ./cntb history objectStorages -o wide
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'USERNAME'
  assert_output --partial 'CHANGEDBY'
  assert_output --partial 'REQUESTID'
  assert_output --partial 'TRACEID'
  assert_output --partial 'CHANGES'
}

@test 'get object storage history : ok : format json' {
  run ./cntb history objectStorages -o json
  assert_success
  assert_output --partial '"id"'
  assert_output --partial '"objectStorageId"'
  assert_output --partial '"action"'
  assert_output --partial '"timestamp"'
  assert_output --partial '"username"'
  assert_output --partial '"changedBy"'
  assert_output --partial '"requestId"'
  assert_output --partial '"traceId"'
  assert_output --partial '"changes"'
}

@test 'get object storage history : ok : filter by objectStorageId' {
  deleteObjectStorageIfExisting "EU"

  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 1
  assert_success
  objectStorageId="$output"

  deleteObjectStorage ${objectStorageId}

  run ./cntb history objectStorages --objectStorageId "$objectStorageId"
  assert_success

  # header
  assert_output --partial 'ID'
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
  # content
  assert_output --partial 'CREATED'
  assert_output --partial 'DELETED'
}