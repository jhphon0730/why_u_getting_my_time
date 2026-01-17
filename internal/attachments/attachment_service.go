package attachments

import (
	"github.com/jhphon0730/action_manager/internal/testcases"
	"github.com/jhphon0730/action_manager/internal/testresults"
)

type AttachmentService interface {
	Create(req CreateAttachmentRequest, projectID uint, files []string) error
}

type attachmentService struct {
	attachmentRepo    AttachmentRepository
	testCaseService   testcases.TestCaseService
	testResultService testresults.TestResultService
}

func NewAttachmentService(attachmentRepo AttachmentRepository, testCaseService testcases.TestCaseService, testResultService testresults.TestResultService) AttachmentService {
	return &attachmentService{
		attachmentRepo:    attachmentRepo,
		testCaseService:   testCaseService,
		testResultService: testResultService,
	}
}

func (s *attachmentService) isValidTargetType(targetType string) bool {
	return targetType == "test_case" || targetType == "test_result"
}

func (s *attachmentService) Create(req CreateAttachmentRequest, projectID uint, files []string) error {
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

	// TODO : 저장 로직 필요

	return nil
}
