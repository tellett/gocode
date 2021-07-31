load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "4698e467ed72b09227afcb7b8d289f8ec0eaead3494f7b80aefc6610e0ae855b",
    strip_prefix = "buildtools-master",
    url = "https://github.com/bazelbuild/buildtools/archive/master.zip",
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "3093c29fcdc702b56aef93394ec3666fd90330214b63c81b28ab9249428f2eba",
    strip_prefix = "protobuf-master",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/master.zip"],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "59d5b42ac315e7eadffa944e86e90c2990110a1c8075f1cd145f487e999d22b3",
    strip_prefix = "rules_docker-0.17.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.17.0/rules_docker-v0.17.0.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# See https://github.com/bazelbuild/rules_docker/issues/1847

go_repository(
    name = "com_github_vdemeester_k8s_pkg_credentialprovider",
    importpath = "github.com/vdemeester/k8s-pkg-credentialprovider",
    sum = "h1:7Ajl3rjeYoB5V47jPknnLbyxYlhMXTTJiQsye5aT7f0=",
    version = "v1.21.0-1",
)

go_repository(
    name = "org_golang_x_tools",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/tools",
    sum = "h1:po9/4sTYwZU9lPhi1tOrb4hCv3qrhiQ77LZfGa2OjwY=",
    version = "v0.1.0",
    build_extra_args = [
        "-exclude=**/testdata",
        "-exclude=go/packages/packagestest",
    ],
)

go_rules_dependencies()

go_register_toolchains(
    nogo = "@//:nogo_vet",
    version = "1.16.5",
)

gazelle_dependencies()

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()