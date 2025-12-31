# Database

## User

사용자 정보

- user_id
- name
- email
- created_at

## Project

프로젝트

- project_id
- name
- description
- created_at

## ProjectMember

프로젝트 참여자

- project_member_id
- project_id
- user_id
- project_role (MANAGER / MEMBER)
- joined_at

## TestCase

테스트 케이스

- test_case_id
- project_id
- title
- description
- current_status_id
- current_assignee_id
- due_date
- created_at
- updated_at

## TestStatus

프로젝트별 테스트 상태

- status_id
- project_id
- name
- is_active

## TestResult

테스트 케이스 테스트 실행 기록

- test_result_id
- test_case_id
- result (PASS / FAIL / BLOCKED)
- comment
- created_by

- - 동일 테스트에 대해서 여러 번 테스트가 실행될 수 있기에 분리

## TestStatusHistory

테스트케이스 상태 변경 이력

- history_id
- test_case_id
- from_status_id
- to_status_id
- changed_by
- changed_at

## TestAssigneeHistory

테스트 케이스 담당자 변경 이력

- assignee_history_id
- test_case_id
- from_user_id
- to_user_id
- changed_by

## Attachment

테스트 케이스 증빙 파일

- attachment_id
- target_type (TEST_RESULT / TEST_CASE)
- target_id
- file_path
- uploaded_at

## ERD

[URL](https://dbdiagram.io/d/6953503939fa3db27bc7bc82)

---

# Flow
