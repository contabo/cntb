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
### arguments
###

@test 'resend email verification : ok' {
  run ./cntb resendEmailVerification user "$USER_ID"
  assert_success
}

@test 'resend email verification invalid user id : nok' {
  run ./cntb resendEmailVerification user 32163657126573126535612
  assert_failure
}

@test 'resend email verification wrong user id : nok' {
  run ./cntb resendEmailVerification user 78953218-1ba4-4b9b-ae32-1c5ebf037ac0
  assert_failure
}
