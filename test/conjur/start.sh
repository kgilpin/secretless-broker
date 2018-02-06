#!/bin/bash -ex

docker-compose build
docker-compose up -d conjur
docker-compose exec conjur conjurctl wait

admin_api_key=$(docker-compose exec conjur conjurctl role retrieve-key dev:user:admin | tr -d '\r')
export CONJUR_AUTHN_API_KEY=$admin_api_key

conjur_host_port=$(docker-compose port conjur 80)
conjur_port=$(echo "$conjur_host_port" | go run ../util/parse_port.go)

rm -rf tmp
mkdir -p tmp

cat <<CONJURRC > tmp/.conjurrc
url: http://localhost:$conjur_port
account: dev
api_key: $admin_api_key
CONJURRC

docker-compose up -d secretless

sleep 2

docker-compose run \
  --rm \
  --no-deps \
  -e http_proxy=http://secretless:80 \
  test conjur variable values add db/password secret