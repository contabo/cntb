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

@test "create objectStorage with auto scaling normal: ok" {
  deleteObjectStorageIfExisting "EU"

  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 1
  assert_success
  objectStorageId="$output"
}

@test "create objectStorage with scalingLimit when scalingState inactive: ok" {
  deleteObjectStorageIfExisting "EU"

  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 10
  assert_success
  objectStorageId="$output"
}

@test "create objectStorage with non existing region : nok" {
  run ./cntb create objectStorage --region "FOO" --totalPurchasedSpaceTB 1 --scalingState "disabled"
  assert_failure
}


@test "create objectStorage missing totalPurchasedSpace: nok" {
  run ./cntb create objectStorage --productId 1  --region "EU" --scalingState "enabled"
  assert_failure
}

@test "create objectStorage missing scalingLimit when scalingState active: nok" {
  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled"
  assert_failure
}
