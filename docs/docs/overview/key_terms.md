---
title: Key Terms
id: key_terms
layout: docs
description: Secretless Broker Documentation
permalink: docs/overview/key_terms.html
---

## General-Purpose Terminology

- ### Target service

  A data storage server, compute resource, or API server which contains or provides access to sensitive information. Examples: SQL database, HTTP web service, server or VM accessed via SSH.

- ### Secret

  A random piece of data which is presented in order to gain access to a target service. Presentation of a secret (sometimes two) is usually sufficient to access protected data (for legitimate or illegitimate purposes).

- ### Connection broker

  Software which negotiates connections to target services, such as web APIs and databases, without exposing the implementation details to the client.

- ### Vault

  A vault is a special type of database that is designed to store secret data. All the secret data in the vault is encrypted. The vault provides an API (programmatic interface) that clients (users and applications) can use to request access to secret data. The API requires clients to be both authenticated and authorized.

  When a client uses the API to interact with the vault, this interaction is written to a separate database called the audit. The audit can be used to search and report on historical interactions with secret data. Vaults may also be capable of automatically rotating the secrets. Or rotation may be performed by a separate application, acting through the Vault API.

## Secretless Broker-Specific Terminology

- ### <a href="/docs/reference/listeners.html">Listeners</a>

  A Listener listens on a TCP port or Unix socket for connections from clients. When a connection is received, it's handed off to a Handler for further processing.

- ### <a href="/docs/reference/providers/overview.html">Credential Providers</a>

  The Secretless Broker can obtain credentials from a variety of vaults. For each kind of vault, the Secretless Broker contains a Credential Provider which can connect to that vault and obtain credentials. By adding new Credential Providers, the Secretless Broker can be extended to support new vaults.

- ### <a href="/docs/reference/handlers/overview.html">Handlers</a>

  A Handler uses a Provider to obtain credentials to a target service. It uses the credentials to establish a connection to the target service, and then brokers the traffic between the client and the target.  Each Handler understands how to negotiate authentication with one specific type of target service, such as a MySQL database or SSH server. By adding new Handlers, the Secretless Broker can be extended to support new types of target services.

- ### <a href="/docs/reference/config-managers/overview.html">Configuration Managers</a>

  The Secretless Broker configuration defines the Listeners, Handlers, and Credential Providers the Broker should use.
  The Secretless Broker comes built-in with a couple of different Configuration Managers and a plugin API so that almost
  any configuration source can be easily integrated.
