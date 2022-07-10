package utils

import "fmt"

func GenerateLinkImage(fileName string) string {
	return fmt.Sprintf("https://jgjyjvyldoamqndazixl.supabase.co/storage/v1/object/public/foto-proker/%s", fileName)
}
