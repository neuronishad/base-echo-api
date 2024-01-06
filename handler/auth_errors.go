package handler

type Auth_ProviderNotAvailable struct {
	*Err
}

func providerNotAvailableErr() Auth_ProviderNotAvailable {
	return Auth_ProviderNotAvailable{
		Err: &Err{
			Msg: "provider not avialable",
		},
	}
}
