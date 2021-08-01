package main

import "testing"

func TestList_remove(t *testing.T) {
	type fields struct {
		begin  *ListNode
		end    *ListNode
		length int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			to := &List{
				begin:  tt.fields.begin,
				end:    tt.fields.end,
				length: tt.fields.length,
			}
		})
	}
}
