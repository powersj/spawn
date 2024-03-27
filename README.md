# Spawn

## TODO

* Create standard config options
  * logger
  * interval

## Preflight Checks

* Only one agent section
* All generators have a unique, non-empty ID
* All serializers have a unique, non-empty ID
* Serializer template references defined generator ID
* Output references defined serializer ID
* Warning if a generator is not referenced by a serializer
* Warning if a serializer is not referenced by an output

## Outputs

* http
* socket

## Generators

* Figure out what generators to create
