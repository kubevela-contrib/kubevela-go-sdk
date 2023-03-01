package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chivalryq/vela-go-sdk/pkg/client"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	// TODO(chivalryq) pre-import all types so they can be all registered
	"github.com/chivalryq/vela-go-sdk/pkg/apis/component/webservice"
	_ "github.com/chivalryq/vela-go-sdk/pkg/apis/trait/init-container"
	_ "github.com/chivalryq/vela-go-sdk/pkg/apis/trait/resource"
)

func main() {
	clt := client.NewDefaultOrDie()
	ctx := context.Background()
	app, err := clt.Get(ctx, runtimeclient.ObjectKey{Name: "test-app", Namespace: "default"})
	if err != nil {
		panic(err)
	}
	marshal, err := json.Marshal(app.Build())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))

	wss := app.GetComponentsByType(webservice.WebserviceType)
	ws := wss[0]
	w, ok := ws.(*webservice.WebserviceComponent)
	if !ok {
		panic("not webservice")
	}
	w.SetImage("nginx:1.20")
	app.SetComponents(w)
	// Then you can re-apply the application

	err = clt.Update(ctx, app)
	if err != nil {
		panic(err)
	}

	fmt.Println("now the application")
	now, err := clt.Get(ctx, runtimeclient.ObjectKey{Name: "test-app", Namespace: "default"})
	if err != nil {
		panic(err)
	}

	marshal, err = json.Marshal(now.Build())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
