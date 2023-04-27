# application-config

Chart that is responsible to create secrets 
that hold common environment variables and configuration files for all application pods

## Usage

Add this chart as dependency to your application chart:
```yaml
dependencies:
  - name: application-config
    version: 0.0.0
    repository: "https://matic-insurance.github.io/helm-charts"
```

In application chart create template `config-files.yaml` with following content (More on this step read below in [Helm Files Shenanigans]() section):
```yaml
{{ include "application-helpers.config-files.tpl" . }}
```

Chart using common global values that control application deployment:
```yaml
global:
  application:
    product: olb
    environment: "staging"
    version: "1.2.3"
```

Environment variables and config files are also configured in global section:

```yaml
global:
  application:
    environmentVariables:
      RAILS_ENV: "production"
      RAILS_SERVE_ASSETS: "true"

    configFiles:
      app_settings: /app/config/app_settings.local.yaml
      variables: /app/config/database.yaml
```

As result of above configuration k8s will:
- Have 3 secrets
  - `olb-configs-environment` - Environment variables
  - `olb-configs-migration` - Configuration files for migration job
  - `olb-configs-runtime` - Configuration files for pod runtime

Every application pod (including migration) will have:
- Specified environment variables in runtime (`RAILS_ENV`, `RAILS_SERVE_ASSETS`)
- Mounted configuration files at specified location (`/app/config/app_settings.local.yaml`, `/app/config/database.yaml`)

During template rendering - chart will read configuration files from: 
- `CHART/configs/{{.Values.global.application.environment}}/app/config/app_settings.local.yaml` 
- `CHART/configs/{{.Values.global.application.environment}}/app/config/database.yaml` 

## Secret versioning during deployment

TBD

## Helm Files Shenanigans

Due to helm [security model](https://helm.sh/docs/chart_template_guide/accessing_files/)
 it is not possible to read file outside of chart folder. 
This means that sub chart like this - cannot create secrets that are contained in parent chart

In order to reduce amount of code duplication and invisible naming dependencies - 
we using [application-helpers](../application-helpers) library chart that defines common functions/templates

One of the template - renders config files listed in `.Values.global.applicaiton.configFiles`
as secrets.  

Maybe in foreseeable future - helm will allow to include additional folders that chart has access to. 
In such case - we will be able to reduce complexity of current code. 
Couple of Helm PRs are already in progress to solve this:
- [https://github.com/helm/helm/issues/3276]()
- [https://github.com/helm/helm/pull/10077]()
