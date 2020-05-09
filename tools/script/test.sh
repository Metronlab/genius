#!/bin/bash -e

# The -o pipefail option is important for the trap to be executed if the "go examples" command fails
set -o pipefail

: ${TEST_RESULTS:=/tmp/test-results}
: ${COVER_RESULTS:=/tmp/cover-results}
: ${PKG:=./...}
: ${SHORT:=false}
: ${FAILFAST:=false}
: ${REPORT:=true}
: ${TIMEOUT:="60s"}
: ${RUN:=".*"}

mkdir -p ${COVER_RESULTS}
mkdir -p ${TEST_RESULTS}

if [ "$REPORT" == true ]; then
  trap "go-junit-report <${TEST_RESULTS}/go-test.out > ${TEST_RESULTS}/go-test-report.xml" EXIT
fi

failfast_flag=
if [ "$FAILFAST" == true ]; then
  failfast_flag='-failfast'
fi

go test ${PKG} -v ${failfast_flag} -short=${SHORT} -timeout ${TIMEOUT} -race -cover -covermode=atomic -coverprofile=${COVER_RESULTS}/coverage.cover -run ${RUN} \
    | tee ${TEST_RESULTS}/go-examples.out \
    | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' \
    | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' \
    | sed ''/RUN/s//$(printf "\033[34mRUN\033[0m")/''

if [ "$REPORT" == true ]; then
  go tool cover -html=${COVER_RESULTS}/coverage.cover -o ${COVER_RESULTS}/coverage.html

  echo "To open the html coverage file use one of the following commands:"
  echo "open file://$COVER_RESULTS/coverage.html on mac"
  echo "xdg-open file://$COVER_RESULTS/coverage.html on linux"
fi
