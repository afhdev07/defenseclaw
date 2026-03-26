#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

GATEWAY="${ROOT_DIR}/defenseclaw-gateway"

ENABLE_TELEMETRY=false
ENABLE_LOGS=false
ACCESS_KEY=""
REALM="us1"
SERVICE_NAME=""

usage() {
    cat <<EOF
Usage: $(basename "$0") [OPTIONS]

Start the DefenseClaw gateway, optionally sending OTel telemetry directly
to Splunk Observability Cloud (traces + metrics via OTLP HTTP) and/or
bootstrapping a local Splunk Enterprise instance for HEC logs.

Options:
  --enable-telemetry    Enable direct-to-Splunk OTel telemetry
  --access-key TOKEN    Splunk access token (required with --enable-telemetry)
  --realm REALM         Splunk ingest realm (default: us1)
  --app-name NAME       OTel service name (default: defenseclaw)
  --enable-logs         Start local Splunk bridge before gateway
  -h, --help            Show this help message

Examples:
  $(basename "$0")                                             # gateway only
  $(basename "$0") --enable-telemetry --access-key <token>     # Splunk O11y (us1)
  $(basename "$0") --enable-telemetry --access-key <token> --realm us0
  $(basename "$0") --enable-logs                               # local Splunk + gateway
  $(basename "$0") --enable-telemetry --access-key <token> --enable-logs  # both
EOF
    exit 0
}

while [[ $# -gt 0 ]]; do
    case "$1" in
        --enable-telemetry) ENABLE_TELEMETRY=true; shift ;;
        --access-key)       ACCESS_KEY="$2"; shift 2 ;;
        --realm)            REALM="$2"; shift 2 ;;
        --app-name)         SERVICE_NAME="$2"; shift 2 ;;
        --enable-logs)      ENABLE_LOGS=true; shift ;;
        -h|--help)          usage ;;
        *) echo "Unknown option: $1" >&2; usage ;;
    esac
done

cleanup() {
    echo ""
    echo "Shutting down..."
    if [[ -n "${GATEWAY_PID:-}" ]]; then
        kill "$GATEWAY_PID" 2>/dev/null || true
        wait "$GATEWAY_PID" 2>/dev/null || true
        echo "Gateway stopped"
    fi
}
trap cleanup EXIT INT TERM

if [[ "$ENABLE_LOGS" == "true" ]]; then
    BRIDGE=""
    if [[ -x "${HOME}/.defenseclaw/splunk-bridge/bin/splunk-claw-bridge" ]]; then
        BRIDGE="${HOME}/.defenseclaw/splunk-bridge/bin/splunk-claw-bridge"
    elif [[ -x "${ROOT_DIR}/bundles/splunk_local_bridge/bin/splunk-claw-bridge" ]]; then
        BRIDGE="${ROOT_DIR}/bundles/splunk_local_bridge/bin/splunk-claw-bridge"
    fi

    if [[ -z "$BRIDGE" ]]; then
        echo "Error: Splunk bridge not found. Run 'defenseclaw init' first." >&2
        exit 1
    fi

    echo "Starting local Splunk bridge..."
    "$BRIDGE" up --output text
    echo ""
fi

if [[ "$ENABLE_TELEMETRY" == "true" ]]; then
    if [[ -z "$ACCESS_KEY" ]]; then
        echo "Error: --access-key is required when --enable-telemetry is set" >&2
        exit 1
    fi

    export DEFENSECLAW_OTEL_ENABLED=true
    export SPLUNK_ACCESS_TOKEN="$ACCESS_KEY"

    export DEFENSECLAW_OTEL_TRACES_ENDPOINT="ingest.${REALM}.observability.splunkcloud.com"
    export DEFENSECLAW_OTEL_TRACES_PROTOCOL="http"
    export DEFENSECLAW_OTEL_TRACES_URL_PATH="/v2/trace/otlp"

    export DEFENSECLAW_OTEL_METRICS_ENDPOINT="ingest.${REALM}.observability.splunkcloud.com"
    export DEFENSECLAW_OTEL_METRICS_PROTOCOL="http"
    export DEFENSECLAW_OTEL_METRICS_URL_PATH="/v2/datapoint/otlp"

    if [[ -n "$SERVICE_NAME" ]]; then
        export OTEL_SERVICE_NAME="$SERVICE_NAME"
    fi

    echo "Telemetry enabled (realm=$REALM)"
    echo "  Traces  → HTTP  ingest.${REALM}.observability.splunkcloud.com/v2/trace/otlp"
    echo "  Metrics → HTTP  ingest.${REALM}.observability.splunkcloud.com/v2/datapoint/otlp"
fi

if [[ ! -x "$GATEWAY" ]]; then
    echo "Error: gateway not found at $GATEWAY" >&2
    echo "Run 'make gateway' first." >&2
    exit 1
fi

echo "Starting gateway..."
"$GATEWAY" &
GATEWAY_PID=$!

wait "$GATEWAY_PID"
