package query

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/lib/pq"
	"golang.org/x/text/language"

	"github.com/caos/zitadel/internal/domain"
	errs "github.com/caos/zitadel/internal/errors"
)

var (
	userQuery = `SELECT zitadel.projections.users.id,` +
		` zitadel.projections.users.creation_date,` +
		` zitadel.projections.users.change_date,` +
		` zitadel.projections.users.resource_owner,` +
		` zitadel.projections.users.sequence,` +
		` zitadel.projections.users.state,` +
		` zitadel.projections.users.type,` +
		` zitadel.projections.users.username,` +
		` login_names.loginnames,` +
		` preferred_login_name.login_name,` +
		` zitadel.projections.users_humans.user_id,` +
		` zitadel.projections.users_humans.first_name,` +
		` zitadel.projections.users_humans.last_name,` +
		` zitadel.projections.users_humans.nick_name,` +
		` zitadel.projections.users_humans.display_name,` +
		` zitadel.projections.users_humans.preferred_language,` +
		` zitadel.projections.users_humans.gender,` +
		` zitadel.projections.users_humans.avatar_key,` +
		` zitadel.projections.users_humans.email,` +
		` zitadel.projections.users_humans.is_email_verified,` +
		` zitadel.projections.users_humans.phone,` +
		` zitadel.projections.users_humans.is_phone_verified,` +
		` zitadel.projections.users_machines.user_id,` +
		` zitadel.projections.users_machines.name,` +
		` zitadel.projections.users_machines.description` +
		` FROM zitadel.projections.users` +
		` LEFT JOIN zitadel.projections.users_humans ON zitadel.projections.users.id = zitadel.projections.users_humans.user_id` +
		` LEFT JOIN zitadel.projections.users_machines ON zitadel.projections.users.id = zitadel.projections.users_machines.user_id` +
		` LEFT JOIN` +
		` (SELECT login_names.user_id, ARRAY_AGG(login_names.login_name) as loginnames` +
		` FROM zitadel.projections.login_names as login_names` +
		` GROUP BY login_names.user_id) as login_names` +
		` on login_names.user_id = zitadel.projections.users.id` +
		` LEFT JOIN` +
		` (SELECT preferred_login_name.user_id, preferred_login_name.login_name FROM zitadel.projections.login_names as preferred_login_name WHERE preferred_login_name.is_primary = $1) as preferred_login_name` +
		` on preferred_login_name.user_id = zitadel.projections.users.id`
	userCols = []string{
		"id",
		"creation_date",
		"change_date",
		"resource_owner",
		"sequence",
		"state",
		"type",
		"username",
		"loginnames",
		"login_name",
		//human
		"user_id",
		"first_name",
		"last_name",
		"nick_name",
		"display_name",
		"preferred_language",
		"gender",
		"avatar_key",
		"email",
		"is_email_verified",
		"phone",
		"is_phone_verified",
		//machine
		"user_id",
		"name",
		"description",
	}
	profileQuery = `SELECT zitadel.projections.users.id,` +
		` zitadel.projections.users.creation_date,` +
		` zitadel.projections.users.change_date,` +
		` zitadel.projections.users.resource_owner,` +
		` zitadel.projections.users.sequence,` +
		//` zitadel.projections.users.state,` + //TODO:
		` zitadel.projections.users_humans.user_id,` +
		` zitadel.projections.users_humans.first_name,` +
		` zitadel.projections.users_humans.last_name,` +
		` zitadel.projections.users_humans.nick_name,` +
		` zitadel.projections.users_humans.display_name,` +
		` zitadel.projections.users_humans.preferred_language,` +
		` zitadel.projections.users_humans.gender,` +
		` zitadel.projections.users_humans.avatar_key` +
		` FROM zitadel.projections.users` +
		` LEFT JOIN zitadel.projections.users_humans ON zitadel.projections.users.id = zitadel.projections.users_humans.user_id`
	profileCols = []string{
		"id",
		"creation_date",
		"change_date",
		"resource_owner",
		"sequence",
		//"state", //TODO:
		"user_id",
		"first_name",
		"last_name",
		"nick_name",
		"display_name",
		"preferred_language",
		"gender",
		"avatar_key",
	}
	emailQuery = `SELECT zitadel.projections.users.id,` +
		` zitadel.projections.users.creation_date,` +
		` zitadel.projections.users.change_date,` +
		` zitadel.projections.users.resource_owner,` +
		` zitadel.projections.users.sequence,` +
		//` zitadel.projections.users.state,` + //TODO:
		` zitadel.projections.users_humans.user_id,` +
		` zitadel.projections.users_humans.email,` +
		` zitadel.projections.users_humans.is_email_verified` +
		` FROM zitadel.projections.users` +
		` LEFT JOIN zitadel.projections.users_humans ON zitadel.projections.users.id = zitadel.projections.users_humans.user_id`
	emailCols = []string{
		"id",
		"creation_date",
		"change_date",
		"resource_owner",
		"sequence",
		//"state", //TODO:
		"user_id",
		"email",
		"is_email_verified",
	}
	phoneQuery = `SELECT zitadel.projections.users.id,` +
		` zitadel.projections.users.creation_date,` +
		` zitadel.projections.users.change_date,` +
		` zitadel.projections.users.resource_owner,` +
		` zitadel.projections.users.sequence,` +
		//` zitadel.projections.users.state,` + //TODO:
		` zitadel.projections.users_humans.user_id,` +
		` zitadel.projections.users_humans.phone,` +
		` zitadel.projections.users_humans.is_phone_verified` +
		` FROM zitadel.projections.users` +
		` LEFT JOIN zitadel.projections.users_humans ON zitadel.projections.users.id = zitadel.projections.users_humans.user_id`
	phoneCols = []string{
		"id",
		"creation_date",
		"change_date",
		"resource_owner",
		"sequence",
		//"state", //TODO:
		"user_id",
		"phone",
		"is_phone_verified",
	}

	userUniqueQuery = `SELECT zitadel.projections.users.id,` +
		` zitadel.projections.users.state,` +
		` zitadel.projections.users.username,` +
		//` login_names.login_names,` +
		//` preferred_login_name.login_name,` +
		` zitadel.projections.users_humans.user_id,` +
		` zitadel.projections.users_humans.email,` +
		` zitadel.projections.users_humans.is_email_verified` +
		` FROM zitadel.projections.users` +
		` LEFT JOIN zitadel.projections.users_humans ON zitadel.projections.users.id = zitadel.projections.users_humans.user_id`
	//` LEFT JOIN` +
	//` (SELECT login_names.user_id, ARRAY_AGG(login_names.login_name) as login_names` +
	//` FROM zitadel.projections.login_names as login_names` +
	//` GROUP BY login_names.user_id) as login_names` +
	//` on login_names.user_id = zitadel.projections.users.id` +
	//` LEFT JOIN` +
	//` (SELECT preferred_login_name.user_id, preferred_login_name.login_name FROM zitadel.projections.login_names as preferred_login_name WHERE preferred_login_name.is_primary = $1) as preferred_login_name` +
	//` on preferred_login_name.user_id = zitadel.projections.users.id`
	userUniqueCols = []string{
		"id",
		"state",
		"username",
		//"login_names",
		//"login_name",
		//human
		"user_id",
		"email",
		"is_email_verified",
	}
	usersQuery = `SELECT zitadel.projections.users.id,` +
		` zitadel.projections.users.creation_date,` +
		` zitadel.projections.users.change_date,` +
		` zitadel.projections.users.resource_owner,` +
		` zitadel.projections.users.sequence,` +
		` zitadel.projections.users.state,` +
		` zitadel.projections.users.type,` +
		` zitadel.projections.users.username,` +
		` login_names.loginnames,` +
		` preferred_login_name.login_name,` +
		` zitadel.projections.users_humans.user_id,` +
		` zitadel.projections.users_humans.first_name,` +
		` zitadel.projections.users_humans.last_name,` +
		` zitadel.projections.users_humans.nick_name,` +
		` zitadel.projections.users_humans.display_name,` +
		` zitadel.projections.users_humans.preferred_language,` +
		` zitadel.projections.users_humans.gender,` +
		` zitadel.projections.users_humans.avatar_key,` +
		` zitadel.projections.users_humans.email,` +
		` zitadel.projections.users_humans.is_email_verified,` +
		` zitadel.projections.users_humans.phone,` +
		` zitadel.projections.users_humans.is_phone_verified,` +
		` zitadel.projections.users_machines.user_id,` +
		` zitadel.projections.users_machines.name,` +
		` zitadel.projections.users_machines.description,` +
		` COUNT(*) OVER ()` +
		` FROM zitadel.projections.users` +
		` LEFT JOIN zitadel.projections.users_humans ON zitadel.projections.users.id = zitadel.projections.users_humans.user_id` +
		` LEFT JOIN zitadel.projections.users_machines ON zitadel.projections.users.id = zitadel.projections.users_machines.user_id` +
		` LEFT JOIN` +
		` (SELECT login_names.user_id, ARRAY_AGG(login_names.login_name) as loginnames` +
		` FROM zitadel.projections.login_names as login_names` +
		` GROUP BY login_names.user_id) as login_names` +
		` on login_names.user_id = zitadel.projections.users.id` +
		` LEFT JOIN` +
		` (SELECT preferred_login_name.user_id, preferred_login_name.login_name FROM zitadel.projections.login_names as preferred_login_name WHERE preferred_login_name.is_primary = $1) as preferred_login_name` +
		` on preferred_login_name.user_id = zitadel.projections.users.id`
	usersCols = []string{
		"id",
		"creation_date",
		"change_date",
		"resource_owner",
		"sequence",
		"state",
		"type",
		"username",
		"loginnames",
		"login_name",
		//human
		"user_id",
		"first_name",
		"last_name",
		"nick_name",
		"display_name",
		"preferred_language",
		"gender",
		"avatar_key",
		"email",
		"is_email_verified",
		"phone",
		"is_phone_verified",
		//machine
		"user_id",
		"name",
		"description",
		"count",
	}
)

func Test_UserPrepares(t *testing.T) {
	type want struct {
		sqlExpectations sqlExpectation
		err             checkErr
	}
	tests := []struct {
		name    string
		prepare interface{}
		want    want
		object  interface{}
	}{
		{
			name:    "prepareUserQuery no result",
			prepare: prepareUserQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(userQuery),
					nil,
					nil,
				),
				err: func(err error) (error, bool) {
					if !errs.IsNotFound(err) {
						return fmt.Errorf("err should be zitadel.NotFoundError got: %w", err), false
					}
					return nil, true
				},
			},
			object: (*User)(nil),
		},
		{
			name:    "prepareUserQuery human found",
			prepare: prepareUserQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(userQuery),
					userCols,
					[]driver.Value{
						"id",
						testNow,
						testNow,
						"resource_owner",
						uint64(20211108),
						domain.UserStateActive,
						domain.UserTypeHuman,
						"username",
						pq.StringArray{"login_name1", "login_name2"},
						"login_name1",
						//human
						"id",
						"first_name",
						"last_name",
						"nick_name",
						"display_name",
						"de",
						domain.GenderUnspecified,
						"avatar_key",
						"email",
						true,
						"phone",
						true,
						//machine
						nil,
						nil,
						nil,
					},
				),
			},
			object: &User{
				ID:                 "id",
				CreationDate:       testNow,
				ChangeDate:         testNow,
				ResourceOwner:      "resource_owner",
				Sequence:           20211108,
				State:              domain.UserStateActive,
				Type:               domain.UserTypeHuman,
				Username:           "username",
				LoginNames:         []string{"login_name1", "login_name2"},
				PreferredLoginName: "login_name1",
				Human: &Human{
					FirstName:         "first_name",
					LastName:          "last_name",
					NickName:          "nick_name",
					DisplayName:       "display_name",
					AvatarKey:         "avatar_key",
					PreferredLanguage: language.German,
					Gender:            domain.GenderUnspecified,
					Email:             "email",
					IsEmailVerified:   true,
					Phone:             "phone",
					IsPhoneVerified:   true,
				},
			},
		},
		{
			name:    "prepareUserQuery machine found",
			prepare: prepareUserQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(userQuery),
					userCols,
					[]driver.Value{
						"id",
						testNow,
						testNow,
						"resource_owner",
						uint64(20211108),
						domain.UserStateActive,
						domain.UserTypeMachine,
						"username",
						pq.StringArray{"login_name1", "login_name2"},
						"login_name1",
						//human
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						//machine
						"id",
						"name",
						"description",
					},
				),
			},
			object: &User{
				ID:                 "id",
				CreationDate:       testNow,
				ChangeDate:         testNow,
				ResourceOwner:      "resource_owner",
				Sequence:           20211108,
				State:              domain.UserStateActive,
				Type:               domain.UserTypeMachine,
				Username:           "username",
				LoginNames:         []string{"login_name1", "login_name2"},
				PreferredLoginName: "login_name1",
				Machine: &Machine{
					Name:        "name",
					Description: "description",
				},
			},
		},
		{
			name:    "prepareUserQuery sql err",
			prepare: prepareUserQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(userQuery),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
		{
			name:    "prepareProfileQuery no result",
			prepare: prepareProfileQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(profileQuery),
					nil,
					nil,
				),
				err: func(err error) (error, bool) {
					if !errs.IsNotFound(err) {
						return fmt.Errorf("err should be zitadel.NotFoundError got: %w", err), false
					}
					return nil, true
				},
			},
			object: (*Profile)(nil),
		},
		{
			name:    "prepareProfileQuery human found",
			prepare: prepareProfileQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(profileQuery),
					profileCols,
					[]driver.Value{
						"id",
						testNow,
						testNow,
						"resource_owner",
						uint64(20211108),
						"id",
						"first_name",
						"last_name",
						"nick_name",
						"display_name",
						"de",
						domain.GenderUnspecified,
						"avatar_key",
					},
				),
			},
			object: &Profile{
				ID:                "id",
				CreationDate:      testNow,
				ChangeDate:        testNow,
				ResourceOwner:     "resource_owner",
				Sequence:          20211108,
				FirstName:         "first_name",
				LastName:          "last_name",
				NickName:          "nick_name",
				DisplayName:       "display_name",
				AvatarKey:         "avatar_key",
				PreferredLanguage: language.German,
				Gender:            domain.GenderUnspecified,
			},
		},
		{
			name:    "prepareProfileQuery not human found (error)",
			prepare: prepareProfileQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(profileQuery),
					profileCols,
					[]driver.Value{
						"id",
						testNow,
						testNow,
						"resource_owner",
						uint64(20211108),
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
					},
				),
				err: func(err error) (error, bool) {
					if !errs.IsPreconditionFailed(err) {
						return fmt.Errorf("err should be zitadel.PredconditionError got: %w", err), false
					}
					return nil, true
				},
			},
			object: (*Profile)(nil),
		},
		{
			name:    "prepareProfileQuery sql err",
			prepare: prepareProfileQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(profileQuery),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
		{
			name:    "prepareEmailQuery no result",
			prepare: prepareEmailQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(emailQuery),
					nil,
					nil,
				),
				err: func(err error) (error, bool) {
					if !errs.IsNotFound(err) {
						return fmt.Errorf("err should be zitadel.NotFoundError got: %w", err), false
					}
					return nil, true
				},
			},
			object: (*Email)(nil),
		},
		{
			name:    "prepareEmailQuery human found",
			prepare: prepareEmailQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(emailQuery),
					emailCols,
					[]driver.Value{
						"id",
						testNow,
						testNow,
						"resource_owner",
						uint64(20211108),
						//domain.UserStateActive,
						"id",
						"email",
						true,
					},
				),
			},
			object: &Email{
				ID:            "id",
				CreationDate:  testNow,
				ChangeDate:    testNow,
				ResourceOwner: "resource_owner",
				Sequence:      20211108,
				//State:              domain.UserStateActive,
				Email:      "email",
				IsVerified: true,
			},
		},
		{
			name:    "prepareEmailQuery not human found (error)",
			prepare: prepareEmailQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(emailQuery),
					emailCols,
					[]driver.Value{
						"id",
						testNow,
						testNow,
						"resource_owner",
						uint64(20211108),
						//domain.UserStateActive,
						nil,
						nil,
						nil,
					},
				),
				err: func(err error) (error, bool) {
					if !errs.IsPreconditionFailed(err) {
						return fmt.Errorf("err should be zitadel.PredconditionError got: %w", err), false
					}
					return nil, true
				},
			},
			object: (*Email)(nil),
		},
		{
			name:    "prepareEmailQuery sql err",
			prepare: prepareEmailQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(emailQuery),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
		{
			name:    "preparePhoneQuery no result",
			prepare: preparePhoneQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(phoneQuery),
					nil,
					nil,
				),
				err: func(err error) (error, bool) {
					if !errs.IsNotFound(err) {
						return fmt.Errorf("err should be zitadel.NotFoundError got: %w", err), false
					}
					return nil, true
				},
			},
			object: (*Phone)(nil),
		},
		{
			name:    "preparePhoneQuery human found",
			prepare: preparePhoneQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(phoneQuery),
					phoneCols,
					[]driver.Value{
						"id",
						testNow,
						testNow,
						"resource_owner",
						uint64(20211108),
						//domain.UserStateActive,
						"id",
						"phone",
						true,
					},
				),
			},
			object: &Phone{
				ID:            "id",
				CreationDate:  testNow,
				ChangeDate:    testNow,
				ResourceOwner: "resource_owner",
				Sequence:      20211108,
				//State:              domain.UserStateActive,
				Phone:      "phone",
				IsVerified: true,
			},
		},
		{
			name:    "preparePhoneQuery not human found (error)",
			prepare: preparePhoneQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(phoneQuery),
					phoneCols,
					[]driver.Value{
						"id",
						testNow,
						testNow,
						"resource_owner",
						uint64(20211108),
						//domain.UserStateActive,
						nil,
						nil,
						nil,
					},
				),
				err: func(err error) (error, bool) {
					if !errs.IsPreconditionFailed(err) {
						return fmt.Errorf("err should be zitadel.PredconditionError got: %w", err), false
					}
					return nil, true
				},
			},
			object: (*Phone)(nil),
		},
		{
			name:    "preparePhoneQuery sql err",
			prepare: preparePhoneQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(phoneQuery),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
		{
			name:    "prepareUserUniqueQuery no result",
			prepare: prepareUserUniqueQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(userUniqueQuery),
					nil,
					nil,
				),
				err: func(err error) (error, bool) {
					if !errs.IsNotFound(err) {
						return fmt.Errorf("err should be zitadel.NotFoundError got: %w", err), false
					}
					return nil, true
				},
			},
			object: true,
		},
		{
			name:    "prepareUserUniqueQuery found",
			prepare: prepareUserUniqueQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(userUniqueQuery),
					userUniqueCols,
					[]driver.Value{
						"id",
						domain.UserStateActive,
						"username",
						"id",
						"email",
						true,
					},
				),
			},
			object: false,
		},
		{
			name:    "prepareUserUniqueQuery sql err",
			prepare: prepareUserUniqueQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(userUniqueQuery),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
		{
			name:    "prepareUsersQuery no result",
			prepare: prepareUsersQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(usersQuery),
					nil,
					nil,
				),
				err: func(err error) (error, bool) {
					if !errs.IsNotFound(err) {
						return fmt.Errorf("err should be zitadel.NotFoundError got: %w", err), false
					}
					return nil, true
				},
			},
			object: &Users{Users: []*User{}},
		},
		{
			name:    "prepareUsersQuery one result",
			prepare: prepareUsersQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(usersQuery),
					usersCols,
					[][]driver.Value{
						{
							"id",
							testNow,
							testNow,
							"resource_owner",
							uint64(20211108),
							domain.UserStateActive,
							domain.UserTypeHuman,
							"username",
							pq.StringArray{"login_name1", "login_name2"},
							"login_name1",
							//human
							"id",
							"first_name",
							"last_name",
							"nick_name",
							"display_name",
							"de",
							domain.GenderUnspecified,
							"avatar_key",
							"email",
							true,
							"phone",
							true,
							//machine
							nil,
							nil,
							nil,
						},
					},
				),
			},
			object: &Users{
				SearchResponse: SearchResponse{
					Count: 1,
				},
				Users: []*User{
					{
						ID:                 "id",
						CreationDate:       testNow,
						ChangeDate:         testNow,
						ResourceOwner:      "resource_owner",
						Sequence:           20211108,
						State:              domain.UserStateActive,
						Type:               domain.UserTypeHuman,
						Username:           "username",
						LoginNames:         []string{"login_name1", "login_name2"},
						PreferredLoginName: "login_name1",
						Human: &Human{
							FirstName:         "first_name",
							LastName:          "last_name",
							NickName:          "nick_name",
							DisplayName:       "display_name",
							AvatarKey:         "avatar_key",
							PreferredLanguage: language.German,
							Gender:            domain.GenderUnspecified,
							Email:             "email",
							IsEmailVerified:   true,
							Phone:             "phone",
							IsPhoneVerified:   true,
						},
					},
				},
			},
		},
		{
			name:    "prepareUsersQuery multiple results",
			prepare: prepareUsersQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(usersQuery),
					usersCols,
					[][]driver.Value{
						{
							"id",
							testNow,
							testNow,
							"resource_owner",
							uint64(20211108),
							domain.UserStateActive,
							domain.UserTypeHuman,
							"username",
							pq.StringArray{"login_name1", "login_name2"},
							"login_name1",
							//human
							"id",
							"first_name",
							"last_name",
							"nick_name",
							"display_name",
							"de",
							domain.GenderUnspecified,
							"avatar_key",
							"email",
							true,
							"phone",
							true,
							//machine
							nil,
							nil,
							nil,
						},
						{
							"id",
							testNow,
							testNow,
							"resource_owner",
							uint64(20211108),
							domain.UserStateActive,
							domain.UserTypeMachine,
							"username",
							pq.StringArray{"login_name1", "login_name2"},
							"login_name1",
							//human
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							//machine
							"id",
							"name",
							"description",
						},
					},
				),
			},
			object: &Users{
				SearchResponse: SearchResponse{
					Count: 2,
				},
				Users: []*User{
					{
						ID:                 "id",
						CreationDate:       testNow,
						ChangeDate:         testNow,
						ResourceOwner:      "resource_owner",
						Sequence:           20211108,
						State:              domain.UserStateActive,
						Type:               domain.UserTypeHuman,
						Username:           "username",
						LoginNames:         []string{"login_name1", "login_name2"},
						PreferredLoginName: "login_name1",
						Human: &Human{
							FirstName:         "first_name",
							LastName:          "last_name",
							NickName:          "nick_name",
							DisplayName:       "display_name",
							AvatarKey:         "avatar_key",
							PreferredLanguage: language.German,
							Gender:            domain.GenderUnspecified,
							Email:             "email",
							IsEmailVerified:   true,
							Phone:             "phone",
							IsPhoneVerified:   true,
						},
					},
					{
						ID:                 "id",
						CreationDate:       testNow,
						ChangeDate:         testNow,
						ResourceOwner:      "resource_owner",
						Sequence:           20211108,
						State:              domain.UserStateActive,
						Type:               domain.UserTypeMachine,
						Username:           "username",
						LoginNames:         []string{"login_name1", "login_name2"},
						PreferredLoginName: "login_name1",
						Machine: &Machine{
							Name:        "name",
							Description: "description",
						},
					},
				},
			},
		},
		{
			name:    "prepareUsersQuery sql err",
			prepare: prepareUsersQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(usersQuery),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertPrepare(t, tt.prepare, tt.object, tt.want.sqlExpectations, tt.want.err)
		})
	}
}
