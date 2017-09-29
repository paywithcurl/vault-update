# Vault-update

## Latest build

```
wget https://github.com/paywithcurl/vault-update/releases/download/v1.0/vault-update
```

## Overview

The vault client does not allow for a single key in a secret to be updated, the current work around is

* Dump the existing secret to a json file
```
vault -format=json read path/to/secret | jq .data > secret.json
```
* Update the json file
* Write the secret back to vault
```
vault write path/to/secret @secret.json
```

Or if the secret doesn't exist yet, you need to first write to it.

This tool allows you to all steps in a single command and not have to worry whether the secret exists yet when updating a key.

## Usage

* To update a key (or create a new secret with that key/value if not already present)
```
vault-update path/to/secret key=value
```
* To delete a key (or ignore if it doesn't exist)
```
vault-update path/to/secret --delete key
```

## Limitations

Just like reading and rewriting secrets with the official client, there is a risk of race condition if the secret changes between the read and write.

This is due to the vault internal API which only allows updating the whole secret at once.

Once vault includes a PATCH api, this tool won't be needed anymore. See [Vault Issue 1468](https://github.com/hashicorp/vault/issues/1468) to follow progress on the issue.
