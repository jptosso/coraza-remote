
## HTTP
### GET /{apiVersion}/waf/{wafId}
Downloads WAF .tar.gz file

**Compatibility:** v1.0+

**Response codes:**

- **304:** WAF not modified, nothing to download.
- **403:** User not authorized to access WAF.
- **404:** WAF not found.
- **500:** Internal error.

### POST /{apiVersion}/waf/{wafId}
Upload WAF .tar.gz file

**Compatibility:** v1.0+

**Response codes:**

- **403:** User not authorized to access WAF.
- **404:** WAF not found.
- **500:** Internal error.

### GET /{apiVersion}/waf
Get list of waf instances. Only WAFs that the user is authorized to access are returned.

**Compatibility:** v1.0+

**Response codes:**

- **403:** User not authorized to access WAFs
- **500:** Internal error.

## Package
The downloaded package is a .tar.gz file containing all config, data, and scripts. The filename is: `coraza_{tag}_{timestamp}.tar.gz`