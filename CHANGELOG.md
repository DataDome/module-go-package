# DataDome Go Module

## 2.5.1 (2026-09-29)

- Fix an issue with the `DatadomeHandler` function

## 2.5.0 (2026-07-28)

- Introduce `GraphQLEndpoint` parameter to specify paths for the GraphQL endpoint (defaults to `/graphql`)

## 2.4.1 (2026-06-18)

- Sanitize payloads sent to Protection API

## 2.4.0 (2026-06-01)

- Remove hard-coded status codes in favor of DataDome Protection API response headers, enabling seamless support for upcoming features
- Increase length limit from 128 to 512 characters for DataDome cookie to support upcoming features
- Reduce log verbosity in `httpClient`

## 2.3.1 (2026-05-13)

- Reduce log verbosity in `client.handler`

## 2.3.0 (2026-04-28)

- Collect [Sec-Fetch-Storage-Access header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Sec-Fetch-Storage-Access) from requests

## 2.2.1 (2025-12-02)

- Remove `go-querystring` dependency and reimplement URL encoding for payloads

## 2.2.0 (2025-06-05)

- Add `CookiesList` to payloads sent to Protection API

## 2.1.0 (2025-05-12)

- Add `UseXForwardedHost` setting to support `host` override via `X-Forwarded-Host` header

## 2.0.0 (2025-03-04)

### Breaking changes

- Rename `DataDomeStruct` structure to `Client`
- Remove `DataDome` prefix on fields of the `Client` structure
- Replace sub-packages with a single `modulego` package
- Update the `NewClient` signature to use the functional options pattern
- Handle configuration errors during the client's instantiation

### General changes

- Add support of 301/302 redirections returned by the Protection API
- Add `Logger` field to the `Client` structure
- Enhance code documentation

## 1.3.0 (2024-12-18)

- Add `EnableReferrerRestoration` field to enable the referrer restoration

## 1.2.0 (2024-10-30)

- Add GraphQL support for POST requests
  - Add `EnableGraphQLSupport` field to enable GraphQL support
  - Add `MaximumBodySize` field to define maximum amount of data to read on GraphQL requests
- Add debug logs and enhance log outputs
  - Add `Debug` field to enable debug mode

## 1.1.2 (2024-08-27)

- Update `TimeRequest` value to a timestamp in microseconds without floating point to comply with the API contract
- Update inclusion/exclusion regex matching to apply to the complete URL, making configuration simpler and more secure
- Update default URL pattern exclusion regex to ensure consistent regex format across all platforms
- Update truncation limits for the data sent to the API Server

## 1.1.1 (2023-12-04)

- Fix hostname for DataDome endpoint

## 1.1.0 (2023-11-27)

- Use hostname for endpoint configuration

## 1.0.0 (2023-11-16)

- First release
