package attachments

import (
	"github.com/jhphon0730/action_manager/internal/model"
	"github.com/jhphon0730/action_manager/internal/storage"
	"github.com/jhphon0730/action_manager/internal/testcases"
	"github.com/jhphon0730/action_manager/internal/testresults"
	"gorm.io/gorm"
)

type AttachmentService interface {
	Create(req CreateAttachmentRequest, projectID uint, uploadedBy uint, files []UploadedFile) error
	FindOne(targetType string, projectID, targetID, attachmentID uint) (*model.Attachment, error)
}

type attachmentService struct {
	attachmentRepo    AttachmentRepository
	testCaseService   testcases.TestCaseService
	testResultService testresults.TestResultService
	fileStorage       storage.FileStorage
}

func NewAttachmentService(attachmentRepo AttachmentRepository, testCaseService testcases.TestCaseService, testResultService testresults.TestResultService, fileStorage storage.FileStorage) AttachmentService {
	return &attachmentService{
		attachmentRepo:    attachmentRepo,
		testCaseService:   testCaseService,
		testResultService: testResultService,
		fileStorage:       fileStorage,
	}
}

func (s *attachmentService) isValidTargetType(targetType string) bool {
	return targetType == "test_case" || targetType == "test_result"
}

func (s *attachmentService) Create(req CreateAttachmentRequest, projectID uint, uploadedBy uint, files []UploadedFile) error {
	if !s.isValidTargetType(req.TargetType) {
		return ErrInvalidTargetType
	}

	switch req.TargetType {
	case "test_case":
		if testCase, err := s.testCaseService.FindOne(projectID, req.TargetID); err != nil || testCase == nil {
			return ErrNotFoundTestCase
		}
	case "test_result":
		if testResult, err := s.testResultService.FindOneByID(req.TargetID); err != nil || testResult == nil {
			return ErrNotFoundTestResult
		}
	}

	if len(files) == 0 {
		return ErrNoFilesProvided
	}

	return s.attachmentRepo.WithTx(func(tx *gorm.DB) error {
		var savedPaths []string

		for _, f := range files {
			// 이미지 저장
			path, err := s.fileStorage.Save(f.Bytes, map[string]string{
				"filename": f.Filename,
			})
			if err != nil {
				// 이미지 저장 실패 시에 이전에 저장되었던 모든 이미지 삭제
				for _, p := range savedPaths {
					_ = s.fileStorage.Delete(p)
				}
				return err
			}
			// 저장 성공 목록에 추가
			savedPaths = append(savedPaths, path)

			// DB 추가
			if err := s.attachmentRepo.Create(&model.Attachment{
				TargetType: req.TargetType,
				TargetID:   req.TargetID,
				FilePath:   path,
				UploadedBy: uploadedBy,
			}); err != nil {
				// DB 저장 실패 시에도 이미지 삭제
				for _, p := range savedPaths {
					_ = s.fileStorage.Delete(p)
				}

				return err
			}
		}
		return nil
	})
}

func (s *attachmentService) FindOne(targetType string, projectID, targetID, attachmentID uint) (*model.Attachment, error) {
	if !s.isValidTargetType(targetType) {
		return nil, ErrInvalidTargetType
	}

	switch targetType {
	case "test_case":
		if testCase, err := s.testCaseService.FindOne(projectID, targetID); err != nil || testCase == nil {
			return nil, ErrNotFoundTestCase
		}
	case "test_result":
		if testResult, err := s.testResultService.FindOneByID(targetID); err != nil || testResult == nil {
			return nil, ErrNotFoundTestResult
		}
	}

	return s.attachmentRepo.FindOne(targetType, targetID, attachmentID)
}
