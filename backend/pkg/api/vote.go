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
