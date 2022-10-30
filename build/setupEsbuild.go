package build

//* 英語 - English
//* This file is a config file of Esbuild, a js bundler(https://esbuild.github.io/)
//*	日本語 - Japanese
//* これはEsbuild(Javascriptバンドラー)のcnfigファイルになります。詳しくはこちら(https://esbuild.github.io/)

import (
	"io/ioutil"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func setupEsbuild() {
	ioutil.WriteFile("app.js", []byte("let x: number = 1"), 6440)

	result := api.Build(api.BuildOptions{
		Color:               0,
		LogLevel:            0,
		LogLimit:            0,
		LogOverride:         map[string]api.LogLevel{},
		Sourcemap:           0,
		SourceRoot:          "",
		SourcesContent:      0,
		Target:              0,
		Engines:             []api.Engine{{Name: api.EngineNode, Version: "16.18.0"}, {Name: api.EngineChrome, Version: "58"}, {Name: api.EngineFirefox, Version: "57"}, {Name: api.EngineSafari, Version: "11"}, {Name: api.EngineEdge, Version: "16"}},
		Supported:           map[string]bool{},
		MangleProps:         "",
		ReserveProps:        "",
		MangleQuoted:        0,
		MangleCache:         map[string]interface{}{},
		Drop:                0,
		MinifyWhitespace:    true,
		MinifyIdentifiers:   true,
		MinifySyntax:        true,
		Charset:             0,
		TreeShaking:         0,
		IgnoreAnnotations:   false,
		LegalComments:       0,
		JSXMode:             0,
		JSXFactory:          "",
		JSXFragment:         "",
		JSXImportSource:     "",
		JSXDev:              false,
		JSXSideEffects:      false,
		Define:              map[string]string{},
		Pure:                []string{},
		KeepNames:           false,
		GlobalName:          "",
		Bundle:              true,
		PreserveSymlinks:    false,
		Splitting:           false,
		Metafile:            false,
		Outdir:              "",
		Outbase:             "",
		AbsWorkingDir:       "",
		Platform:            api.PlatformNode,
		Format:              0,
		External:            []string{},
		MainFields:          []string{},
		Conditions:          []string{},
		Loader:              map[string]api.Loader{},
		ResolveExtensions:   []string{},
		Tsconfig:            "",
		OutExtensions:       map[string]string{},
		PublicPath:          "",
		Inject:              []string{},
		Banner:              map[string]string{},
		Footer:              map[string]string{},
		NodePaths:           []string{},
		EntryNames:          "",
		ChunkNames:          "",
		AssetNames:          "",
		EntryPoints:         []string{"app.js"},
		Outfile:             "out.js",
		EntryPointsAdvanced: []api.EntryPoint{},
		Stdin:               &api.StdinOptions{},
		Write:               true,
		AllowOverwrite:      false,
		Incremental:         false,
		Plugins:             []api.Plugin{},
		Watch:               &api.WatchMode{},
	})

	if len(result.Errors) > 0 {
		os.Exit(1)
	}
	println("watching...\n")

	<-make(chan bool)
}
