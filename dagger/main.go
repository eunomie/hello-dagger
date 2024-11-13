// A generated module for HelloDagger functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/hello-dagger/internal/dagger"
)

type HelloDagger struct{}

//func (m *HelloDagger) Publish(ctx context.Context, source *dagger.Directory) (string, error) {
//	_, err := m.Test(ctx, source)
//	if err != nil {
//		return "", err
//	}
//	return m.Build(source).Publish(ctx, fmt.Sprintf("ttl.sh/hello-dagger-%.0f", math.Floor(rand.Float64()*10000000)))
//}
//
//func (m *HelloDagger) Build(source *dagger.Directory) *dagger.Container {
//	build := m.BuildEnv(source).
//		WithExec([]string{"npm", "run", "build"}).
//		Directory("./dist")
//	return dag.Container().From("nginx:1.25-alpine").
//		WithDirectory("/usr/share/ngingx/html", build).
//		WithExposedPort(80)
//}

func (m *HelloDagger) Test(ctx context.Context, source *dagger.Directory) (string, error) {
	return m.BuildEnv(source).
		WithExec([]string{"npm", "run", "test:unit", "run"}).
		Stdout(ctx)
}

func (m *HelloDagger) BuildEnv(source *dagger.Directory) *dagger.Container {
	//nodeCache := dag.CacheVolume("node")
	return dag.Container().
		From("node:21-slim").
		WithDirectory("/src", source).
		//WithMountedCache("/root/.npm", nodeCache).
		WithWorkdir("/src").
		WithExec([]string{"npm", "install"})
}

//func (m *HelloDagger) Term() *dagger.Container {
//	return dag.Container().
//		From("alpine:latest").
//		Terminal().
//		WithExec([]string{"sh", "-c", "echo hello world > /foo && cat /foo"}).
//		Terminal()
//}
//
//func (m *HelloDagger) SimpleDirectory(ctx context.Context) (string, error) {
//	return dag.
//		Git("https://github.com/dagger/dagger.git").
//		Head().
//		Tree().
//		Terminal().
//		File("README.md").
//		Contents(ctx)
//}
//
//func (m *HelloDagger) AdvancedDirectory(ctx context.Context) (string, error) {
//	return dag.
//		Git("https://github.com/dagger/dagger.git").
//		Head().
//		Tree().
//		Terminal(dagger.DirectoryTerminalOpts{
//			Container: dag.Container().From("ubuntu"),
//			Cmd:       []string{"/bin/bash"},
//		}).
//		File("README.md").
//		Contents(ctx)
//}
