package basics

import "testing"

func Test_ioFunc(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		want    int
		args    args
		wantErr bool
	}{
		{
			name: "no error",
			args: args{
				"test",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "should return 0",
			args: args{
				"",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ioFunc(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ioFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ioFunc() got = %v, want %v", got, tt.want)
			}
		})
	}
}
