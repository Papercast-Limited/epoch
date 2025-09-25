# Traefik Plugin: Epoch Middleware

It can either always respond with JSON or only intercept requests on a given path when passthrough is enabled.

## Example Responses

**Milliseconds (default):**
```json
{ "epoch": 1758711130123 }
```
**Seconds**
```json
{ "epoch": 1758711130 }
```
**Nanoseconds**
```json
{ "epoch": 1758711130123456789 }
```
**All**
```json
{
  "epoch": 1758711130123,
  "epoch_s": 1758711130,
  "epoch_ns": 1758711130123456789,
  "rfc3339": "2025-09-24T15:04:05Z"
}
```
| Name         | Type   | Default  | Description                                                                 |
|--------------|--------|----------|-----------------------------------------------------------------------------|
| passthrough  | bool   | false    | If true, only intercepts requests that match `matchPath`. Otherwise always responds. |
| matchPath    | string | /epoch   | Path prefix to intercept when passthrough is enabled.                       |
| keyName      | string | epoch    | JSON key name when returning a single format. Ignored for `all`.            |
| format       | string | epoch    | Output format. Options: `epoch` (ms), `epoch_s` (s), `epoch_ns` (ns), `rfc3339`, `all`. |
