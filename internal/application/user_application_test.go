package application

import (
	"github.com/golang/mock/gomock"
	"github.com/ningzining/cotton-pavilion/internal/domain/model"
	"github.com/ningzining/cotton-pavilion/internal/domain/repository"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/util/cryptoutil"
	"github.com/ningzining/cotton-pavilion/pkg/code"
	"github.com/ningzining/cotton-pavilion/pkg/errors"
	"gorm.io/gorm"
	"reflect"
	"testing"

	"github.com/ningzining/cotton-pavilion/internal/application/types"
	"github.com/ningzining/cotton-pavilion/internal/domain/service"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/store"
)

func Test_userApplication_Login(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	storeFactory := store.NewMockFactory(controller)

	type fields struct {
		Store   store.Factory
		Service service.Service
	}
	type args struct {
		dto types.LoginDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.LoginRet
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Store: storeFactory,
			},
			args: args{
				dto: types.LoginDTO{
					Password: "123456",
					Mobile:   "123456",
				},
			},
			want: &types.LoginRet{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tL25pbmd6aW5pbmcvY290dG9uLXBhdmlsaW9uIiwic3ViIjoiYXV0aCIsImlhdCI6MTcwMzMxMTU2MSwiVXNlciI6eyJJZCI6MCwiVXNlcm5hbWUiOiJtb2NrX3VzZXIifX0.AnHpSvvH2VDrnNFOpuWli0VXw3FDWG6hnYFR0-CHft8",
			},
			wantErr: false,
		},
		{
			name: "password error",
			fields: fields{
				Store: storeFactory,
			},
			args: args{
				dto: types.LoginDTO{
					Password: "123456",
					Mobile:   "123456789",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepository := repository.NewMockUserRepository(controller)
			storeFactory.EXPECT().UserRepository().Return(userRepository)

			userRepository.EXPECT().FindByMobile(tt.args.dto.Mobile).Return(&model.User{
				Model:    gorm.Model{},
				Username: "mock_user",
				Mobile:   tt.args.dto.Mobile,
				Email:    "mock@example.com",
				Password: "ea48576f30be1669971699c09ad05c94",
			}, nil)

			u := &userApplication{
				Store:   tt.fields.Store,
				Service: tt.fields.Service,
			}
			got, err := u.Login(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("userApplication.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Token != "" {

			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("userApplication.Login() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_userApplication_Register(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	storeFactory := store.NewMockFactory(controller)

	type fields struct {
		Store   store.Factory
		Service service.Service
	}
	type args struct {
		dto types.RegisterDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "register success",
			fields: fields{
				Store:   storeFactory,
				Service: nil,
			},
			args: args{
				dto: types.RegisterDTO{
					Email:    "mock@example.com",
					Password: "123456789",
					Username: "mock_user",
					Mobile:   "123456789",
				},
			},
			wantErr: false,
		},
		{
			name: "register error",
			fields: fields{
				Store:   storeFactory,
				Service: nil,
			},
			args: args{
				dto: types.RegisterDTO{
					Email:    "mock@example.com",
					Password: "123456",
					Username: "mock_user",
					Mobile:   "123456",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepository := repository.NewMockUserRepository(controller)
			storeFactory.EXPECT().UserRepository().Return(userRepository)
			user := &model.User{
				Username: tt.args.dto.Username,
				Mobile:   tt.args.dto.Mobile,
				Email:    tt.args.dto.Email,
				Password: cryptoutil.Md5Password(tt.args.dto.Mobile, tt.args.dto.Password),
			}
			if tt.args.dto.Mobile == "123456" {
				userRepository.EXPECT().Save(user).Return(errors.WithCode(code.ErrDatabase, "数据库操作异常"))
			} else {
				userRepository.EXPECT().Save(user).Return(nil)
			}

			u := &userApplication{
				Store:   tt.fields.Store,
				Service: tt.fields.Service,
			}
			if err := u.Register(tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("userApplication.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
