load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test","go_binary")
go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
    ],
    importpath = "coverage/src",
    visibility = ["//visibility:public"],
    deps = ["@org_uber_go_zap//:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["main_test.go"],
    embed = [":go_default_library"],
)

go_binary(
    name = "main",
    srcs=["main.go"],
    visibility = ["//visibility:public"],
    deps=[]
)

# container_image(
#     name = "app",
#     base = "@official_ubuntu//image",
#     cmd=["main"]
# )
