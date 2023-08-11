package model

func Tables() (t []interface{}) {
	t = append(t,
		new(OpenApp),
		new(CredentialsApp),
	)
	return
}
