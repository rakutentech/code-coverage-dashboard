package repositories

import (
	"net/http"
	"time"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/rakutentech/code-coverage-dashboard/app"
	"github.com/rakutentech/code-coverage-dashboard/config"
	"github.com/rakutentech/code-coverage-dashboard/models"
	"gorm.io/gorm"
)

// CoveragesRepository is the respository for the sales data
type CoveragesRepository struct {
	db *gorm.DB
}

// NewCoveragesRepository the constructor for NewCoveragesRepository
func NewCoveragesRepository() *CoveragesRepository {
	return &CoveragesRepository{
		db: app.NewDB(),
	}
}

// NewCoverage returns first sale
func (r *CoveragesRepository) NewCoverage(orgName,
	repoName,
	branchName,
	commitHash,
	commitAuthor,
	language string,
	prNumber int,
	percentage float64) (*models.Coverage, error) {
	now := time.Now()
	query := `org_name = ?
			AND repo_name = ?
			AND branch_name = ?
			AND language = ?
			AND deleted_at IS NULL`
	deleted := r.db.Model(&models.Coverage{}).Where(query, orgName, repoName, branchName, language).Update("deleted_at", &now)
	if deleted.Error != nil {
		return nil, deleted.Error
	}

	coverage := &models.Coverage{
		OrgName:      orgName,
		RepoName:     repoName,
		BranchName:   branchName,
		CommitHash:   commitHash,
		PRNumber:     prNumber,
		CommitAuthor: commitAuthor,
		Language:     language,
		Percentage:   percentage,
		DeletedAt:    nil,
	}

	updated := r.db.Model(&coverage).Create(coverage)
	if updated.Error != nil {
		return nil, updated.Error
	}

	return coverage, nil
}

// PaginateCoverages...
func (r *CoveragesRepository) PaginateCoverages(request *http.Request, orgName string, repoName string, commitAuthor string, keyword string, full bool, perPage int) (*pagination.Paginator, []models.Coverage, error) {
	c := config.NewConfig()
	var coverages = []models.Coverage{}
	query := ""
	if full {
		query += `1 = 1`
	} else {
		query += `deleted_at IS NULL`
	}
	if orgName != "" {
		query += ` AND org_name = @orgName`
	} else {
		query += ` AND org_name != @orgName`
	}
	if repoName != "" {
		query += ` AND repo_name = @repoName`
	} else {
		query += ` AND repo_name != @repoName`
	}
	if commitAuthor != "" {
		query += ` AND commit_author = @commitAuthor`
	} else {
		query += ` AND commit_author != @commitAuthor`
	}
	if keyword != "" {
		if c.DBConfig.DBConnection == "sqlite" {
			query += ` AND org_name || '/' || repo_name LIKE @keyword`
		} else {
			// mysql
			query += ` AND CONCAT(org_name, '/', repo_name) LIKE @keyword`
		}
	}

	var total int64
	var paginator = &pagination.Paginator{}
	var queryMap = map[string]interface{}{
		"orgName":      orgName,
		"repoName":     repoName,
		"commitAuthor": commitAuthor,
		"keyword":      "%" + keyword + "%",
	}
	r.db.Where(query, queryMap).Find(&coverages).Count(&total)

	// the implementation is is incomplete, this should paginate by org name instead of branches per repo
	paginator = pagination.NewPaginator(request, perPage, total)
	offset := paginator.Offset()

	r.db.Limit(perPage).Order("created_at asc").Offset(offset).Where(query, queryMap).Find(&coverages)
	return paginator, coverages, nil
}

// ListCoveragesByOrgRepo...
func (r *CoveragesRepository) ListCoveragesByOrgRepo(orgName, repoName string) ([]models.Coverage, error) {
	var coverages = []models.Coverage{}
	query := `org_name = ?
		AND repo_name = ?
		AND deleted_at IS NULL`
	found := r.db.Where(query, orgName, repoName).Find(&coverages)
	if found.Error != nil {
		return nil, found.Error
	}
	return coverages, nil
}

// FindCoverage...
func (r *CoveragesRepository) FindCoverage(orgName, repoName, branchName, lanauge string) (models.Coverage, error) {

	var coverage = models.Coverage{}
	query := `org_name = ?
		AND repo_name = ?
		AND branch_name = ?
		AND language = ?
		AND deleted_at IS NULL`

	found := r.db.Where(query, orgName, repoName, branchName, lanauge).First(&coverage)
	if found.Error != nil {
		return coverage, found.Error
	}
	return coverage, nil
}

// SyncBranches
func (r *CoveragesRepository) DeleteStaleBranches(orgName, repoName string, branches []string) ([]models.Coverage, error) {
	var coverages = []models.Coverage{}
	query := `org_name = ?
		AND repo_name = ?
		AND branch_name NOT IN (?)`
	found := r.db.Where(query, orgName, repoName, branches).Find(&coverages)
	if found.Error != nil {
		return nil, found.Error
	}
	deleted := r.db.Where(query, orgName, repoName, branches).Delete(&models.Coverage{})
	if deleted.Error != nil {
		return nil, deleted.Error
	}
	return coverages, nil
}
