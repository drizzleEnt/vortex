package repository

type ClientRepository interface {
	AddClient()
	DeleteClient()
	UpdateClient()
	UpdateAlgorithmStatus()
}
