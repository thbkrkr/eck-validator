# eck validator

Is my [ECK](https://github.com/elastic/cloud-on-k8s) Elasticsearch YAML manifest for valid?

```sh
> cat examples/ok.yml | eck-validator
OK
```

```sh
> cat examples/ko.yml | eck-validator
KO unknown field "type"
```
