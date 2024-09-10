# idp

Here is a dev authorization server we can use to issue jwt tokens and expose a JWKS endpoint.

## Env file

You need a .env in the idp folder like this

```bash
POSTGRES_DB=keycloak_db
POSTGRES_USER=keycloak_db_user
POSTGRES_PASSWORD=keycloak_db_user_password
KEYCLOAK_ADMIN=admin
KEYCLOAK_ADMIN_PASSWORD=password
KEYCLOAK_HOSTNAME=auth-192-168-88-88.nip.io # or localhost running docker-compose locally
KEYCLOAK_HOSTNAME_ADMIN=auth-192-168-88-88.nip.io # or localhost running docker-compose locally
```

## Startup

Then run

```bash
docker-compose up -d
```

## Admin Console

You can access via <http://localhost:8080/admin/master/console/> if running docker-compose locally.

If you do not have docker-compose locally and are using vagrant then you can use the following assuming the vagrant host
only ip is 192.168.88.88

<http://auth-192-168-88-88.nip.io:8080/admin/master/console/>

## Getting a JWT

For testing can just use the admin-cli client with the password grant.

```bash
curl -X POST "http://auth-192-168-88-88.nip.io:8080/realms/master/protocol/openid-connect/token" \
-H "Content-Type: application/x-www-form-urlencoded" \
-d "client_id=admin-cli" \
-d "username=admin" \
-d "password=password" \
-d "grant_type=password"
```

```bash
export JWT=$(curl -X POST "http://auth-192-168-88-88.nip.io:8080/realms/master/protocol/openid-connect/token" -H "Content-Type: application/x-www-form-urlencoded" -d "client_id=admin-cli" -d "username=admin" -d "password=password" -d "grant_type=password" | jq -r .access_token)
echo $JWT
```

## Shutdown/Cleanup

```bash
docker-compose down -v
```

Note: -v will delete volume so any manual config changes will be lost
