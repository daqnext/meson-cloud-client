# Develop

## notice:
- Directory(package) name: Single word, Concat words or with the subdirectory. e.g. `encodingbase64` or `encoding/base64`, not `encoding_base64` or `encodingBase64`
- Single file name: `api_server.go`
- Indentation: `tab`
- Format Tool: `goimports`
- Import group:
```go
import (
	# system
	"log"
	"time"

	# third party
	"github.com/coreservice-io/metrics"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"

	# local
	"daqnext/meson-cloud-client/api"
)
```

### Goland Import Configuration
`Settings -> Editor -> Code Style -> Go -> Imports -> Group -> Current project packages`