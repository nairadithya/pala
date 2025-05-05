package api

type ElectionService interface {
	New(request NewElectionRequest) error
	GetElectionInfo(electionID int) (Election, error)
}

type ElectionRepository interface {
	CreateElection(request NewElectionRequest) error
	GetElection(electionID int) (Election, error)
}

type electionService struct {
	storage ElectionRepository
}

func NewElectionService(repo ElectionRepository) ElectionService {
	return &electionService{storage: repo}
}

func (e *electionService) New(request NewElectionRequest) error {
	if request.ElectionName == 0 {
		return errors.New("election: name is required")
	}

	return e.storage.CreateElection(request)
}

func (e *electionService) GetElectionInfo(electionID int) (Election, error) {
	return e.storage.GetElection(electionID)
}
