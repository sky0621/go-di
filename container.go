package app

// appパッケージがinfrastructureパッケージに依存する形がよいかは再考(cyclicインポートを招く可能性)
import "github.com/sky0621/go-di/infrastructure"

// -------------------------------------------------------------------------
// DIコンテナにFactoryメソッドを格納する際のキーと、そのラインナップ
// -------------------------------------------------------------------------

// StoreKey ... コンテナ格納時のキー
type StoreKey int

// CloudSQLAccessor ...
var CloudSQLAccessor StoreKey = 1

// CloudPubSubAccessor ...
var CloudPubSubAccessor StoreKey = 2

// XXXXEndpointAccessor ...
var XXXXEndpointAccessor StoreKey = 3

// -------------------------------------------------------------------------
// DIコンテナに格納するFactory（外部依存ロジックの注入関数）
// -------------------------------------------------------------------------

// Factory ...
type Factory func(c *DIContainer) Accessor

// CloudSQLAccessorFactory ...
func CloudSQLAccessorFactory(c *DIContainer) Accessor {
	// CloudSQLへのアクセッサならではのロジックがあれば、ここで。
	// AccessorをAbstractFactory形式にして、階層化された依存構造を生成して返すロジックにしてもよいかも。
	return &infrastructure.CloudSQLAccessor{}
}

// CloudPubSubAccessorFactory ...
func CloudPubSubAccessorFactory(c *DIContainer) Accessor {
	return &infrastructure.CloudPubSubAccessor{}
}

// -------------------------------------------------------------------------
// CloudSQLや他の外部API等の依存ロジックを既定するインタフェース
// -------------------------------------------------------------------------

// Accessor ...
type Accessor interface {
	Duck() // マーカーインタフェース用途のため本来は関数不要だが、ダックタイプ形式のため適当な関数を定義
}

// -------------------------------------------------------------------------
// DIコンテナ
// -------------------------------------------------------------------------

// NewDIContainer ...
func NewDIContainer() *DIContainer {
	return &DIContainer{
		container: map[StoreKey]Factory{},
	}
}

// DIContainer ...
type DIContainer struct {
	// 要素追加はmain関数のみで、あとは要素取得のみのため Mutex は使わない
	container map[StoreKey]Factory
}

// RegistFactory ... main関数でのみ呼ばれることを想定しているが、それを強制しているわけではないので堅牢とは言えない
func (c *DIContainer) RegistFactory(k StoreKey, f Factory) {
	c.container[k] = f
}

// GetAccessor ...
func (c *DIContainer) GetAccessor(k StoreKey) Accessor {
	f := c.container[k]
	return f(c)
}
