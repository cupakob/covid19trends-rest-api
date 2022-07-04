# covid19trends-rest-api
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=cupakob_covid19trends-rest-api&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=cupakob_covid19trends-rest-api) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=cupakob_covid19trends-rest-api&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=cupakob_covid19trends-rest-api) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=cupakob_covid19trends-rest-api&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=cupakob_covid19trends-rest-api) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=cupakob_covid19trends-rest-api&metric=coverage)](https://sonarcloud.io/summary/new_code?id=cupakob_covid19trends-rest-api) [![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=cupakob_covid19trends-rest-api&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=cupakob_covid19trends-rest-api) [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=cupakob_covid19trends-rest-api&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=cupakob_covid19trends-rest-api) [![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=cupakob_covid19trends-rest-api&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=cupakob_covid19trends-rest-api) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=cupakob_covid19trends-rest-api&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=cupakob_covid19trends-rest-api) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=cupakob_covid19trends-rest-api&metric=bugs)](https://sonarcloud.io/summary/new_code?id=cupakob_covid19trends-rest-api)

REST API for [covid19-trends.de][1]

## Endpoints
There is only one endpoint currently

`/fetch/{countrycode}`

## Validation

The country code must be 2 chars long and uppercase, e.g.
* valid: 'US'
* invalid: 'us'

## Response

* If you give a country code, which can't be matched to a country, the response status is 404.
* If the data can't be fetched from the remote source, the status is 500
* If the country code is invalid, the response status is 400
* If okay - 200 :)

[1]: https://covid19-trends.de
[2]: https://sonarcloud.io/project/overview?id=cupakob_covid19trends-rest-api
