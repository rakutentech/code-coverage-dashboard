export interface Coverage {
    id: number;
    org_name: string;
    repo_name: string;
    branch_name: string;
    commit_hash: string;
    commit_author: string;
    language: string;
    percentage: number;
    created_at: Date;
    updated_at: Date;
    deleted_at: Date;
}