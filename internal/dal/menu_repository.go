package dal

type MenuRepository struct {
	dataDir string
}

func NewMenuRepository(dataDir string) *MenuRepository {
	return &MenuRepository{
		dataDir: dataDir,
	}
}
