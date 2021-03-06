package dadata

type AddressesCleaner interface {
	CleanAddresses(addresses ...string) ([]Address, error)
}

type PhonesCleaner interface {
	CleanPhones(phones ...string) ([]Phone, error)
}

type NamesCleaner interface {
	CleanNames(names ...string) ([]Name, error)
}

type EmailsCleaner interface {
	CleanEmails(emails ...string) ([]Email, error)
}

type BirthdatesCleaner interface {
	CleanBirthdates(birthdates ...string) ([]Birthdate, error)
}

/*
Public interface. Stubs it for tests.
*/
type Cleaner interface {
	AddressesCleaner
	PhonesCleaner
	NamesCleaner
	EmailsCleaner
	BirthdatesCleaner
}
