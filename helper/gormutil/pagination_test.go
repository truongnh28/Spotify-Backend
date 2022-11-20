package gormutil

import "testing"

func TestPagination_calculateOffsetLimit(t *testing.T) {
	type fields struct {
		Size   int
		Number int
	}
	tests := []struct {
		name       string
		fields     fields
		wantOffset int
		wantLimit  int
	}{
		{
			name: "test case #1: page size is zero",
			fields: fields{
				Size:   0,
				Number: 1,
			},
			wantOffset: 0,
			wantLimit:  -1,
		},
		{
			name: "test case #3: page size over max limit",
			fields: fields{
				Size:   1001,
				Number: 1,
			},
			wantOffset: 0,
			wantLimit:  1001,
		},
		{
			name: "test case #4: page size less than min page size",
			fields: fields{
				Size:   -2,
				Number: 2,
			},
			wantOffset: 1,
			wantLimit:  -1,
		},
		{
			name: "test case #5: invalid page number",
			fields: fields{
				Size:   2,
				Number: -1,
			},
			wantOffset: 0,
			wantLimit:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paging := &Pagination{
				Size:   tt.fields.Size,
				Number: tt.fields.Number,
			}
			gotOffset, gotLimit := paging.calculateOffsetLimit()
			if gotOffset != tt.wantOffset {
				t.Errorf("Pagination.calculateOffsetLimit() gotOffset = %v, want %v", gotOffset, tt.wantOffset)
			}
			if gotLimit != tt.wantLimit {
				t.Errorf("Pagination.calculateOffsetLimit() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
		})
	}
}
