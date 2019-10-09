### Features

- Sending message to phone number;
- Sending message to many phone numbers together;

# Description

![](https://turbosms.ua/img/logo.png)
Is a golang client for SMS Gateway service (turbosms.com.ua)

**Usage examples**

### Creating a new instance
```go
	var (
		turboSMS  *TurboSMS
		err       error
	)

	// username, password must be type string
	if turboSMS, err = New(username, password); err != nil {
		t.Fatal(err)
	}
```

### Build a message
```go
	m := NewMessageBuilder()
	message, err = m.SetSign("SignExample").
		SetBody("Message Body").
		SetDestination("+380999999111").
		Build()
```

### Send a message
```go
	var (
		turboSMS  *TurboSMS
		messageID string
		err       error
	)
    // username, password must be type string
	if turboSMS, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	// m = was build by message builder
	if messageID, err = turboSMS.SendMessage(m); err != nil {
		t.Fatal(err)
	}
```

## TODO
- ##### Balance checking functionality
- ##### Getting status of the message functionality

# End