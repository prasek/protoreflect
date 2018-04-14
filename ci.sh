#!/usr/bin/env bash
set +e

gover="$(go version | awk '{ print $3 }')"

# We don't run gofmt for devel versions because it
# changed circa 11/2017. So code that passes the gofmt
# check for other versions will fail for devel version.
# For now, just skip the check for devel versions.

# The second term removes "devel" prefix, so if the two
# strings are equal, it does not have that prefix, and
# thus this is not a devel version.
if [[ ${gover} == ${gover#devel*} ]]; then
  fmtdiff="$(gofmt -s -l ./)"
  if [[ -n "$fmtdiff" ]]; then
    gofmt -s -l ./ >&2
    echo "Run gofmt on the above files!" >&2
    exit 1
  fi
fi


RED=$'\033[31m'
GREEN=$'\033[32m'
TEXTRESET=$'\033[0m' # reset the foreground colour

REPO=github.com/jhump/protoreflect

# This helper function walks the current directory looking for directories
# holding certain files ($1 parameter), and prints their paths on standard
# output, one per line.
find_dirs() {
    find . -not \( \
         \( \
         -path './.git/*' \
         -o -path './vendor/*' \
         \) \
         -prune \
         \) -name "$1" -print0 | xargs -0n1 dirname | sort -u
}

if [ -z "$TESTDIRS" ]; then
    TESTDIRS=$(find_dirs '*_test.go')
fi

if [ -z "$VETDIRS" ]; then
    VETDIRS=$(find_dirs '*.go')
fi

declare -A TESTS_FAILED

TESTFLAGS="-v -race"

for dir in $VETDIRS; do
    echo '+ skipping go vet' "./${dir#./}"
    #go vet "${REPO}/${dir#./}"
    if [ $? != 0 ]; then
        TESTS_FAILED["${dir}"]="go vet failed"
        echo
        echo "${RED}go vet failed: $dir${TEXTRESET}"
    fi

done

DIR_ID=0
for dir in $TESTDIRS; do
    DIR_ID=$(expr $DIR_ID + 1)

    echo '+ go test' $TESTFLAGS "./${dir#./}"
    go test -i "${REPO}/${dir#./}" >& /dev/null # install dependencies, don't execute
    go test ${TESTFLAGS} "${REPO}/${dir#./}" >& test.log

    if [ $? != 0 ]; then
        cat test.log
        TESTS_FAILED["${dir}"]="${DIR_ID}"
        echo
        echo "${RED}Tests failed: $dir${TEXTRESET}"
        continue
    fi
done

echo

for FAILED in "${!TESTS_FAILED[@]}"
do
    text=${TESTS_FAILED[${FAILED}]}
    echo ${RED}${FAILED}${TEXTRESET}
    if [ "${text}" == "go vet failed" ]
    then
        echo "    "${RED}${text}${TEXTRESET}
    fi
done

echo

# if some tests fail, we want the bundlescript to fail, but we want to
# try running ALL the tests first, hence TESTS_FAILED
if [ "${#TESTS_FAILED[@]}" -gt 0 ]; then
    echo "${RED}Test failures in: ${!TESTS_FAILED[@]}${TEXTRESET}"
    echo
    exit 1
else
    echo "${GREEN}Test success${TEXTRESET}"
    echo
    true
fi
