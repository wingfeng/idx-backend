package base

import "testing"

func TestPage_ValidateFilters(t *testing.T) {
	tests := []struct {
		name    string
		p       *Page
		wantErr bool
	}{
		{"test1", &Page{Filters: []string{"id = ?"}}, false},
		{"test2", &Page{Filters: []string{"i d = 1;1select?"}}, true},
		{"test2", &Page{Filters: []string{"1=1; select * from ?"}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.ValidateFilters(); (err != nil) != tt.wantErr {
				t.Errorf("Page.ValidateFilters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
