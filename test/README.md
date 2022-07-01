# Tests

## Integration API tests with Venom

The tool running the test suite: https://github.com/ovh/venom

```
curl https://github.com/ovh/venom/releases/download/v1.0.1/venom.linux-amd64 -L -o /usr/local/bin/venom && chmod +x /usr/local/bin/venom
```
(or install as go package)

### CSV application import

Uploads sample CSV file (should be in sync with the tackle2-ui template) and ensures there are expected applications and dependencies via tackle2 API.

```
venom run api-csv-import.yml 
```