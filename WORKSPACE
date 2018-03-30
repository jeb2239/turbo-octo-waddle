# local_repository(
#     name = "io_bazel_rules_go",
#     path = "/home/barriosj/real_rules_go",
# )

# local_repository(
#     name = "bazel_gazelle",
#     path = "/home/barriosj/bazel_gazelle",
# )


http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.10.2/rules_go-0.10.2.tar.gz",
    sha256 = "4b2c61795ac2eefcb28f3eb8e1cb2d8fb3c2eafa0f6712473bc5f93728f38758",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains","go_repository")
# go_rules_dependencies()
# go_register_toolchains()
# load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    tag = "v1.0.4",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "85f98707c97e11569271e4d9b3d397e079c4f4d0",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "dd2ff4accc098aceecb86b36eaa7829b2a17b1c9",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "bazel_gopath",

    # bazel-gopath ships with a proto file and also a precompiled version of it.
    # The proto file does not include the right Go options, which confuses
    # Gazelle, so prefer the precompiled version.
    build_file_proto_mode = "disable",
    commit = "f83ae8d403d0c826335a5edf4bbd1f7b0cf176e4",
    importpath = "github.com/DarkDNA/bazel-gopath",
)

go_repository(
    name = "golint",
    commit = "6aaf7c34af0f4c36a57e0c429bace4d706d8e931",
    importpath = "github.com/golang/lint",
)

gazelle_dependencies()


go_repository(
    name = "org_uber_go_zap",
    commit = "f85c78b1dd998214c5f2138155b320a4a43fbe36",
    importpath = "go.uber.org/zap",
)

go_repository(
    name = "org_uber_go_atomic",
    commit = "8474b86a5a6f79c443ce4b2992817ff32cf208b8",
    importpath = "go.uber.org/atomic",
)

go_repository(
    name = "org_uber_go_multierr",
    commit = "ddea229ff1dff9e6fe8a6c0344ac73b09e81fce5",
    importpath = "go.uber.org/multierr",
)

go_repository(
    name = "org_uber_go_fx",
    commit = "71ef1b9dd628e2a1ab593a5fcb8c679208b291be",
    importpath = "go.uber.org/fx",
)
