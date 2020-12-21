#! /bin/sh

PROTOC_GEN_TS_PATH="../../frontend/webapp/node_modules/.bin/protoc-gen-ts"
TS_OUT="../../frontend/webapp/src/types"
GO_OUT="../.."

protoc \
    -I=. \
    --go_out=${GO_OUT} \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --js_out="import_style=commonjs:${TS_OUT}" \
    --ts_out="${TS_OUT}" \
    deployment.proto \
    domainConfiguration.proto