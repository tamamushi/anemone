/* vim: set ts=4 sw=4: */

/*
Entry は実行時のブートストラップが含まれる。
CleanArchitectureのInterface Adaptersにおける振る舞いを
実装しています。
*/
package entry

import (
	"anemone/adapter"
)

func CLI() {
	adapter.Bootstrap()
}
