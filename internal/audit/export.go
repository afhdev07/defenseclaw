package audit

import "fmt"

func (s *Store) ExportJSON(_ string) error {
	return fmt.Errorf("audit: JSON export not yet implemented — coming in iteration 5")
}

func (s *Store) ExportCSV(_ string) error {
	return fmt.Errorf("audit: CSV export not yet implemented — coming in iteration 5")
}
