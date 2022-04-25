package services

import (
	"bufio"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/c4milo/unpackit"
)

// ArchiveTarName is the name of tar that will be saved on the disk
// If user uploads xx.tar.gz, yy.tar.gz or any.tar.gz
// it will be saved as archive.tar.gz
const ArchiveTarName = "archive.tar.gz"

//ForwardSlashReplace is to replace / with : for branchname
// Note  : is an illegal characters for branch name
// so a branch name will never have a : in it
// replacing it allows to create a dir for branch
const ForwardSlashReplace = "_fs_"

// UploadService is the service for Uploading coverage assets from user
type UploadService struct {
}

// NewUploadService creates a new UploadService
func NewUploadService() *UploadService {
	return &UploadService{}
}

// RemoveArchiveDirs
func (s *UploadService) RemoveArchiveDirs(assetsDir, orgName, repoName, branchName string) error {
	// replace / with ; so that dirs are created correctly
	branchName = strings.Replace(branchName, "/", ForwardSlashReplace, -1)

	log.Println("branchName: ", branchName)

	// Destination Path where tar.gz file will be saved
	dstDir := assetsDir + orgName + "/" + repoName + "/" + branchName

	// remove dir
	err := os.RemoveAll(dstDir)
	if err != nil {
		log.Print("Error: ", err)
		return err
	}
	return nil
}

// MakeArchiveDirs
func (s *UploadService) MakeArchiveDirs(assetsDir, orgName, repoName, branchName string) (string, error) {
	// replace / with ; so that dirs are created correctly
	branchName = strings.Replace(branchName, "/", ForwardSlashReplace, -1)

	log.Println("branchName: ", branchName)

	// Destination Path where tar.gz file will be saved
	dstDir := assetsDir + orgName + "/" + repoName + "/" + branchName

	log.Println("dstDir: ", dstDir)

	// remove dir
	err := os.RemoveAll(dstDir)
	if err != nil {
		log.Print("Error: ", err)
		return "", err
	}

	// make dir if not exists
	err = os.MkdirAll(dstDir, os.ModeDir|os.ModePerm)
	if err != nil {
		log.Print("Error: ", err)
		return "", err
	}
	return dstDir, nil
}

func (s *UploadService) SaveToAssets(tarFile *multipart.FileHeader, dstDir string) (string, error) {
	// Source
	src, err := tarFile.Open()
	if err != nil {
		log.Print("ERROR: ", err.Error())
		return "", err
	}
	defer src.Close()

	// make dir if not exists
	err = os.MkdirAll(dstDir, os.ModeDir|os.ModePerm)
	if err != nil {
		log.Print("Error: ", err)
		return "", err
	}
	dstPath := dstDir + "/" + ArchiveTarName

	//#nosec G304
	dst, err := os.Create(dstPath)
	if err != nil {
		log.Print("ERROR: ", err.Error())
		return "", err
	}
	wt := bufio.NewWriter(dst)

	// Copy tar to the destination
	// override if exists
	if _, err = io.Copy(dst, src); err != nil {
		log.Print("ERROR: ", err.Error())
		return "", err
	}
	err = wt.Flush()
	if err != nil {
		log.Print("ERROR: ", err.Error())
		return "", err
	}
	log.Print("INFO: File saved to: ", dstPath)
	return dstPath, nil
}

func (s *UploadService) ExtractTarGz(srcTar string) (string, error) {
	//#nosec G304
	file, err := os.Open(srcTar)
	if err != nil {
		log.Print("ERROR: ", err.Error())
		return "", err
	}
	log.Println("srcTar: ", srcTar)
	//replace ".tar.gz" with "" in a string
	// to get the name of the folder
	// where the tar.gz file will be extracted to
	// e.g. if srcTar is "a.tar.gz"
	// then extractToFolder will be "a"
	extractToFolder := strings.Replace(srcTar, ".tar.gz", "", -1)
	_, err = unpackit.Unpack(file, extractToFolder)
	if err != nil {
		log.Print("ERROR: ", err.Error())
		return "", err
	}

	return extractToFolder, nil
}
