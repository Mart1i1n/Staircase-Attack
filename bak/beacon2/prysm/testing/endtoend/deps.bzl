load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")  # gazelle:keep

lighthouse_version = "v3.1.2"
lighthouse_archive_name = "lighthouse-%s-x86_64-unknown-linux-gnu-portable.tar.gz" % lighthouse_version

def e2e_deps():
    http_archive(
        name = "web3signer",
        urls = ["https://artifacts.consensys.net/public/web3signer/raw/names/web3signer.tar.gz/versions/22.8.1/web3signer-22.8.1.tar.gz"],
        sha256 = "ec888222484c4d1b6203bd6d248890adf713f8bf47fb362fb36e8d47a98cb401",
        build_file = "@prysm//testing/endtoend:web3signer.BUILD",
        strip_prefix = "web3signer-22.8.1",
    )

    http_archive(
        name = "lighthouse",
        sha256 = "172bb132d5fdc5bd257d5a66e98d0799498f08cb60502f93a5f4437a70d9c5e0",
        build_file = "@prysm//testing/endtoend:lighthouse.BUILD",
        url = ("https://github.com/sigp/lighthouse/releases/download/%s/" + lighthouse_archive_name) % lighthouse_version,
    )
