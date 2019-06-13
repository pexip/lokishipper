# Lokishipper

A standalone Loki client slightly modified from the official promtail client from
https://github.com/grafana/loki.

See the GoDoc: https://godoc.org/github.com/1lann/lokishipper

## License
Licensed under Apache 2.0 as a derivative work of Loki, which is also licensed
under Apache 2.0.

```
Copyright 2019 Loki Authors (https://github.com/grafana/loki)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

## Changes made

- Isolated the promtail client.
- Reduced dependencies by lifting code into the package directly.
