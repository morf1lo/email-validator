# Email Validator

### Example:
#### First you need to install:
```
go get github.com/morf1lo/email-validator
```
---
```
email := "blablabla@example.lol"
if err := emailvalidator.Valid(email); err != nil {
	fmt.Println(err)
	return
}
fmt.Println("Email is OK")
```
