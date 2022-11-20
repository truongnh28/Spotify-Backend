#!/usr/bin/env bash
set -e

COVERAGE_FILE="coverage.out"
RUN_COVERAGE=false
DISPLAY_COVERAGE=false

# Default path
SOURCE_DIRECTORY='.'
TEST_LOG_PATH="./test-logs"
REPORT_PATH="./test-reports"

GO_TEST="go test -v -covermode=count $(go list ./...)"

ENABLE_TIMESTAMP=true
log_msg() {
    local prefix="$1"
    local msg="$2"
    local output="[$prefix] $msg"
    test $ENABLE_TIMESTAMP && output=$(date "+%F %T")" $output"
    echo "$output"
}

log_info() {
    local prefix='INFO'
    local msg="$1"
    log_msg "$prefix" "$msg"
}

log_error() {
    local prefix='ERROR'
    local msg="$1"
    >&2 log_msg "$prefix" "$msg"
}

die() {
    log_error "$1"
    exit 1
}

boolean() {
    case $1 in
    TRUE | true) echo true ;;
    FALSE | false) echo false ;;
    *)
        die "unknown boolean value \"$1\"" 1>&2
        ;;
    esac
}

file_name_validate() {
    local file_name=$1
    if [[ -z $1 ]] || [[ ${file_name:0:1} == "-" ]] || grep '\\' <<< $1 ; then
        die "unvalid filename \"$1\"" 1>&2
    else
        echo $1
    fi
}

parse_option() {
    while [[ $# -gt 0 ]]; do
        option=$1
        shift
        case $option in
        -s | --source-directory)
            SOURCE_DIRECTORY=$1
            shift
            ;;
        --coverage)
            RUN_COVERAGE=$(boolean $1)
            shift
            ;;
        --display)
            DISPLAY_COVERAGE=$(boolean $1)
            shift
            ;;
        --html-report)
            HTML_REPORT=$(file_name_validate $1)
            shift
            ;;
        --junit-report)
            JUNIT_REPORT=$(file_name_validate $1)
            shift
            ;;
        --cobertura-report)
            COBERTURA_REPORT=$(file_name_validate $1)
            shift
            ;;
        --report-path)
            REPORT_PATH=$1
            shift
            ;;
        *)
            die "Unknown option $option"
            ;;
        esac
    done
}

# Check if option --html-report --junit-report --cobertura-report is empty return 0
# Else return true
check_report_options() {
    local REPORT_OPTION=false
    if ! [[ -z ${HTML_REPORT} ]] || ! [[ -z ${JUNIT_REPORT} ]] || ! [[ -z ${COBERTURA_REPORT} ]]; then
        REPORT_OPTION=true
    fi
    echo $REPORT_OPTION
}

run_test() {
    pushd ${SOURCE_DIRECTORY}
    if $(check_report_options) || ${RUN_COVERAGE} || ${DISPLAY_COVERAGE}; then
        COVERDIR=$(mktemp -d /tmp/coverage.XXXXXXXXXX)
        PROFILE="${COVERDIR}/${COVERAGE_FILE}"
        run_test_with_option
    else
        $GO_TEST
    fi
    popd
}

run_test_with_option() {
    if $(check_report_options); then
        run_test_and_report
    # If disable report options
    else
        $GO_TEST --coverprofile=${PROFILE}
    fi
}

run_test_and_report() {
    create_test_log_path
    $GO_TEST --coverprofile=${PROFILE} | tee ${TEST_LOG_PATH}/test.log
    cp ${COVERDIR}/coverage.out ${TEST_LOG_PATH}
    (grep -v "/mocks\|/test") <${TEST_LOG_PATH}/coverage.out >${TEST_LOG_PATH}/coverage-profile.out

    run_report
}

create_test_log_path() {
    mkdir -p ${TEST_LOG_PATH}
    mkdir -p ${REPORT_PATH}
}

run_report() {
    get_html_report
    get_junit_test_report
    get_cobertura_report
}

get_html_report() {
    if ! [[ -z ${HTML_REPORT} ]]; then
        log_info "HTML coverage report"
        touch ${REPORT_PATH}/${HTML_REPORT}
        go tool cover -html ${TEST_LOG_PATH}/coverage-profile.out -o ${REPORT_PATH}/${HTML_REPORT}
    fi
}

get_junit_test_report() {
    if ! [[ -z ${JUNIT_REPORT} ]]; then
        log_info "Unit test report"
        touch ${REPORT_PATH}/${JUNIT_REPORT}
        cat ${TEST_LOG_PATH}/test.log | go run github.com/jstemmer/go-junit-report >${REPORT_PATH}/${JUNIT_REPORT}
    fi
}

get_cobertura_report() {
    if ! [[ -z ${COBERTURA_REPORT} ]]; then
        log_info "Cobertura report"
        touch ${REPORT_PATH}/${COBERTURA_REPORT}
        go run github.com/boumenot/gocover-cobertura <${TEST_LOG_PATH}/coverage-profile.out >${REPORT_PATH}/${COBERTURA_REPORT}
    fi
}

display_coverage() {
    if ${DISPLAY_COVERAGE}; then
        go tool cover -func ${PROFILE}
    fi
}

main() {
    parse_option $@
    run_test
    display_coverage
}

main $@
