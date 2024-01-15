# Email Validator

### Installation
```
go get github.com/morf1lo/email-validator
```

### Example:
```go
package main

import (
	"fmt"

	"github.com/morf1lo/email-validator"
)

func main() {
	email := "blablabla@example.lol" // Valid email format
	if err := emailvalidator.Valid(email); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email is OK")
}

```
