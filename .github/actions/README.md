# Hello world docker action

This action prints "Hello World" or "Hello" + the name of a person to greet to the log.

## Inputs

### `access-token`

**Required** The name of the person to greet. Default `"default"`.

## Example usage

uses: ./
with:
  access-token: 'Mona the Octocat'
