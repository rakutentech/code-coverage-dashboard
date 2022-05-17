package handlers

import (
	"log"
	"math"
	"net/http"

	"github.com/k0kubun/pp"
	echo "github.com/labstack/echo/v4"
	"github.com/rakutentech/code-coverage-dashboard/config"
	"github.com/rakutentech/code-coverage-dashboard/models"
	"github.com/rakutentech/code-coverage-dashboard/models/repositories"
	"github.com/rakutentech/code-coverage-dashboard/services"
)

// BadgeRequest for the /badge
type BadgeRequest struct {
	OrgName    string `form:"org_name" query:"org_name" json:"org_name"  validate:"required" message:"org_name is required"`
	RepoName   string `form:"repo_name" query:"repo_name" json:"repo_name"  validate:"required" message:"repo_name is required"`
	BranchName string `form:"branch_name" query:"branch_name" json:"branch_name"  validate:"required" message:"branch_name is required"`
	Language   string `form:"language" query:"language" json:"language"  validate:"required,oneof=go php js" message:"language must be either go or php or js"`
	Subtitle   string `form:"subtitle" query:"subtitle" json:"subtitle"`
}

// CoveragesPaginateRequest for the /Coverages
type CoveragesPaginateRequest struct {
	OrgName  string `form:"org_name" query:"org_name" json:"org_name" hint:"To filter by org name"`
	RepoName string `form:"repo_name" query:"repo_name" json:"repo_name" hint:"To filter by repository name"`
	Full     bool   `form:"full" query:"full" json:"full" hint:"To include all history for trends"`
	PerPage  int    `form:"per_page" query:"per_page" json:"per_page" hint:"Limit for per page"`
	Page     int64  `form:"p" query:"p" json:"p" validate:"gte=0"  message:"p greater than 0" hint:"Page number for pagination"`
}

type CoveragesPaginateResponse struct {
	Data    map[string][]models.Coverage `json:"data"`
	HasNext bool                         `json:"has_next"`
}

// CoveragesUploadRequest for the /Coverages
type CoveragesUploadRequest struct {
	OrgName             string `form:"org_name" query:"org_name" json:"org_name"  validate:"required" message:"org_name is required"`
	GithubApiURL        string `form:"github_api_url" query:"github_api_url" json:"github_api_url"  validate:"required" message:"github_api_url is required"`
	RepoName            string `form:"repo_name" query:"repo_name" json:"repo_name"  validate:"required" message:"repo_name is required"`
	BranchName          string `form:"branch_name" query:"branch_name" json:"branch_name"  validate:"required" message:"branch_name is required"`
	CommitHash          string `form:"commit_hash" query:"commit_hash" json:"commit_hash"  validate:"required" message:"commit_hash is required"`
	CommitAuthor        string `form:"commit_author" query:"commit_author" json:"commit_author"  validate:"required" message:"commit_author is required"`
	Language            string `form:"language" query:"language" json:"language"  validate:"required,oneof=go php js" message:"language must be either go or php or js"`
	CoverageXMLFileName string `form:"coverage_xml_file_name" query:"coverage_xml_file_name" json:"coverage_xml_file_name"  validate:"required" message:"please provide the coverage.xml filename"`
}

type CoveragesUploadResponse struct {
	Coverage *models.Coverage  `json:"coverage"`
	Data     []models.Coverage `json:"data"`
}

// CoveragesUpdateBranchesRequest for the /Coverages
type CoveragesUpdateBranchesRequest struct {
	OrgName        string   `form:"org_name" query:"org_name" json:"org_name"  validate:"required" message:"org_name is required"`
	GithubApiURL   string   `form:"github_api_url" query:"github_api_url" json:"github_api_url"  validate:"required" message:"github_api_url is required"`
	CommitHash     string   `form:"commit_hash" query:"commit_hash" json:"commit_hash"  validate:"required" message:"commit_hash is required"`
	RepoName       string   `form:"repo_name" query:"repo_name" json:"repo_name"  validate:"required" message:"repo_name is required"`
	ActiveBranches []string `json:"active_branches" form:"active_branches" query:"active_branches" validate:"required" message:"active_branches is required"`
}

type CoveragesSuccessResponse struct {
	Success string `json:"success"`
}

// CoveragesHandler injected with dependendencies to be used in the controller
// Passed to routes and defined methods
type CoveragesHandler struct {
	coveragesRepository *repositories.CoveragesRepository
	uploadService       *services.UploadService
	coverageService     *services.CoverageService
	validatorService    *services.ValidatorService
	githubOAuthService  *services.GithubOAuthService
}

// NewCoveragesHandler returns a new CoveragesHandler interface
func NewCoveragesHandler() CoveragesHandler {
	return CoveragesHandler{
		coveragesRepository: repositories.NewCoveragesRepository(),
		uploadService:       services.NewUploadService(),
		coverageService:     services.NewCoverageService(),
		validatorService:    services.NewValidatorService(),
		githubOAuthService:  services.NewGithubOAuthService(),
	}
}

func (h *CoveragesHandler) CoveragesUpload(c echo.Context) error {
	// --------------
	// Handle Request
	// --------------
	request := &CoveragesUploadRequest{}
	response := &CoveragesUploadResponse{}
	err := c.Bind(request)
	log.Print(pp.Sprint(request))
	if err != nil {
		log.Print("error: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate request
	msgs, err := services.ValidateRequest(request)
	if err != nil {
		log.Print("Validation Errors: ", err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	// get authorization header from context
	githubToken := c.Request().Header.Get("Authorization")
	err = h.githubOAuthService.VerifyGithubToken(request.GithubApiURL, githubToken, request.OrgName, request.RepoName, request.CommitHash)
	if err != nil {
		log.Print("Authentication Errors: ", err.Error())
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	// process request
	tarFile, err := c.FormFile("file")
	if err != nil {
		log.Print("ERROR: ", err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// ---------------
	// Process Prepare
	// ----------------
	conf := config.NewConfig()

	dstDir, err := h.uploadService.MakeArchiveDirs(conf.AppConfig.AssetsDir, request.OrgName, request.RepoName, request.BranchName)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	savedTar, err := h.uploadService.SaveToAssets(tarFile, dstDir)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	extractedToFolder, err := h.uploadService.ExtractTarGz(savedTar)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	log.Println("extractedToFolder: ", extractedToFolder)

	// ------------------
	// Calculate Coverage
	// -------------------

	coverageXMLPath, err := h.coverageService.FindCoverageXMLPath(extractedToFolder, request.CoverageXMLFileName)
	log.Println("coverageXMLPath: ", coverageXMLPath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	coveragePercentage, err := h.coverageService.ParseCoveragePercentage(coverageXMLPath)
	log.Println("coveragePercentage: ", coveragePercentage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// ------------------
	// Process DB Update
	// -------------------
	coverage, err := h.coveragesRepository.NewCoverage(
		request.OrgName,
		request.RepoName,
		request.BranchName,
		request.CommitHash,
		request.CommitAuthor,
		request.Language,
		coveragePercentage,
	)
	log.Println(pp.Sprint(coverage))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	coverages, err := h.coveragesRepository.ListCoveragesByOrgRepo(request.OrgName, request.RepoName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response.Data = coverages
	response.Coverage = coverage
	return c.JSON(http.StatusOK, response)
}

func (h *CoveragesHandler) CoveragesPaginate(c echo.Context) error {
	request := &CoveragesPaginateRequest{}
	response := &CoveragesPaginateResponse{}
	err := c.Bind(request)
	log.Print(pp.Sprint(request))
	if err != nil {
		log.Print("error: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate request
	msgs, err := services.ValidateRequest(request)
	if err != nil {
		log.Print("Validation Errors: ", err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	// process request
	// TODO: after Github Auth, verify user's org and repo and only return those results
	paginator, coverages, err := h.coveragesRepository.PaginateCoverages(c.Request(), request.OrgName, request.RepoName, request.Full, request.PerPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// map of string and array
	response.HasNext = paginator.HasNext()
	response.Data = make(map[string][]models.Coverage)
	for _, coverage := range coverages {
		// check if key exists
		key := coverage.OrgName + "/" + coverage.RepoName
		if _, ok := response.Data[key]; !ok {
			response.Data[key] = make([]models.Coverage, 0)
		}
	}
	for _, coverage := range coverages {
		key := coverage.OrgName + "/" + coverage.RepoName
		coverage.Percentage = math.Floor(coverage.Percentage*100) / 100
		response.Data[key] = append(response.Data[key], coverage)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *CoveragesHandler) CoveragesUpdateBranches(c echo.Context) error {
	request := &CoveragesUpdateBranchesRequest{}
	response := &CoveragesSuccessResponse{}

	err := c.Bind(request)
	log.Print(pp.Sprint(request))
	if err != nil {
		log.Print("error: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// validate request
	// validate request
	msgs, err := services.ValidateRequest(request)
	if err != nil {
		log.Print("Validation Errors: ", err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	// get authorization header from context
	githubToken := c.Request().Header.Get("Authorization")
	err = h.githubOAuthService.VerifyGithubToken(request.GithubApiURL, githubToken, request.OrgName, request.RepoName, request.CommitHash)
	if err != nil {
		log.Print("Authentication Errors: ", err.Error())
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	// process request
	staleCoverages, err := h.coveragesRepository.DeleteStaleBranches(request.OrgName, request.RepoName, request.ActiveBranches)
	log.Print("Stale Coverages: ", pp.Sprint(staleCoverages))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	conf := config.NewConfig()
	for _, staleCoverage := range staleCoverages {
		err := h.uploadService.RemoveArchiveDirs(conf.AppConfig.AssetsDir, staleCoverage.OrgName, staleCoverage.RepoName, staleCoverage.BranchName)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
	}
	response.Success = "1"
	return c.JSON(http.StatusOK, response)
}
