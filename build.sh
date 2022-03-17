#!/bin/bash -x
usage() { echo "Usage: $0 -n <build number, default is 1> -u <build user, default is current user> -p <goproxy url, default is \"direct\">" 1>&2; exit 1; }

while getopts ":H:u:p:" o; do
    case "${o}" in
        n)
            BUILD_NUM=${OPTARG}
            ;;
        u)
            BUILD_USER=${OPTARG}
            ;;
        p)
            GOPROXY=${OPTARG}
            ;;
        *)
            usage
            ;;
    esac
done
shift $((OPTIND-1))

if [ -z "${BUILD_NUM}" ]; then
    BUILD_NUM=1
fi

if [ -z "${BUILD_USER}" ]; then
    BUILD_USER=`whoami`
fi
if [ -z "${GOPROXY}" ]; then
    GOPROXY=direct
fi

# build artifactor
make go-download build-dep fmt check-fmt lint gosec build BUILD_NUM=${BUILD_NUM} BUILD_USER=$BUILD_USER GOPROXY=${GOPROXY}
