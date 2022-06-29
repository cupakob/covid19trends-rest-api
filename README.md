# covid19trends-rest-api
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

## SonarCloud

You can find some stats for the project [here][2]

[1]: https://covid19-trends.de
[2]: https://sonarcloud.io/project/overview?id=cupakob_covid19trends-rest-api
