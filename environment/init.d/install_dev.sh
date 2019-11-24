#!/usr/bin/env bash
DIR=$(cd `dirname ${BASH_SOURCE[0]}` && pwd)
ENV_DIR=$(dirname ${DIR})
ROOT_DIR=$(dirname ${ENV_DIR})

cp -Rf ${ENV_DIR}/dev/* ${ROOT_DIR}
cp -Rf ${ENV_DIR}/dev/.env ${ROOT_DIR}/.env
cp -Rf ${ENV_DIR}/dev/.config.toml ${ROOT_DIR}/configs/config.toml
