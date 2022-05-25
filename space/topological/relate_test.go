package topological

import (
	"testing"

	"github.com/spatial-go/geoos/space"
	"github.com/spatial-go/geoos/space/topograph"
)

func TestTopological_Within(t *testing.T) {
	type args struct {
		A space.Geometry
		B space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"within 1", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := topograph.GetRelationship(topograph.NewTopological)
			got, err := tr.Within(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.Within() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.Within() = %v, want %v", got, tt.want)
			}
		})
	}
}