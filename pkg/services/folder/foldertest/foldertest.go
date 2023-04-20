package foldertest

import (
	"context"

	"github.com/grafana/grafana/pkg/services/folder"
)

type FakeService struct {
	ExpectedFolders        []*folder.Folder
	ExpectedFolder         *folder.Folder
	ExpectedError          error
	ExpectedChildrenCounts map[string]int64
}

func NewFakeService() *FakeService {
	return &FakeService{}
}

var _ folder.Service = (*FakeService)(nil)

func (s *FakeService) GetChildren(ctx context.Context, cmd *folder.GetChildrenQuery) ([]*folder.Folder, error) {
	return s.ExpectedFolders, s.ExpectedError
}

func (s *FakeService) GetParents(ctx context.Context, q folder.GetParentsQuery) ([]*folder.Folder, error) {
	return s.ExpectedFolders, s.ExpectedError
}

func (s *FakeService) Create(ctx context.Context, cmd *folder.CreateFolderCommand) (*folder.Folder, error) {
	return s.ExpectedFolder, s.ExpectedError
}
func (s *FakeService) Get(ctx context.Context, cmd *folder.GetFolderQuery) (*folder.Folder, error) {
	return s.ExpectedFolder, s.ExpectedError
}
func (s *FakeService) Update(ctx context.Context, cmd *folder.UpdateFolderCommand) (*folder.Folder, error) {
	return s.ExpectedFolder, s.ExpectedError
}
func (s *FakeService) Delete(ctx context.Context, cmd *folder.DeleteFolderCommand) error {
	return s.ExpectedError
}
func (s *FakeService) MakeUserAdmin(ctx context.Context, orgID int64, userID, folderID int64, setViewAndEditPermissions bool) error {
	return s.ExpectedError
}

func (s *FakeService) Move(ctx context.Context, cmd *folder.MoveFolderCommand) (*folder.Folder, error) {
	return s.ExpectedFolder, s.ExpectedError
}

func (s *FakeService) RegisterService(service folder.RegistryService) error {
	return s.ExpectedError
}

func (s *FakeService) GetFolderChildrenCounts(ctx context.Context, cmd *folder.GetFolderChildrenCountsQuery) (folder.FolderChildrenCounts, error) {
	return s.ExpectedChildrenCounts, s.ExpectedError
}
