# Progress

## Middleware

- AuthMiddleware
- RequireProjectManager
- RequireProjectMember

## Users

- Sign up
- Sign in
- Sign out

## Projects

- Create
- GetAll

## Project-Members

- Create
- Find (projectID)
- Delete
- Update Role To Member
- Update Role To Manager

## Test Status

- Create Default (프로젝트 생성 시에 기본 Status 생성)
- Find (projectID)
- Create
- Delete

## Test Case

- Create ( Check Status ID: O / Check Assigned User ID : O )
- FindOne(projectID, testCaseID)
- Find (projectID)
- Update CurrentStatusID ( with log )
- Update CurrentAssigneeID ( with log )

## Test Result

- Create
- FindOne(projectID, testCaseID, testResultID)
- Find(projectID, testCaseID)
- FindOneByID(testResultID)

## Attachment

- Create (form-xxx) ( required files )
- FindOne (target_type, targetID, projectID, attachmentID)
- Find (target_type, targetID, projectID)
- Download (target_type, targetID, projectID. attachmentID) -> (FindOne)
