#!/usr/bin/python3

import subprocess as sh
from os import environ, mkdir


VARIANTS = [
    ('darwin', 'amd64'),
    ('linux', '386'),
    ('linux', 'amd64'),
    ('windows', '386'),
    ('windows', 'amd64'),
]


def extension(os):
    return '.exe' if os == 'windows' else ''


def executable_name(os, arch):
    return f'./bin/{os}_{arch}_sema{extension(os)}'


def compile():
    for os, arch in VARIANTS:
        environ['GOOS'] = os
        environ['GOARCH'] = arch
        sh.call(['go', 'build', '-o', executable_name(os, arch), '.'])


if __name__ == '__main__':
    try:
        mkdir('bin')
    except FileExistsError:
        pass
    compile()
