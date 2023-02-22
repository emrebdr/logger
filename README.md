# logger

'logger' is a simple logging library for Go.

## Installation

    go get github.com/emrebdr/logger

## Usage
``` go
import(
    "github.com/emrebdr/logger"
)

logger.Log("this is a debug message")
logger.LogInfo("this is an info message")
logger.LogWarning("this is a warning message")
logger.LogError("this is an error message")
logger.LogSuccess("this is a success message")
```

## License

MIT