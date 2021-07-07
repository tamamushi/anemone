/* vim: set ts=4 sw=4: */

/*
Entry は実行時のブートストラップが含まれる。
CleanArchitectureのInterface Adaptersにおける振る舞いを
実装しています。
*/
package entry

import (
	"context"

	"anemone/adapter"
	"anemone/adapter/presenter"
)

type Event struct {
	Arguments map[string]string `json:"arguments"`
	Identity  string            `json:"identity"`
	FieldName string            `json:"fieldName"`
	TypeName  string            `json:"typeName"`
}

func Lambda(ctx context.Context, event Event) {

	p := presenter.NewLambdaPresenter(event.FieldName)
	// FieldNameで"user create"の決定
	// gatewayをSingletonで抽出
	// event.FieldNameにより対応する
	// responseを gatewayはinjectionされる
	p.SetArgument(event.Arguments)
	adapter.Bootstrap(p)
	//return gateway.output(), nil
}
