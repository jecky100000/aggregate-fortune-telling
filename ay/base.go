/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package ay

type Base struct {
}

var (
	Domain string
)

func init() {
	var yaml Yaml
	yaml.GetConf()
	Domain = yaml.Domain
}
