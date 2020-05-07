package section1

import "fmt"

func Template(x int, y string, z float64) string {
	return fmt.Sprintf("%d日の%sは%.1f", x, y, z)
}
