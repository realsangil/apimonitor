package repositories

import (
	"github.com/realsangil/apimonitor/models"
	"github.com/realsangil/apimonitor/pkg/rsdb"
	"github.com/realsangil/apimonitor/pkg/rsmodels"
)

type TestResultRepository interface {
	rsdb.Repository
	GetResultList(conn rsdb.Connection, test *models.WebServiceTest, request models.TestResultListRequest) (*rsmodels.PaginatedList, error)
}

type TestResultRepositoryImp struct {
	rsdb.Repository
}

func (repository *TestResultRepositoryImp) CreateTable(conn rsdb.Connection) error {
	m := &models.TestResult{}
	if err := conn.Conn().AutoMigrate(m).Error; err != nil {
		return rsdb.HandleSQLError(err)
	}
	if err := conn.Conn().Model(m).
		AddForeignKey("test_id", "web_service_tests(id)", "CASCADE", "CASCADE").Error; err != nil {
		return rsdb.HandleSQLError(err)
	}
	if err := conn.Conn().Model(m).AddIndex("idx_tested_at_status_code_is_success", "tested_at", "status_code", "is_success").Error; err != nil {
		return rsdb.HandleSQLError(err)
	}
	return nil
}

func (repository *TestResultRepositoryImp) GetResultList(
	conn rsdb.Connection, test *models.WebServiceTest, request models.TestResultListRequest) (*rsmodels.PaginatedList, error) {
	sql := conn.Conn().Table("test_results AS tr").Select("*")
	sql = sql.Joins("INNER JOIN web_service_tests AS t ON tr.web_service_test_id=t.id AND t.id=?", test.Id)

	query := rsdb.NewEmptyQuery()
	if !request.IsSuccess.IsBoth() {
		q, _ := rsdb.NewQuery("tr.is_success = ?", request.IsSuccess)
		query = query.And(q)
	}

	switch {
	case !request.StartTestedAt.IsZero():
		q, _ := rsdb.NewQuery("tr.tested_at > ?", request.StartTestedAt)
		query = query.And(q)
	case !request.EndTestedAt.IsZero():
		q, _ := rsdb.NewQuery("tr.tested_at < ?", request.EndTestedAt)
		query = query.And(q)
	}

	sql = sql.Where(query.Where(), query.Values()...)

	listFilter := rsdb.ListFilter{
		Page:       request.Page,
		NumItem:    request.NumItem,
		Conditions: nil,
	}

	var totalCount int
	if err := sql.Count(&totalCount).Error; err != nil {
		return nil, rsdb.HandleSQLError(err)
	}

	items := make([]*models.TestResult, 0)
	if err := sql.Order("tr.tested_at DESC").Offset(listFilter.Offset()).Limit(listFilter.NumItem).Find(&items).Error; err != nil {
		return nil, rsdb.HandleSQLError(err)
	}

	return &rsmodels.PaginatedList{
		CurrentPage: request.Page,
		NumItem:     request.NumItem,
		TotalCount:  totalCount,
		Items:       items,
	}, nil
}

func NewTestResultRepository() TestResultRepository {
	return &TestResultRepositoryImp{
		&rsdb.DefaultRepository{},
	}
}
