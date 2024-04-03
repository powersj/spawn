# TODO

## Logging

Pick library and/or some standard

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

* Numeric
  * RandInt
  * RandInt16
  * RandInt32
  * RandInt64
  * RandUint
  * RandUint16
  * RandUint32
  * RandUint64
  * RandFloat32
  * RandFloat64
  * RandBool
  * RandComplex64
  * RandComplex128
* Geography
  * USZipCode
  * ZipCodeInternational
  * USState
  * City
  * Country
  * CountryCode
  * CurrencyCode
  * Latitude
  * Longitude
  * Address
  * GeoCoordinate
  * Mountain
  * River
  * Ocean
  * Sea
* Dates and Time
  * Timezone
  * Unix Timestamp
  * Timestamp
  * DateTime
  * PastDate
  * FutureDate
* Strings
  * PhoneNumberAreaCode
  * Email Address
  * Username
* Financial
  * CreditCardNumber
  * RoutingNumber
  * AccountNumber
* Internet
  * IPv4
  * IPv6
  * MAC Address
  * UserAgent
  * ProgrammingLanguage
  * Technology
  * UUID
* Text/Language
  * RandWord
  * RandSentence
  * RandParagraph
  * LoremIpsum
* Codes and Identifiers
  * Barcode
  * QRCode
* Biomedical
  * Gene
  * Protein
  * BiologicalProcess
* Demographic
  * Name
  * FirstName
  * LastName
  * Gender
  * Age
  * JobTitle
* Science
  * Element
  * Compound
  * ScientificNumber
* Color
  * HEXColor
