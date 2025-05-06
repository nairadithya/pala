package api

type VoteService interface {
	New(request NewVoteRequest) error
	GetVoteByVoter(voterID int) (Vote, error)
}

type VoteRepository interface {
	CreateVoteEntry(request NewVoteRequest) error
	GetVoteByVoter(voterID int) (Vote, error)
}

type voteService struct {
	storage VoteRepository
}

func NewVoteService(voteRepo VoteRepository) VoteService {
	return &voteService{storage: voteRepo}
}

func (w *voteService) New(request NewVoteRequest) error {
	err := w.storage.CreateVoteEntry(request)

	if err != nil {
		return err
	}

	return nil
}

func (w *voteService) GetVoteByVoter(voteID int) (Vote, error) {
	vote, err := w.storage.GetVoteByVoter(voteID)

	if err != nil {
		return Vote{}, err
	}

	return vote, nil
}
