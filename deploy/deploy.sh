#!/bin/bash

readonly CHART_NAME=notifications
readonly CHART_DIR=./deploy/helm

CONSUL_ADDR=${CONSUL_ADDR:=127.0.0.1:8500}
ENV=${ENV:=dev}
VERSION=${VERSION:=`git describe --abbrev=0`-`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD`}

function log {
  local readonly level="$1"
  local readonly message="$2"
  local readonly timestamp=$(date +"%Y-%m-%d %H:%M:%S")
  >&2 echo -e "${timestamp} [${level}] [$SCRIPT_NAME] ${message}"
}

function log_info {
  local readonly message="$1"
  log "INFO" "$message"
}

function log_warn {
  local readonly message="$1"
  log "WARN" "$message"
}

function log_error {
  local readonly message="$1"
  log "ERROR" "$message"
}

function update_deps() {
    log_info "Syncing dependencies..."
    helm dependencies update --kube-context ${KUBE_CONTEXT} ${CHART_DIR}
}

function has_jq {
  [ -n "$(command -v jq)" ]
}

function has_consul {
  [ -n "$(command -v consul)" ]
}

function has_helm {
  [ -n "$(command -v helm)" ]
}

function get_vars() {
    log_info "Getting variables..."
    readonly KUBE_CONTEXT=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/common/kube_context`
    readonly MQ_URI=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/mqUri`
    readonly SENDGRID_API_KEY=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/sendgridApiKey`
    readonly CENT_API_ADDR=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/centApiAddr`
    readonly CENT_API_KEY=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/centApiKey`
    readonly SENTRY_DSN=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/sentryDsn`
}

function deploy() {
    log_info "Deploying ${CHART_NAME} version ${VERSION}"
    helm upgrade \
        --kube-context "${KUBE_CONTEXT}" \
        --install \
        --set image.tag="${VERSION}" \
        --set config.env="${ENV}" \
        --set secrets.mqUri="${MQ_URI}" \
        --set secrets.sendgridApiKey="${SENDGRID_API_KEY}" \
        --set secrets.centApiAddr="${CENT_API_ADDR}" \
        --set secrets.centApiKey="${CENT_API_KEY}" \
        --set secrets.sentryDsn="${SENTRY_DSN}" \
        --wait ${CHART_NAME} ${CHART_DIR}
}

if ! $(has_jq); then
    log_error "Could not find jq"
    exit 1
fi

if ! $(has_consul); then
    log_error "Could not find consul"
    exit 1
fi

if ! $(has_helm); then
    log_error "Could not find helm"
    exit 1
fi


get_vars
update_deps
deploy

exit $?