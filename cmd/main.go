package main

import (
	app "github.com/sky0621/go-di"
	"github.com/sky0621/go-di/otherpackage"
)

// プログラム内のどこからでもアクセスできるDIコンテナとしてもよかったが、呼び出し関係の都合上、cyclicインポートを防ぐため引数での引き回しを採用
// mainパッケージに置くべきかも再考
var dicon *app.DIContainer

func init() {
	// テストコードでは、ここをモックに差し替えることで冪等性を担保したコードにする
	dicon = app.NewDIContainer()
	dicon.RegisterFactory(app.CloudSQLAccessor, app.CloudSQLAccessorFactory)
	dicon.RegisterFactory(app.CloudPubSubAccessor, app.CloudPubSubAccessorFactory)
}

func main() {
	otherpackage.Logic(dicon)
}
