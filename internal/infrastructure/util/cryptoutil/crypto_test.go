package cryptoutil

import "testing"

func TestMd5(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Md5_success",
			args: args{
				source: "1234456",
			},
			want: "4266bf8d3dc65bc84fd3badf2edfdbe7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.source); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5Password(t *testing.T) {
	type args struct {
		mobile   string
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Md5Password_Success",
			args: args{
				mobile:   "123456",
				password: "123456",
			},
			want: "ea48576f30be1669971699c09ad05c94",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Password(tt.args.mobile, tt.args.password); got != tt.want {
				t.Errorf("Md5Password() = %v, want %v", got, tt.want)
			}
		})
	}
}
