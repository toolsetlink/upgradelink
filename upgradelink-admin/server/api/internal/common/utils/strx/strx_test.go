package strx

import "testing"

func TestStructToString(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				v: struct {
					Name      string
					Age       int
					Sex       bool
					Weight    float64
					CourseIds []int
					Identify  struct {
						Id   int
						Name string
					}
				}{
					Name:      "Jacky",
					Age:       18,
					Sex:       true,
					Weight:    80.5,
					CourseIds: []int{1, 2, 3},
					Identify: struct {
						Id   int
						Name string
					}{
						Id:   1,
						Name: "Jack",
					},
				},
			},
			want: "Jacky_18_true_80.5_[1 2 3]_1_Jack",
		},
		{
			name: "test2",
			args: args{
				v: struct{}{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToString(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StructToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
