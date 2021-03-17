## NGSI Go tutorial for Keyrock

## Identity Management

### Get settings of alias

```console
ngsi server get --host keyrock
```

```json
ngsi server get --host keyrock
serverType keyrock
serverHost http://keyrock:3000
IdmType idm
IdmHost http://keyrock:3000/v1/auth/tokens
Username keyrock@letsfiware.jp
Password 1234
```

### Get a token

```console
ngsi token --host keyrock
```

```console
871d471b-1ed2-4c80-b8f5-a4439cf26f51
```

### Create a user

```console
ngsi users create --username alice --email alice@test.com --password test --pretty
```

```json
{
  "user": {
    "id": "8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9",
    "image": "default",
    "gravatar": false,
    "enabled": true,
    "admin": false,
    "starters_tour_ended": false,
    "eidas_id": null,
    "username": "alice",
    "email": "alice@test.com",
    "salt": "94416345423767ac",
    "date_password": "2021-02-21T01:07:57.806Z"
  }
}
```

```console
ngsi users create --username bob --email bob-the-manager@test.com --password test
```

```console
9813fa3e-bc32-4466-8c97-a60a4e61735b
```

```console
ngsi users create --username charlie --email charlie-security@test.com --password test
```

```console
cb4242ca-8f8f-4162-9fc8-3a8a9dd10977
```

```console
ngsi users create --username manager1 --email manager1@test.com --password test
```

```console
86d68660-7b83-454a-b49b-b4cba1842542
```

```console
ngsi users create --username manager2 --email manager2@test.com --password test
```

```console
a66bb706-60f7-413a-8380-d2ced6e295ab
```

```console
ngsi users create --username detective1 --email detective1@test.com --password test
```

```console
fdd692b8-bc08-4b88-a189-8dce222ed770
```

```console
ngsi users create --username detective2 --email detective2@test.com --password test
```

```console
ddb88540-b7cc-4c9e-b927-a6ae65b7e36f
```

### Get all users

```console
ngsi users list --pretty
```

```json
{
  "users": [
    {
      "scope": [],
      "id": "86d68660-7b83-454a-b49b-b4cba1842542",
      "username": "manager1",
      "email": "manager1@test.com",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-21T01:10:40.000Z",
      "description": null,
      "website": null
    },
    {
      "scope": [],
      "id": "8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9",
      "username": "alice",
      "email": "alice@test.com",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-21T01:07:57.000Z",
      "description": null,
      "website": null
    },
    {
      "scope": [],
      "id": "9813fa3e-bc32-4466-8c97-a60a4e61735b",
      "username": "bob",
      "email": "bob-the-manager@test.com",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-21T01:09:21.000Z",
      "description": null,
      "website": null
    },
    {
      "scope": [],
      "id": "a66bb706-60f7-413a-8380-d2ced6e295ab",
      "username": "manager2",
      "email": "manager2@test.com",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-21T01:10:46.000Z",
      "description": null,
      "website": null
    },
    {
      "scope": [],
      "id": "admin",
      "username": "admin",
      "email": "keyrock@letsfiware.jp",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-21T01:05:37.000Z",
      "description": null,
      "website": null
    },
    {
      "scope": [],
      "id": "cb4242ca-8f8f-4162-9fc8-3a8a9dd10977",
      "username": "charlie",
      "email": "charlie-security@test.com",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-21T01:09:57.000Z",
      "description": null,
      "website": null
    },
    {
      "scope": [],
      "id": "ddb88540-b7cc-4c9e-b927-a6ae65b7e36f",
      "username": "detective2",
      "email": "detective2@test.com",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-21T01:11:15.000Z",
      "description": null,
      "website": null
    },
    {
      "scope": [],
      "id": "fdd692b8-bc08-4b88-a189-8dce222ed770",
      "username": "detective1",
      "email": "detective1@test.com",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-21T01:11:09.000Z",
      "description": null,
      "website": null
    }
  ]
}
```

### Get a user

```console
ngsi users get --uid 8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9 --pretty
```

```json
{
  "user": {
    "scope": [],
    "id": "8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9",
    "username": "alice",
    "email": "alice@test.com",
    "enabled": true,
    "admin": false,
    "image": "default",
    "gravatar": false,
    "date_password": "2021-02-21T01:07:57.000Z",
    "description": null,
    "website": null
  }
}
```

### Update a user

```console
ngsi users update --uid 8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9 --description "Alice works for FIWARE"
```

```json
{
  "values_updated": {
    "description": "Alice works for FIWARE"
  }
}
```

```console
ngsi users get --uid 8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9 --pretty
```

```json
{
  "user": {
    "scope": [],
    "id": "8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9",
    "username": "alice",
    "email": "alice@test.com",
    "enabled": true,
    "admin": false,
    "image": "default",
    "gravatar": false,
    "date_password": "2021-02-21T01:07:57.000Z",
    "description": "Alice works for FIWARE",
    "website": null
  }
}
```

### Create an organization

```console
ngsi organizations create --name Security --description "This group is for the store detectives"
```

```console
c2ad7373-a166-4b6f-9f51-a9af7ecf919e
```

#### Get an organization

```console
ngsi organizations get --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e --pretty
```

```json
{
  "organization": {
    "id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
    "name": "Security",
    "description": "This group is for the store detectives",
    "website": null,
    "image": "default"
  }
}
```

### List all organizations

```console
ngsi organizations list --pretty
```

```json
{
  "organizations": [
    {
      "role": "owner",
      "Organization": {
        "id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
        "name": "Security",
        "description": "This group is for the store detectives",
        "image": "default",
        "website": null
      }
    }
  ]
}
```

### Update an organization

```
ngsi organizations update --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e \
  --name "FIWARE Security" \
  --description "The FIWARE Foundation is the .." \
  --website "https://fiware.org"
```

```console
{
  "values_updated": {
    "name": "FIWARE Security",
    "description": "The FIWARE Foundation is the ..",
    "website": "https://fiware.org"
  }
}
```

```console
ngsi organizations get --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e --pretty
```

```json
{
  "organization": {
    "id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
    "name": "FIWARE Security",
    "description": "The FIWARE Foundation is the ..",
    "website": "https://fiware.org",
    "image": "default"
  }
}
```

### Add a user as a member of an organization

```console
ngsi organizations users --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e list
```

```console
admin
```

ngsi organizations users --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e \
  create --uid 8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9 --orid member --pretty

```json
{
  "user_organization_assignments": {
    "role": "member",
    "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
    "user_id": "8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9"
  }
}
```

```console
ngsi organizations users --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e list --pretty
```

```json
{
  "organization_users": [
    {
      "user_id": "admin",
      "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
      "role": "owner"
    },
    {
      "user_id": "8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9",
      "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
      "role": "member"
    }
  ]
}
```

### Add a user as an owner of an organization

```console
ngsi organizations users --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e \
  add --uid 9813fa3e-bc32-4466-8c97-a60a4e61735b --orid owner --pretty

```json
{
  "user_organization_assignments": {
    "role": "owner",
    "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
    "user_id": "9813fa3e-bc32-4466-8c97-a60a4e61735b"
  }
}
```

```console
ngsi organizations users --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e list --pretty
```

```json
{
  "organization_users": [
    {
      "user_id": "admin",
      "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
      "role": "owner"
    },
    {
      "user_id": "8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9",
      "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
      "role": "member"
    },
    {
      "user_id": "9813fa3e-bc32-4466-8c97-a60a4e61735b",
      "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
      "role": "owner"
    }
  ]
}
```

### Read user roles within an organizatioN

```console
ngsi organizations users --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e \
  get --uid 9813fa3e-bc32-4466-8c97-a60a4e61735b --pretty
```

```json
{
  "organization_user": {
    "user_id": "9813fa3e-bc32-4466-8c97-a60a4e61735b",
    "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
    "role": "owner"
  }
}
```

### Remove a user from an organization

```console
ngsi organizations users --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e \
  remove --uid 9813fa3e-bc32-4466-8c97-a60a4e61735b --orid owner
```

```console
ngsi organizations users --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e list --pretty
```

```json
{
  "organization_users": [
    {
      "user_id": "admin",
      "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
      "role": "owner"
    },
    {
      "user_id": "8f0f155a-3ea1-4871-99ea-a1eb61bbf9f9",
      "organization_id": "c2ad7373-a166-4b6f-9f51-a9af7ecf919e",
      "role": "member"
    }
  ]
}
```

### Delete an organizatioN

```console
ngsi organizations delete --oid c2ad7373-a166-4b6f-9f51-a9af7ecf919e
```

```console
ngsi organizations list
```

```console
Organizations not found
```
