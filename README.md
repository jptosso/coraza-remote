## Directives

## Installation with OWASP Core Ruleset

1. Server setup (coraza-server.yourcompany.com:2022)
```sh
# Install server
go install github.com/jptosso/coraza-center/cmd/coraza-server@latest
# Create a new user
coraza-server user new -u admin -p secret-token --admin
# Start the server
coraza-server serve
```

2. Admin setup

```sh
# Install client
go install github.com/jptosso/coraza-center/cmd/coraza-cli@latest
# Login to the server, use the same credentials as "coraza-server user new"
coraza-cli login
# Create the project path
mkdir crs/ && cd crs/
# Initialize waf settings using the CRS template
coraza-cli init --template coreruleset some-tag-name
# Deploy the waf to the server
$ coraza-cli deploy
```

3. Coraza setup, use this as your configuration directives

```
SecRemoteAdmin coraza-server.yourcompany.com:2022 admin:secret-token
```

## Best Practices

### Store your settings under SCM
Always store your coraza projects under SCM so that you can easily roll back to previous versions. 

### Nomenclatures

Use WAF tags you can easily remember and understand.
- billing-api/dev
- erp/report-service/prod

Use user tags you can easily remember and understand.
- admin/jtosso
- admin/ti/jtosso
- billing/service

### Limit admin access

Only allow internal resources to perform admin actions like deploying a WAF config. Use `coraza-server serve --admin-acl 192.168.0.0/24` to limit admin access to a subnet.

> In case of CI deployments, use strong passwords and restrict access to the CI server.

### Using Core Ruleset


### Additional security best practices
- Create one user per production waf instance
- Don't use admin accounts on production waf instances

## Server Commands:

**serve:** Creates a coraza-center server instance. By default the system will bind 127.0.0.1:2022

**user**: User context.

- **new**: Adds a user to the system.
- **passwd**: Changes the password of a user.
- **del**: Removes a user from the system.
- **list**: Lists all users in the system.
- **show**: Shows the details of a user.

**waf**: WAF context.

- **new**: Adds a WAF to the system.
- **del**: Removes a WAF from the system.
- **useradd**: Adds a user to a WAF.
- **commit**: Commits changes to a WAF.
- **revisions**: Lists all revisions of a WAF.
- **rollback**: Rolls back to a previous revision of a WAF.