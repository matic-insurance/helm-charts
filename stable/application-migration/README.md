# application-migration

Chart that is responsible to run arbitrary jobs before deployment 
such as migration.

## Usage

Add this chart as dependency to your application chart:
```yaml
dependencies:
  - name: application-migration
    version: 0.0.0
    repository: "https://matic-insurance.github.io/helm-charts"
    alias: "migrations"
```

Chart using common global values that control application deployment:
```yaml
global:
  application:
    product: olb
    environment: "staging"
    version: "1.2.3"

  applicationImage:
    repository: "maticinsurance/olb"
```

The command to run is configured vua command parameter

```yaml
migrations:
  command: "rails db:migrate"
```

As result of above configuration helm will run command `rails db:migrate` 
in the docker image `maticinsurance/olb:1.2.3` just before rolling out new deployment 
