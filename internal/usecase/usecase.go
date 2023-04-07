package usecase

type Usecase struct {
	StringFinderUsecase
	EmailCheckerUsecase
	IinGetterUsecase
}

func NewUsecase() *Usecase {
	return &Usecase{
		NewStringFinder(),
		NewEmailChecker(),
		NewIinGetter(),
	}
}
