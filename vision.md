# Vision

## Goals

- Color palette
  - Define
  - Share
  - Persistent / Searchable
- Automated deployment
- Cheap
- simple and easy to use interface
- mobile friendly
- CI
  - Testing / code coverage
  - code analysis (go report)

## Constraints

- 1.5 days
- must use golang, aws lambda

## Risks

- aws lambda cold start times are not stellar
- a lot of unknowns
  - best practices of this architecture
  - golang best practices (not a big deal because its a small app)
  - ?

## Non Goals (Possible ideas for next release)

- color utilities (gradient, box-shadow, etc)
- PWA ? could be useful when working in a laptop
  - color utilities work offline
  - cache prev view palettes
- Importing from image/scss/css/url
- Exporting to scss/css
- social media integration
- auth
- end to end testing of the whole stack