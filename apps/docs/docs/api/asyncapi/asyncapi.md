# user 0.0.0 documentation

* License: MIT
* Default content type: [application/json](https://www.iana.org/assignments/media-types/application/json)
* Support: [dicetrace.io support](https://docs.dicetrace.io/asyncapi)
* Email support: [support@dicetrace.io](mailto:support@dicetrace.io)

Add your multiline description here.

##### Specification tags

| Name | Description | Documentation |
|---|---|---|
| user | - | - |
| new | - | - |


## Table of Contents

* [Servers](#servers)
  * [dev](#dev-server)
  * [prod](#prod-server)
* [Operations](#operations)
  * [SUB user.new](#sub-usernew-operation)

## Servers

### `dev` Server

* URL: `https://kafka.dev.dicetrace.io`
* Protocol: `kafka`

Production Kafka Broker

#### Security

##### Security Requirement 1

  * security.protocol: PLAINTEXT







### `prod` Server

* URL: `https://kafka.dicetrace.io`
* Protocol: `kafka`

Production Kafka Broker

#### Security

##### Security Requirement 1

  * security.protocol: PLAINTEXT







## Operations

### SUB `user.new` Operation

* Operation ID: `user.new`

create a new user

listen

#### Message UserNew `user.new`

*Action to create a new user*

* Message ID: `user.new`
* Content type: [application/json](https://www.iana.org/assignments/media-types/application/json)

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| user | object | - | - | - | **additional properties are allowed** |
| user.username | string | - | - | - | - |
| user.email | string | - | - | format (`email`) | - |
| user.name | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "user": {
    "username": "string",
    "email": "user@example.com",
    "name": "string"
  }
}
```


##### Message tags

| Name | Description | Documentation |
|---|---|---|
| user | - | - |
| new | - | - |


