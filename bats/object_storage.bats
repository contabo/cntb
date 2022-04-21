#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load cleanup-object-storage.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
  store_config_files
  ensureTestConfig
  deleteCache
}

function teardown_file() {
  restore_config_files
}

## Positive Tests

## create object storage
@test "create object storage with auto scaling normal: ok" {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  deleteObjectStorageIfExisting "EU"

  sleep 10

  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 1
  assert_success
  objectStorageId="$output"
}

## get object storage normal output
@test 'get object storage normal output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  objectStorageId=$(getObjectStorageInRegion "EU")

  run ./cntb get objectStorage "${objectStorageId}"
  assert_success
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'DATACENTER'
  assert_output --partial 'REGION'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'TOTALPURCHASEDSPACETB'
}

## get object storage wide output
@test 'get object storage wide output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi
  
  objectStorageId=$(getObjectStorageInRegion "EU")

  run ./cntb get objectStorage "${objectStorageId}" -o wide
  assert_success
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'DATACENTER'
  assert_output --partial 'REGION'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'STATUS'
  assert_output --partial 'TOTALPURCHASEDSPACETB'
  assert_output --partial 'S3URL'
}

## list object storages normal output
@test 'list object storages normal output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb get objectStorages -s 2
  assert_success
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'DATACENTER'
  assert_output --partial 'REGION'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'TOTALPURCHASEDSPACETB'
}

## list object storages wide output
@test 'list object storages wide output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb get objectStorages -o wide -s 2
  assert_success
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'DATACENTER'
  assert_output --partial 'REGION'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'STATUS'
  assert_output --partial 'TOTALPURCHASEDSPACETB'
  assert_output --partial 'S3URL'
}

## get object storage stats normal output
@test 'get object storage stats normal output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi
  
  objectStorageId=$(getObjectStorageInRegion "EU")

  run ./cntb stats objectStorage "${objectStorageId}"
  assert_success
  assert_output --partial 'USEDSPACETB'
  assert_output --partial 'USEDSPACEPERCENTAGE'
  assert_output --partial 'NUMBEROFOBJECTS'
}

## get object storage stats json output
@test 'get object storage stats json output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi
  
  objectStorageId=$(getObjectStorageInRegion "EU")

  run ./cntb stats objectStorage "${objectStorageId}" -o json
  assert_success
  assert_output --partial '"customerId"'
  assert_output --partial '"numberOfObjects"'
  assert_output --partial '"tenantId"'
  assert_output --partial '"usedSpacePercentage"'
  assert_output --partial '"usedSpaceTB"'
}

## get object storages audit trail normal output
@test 'get object storages history: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb history objectStorages
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'USERNAME'
  assert_output --partial 'TIMESTAMP'
}

## get object storages audit trail wide output
@test 'get object storages history wide output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

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

## get object storages audit trail json output
@test 'get object storages history json output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

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

## get object storages audit trail filter by object storage id
@test 'get object storages history filtered by object storage id: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  objectStorageId=$(getObjectStorageInRegion "EU")

  run ./cntb history objectStorages --objectStorageId "$objectStorageId"
  assert_success

  # header
  assert_output --partial 'ID'
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
  # content
  assert_output --partial 'CREATED'
}

## create bucket
@test "create bucket: ok" {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
      skip "Skip: test env has no CMS backend"
  fi

  run ./cntb create bucket EU "${TEST_BUCKET_NAME}"
  assert_success
}

## get buckets normal output
@test 'get buckets normal output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
      skip "Skip: test env has no CMS backend"
  fi

  run ./cntb get buckets
  assert_success
  assert_output --partial 'NAME'
}

## get buckets wide output
@test 'get buckets wide output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb get buckets -o wide
  assert_success

  assert_output --partial 'NAME'
  assert_output --partial 'CREATIONDATE'
}

## get buckets json output
@test 'get buckets json output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb get buckets -o json
  assert_success

  assert_output --partial '"creationDate"'
  assert_output --partial '"name"'
}

## get buckets yaml format
@test 'get buckets yaml output: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb get buckets -o yaml
  assert_success

  assert_output --partial 'creationDate:'
  assert_output --partial 'name:'
}

## get buckets in a specificed version
@test 'get buckets filtered by region: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb get buckets -r EU
  assert_success
  assert_output --partial 'NAME'
}

## create folder
@test 'create object: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb create object --region "EU" --bucket "${TEST_BUCKET_NAME}" --prefix '/test/folder'
  assert_success
}

## get folder
@test 'get object: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb get object --region "EU" --bucket "${TEST_BUCKET_NAME}" --path '/test/folder'
  assert_success
}

## get all folder
@test 'get objects: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb get objects --region "EU" --bucket "${TEST_BUCKET_NAME}"
  assert_success
  assert_output --partial 'NAME'
  assert_output --partial 'SIZE'
  assert_output --partial 'LASTMODIFIED'
}

## upload file to folder
@test 'upload file: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb create object --region "EU" --bucket "${TEST_BUCKET_NAME}" --prefix '/test/folder' --path "go.sum"
  assert_success
}

## delete file from folder
@test 'delete file: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb delete object --region "EU" --bucket "${TEST_BUCKET_NAME}" --path 'test/folder/go.sum'
  assert_success
}

## delete folder
@test 'delete object: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb delete object --region "EU" --bucket "${TEST_BUCKET_NAME}" --path 'test/folder'
  assert_success
}

## delete bucket
@test "delete bucket: ok" {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  run ./cntb delete bucket EU "${TEST_BUCKET_NAME}"
  assert_success
}

## resize object storage with lower limit
@test 'resize object storage where size limit is lower: nok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi
  
  objectStorageId=$(getObjectStorageInRegion "EU")

  run ./cntb resize objectStorage "${objectStorageId}" --totalPurchasedSpaceTB 1 --scalingLimitTB 1 --scalingState "enabled"
  assert_failure
  assert_output --partial 'Error while updating object storage: 400 - Bad Request totalPurchasedSpaceTB size limit should be bigger than the existing one.\n'
}

## resize object storage
@test 'resize object storage: ok' {
  if [ ${INT_ENVIRONMENT} == 'test' ]; then
    skip "Skip: test env has no CMS backend"
  fi

  objectStorageId=$(getObjectStorageInRegion "EU")

  run ./cntb resize objectStorage "${objectStorageId}" --totalPurchasedSpaceTB 3 --scalingLimitTB 3
  assert_success
}

## Negative Tests

@test "create object storage with non existing region: nok" {
  run ./cntb create objectStorage --region "FOO" --totalPurchasedSpaceTB 1 --scalingState "disabled"
  assert_failure
}

@test "create object storage missing totalPurchasedSpace: nok" {
  run ./cntb create objectStorage --productId 1  --region "EU" --scalingState "enabled"
  assert_failure
}

@test "create object storage missing scalingLimit when scalingState active: nok" {
  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled"
  assert_failure
}

@test 'create bucket with invalid name: nok' {
  run ./cntb create bucket EU 'T..${TEST_BUCKET_NAME}'
  assert_failure
  assert_output --partial 'Could not create bucket T..${TEST_BUCKET_NAME}. Name may contain unaccepted characters.'
}

@test 'create bucket with invalid region: nok' {
  run ./cntb create bucket DE 'T..${TEST_BUCKET_NAME}'
  assert_failure
  assert_output --partial 'No Object Storage could be found in this region.'
}

@test 'create bucket with missing region: nok' {
  run ./cntb create bucket  'T..${TEST_BUCKET_NAME}'
  assert_failure
  assert_output --partial 'Please provide a region and a bucketName'
}

@test 'create object with missing arguments: nok' {
  run ./cntb create object  --bucket ${TEST_BUCKET_NAME} --prefix '/test/${TEST_BUCKET_NAME}' --path "go.sum"
  assert_failure
  assert_output --partial 'Argument region is empty.'

  run ./cntb create object --region "EU"  --prefix '/test/${TEST_BUCKET_NAME}' --path "go.sum"
  assert_failure
  assert_output --partial 'Argument bucket is empty.'

  run ./cntb create object --region "EU" --bucket ${TEST_BUCKET_NAME}  --path "go.sum"
  assert_failure
  assert_output --partial 'Argument prefix is empty.'
}

@test 'delete object with missing arguments: nok' {
  run ./cntb delete object  --bucket ${TEST_BUCKET_NAME} --path '/test/${TEST_BUCKET_NAME}'
  assert_failure
  assert_output --partial 'Argument region is empty.'

  run ./cntb delete object --region "EU"  --path '/test/${TEST_BUCKET_NAME}'
  assert_failure
  assert_output --partial 'Argument bucket is empty.'

  run ./cntb delete object --region "EU" --bucket ${TEST_BUCKET_NAME}
  assert_failure
  assert_output --partial 'Argument path is empty.'
}

@test 'delete bucket which not exists: nok' {
  run ./cntb delete bucket EU a19b8645-79d2-4fbb-ac56-a927d69b8d2b
  assert_failure
  assert_output --partial 'The specified bucket does not exist.'
}

@test 'delete bucket with missing arguments: nok' {
  run ./cntb delete bucket EU
  assert_failure
  assert_output --partial 'Please provide a region and a bucketName'

  run ./cntb delete bucket bucket123
  assert_failure
  assert_output --partial 'Please provide a region and a bucketName'
}

@test 'delete buckets with invalid region: nok' {
  run ./cntb delete bucket USA "${TEST_BUCKET_NAME}"
  assert_failure
  assert_output --partial 'No Object Storage could be found in this region.'
}

@test 'get buckets with invalid region: nok' {
  run ./cntb get buckets -r EUO
  assert_failure
  assert_output --partial 'No Object Storage could be found in this region'
}

@test 'get object with missing arguments: nok' {
  run ./cntb get object --bucket ${TEST_BUCKET_NAME} --path ${TEST_BUCKET_NAME}
  assert_failure
  assert_output --partial 'Argument region is empty.'

  run ./cntb get object --region "EU" --path ${TEST_BUCKET_NAME}
  assert_failure
  assert_output --partial 'Argument bucket is empty.'

  run ./cntb get object --region "EU" --bucket ${TEST_BUCKET_NAME}
  assert_failure
  assert_output --partial 'Argument path is empty'
}

@test 'get objects with missing arguments: nok' {
  run ./cntb get objects --bucket ${TEST_BUCKET_NAME}
  assert_failure
  assert_output --partial 'Argument region is empty.'

  run ./cntb get objects --region "EU"
  assert_failure
  assert_output --partial 'Argument bucket is empty.'
}

@test 'get object storage stats with no object storage id: nok' {
  run ./cntb stats objectStorage
  assert_failure
}

@test 'get object storage stats with wrong object storage id: nok' {
  run ./cntb stats objectStorage 3162351f-456b-423a-b127-ab77baaa4e48
  assert_failure
}

@test 'get object storage with not an uuid object storage id: nok' {
  run ./cntb get objectStorage 100
  assert_failure
  assert_output --partial 'Error while retrieving object storage: 400 - [objectStorageId must be a UUID]'
}

@test 'resize object storage that not exists: nok' {
  run ./cntb resize objectStorage "23729f17-85d3-4f0c-b4b8-b8b90225d167" --totalPurchasedSpaceTB 40
  assert_failure
  assert_output --partial 'Error while updating object storage: 404 - Entry ObjectStorage not found by objectStorageId'
}