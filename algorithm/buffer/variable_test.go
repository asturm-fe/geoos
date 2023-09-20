// Package buffer define geomtry matrix conversion.
package buffer

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestVariableBuffer(t *testing.T) {
	type args struct {
		lineBuffer    *VariableLineBuffer
		startDistance float64
		endDistance   float64
	}
	tests := []struct {
		name string
		args args
		want matrix.Steric
	}{
		{"simple case", args{&VariableLineBuffer{matrix.LineMatrix{{2.073333263397217, 48.81027603149414}, {1.5225944519042969, 48.45795440673828}}, 8},
			0.001, 0.0012},
			matrix.PolygonMatrix{{{2.0738724106666613, 48.80943381998754},
				{1.5232414286276303, 48.45694375293036},
				{1.523053672023135, 48.45684575129927},
				{1.5228285602907163, 48.456777464401796},
				{1.5225944519042969, 48.456754406738284},
				{1.5223603435178774, 48.456777464401796},
				{1.5221352317854588, 48.45684575129927},
				{1.5219277676246734, 48.45695664320352},
				{1.5217459237668731, 48.457105878600856},
				{1.5215966883695338, 48.45728772245866},
				{1.5214857964652833, 48.45749518661944},
				{1.521417509567813, 48.45772029835186},
				{1.5213944519042968, 48.45795440673828},
				{1.521417509567813, 48.458188515124704},
				{1.5214857964652833, 48.45841362685712},
				{1.5215966883695338, 48.4586210910179},
				{1.5217459237668731, 48.458802934875706},
				{1.5219277676246734, 48.458952170273044},
				{1.5219480936357215, 48.45896545618744},
				{2.072794631506731, 48.81111857270177},
				{2.0729505799648518, 48.811199911026655},
				{2.0731381730752005, 48.811256816774545},
				{2.073333263397217, 48.81127603149414},
				{2.073528353719233, 48.811256816774545},
				{2.073715946829582, 48.811199911026655},
				{2.0738888336302366, 48.81110750110644},
				{2.074040370178403, 48.810983138275326},
				{2.0741647330095194, 48.81083160172716},
				{2.0742571429297283, 48.810658714926504},
				{2.07431404867762, 48.81047112181616},
				{2.0743332633972167, 48.81027603149414},
				{2.07431404867762, 48.810080941172124},
				{2.0742571429297283, 48.80989334806178},
				{2.0741647330095194, 48.80972046126112},
				{2.074040370178403, 48.809568924712956},
				{2.0738724106666613, 48.80943381998754},
			},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.lineBuffer.InterpolatedBuffer(tt.args.startDistance, tt.args.endDistance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VariableBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableDistancesBuffer(t *testing.T) {
	type args struct {
		lineBuffer *VariableLineBuffer
		distances  []float64
	}
	tests := []struct {
		name       string
		args       args
		wantBuffer matrix.Steric
	}{
		{"simple case", args{&VariableLineBuffer{matrix.LineMatrix{{2.073333263397217, 48.81027603149414}, {1.5225944519042969, 48.45795440673828}}, 2},
			[]float64{0.001, 0.0012}},
			matrix.PolygonMatrix{{{2.0738724106666613, 48.80943381998754},
				{1.5232414286276303, 48.45694375293036},
				{1.5225944519042969, 48.456754406738284},
				{1.5217459237668731, 48.457105878600856},
				{1.5213944519042968, 48.45795440673828},
				{1.5217459237668731, 48.458802934875706},
				{1.5219480936357215, 48.45896545618744},
				{2.072794631506731, 48.81111857270177},
				{2.073333263397217, 48.81127603149414},
				{2.074040370178403, 48.810983138275326},
				{2.0743332633972167, 48.81027603149414},
				{2.0738724106666613, 48.80943381998754},
			},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBuffer := tt.args.lineBuffer.DistancesBuffer(tt.args.distances); !reflect.DeepEqual(gotBuffer, tt.wantBuffer) {
				t.Errorf("VariableDistancesBuffer() = %v, want %v", gotBuffer, tt.wantBuffer)
			}
		})
	}
}