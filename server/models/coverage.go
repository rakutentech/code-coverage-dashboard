package models

import "time"

type Coverage struct {
	ID           int64      `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	OrgName      string     `json:"org_name" gorm:"column:org_name;type:varchar(255); NOT NULL"`
	RepoName     string     `json:"repo_name" gorm:"column:repo_name;type:varchar(255); NOT NULL"`
	BranchName   string     `json:"branch_name" gorm:"column:branch_name;type:varchar(255); NOT NULL"`
	CommitHash   string     `json:"commit_hash" gorm:"column:commit_hash;type:varchar(255); NOT NULL"`
	PRNumber     int        `json:"pr_number" gorm:"column:pr_number;type:int; NOT NULL"`
	CommitAuthor string     `json:"commit_author" gorm:"column:commit_author;type:varchar(255); NOT NULL"`
	Language     string     `json:"language" gorm:"column:language;type:varchar(255); NOT NULL"`
	Percentage   float64    `json:"percentage" gorm:"column:percentage;type:varchar(255); NOT NULL"`
	CreatedAt    *time.Time `gorm:"type:timestamp null" json:"created_at,omitempty"`
	UpdatedAt    *time.Time `gorm:"type:timestamp null" json:"updated_at,omitempty"`
	DeletedAt    *time.Time `gorm:"type:timestamp null" json:"deleted_at,omitempty"`
}
